package private

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/petrolax/chat/pkg/api/private"
	"github.com/petrolax/chat/pkg/plants"
	"go.uber.org/zap"
)

type Controller struct {
	lg      *zap.Logger
	service plants.Service
	private.UnimplementedPlantsApiServer
}

func NewController(lg *zap.Logger, service plants.Service) *Controller {
	return &Controller{
		lg:      lg,
		service: service,
	}
}

func (c *Controller) Endpoints() http.Handler {
	r := mux.NewRouter()

	return r
}
