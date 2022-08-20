package public

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/petrolax/chat/pkg/api/public"
	"go.uber.org/zap"
)

func MakeDeletePlantEndpoint(c *Controller) endpoint.Endpoint {
	req := func(ctx context.Context, request interface{}) (interface{}, error) {
		return c.DeletePlant(ctx, request.(*public.DeletePlantRequest))
	}
	return req
}

func (c *Controller) DeletePlant(ctx context.Context, req *public.DeletePlantRequest) (*public.DeletePlantResponse, error) {
	c.lg.Info("DeletePlantHandler", zap.String("name", req.Name))
	if req.Name == "" {
		return nil, errors.New("validation error")
	}

	err := c.service.DeletePlant(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return &public.DeletePlantResponse{}, nil
}

func (c *Controller) decodeDeletePlant(_ context.Context, r *http.Request) (interface{}, error) {
	var req public.DeletePlantRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		c.lg.Error("Can't decode", zap.Error(err))
		return nil, err
	}
	r.Body.Close()
	return &req, nil
}
