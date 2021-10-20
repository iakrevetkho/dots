package api_spot_v1_test

import (
	// External
	"context"
	"sync"
	"testing"
	"time"

	"github.com/golang/geo/s2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	// Internal
	api_spot_v1 "github.com/iakrevetkho/dots/server/pkg/api/spot/v1"
	proto "github.com/iakrevetkho/dots/server/proto/gen/spot/v1"
)

type MockGetPlayerPositionServer struct {
	SpotUuid             uuid.UUID
	PlayerUuid           uuid.UUID
	MsgCount             uint32
	LastPlayersPositions *proto.GetPlayersStatesResponse
	sync.Mutex
	grpc.ServerStream
}

func (s *MockGetPlayerPositionServer) Send(response *proto.GetPlayersStatesResponse) error {
	print(s.MsgCount)
	s.Lock()
	s.MsgCount -= 1
	s.Unlock()
	if s.MsgCount > 0 {
		return nil
	} else {
		s.LastPlayersPositions = response
		return grpc.ErrServerStopped
	}
}

func TestGetPlayerPosition(t *testing.T) {
	s := api_spot_v1.New(10 * time.Millisecond)

	// Create spot first
	createSpotRet, err := s.CreateSpot(context.Background(), &proto.CreateSpotRequest{
		Position: &proto.Position{
			Latitude:  10,
			Longitude: 20,
		},
		Radius:              100,
		ScanPeriodInSeconds: 10,
		ZonePeriodInSeconds: 30,
	})
	assert.NoError(t, err)

	spotUuid := uuid.MustParse(createSpotRet.SpotUuid)
	playerUuid := uuid.New()
	player2Uuid := uuid.New()
	player3Uuid := uuid.New()

	// Add positions
	spot := s.SpotsMap[spotUuid]
	spot.PlayersStateMap = map[uuid.UUID]api_spot_v1.PlayerState{
		playerUuid: api_spot_v1.PlayerState{
			Position: s2.LatLngFromDegrees(10, 20),
			Health:   88,
		},
		player2Uuid: api_spot_v1.PlayerState{
			Position: s2.LatLngFromDegrees(60, 70),
			Health:   33,
		},
		player3Uuid: api_spot_v1.PlayerState{
			Position: s2.LatLngFromDegrees(80, 90),
			Health:   15,
		},
	}
	s.SpotsMap[spotUuid] = spot

	// Create stream for getting position
	mockServer := MockGetPlayerPositionServer{
		SpotUuid:   spotUuid,
		PlayerUuid: playerUuid,
		MsgCount:   4,
	}
	request := &proto.GetPlayersStatesRequest{
		SpotUuid:   spotUuid.String(),
		PlayerUuid: playerUuid.String(),
	}

	exit := make(chan bool)
	go func() {
		err = s.GetPlayersStates(request, &mockServer)
		assert.Error(t, err)
		exit <- true
	}()

	// Wait channel ready
	for {
		s.SpotsMapMx.Lock()
		spot := s.SpotsMap[spotUuid]
		playerState := spot.PlayersStateMap[playerUuid]
		s.SpotsMapMx.Unlock()
		if playerState.Sub != nil {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}

	// Send data
	spot = s.SpotsMap[spotUuid]
	playerState := spot.PlayersStateMap[playerUuid]
	sub := *playerState.Sub

	sub <- api_spot_v1.PlayerPublicState{
		PlayerUuid: playerUuid,
		Position:   s2.LatLngFromDegrees(10, 20),
		Health:     88,
	}
	sub <- api_spot_v1.PlayerPublicState{
		PlayerUuid: player2Uuid,
		Position:   s2.LatLngFromDegrees(60, 70),
		Health:     33,
	}
	sub <- api_spot_v1.PlayerPublicState{
		PlayerUuid: player3Uuid,
		Position:   s2.LatLngFromDegrees(80, 90),
		Health:     15,
	}
	sub <- api_spot_v1.PlayerPublicState{
		PlayerUuid: playerUuid,
		Position:   s2.LatLngFromDegrees(10, 20),
		Health:     88,
	}

	<-exit

	assert.Equal(t, uint32(0), mockServer.MsgCount)
	assert.NotNil(t, mockServer.LastPlayersPositions)
	assert.Equal(t, float64(10), mockServer.LastPlayersPositions.PlayerState.Position.Latitude)
	assert.Equal(t, float64(20), mockServer.LastPlayersPositions.PlayerState.Position.Longitude)
	assert.Equal(t, int32(88), mockServer.LastPlayersPositions.PlayerState.Health)
}
