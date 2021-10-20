package api_spot_v1

import (
	// External

	"fmt"

	"github.com/google/uuid"

	// Internal
	proto "github.com/iakrevetkho/dots/server/proto/gen/spot/v1"
)

func (s *SpotServiceServer) GetPlayersStates(request *proto.GetPlayersStatesRequest, stream proto.SpotService_GetPlayersStatesServer) error {
	s.log.WithField("request", request.String()).Trace("Get players positions")

	spotUuid, err := uuid.Parse(request.SpotUuid)
	if err != nil {
		return fmt.Errorf("Couldn't parse spot uuid. " + err.Error())
	}

	s.SpotsMapMx.Lock()
	spot, ok := s.SpotsMap[spotUuid]
	s.SpotsMapMx.Unlock()
	if !ok {
		return fmt.Errorf("Spot with uuid '%s' couldn't be found", spotUuid)
	}

	playerUuid, err := uuid.Parse(request.PlayerUuid)
	if err != nil {
		return fmt.Errorf("Couldn't parse user uuid. " + err.Error())
	}

	spot.PlayersStateMapMx.Lock()
	playerState := spot.PlayersStateMap[playerUuid]
	spot.PlayersStateMapMx.Unlock()

	// Check that player hadn't subscription
	if playerState.Sub != nil {
		return fmt.Errorf("User %v already has subscription", playerUuid)
	}

	playerSub := make(chan PlayerPublicState)
	playerState.Sub = &playerSub
	// Update player state
	spot.PlayersStateMapMx.Lock()
	spot.PlayersStateMap[playerUuid] = playerState
	spot.PlayersStateMapMx.Unlock()

	for playerState := range playerSub {
		response := &proto.GetPlayersStatesResponse{
			PlayerState: &proto.PlayerState{
				Position: &proto.Position{
					Latitude:  playerState.Position.Lat.Degrees(),
					Longitude: playerState.Position.Lng.Degrees(),
				},
				Health: int32(playerState.Health),
			},
		}

		s.log.WithField("response", response.String()).Debug("Get players state response")
		if err := stream.Send(response); err != nil {
			// Remove channel from current state
			close(playerSub)
			spot.PlayersStateMapMx.Lock()
			playerState := spot.PlayersStateMap[playerUuid]
			playerState.Sub = nil
			// Update player state
			spot.PlayersStateMap[playerUuid] = playerState
			spot.PlayersStateMapMx.Unlock()
			return err
		}
	}
	return nil
}
