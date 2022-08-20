package public

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/petrolax/chat/pkg/api/public"
	"github.com/petrolax/chat/pkg/transport/helpers"
)

func MakeGetPlantsEndpoint(c *Controller) endpoint.Endpoint {
	req := func(ctx context.Context, request interface{}) (interface{}, error) {
		return c.GetPlants(ctx, request.(*public.GetPlantsRequest))
	}
	return req
}

func (c *Controller) GetPlants(ctx context.Context, req *public.GetPlantsRequest) (*public.GetPlantsResponse, error) {
	c.lg.Info("GetPlantsHandler")

	plants, err := c.service.GetPlants(ctx)
	if err != nil {
		return nil, err
	}

	return &public.GetPlantsResponse{Plants: helpers.ConvertSlicesDTOToPublicPlant(plants)}, nil
}

func (c *Controller) decodeGetPlants(_ context.Context, r *http.Request) (interface{}, error) {
	r.Body.Close()
	return &public.GetPlantsRequest{}, nil
}
