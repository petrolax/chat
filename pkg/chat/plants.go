package plants

import (
	"context"

	"github.com/petrolax/chat/pkg/plants/dto"
)

type Service interface {
	AddPlant(ctx context.Context, name string) (*dto.Plant, error)
	UpdatePlant(ctx context.Context, oldName, newName string) (*dto.Plant, error)
	GetPlantByName(ctx context.Context, name string) (*dto.Plant, error)
	GetPlants(ctx context.Context) ([]*dto.Plant, error)
	DeletePlant(ctx context.Context, name string) error
}
