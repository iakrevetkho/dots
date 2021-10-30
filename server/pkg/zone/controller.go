package zone

import (
	"errors"
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/golang/geo/s2"
	"github.com/google/uuid"
	"github.com/iakrevetkho/archaeopteryx/logger"
	"github.com/iakrevetkho/dots/server/pkg/utils/geo"
	"github.com/sirupsen/logrus"
	"github.com/tjgq/broadcast"
)

const (
	zoneTickPeriod = time.Millisecond * 100
)

// Functions as variable required for unit tests
var (
	randFloat = rand.Float64
	timeNow   = time.Now
	// This mutex is required to prevent race in unit tests
	timeNowMx = sync.RWMutex{}
)

type Controller struct {
	log *logrus.Entry

	minZoneRadiusInM           uint32
	zoneSpeedInMetersPerSecond float64

	prevZone             *Zone
	currentZone          *Zone
	nextZone             *Zone
	nextZoneCreationTime *time.Time

	nextZonePeriod time.Duration
	nextZoneTimer  *time.Timer

	// Delay between next zone creation and zone tick
	nextZoneDelay      time.Duration
	nextZoneDelayTimer *time.Timer

	zoneTicker *time.Ticker

	// Channel for sending zone event (one of NextZone, NextZoneTick, ZoneTickEnd)
	ZoneEventBroadcaster *broadcast.Broadcaster
}

func NewController(spotId uuid.UUID, spotPosition s2.LatLng, spotRadiusInM uint32, minZoneRadiusInM uint32, nextZonePeriod time.Duration, nextZoneDelay time.Duration, zoneSpeedInKmPerH float64) *Controller {
	c := new(Controller)
	c.log = logger.CreateLogger("zone-controller-" + spotId.String())
	c.minZoneRadiusInM = minZoneRadiusInM
	c.zoneSpeedInMetersPerSecond = zoneSpeedInKmPerH * 1000 / 3600
	c.currentZone = NewZone(spotPosition, spotRadiusInM, minZoneRadiusInM)
	c.nextZonePeriod = nextZonePeriod
	c.nextZoneDelay = nextZoneDelay
	c.ZoneEventBroadcaster = broadcast.New(0)

	return c
}

// 200m
// next zone - 100m after period
// Damage - 0.15.. hp / s

// 100m
// next zone - 50m after period
// Damage - 0.3175 hp / s

// 50m
// next zone - 25m after period
// Damage - 0.625 hp / s

// 25m
// next zone - 12.5m after period
// Damage - 1.25 hp / s

// 12.5m
// next zone - 0m after period
// Damage - 2.5 hp / s

// 0m
// next zone - none
// Damage - 5 hp / s

// Create next zone
func (c *Controller) Next(now time.Time) {
	c.nextZone = nextZone(c.currentZone, c.minZoneRadiusInM)
	c.nextZoneCreationTime = &now
	// Also save current zone as previous zone
	c.prevZone = c.currentZone
}

// Approximate zone to next zone
//
// Function returns current zone, flag about last tick (true if it was last tick), error
func (c *Controller) Tick(now time.Time) (*Zone, bool, error) {
	// Check that we have inited new zone
	if c.nextZone == nil {
		return nil, false, errors.New("NextZone is not inited for tick. Call Next() method to init next zone first.")
	}

	// Calc zone overal distance in meters from previous
	overalDistance := geo.AngleToM(c.prevZone.Position.Distance(c.nextZone.Position))

	// Calc zone distance from farrest circle point to next zone in meters
	zoneMaxCircleDistance := float64(c.prevZone.Radius-c.nextZone.Radius) + overalDistance

	// Calc zone time duration in seconds for transition to next zone
	zoneOveralTransDuration := zoneMaxCircleDistance / c.zoneSpeedInMetersPerSecond

	// Calc current zone transition percentage from overal distance to next zone
	secondsFromTickStart := now.Sub(*c.nextZoneCreationTime).Seconds()

	// If we have some time from start
	if secondsFromTickStart != 0 {
		transitionPercentage := 1 - (zoneOveralTransDuration-secondsFromTickStart)/zoneOveralTransDuration

		if transitionPercentage >= 1.0 {
			// Next zone reached

			c.currentZone = c.nextZone
			c.nextZone = nil
			c.prevZone = nil
			c.nextZoneCreationTime = nil

			// This is last tick
			return c.currentZone, true, nil

		} else if transitionPercentage == 0 {
			// Do nothing

		} else {
			// Transition in progress

			// Calc zone current distance in meters from previous
			distance := overalDistance * transitionPercentage

			// Calc zone latitude difference in meters
			latDiff := geo.AngleToM(c.prevZone.Position.Lat - c.nextZone.Position.Lat)

			// Latitude distance = Distance * Latitude diff / Overal Distance
			latDistance := distance * latDiff / overalDistance

			// Calc zone longitude difference in meters
			lngDiff := geo.AngleToM(c.prevZone.Position.Lng - c.nextZone.Position.Lng)

			// Longitude distance = Distance * Longitude diff / Overal Distance
			lngDistance := distance * lngDiff / overalDistance

			lat := c.prevZone.Position.Lat + geo.MToAngle(latDistance)
			lng := c.prevZone.Position.Lng + geo.MToAngle(lngDistance)

			// Zone radius = Next zone radius + (Prev zone radius - Next zone radius) * transition percentage
			radius := float64(c.nextZone.Radius) + float64(c.prevZone.Radius-c.nextZone.Radius)*(1-transitionPercentage)

			c.currentZone = NewZone(s2.LatLng{Lat: lat, Lng: lng}, uint32(radius), 10)
		}
	}

	return c.currentZone, false, nil
}

// Function for starting zone changing workflow
//
// Timer for new zone -> Timer for next zone delay -> Tick zone till next zone reached -> Restart next zone timer
// Workflow will work till zero zone reached
func (c *Controller) Start() error {
	if c.nextZoneTimer != nil {
		return errors.New("nextZoneTimer is already inited")
	}

	go func() {
		// While current zone radius is bigger that 0
		for c.currentZone.Radius > 0 {
			timeNowMx.Lock()
			c.ZoneEventBroadcaster.Send(StartNextZoneTimerEvent{
				CurrentZone:  c.currentZone,
				NextZoneTime: timeNow().UTC().Add(c.nextZonePeriod),
			})
			timeNowMx.Unlock()
			c.nextZoneTimer = time.NewTimer(c.nextZonePeriod)
			<-c.nextZoneTimer.C
			c.nextZoneTimer = nil
			c.log.Debug("Next zone timer fired")

			// Create next zone
			timeNowMx.Lock()
			c.Next(timeNow().UTC())
			timeNowMx.Unlock()

			// Send next zone event to players
			timeNowMx.Lock()
			c.ZoneEventBroadcaster.Send(StartZoneDelayTimerEvent{
				CurrentZone:       c.currentZone,
				NextZone:          c.nextZone,
				ZoneTickStartTime: timeNow().UTC().Add(c.nextZoneDelay),
			})
			timeNowMx.Unlock()
			c.nextZoneDelayTimer = time.NewTimer(c.nextZoneDelay)
			<-c.nextZoneDelayTimer.C
			c.nextZoneDelayTimer = nil
			c.log.Debug("Next zone delay timer fired")

			c.zoneTicker = time.NewTicker(zoneTickPeriod)

		tickerLoop:
			for range c.zoneTicker.C {
				timeNowMx.Lock()
				curZone, lastTick, err := c.Tick(timeNow().UTC())
				timeNowMx.Unlock()
				if err != nil {
					c.log.Error("Couldn't Tick next zone. " + err.Error())
				}

				// Check that it was last tick
				if lastTick {
					c.log.Debug("Last tick. Stop zone ticker")
					c.zoneTicker.Stop()
					c.zoneTicker = nil
					// Go away from ticker loop
					break tickerLoop
				} else {
					c.log.WithFields(logrus.Fields{"curZone": curZone, "lastTick": lastTick}).Debug("Next zone tick")
					c.ZoneEventBroadcaster.Send(ZoneTickEvent{
						CurrentZone: curZone,
						NextZone:    c.nextZone,
						LastTick:    lastTick,
					})
				}
			}
		}
		c.log.Debug("Next zone loop end")
	}()

	return nil
}

// Creates new zone inside current zone
func nextZone(zone *Zone, minZoneRadiusInM uint32) *Zone {
	newR := newRadius(zone.Radius, minZoneRadiusInM)

	// Calc radius of random area
	r := randomR(zone.Radius, newR)
	theta := randFloat() * 2 * math.Pi

	lat := zone.Position.Lat + geo.MToAngle(r*math.Cos(theta))
	lng := zone.Position.Lng + geo.MToAngle(r*math.Sin(theta))

	return NewZone(s2.LatLng{Lat: lat, Lng: lng}, newR, 10)
}

// Creeate random radius as circle position for new zone
func randomR(curZoneR uint32, newZoneR uint32) float64 {
	return float64((curZoneR - newZoneR)) * math.Sqrt(randFloat())
}

func newRadius(radius uint32, minZoneRadiusInM uint32) uint32 {
	if radius > minZoneRadiusInM*2 {
		return radius / 2
	} else {
		return 0
	}
}
