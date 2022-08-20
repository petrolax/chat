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

func MakeUpdatePlantEndpoint(c *Controller) endpoint.Endpoint {
	req := func(ctx context.Context, request interface{}) (interface{}, error) {
		return c.UpdatePlant(ctx, request.(*public.UpdatePlantRequest))
	}
	return req
}

func (c *Controller) UpdatePlant(ctx context.Context, req *public.UpdatePlantRequest) (*public.UpdatePlantResponse, error) {
	c.lg.Info("UpdatePlantHandler", zap.String("old_name", req.OldName), zap.String("new_name", req.NewName))
	if req.OldName == "" {
		return nil, errors.New("validation error")
	}
	if req.NewName == "" {
		return nil, errors.New("validation error")
	}

	plant, err := c.service.UpdatePlant(ctx, req.OldName, req.NewName)
	if err != nil {
		return nil, err
	}

	return &public.UpdatePlantResponse{Plant: helpers.ConvertDTOToPulbicPlant(plant)}, nil
}

func (c *Controller) decodeUpdatePlant(_ context.Context, r *http.Request) (interface{}, error) {
	var req public.UpdatePlantRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		c.lg.Error("Can't decode", zap.Error(err))
		return nil, err
	}
	r.Body.Close()
	return &req, nil
}
