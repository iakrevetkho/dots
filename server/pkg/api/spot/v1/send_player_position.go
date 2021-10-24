package api_spot_v1

import (
	// External
	"fmt"
	"io"
	"time"

	"github.com/golang/geo/s2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	// Internal
	geo_utils "github.com/iakrevetkho/dots/server/pkg/utils/geo"
	proto "github.com/iakrevetkho/dots/server/proto/gen/spot/v1"
)

func (s *SpotServiceServer) SendPlayerPosition(stream proto.SpotService_SendPlayerPositionServer) error {
	s.log.Trace("Open send user position stream")

	for {
		request, err := stream.Recv()
		if err == io.EOF {
			// End of stream
			s.log.Trace("Close user position stream")
			return stream.SendAndClose(&proto.SendPlayerPositionResponse{})
		}
		if err != nil {
			return err
		}
		s.log.WithField("request", request.String()).Trace("Open send user position stream")

		spotUuid, err := uuid.Parse(request.SpotUuid)
		if err != nil {
			return fmt.Errorf("Couldn't parse spot uuid. " + err.Error())
		}

		v, ok := s.SpotsMap.Load(spotUuid)
		if !ok {
			return fmt.Errorf("Spot with uuid '%s' couldn't be found", spotUuid)
		}
		spot := v.(Spot)

		playerUuid, err := uuid.Parse(request.PlayerUuid)
		if err != nil {
			return fmt.Errorf("Couldn't parse user uuid. " + err.Error())
		}

		// Update player state
		v, ok = spot.PlayersStateMap.Load(playerUuid)
		if !ok {
			return fmt.Errorf("Player with uuid '%s' couldn't be found in spot '%s'", playerUuid, spotUuid)
		} else {
			playerState := v.(PlayerState)
			playerState.Position = s2.LatLngFromDegrees(request.Position.Latitude, request.Position.Longitude)

			// Check player health
			playerToSpotDistance := geo_utils.AngleToM(spot.Position.Distance(playerState.Position))
			if playerToSpotDistance > float64(spot.ZoneRadius) {
				s.log.Debugf("Player distance %f > %d zone radius", playerToSpotDistance, spot.ZoneRadius)
				// Start goroutine with ticket for health decreasing
				if !playerState.ZoneDamageActice {
					s.startPlayerZoneDamage(spotUuid, playerUuid)
				}
			} else {
				// Stop player zone damage if needed
				if playerState.ZoneDamageActice {
					playerState.StopZoneDmgCh <- true
				}
			}

			// Update player state
			spot.PlayersStateMap.Store(playerUuid, playerState)

			// Send player state to subscriptions which requires it
			spot.PlayersStateMap.Range(func(k, v interface{}) bool {
				platerState := v.(PlayerState)

				// TODO Add checks for distance, scanning and etc
				// Check that we have subscription
				if platerState.Sub != nil {
					// Send data to subscription channel
					(*platerState.Sub) <- PlayerPublicState{
						PlayerUuid: playerUuid,
						Position:   playerState.Position,
						Health:     playerState.Health,
					}
				}
				return true
			})
		}
		s.SpotsMap.Store(spotUuid, spot)

		s.log.WithFields(logrus.Fields{
			"spotUuid":   spotUuid,
			"playeruuid": playerUuid,
			"position":   request.Position,
		}).Trace("Player position updated")
	}
}

func (s *SpotServiceServer) startPlayerZoneDamage(spotUuid uuid.UUID, playerUuid uuid.UUID) {
	go func() {
		// TODO Move zone damage to consts
		const zoneDamage = 15
		// TODO Move zone damage period to consts
		const zoneDamagePeriod = 1 * time.Second

		v, ok := s.SpotsMap.Load(spotUuid)
		if !ok {
			s.log.Errorf("Spot with uuid '%s' couldn't be found", spotUuid)
			return
		}
		spot := v.(Spot)

		v, ok = spot.PlayersStateMap.Load(playerUuid)
		if !ok {
			s.log.Errorf("Player with uuid '%s' couldn't be found in spot '%s'", playerUuid, spotUuid)
			return
		}
		playerState := v.(PlayerState)

		playerState.ZoneDamageActice = true
		// Update state in spot
		spot.PlayersStateMap.Store(playerUuid, playerState)
		s.SpotsMap.Store(spotUuid, spot)

		damageTicker := time.NewTicker(zoneDamagePeriod)
		for {
			v, ok = s.SpotsMap.Load(spotUuid)
			if !ok {
				s.log.Errorf("Spot with uuid '%s' couldn't be found", spotUuid)
				return
			}
			spot := v.(Spot)

			v, ok = spot.PlayersStateMap.Load(playerUuid)
			if !ok {
				s.log.Errorf("Player with uuid '%s' couldn't be found in spot '%s'", playerUuid, spotUuid)
				return
			}
			playerState := v.(PlayerState)

			select {
			case <-playerState.StopZoneDmgCh:
				playerState.ZoneDamageActice = false
				// Update state in spot
				spot.PlayersStateMap.Store(playerUuid, playerState)
				s.SpotsMap.Store(spotUuid, spot)
				return
			case <-damageTicker.C:
				if playerState.Health <= zoneDamage {
					playerState.Health = 0
					// Update state in spot
					spot.PlayersStateMap.Store(playerUuid, playerState)
					s.SpotsMap.Store(spotUuid, spot)
					// Stop zone damage
					playerState.StopZoneDmgCh <- true
				} else {
					playerState.Health -= zoneDamage
					// Update state in spot
					spot.PlayersStateMap.Store(playerUuid, playerState)
					s.SpotsMap.Store(spotUuid, spot)
				}
			}
		}
	}()
}
