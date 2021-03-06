package zone

import (
	// External

	"testing"
	"time"

	"github.com/golang/geo/s2"
	"github.com/iakrevetkho/dots/server/pkg/utils/geo"
	"github.com/iakrevetkho/dots/server/pkg/utils/mock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	// Internal
)

func TestNewRadius(t *testing.T) {
	assert.Equal(t, float32(100), newRadius(200, 10))
	assert.Equal(t, float32(50), newRadius(100, 10))
	assert.Equal(t, float32(25), newRadius(50, 10))
	assert.Equal(t, float32(12.5), newRadius(25, 10))
	assert.Equal(t, float32(0), newRadius(10, 10))
	assert.Equal(t, float32(0), newRadius(0, 10))
}

func TestRandomR(t *testing.T) {
	// Test min random
	mock.RandFloat = func() float64 { return 0 }
	r := randomR(200, 100)
	assert.Equal(t, float64(0), r)

	// Test max random
	mock.RandFloat = func() float64 { return 1 }
	r = randomR(200, 100)
	assert.Equal(t, float64(100), r)
}

func TestNextZone(t *testing.T) {
	// Test max random
	mock.RandFloat = func() float64 { return 1 }

	c := NewController(s2.LatLng{Lat: 0, Lng: 0}, 200, 10, time.Second*30, time.Second*10, 10.0)

	assert.Equal(t, s2.LatLng{Lat: 0, Lng: 0}, c.currentZone.Position)
	assert.Equal(t, float32(200), c.currentZone.Radius)
	assert.Equal(t, float32(0.25), c.currentZone.Damage)
	assert.Nil(t, c.nextZone)

	c.Next(time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC))

	assert.Equal(t, s2.LatLng{Lat: 0, Lng: 0}, c.currentZone.Position)
	assert.Equal(t, float32(200), c.currentZone.Radius)
	assert.Equal(t, float32(0.25), c.currentZone.Damage)

	assert.Equal(t, s2.LatLng{Lat: 1.5696098420815538e-05, Lng: -3.844435338030709e-21}, c.nextZone.Position)
	assert.Equal(t, float32(100), c.nextZone.Radius)
	assert.Equal(t, float32(0.5), c.nextZone.Damage)

	assert.Equal(t, uint16(99), uint16(geo.AngleToM(c.nextZone.Position.Distance(c.currentZone.Position))))
}

func TestTickZone(t *testing.T) {
	// Test max random
	mock.RandFloat = func() float64 { return 0.5 }

	c := NewController(s2.LatLng{Lat: 0, Lng: 0}, 200, 10, time.Second*30, time.Second*10, 10.0)

	assert.Equal(t, s2.LatLng{Lat: 0, Lng: 0}, c.currentZone.Position)
	assert.Equal(t, float32(200), c.currentZone.Radius)
	assert.Equal(t, float32(0.25), c.currentZone.Damage)
	assert.Nil(t, c.nextZone)

	c.Next(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC))

	assert.Equal(t, s2.LatLng{Lat: 0, Lng: 0}, c.currentZone.Position)
	assert.Equal(t, float32(200), c.currentZone.Radius)
	assert.Equal(t, float32(0.25), c.currentZone.Damage)

	assert.Equal(t, s2.LatLng{Lat: -1.1098817631530128e-05, Lng: 1.359213148677356e-21}, c.nextZone.Position)
	assert.Equal(t, float32(100), c.nextZone.Radius)
	assert.Equal(t, float32(0.5), c.nextZone.Damage)

	assert.Equal(t, uint16(70), uint16(geo.AngleToM(c.nextZone.Position.Distance(c.currentZone.Position))))

	// Zero zone tick
	curZone, lastTick, err := c.Tick(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC))
	assert.NoError(t, err)
	assert.False(t, lastTick)
	assert.Equal(t, float32(0), float32(geo.AngleToM(c.prevZone.Position.Distance(curZone.Position))+float64(c.prevZone.Radius-curZone.Radius)))
	assert.Equal(t, float32(70.71068), float32(geo.AngleToM(c.nextZone.Position.Distance(curZone.Position))))
	assert.Equal(t, float32(200), curZone.Radius)
	assert.Equal(t, float32(0.25), curZone.Damage)

	// 1 second zone tick
	curZone, lastTick, err = c.Tick(time.Date(2000, 1, 1, 0, 0, 1, 0, time.UTC))
	assert.NoError(t, err)
	assert.False(t, lastTick)
	assert.Equal(t, float32(2.777775235650089), float32(geo.AngleToM(c.prevZone.Position.Distance(curZone.Position))+float64(c.prevZone.Radius-curZone.Radius)))
	assert.Equal(t, float32(69.56008), float32(geo.AngleToM(c.nextZone.Position.Distance(curZone.Position))))
	assert.Equal(t, float32(198.37282), curZone.Radius)
	assert.Equal(t, float32(0.25205067), curZone.Damage)

	// 10 seconds zone tick
	curZone, lastTick, err = c.Tick(time.Date(2000, 1, 1, 0, 0, 10, 0, time.UTC))
	assert.NoError(t, err)
	assert.False(t, lastTick)
	assert.Equal(t, float32(27.777782), float32(geo.AngleToM(c.prevZone.Position.Distance(curZone.Position))+float64(c.prevZone.Radius-curZone.Radius)))
	assert.Equal(t, float32(59.204746), float32(geo.AngleToM(c.nextZone.Position.Distance(curZone.Position))))
	assert.Equal(t, float32(183.72815), curZone.Radius)
	assert.Equal(t, float32(0.27214122), curZone.Damage)

	// 30 seconds zone tick
	curZone, lastTick, err = c.Tick(time.Date(2000, 1, 1, 0, 0, 30, 0, time.UTC))
	assert.NoError(t, err)
	assert.False(t, lastTick)
	assert.Equal(t, float32(83.333336), float32(geo.AngleToM(c.prevZone.Position.Distance(curZone.Position))+float64(c.prevZone.Radius-curZone.Radius)))
	assert.Equal(t, float32(36.192883), float32(geo.AngleToM(c.nextZone.Position.Distance(curZone.Position))))
	assert.Equal(t, float32(151.18446), curZone.Radius)
	assert.Equal(t, float32(0.33072183), curZone.Damage)

	// 1 min zone tick
	curZone, lastTick, err = c.Tick(time.Date(2000, 1, 1, 0, 1, 0, 0, time.UTC))
	assert.NoError(t, err)
	assert.False(t, lastTick)
	assert.Equal(t, float32(166.66667), float32(geo.AngleToM(c.prevZone.Position.Distance(curZone.Position))+float64(c.prevZone.Radius-curZone.Radius)))
	assert.Equal(t, float32(1.6750844), float32(geo.AngleToM(c.nextZone.Position.Distance(curZone.Position))))
	assert.Equal(t, float32(102.36893), curZone.Radius)
	assert.Equal(t, float32(0.48842946), curZone.Damage)

	// 1 minute 10 seconds zone tick
	curZone, lastTick, err = c.Tick(time.Date(2000, 1, 1, 0, 1, 10, 0, time.UTC))
	assert.NoError(t, err)
	assert.True(t, lastTick)
	assert.Nil(t, c.prevZone)
	assert.Nil(t, c.nextZone)
	assert.Equal(t, s2.LatLng{Lat: -1.1098817631530128e-05, Lng: 1.359213148677356e-21}, curZone.Position)
	assert.Equal(t, float32(100), curZone.Radius)
	assert.Equal(t, float32(0.5), curZone.Damage)

	// 2 minutes zone tick
	_, _, err = c.Tick(time.Date(2000, 1, 1, 0, 2, 0, 0, time.UTC))
	assert.Error(t, err)
}

func TestStartTickZone(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)

	// Test max random
	mock.RandFloat = func() float64 { return 0.5 }
	mock.TimeNow = func() time.Time { return time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC) }

	c := NewController(s2.LatLng{Lat: 0, Lng: 0}, 200, 10, time.Nanosecond*1, time.Nanosecond*1, 10.0)

	assert.Equal(t, s2.LatLng{Lat: 0, Lng: 0}, c.currentZone.Position)
	assert.Equal(t, float32(200), c.currentZone.Radius)
	assert.Equal(t, float32(0.25), c.currentZone.Damage)
	assert.Nil(t, c.nextZone)

	err := c.Start()
	assert.NoError(t, err)

	zoneEventCh := c.ZoneEventBroadcaster.Listen().Ch

	event := <-zoneEventCh
	assert.IsType(t, StartNextZoneTimerEvent{}, event)
	assert.Equal(t, s2.LatLng{Lat: 0, Lng: 0}, event.(StartNextZoneTimerEvent).CurrentZone.Position)
	assert.Equal(t, float32(200), event.(StartNextZoneTimerEvent).CurrentZone.Radius)
	assert.Equal(t, time.Date(2000, 1, 1, 0, 0, 0, 1, time.UTC), event.(StartNextZoneTimerEvent).NextZoneTime)

	event = <-zoneEventCh
	assert.IsType(t, StartZoneDelayTimerEvent{}, event)
	curZone := event.(StartZoneDelayTimerEvent).CurrentZone
	nextZone := event.(StartZoneDelayTimerEvent).NextZone
	assert.Equal(t, s2.LatLng{Lat: 0, Lng: 0}, curZone.Position)
	assert.Equal(t, float32(200), curZone.Radius)
	assert.Equal(t, float32(70.71068), float32(geo.AngleToM(nextZone.Position.Distance(curZone.Position))))
	assert.Equal(t, float32(100), nextZone.Radius)
	assert.Equal(t, time.Date(2000, 1, 1, 0, 0, 0, 1, time.UTC), event.(StartZoneDelayTimerEvent).ZoneTickStartTime)

	// Tick after 1 second
	mock.TimeNowMx.Lock()
	mock.TimeNow = func() time.Time { return time.Date(2000, 1, 1, 0, 0, 1, 0, time.UTC) }
	mock.TimeNowMx.Unlock()
	event = <-zoneEventCh
	assert.IsType(t, ZoneTickEvent{}, event)
	curZone = event.(ZoneTickEvent).CurrentZone
	nextZone = event.(ZoneTickEvent).NextZone
	assert.False(t, event.(ZoneTickEvent).LastTick)
	assert.Equal(t, float32(198.37282), curZone.Radius)
	assert.Equal(t, float32(69.56008), float32(geo.AngleToM(nextZone.Position.Distance(curZone.Position))))

	// Tick after 10 seconds
	mock.TimeNowMx.Lock()
	mock.TimeNow = func() time.Time { return time.Date(2000, 1, 1, 0, 0, 10, 0, time.UTC) }
	mock.TimeNowMx.Unlock()
	event = <-zoneEventCh
	assert.IsType(t, ZoneTickEvent{}, event)
	curZone = event.(ZoneTickEvent).CurrentZone
	nextZone = event.(ZoneTickEvent).NextZone
	assert.False(t, event.(ZoneTickEvent).LastTick)
	assert.Equal(t, float32(183.72815), curZone.Radius)
	assert.Equal(t, float32(59.204746), float32(geo.AngleToM(nextZone.Position.Distance(curZone.Position))))

	// Tick after 30 seconds
	mock.TimeNowMx.Lock()
	mock.TimeNow = func() time.Time { return time.Date(2000, 1, 1, 0, 0, 30, 0, time.UTC) }
	mock.TimeNowMx.Unlock()
	event = <-zoneEventCh
	assert.IsType(t, ZoneTickEvent{}, event)
	curZone = event.(ZoneTickEvent).CurrentZone
	nextZone = event.(ZoneTickEvent).NextZone
	assert.False(t, event.(ZoneTickEvent).LastTick)
	assert.Equal(t, float32(151.18446), curZone.Radius)
	assert.Equal(t, float32(36.192883), float32(geo.AngleToM(nextZone.Position.Distance(curZone.Position))))

	// Tick after 1 min
	mock.TimeNowMx.Lock()
	mock.TimeNow = func() time.Time { return time.Date(2000, 1, 1, 0, 1, 0, 0, time.UTC) }
	mock.TimeNowMx.Unlock()
	event = <-zoneEventCh
	assert.IsType(t, ZoneTickEvent{}, event)
	curZone = event.(ZoneTickEvent).CurrentZone
	nextZone = event.(ZoneTickEvent).NextZone
	assert.False(t, event.(ZoneTickEvent).LastTick)
	assert.Equal(t, float32(102.36893), curZone.Radius)
	assert.Equal(t, float32(1.6750844), float32(geo.AngleToM(nextZone.Position.Distance(curZone.Position))))

	// Tick after 1 min 10 seconds
	mock.TimeNowMx.Lock()
	mock.TimeNow = func() time.Time { return time.Date(2000, 1, 1, 0, 1, 30, 0, time.UTC) }
	mock.TimeNowMx.Unlock()
	event = <-zoneEventCh
	assert.IsType(t, StartNextZoneTimerEvent{}, event)
	assert.Equal(t, s2.LatLng{Lat: -1.1098817631530128e-05, Lng: 1.359213148677356e-21}, event.(StartNextZoneTimerEvent).CurrentZone.Position)
	assert.Equal(t, float32(0.5), event.(StartNextZoneTimerEvent).CurrentZone.Damage)
	assert.Equal(t, float32(100), event.(StartNextZoneTimerEvent).CurrentZone.Radius)
}
