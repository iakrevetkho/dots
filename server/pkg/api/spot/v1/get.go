package api_spot_v1

import (
	// External
	"context"
	"fmt"

	"github.com/google/uuid"

	// Internal
	proto "github.com/iakrevetkho/dots/server/proto/gen/spot/v1"
)

func (s *SpotServiceServer) GetSpot(ctx context.Context, request *proto.GetSpotRequest) (*proto.GetSpotResponse, error) {
	s.log.WithField("request", request.String()).Debug("Get spot request")

	spotUuid, err := uuid.Parse(request.SpotUuid)
	if err != nil {
		return nil, fmt.Errorf("Couldn't parse spot uuid. " + err.Error())
	}

	spot, ok := s.SpotsMap.Load(spotUuid)
	if !ok {
		return nil, fmt.Errorf("Spot with uuid '%s' couldn't be found", spotUuid)
	}

	response := proto.GetSpotResponse{
		Position: &proto.Position{
			Latitude:  spot.Position.Lat.Degrees(),
			Longitude: spot.Position.Lng.Degrees(),
		},
		RadiusInM:           spot.RadiusInM,
		ScanPeriodInSeconds: int32(spot.ScanPeriod.Seconds()),
		ZonePeriodInSeconds: int32(spot.ZonePeriod.Seconds()),
	}
	s.log.WithField("response", response.String()).Debug("Get spot response")

	return &response, nil
}
