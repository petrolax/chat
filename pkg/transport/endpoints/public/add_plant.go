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

func MakeAddPlantEndpoint(c *Controller) endpoint.Endpoint {
	req := func(ctx context.Context, request interface{}) (interface{}, error) {
		return c.AddPlant(ctx, request.(*public.AddPlantRequest))
	}
	return req
}

func (c *Controller) AddPlant(ctx context.Context, req *public.AddPlantRequest) (*public.AddPlantResponse, error) {
	c.lg.Info("AddPlantHandler", zap.String("name", req.Name))

	if req.Name == "" {
		return nil, errors.New("validation error")
	}

	plant, err := c.service.AddPlant(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return &public.AddPlantResponse{Plant: helpers.ConvertDTOToPulbicPlant(plant)}, nil
}

func (c *Controller) decodeAddPlant(_ context.Context, r *http.Request) (interface{}, error) {
	var req public.AddPlantRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		c.lg.Error("Can't decode", zap.Error(err))
		return nil, err
	}
	r.Body.Close()
	return &req, nil
}
