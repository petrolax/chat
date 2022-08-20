package public

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	kithttp "github.com/go-kit/kit/transport/http"
	kitlog "github.com/go-kit/log"
	"github.com/gorilla/mux"
	"github.com/petrolax/chat/pkg/api/public"
	"github.com/petrolax/chat/pkg/plants"
	"go.uber.org/zap"
)

type Controller struct {
	lg      *zap.Logger
	service plants.Service
	public.UnimplementedPlantsApiServer
}

func NewController(lg *zap.Logger, service plants.Service) *Controller {
	return &Controller{
		lg:      lg,
		service: service,
	}
}

func (c *Controller) Endpoints() http.Handler {
	r := mux.NewRouter()
	options := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(kitlog.NewJSONLogger(os.Stderr)),
	}

	// TODO: path name
	r.Methods("POST").Path("/addPlant").Handler(kithttp.NewServer(
		MakeAddPlantEndpoint(c), // endpoint
		c.decodeAddPlant,        // decode request
		Encode,                  // encode response
		options...,
	))

	r.Methods("POST").Path("/updatePlant").Handler(kithttp.NewServer(
		MakeUpdatePlantEndpoint(c),
		c.decodeUpdatePlant,
		Encode,
		options...,
	))

	r.Methods("POST").Path("/getPlant").Handler(kithttp.NewServer(
		MakeGetPlantEndpoint(c),
		c.decodeGetPlant,
		Encode,
		options...,
	))

	r.Methods("POST").Path("/getPlants").Handler(kithttp.NewServer(
		MakeGetPlantsEndpoint(c),
		c.decodeGetPlants,
		Encode,
		options...,
	))

	r.Methods("POST").Path("/deletePlant").Handler(kithttp.NewServer(
		MakeDeletePlantEndpoint(c),
		c.decodeDeletePlant,
		Encode,
		options...,
	))

	return r
}

func Encode(_ context.Context, w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(resp)
}
