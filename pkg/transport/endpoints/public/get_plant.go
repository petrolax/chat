package public

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/petrolax/chat/pkg/api/public"
	"github.com/petrolax/chat/pkg/transport/helpers"
	"go.uber.org/zap"
)

func MakeGetPlantEndpoint(c *Controller) endpoint.Endpoint {
	req := func(ctx context.Context, request interface{}) (interface{}, error) {
		return c.GetPlant(ctx, request.(*public.GetPlantRequest))
	}
	return req
}

func (c *Controller) GetPlant(ctx context.Context, req *public.GetPlantRequest) (*public.GetPlantResponse, error) {
	c.lg.Info("GetPlantHandler", zap.String("name", req.Name))

	if req.Name == "" {
		return nil, errors.New("validation error")
	}

	plant, err := c.service.GetPlantByName(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return &public.GetPlantResponse{Plant: helpers.ConvertDTOToPulbicPlant(plant)}, nil
}

func (c *Controller) decodeGetPlant(_ context.Context, r *http.Request) (interface{}, error) {
	var req public.GetPlantRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		c.lg.Error("Can't decode", zap.Error(err))
		return nil, err
	}
	r.Body.Close()
	return &req, nil
}
