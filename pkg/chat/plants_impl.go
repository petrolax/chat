package plants

import (
	"context"
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/petrolax/chat/pkg/helpers"
	"github.com/petrolax/chat/pkg/plants/dao"
	"github.com/petrolax/chat/pkg/plants/dto"
	"github.com/petrolax/chat/pkg/repository"
	"go.uber.org/zap"
)

type service struct {
	lg  *zap.Logger
	rep repository.Repository
}

func NewService(lg *zap.Logger, rep repository.Repository) *service {
	return &service{
		lg:  lg,
		rep: rep,
	}
}

func (s *service) AddPlant(ctx context.Context, name string) (*dto.Plant, error) {
	if name == "" {
		s.lg.Error("New plant name is empty", zap.String("plant_name", name))
		return nil, errors.New("validation error")
	}

	s.lg.Info("Add new plant", zap.String("plant_name", name))

	plant, err := s.rep.GetPlantByName(ctx, name)
	if err != nil && err != gorm.ErrRecordNotFound {
		s.lg.Error("Can't check plant name", zap.Error(err))
		return nil, errors.New("internal error")
	}
	if plant != nil && plant.Name != "" {
		s.lg.Error("Plant already exist", zap.String("plant_name", name))
		return nil, errors.New("already exist")
	}

	plant = new(dao.Plant)
	plant.Name = name
	if err = s.rep.AddPlant(ctx, plant); err != nil {
		s.lg.Error("Can't add new plant", zap.Error(err), zap.String("plant_name", name))
		return nil, errors.New("internal error")
	}
	return &dto.Plant{
		ID:   plant.ID,
		Name: plant.Name,
	}, nil
}

func (s *service) UpdatePlant(ctx context.Context, oldName, newName string) (*dto.Plant, error) {
	if oldName == "" {
		s.lg.Error("Plant old name is empty", zap.String("plant_old_name", oldName))
		return nil, errors.New("validation error")
	}
	if newName == "" {
		s.lg.Error("Plant new name is empty", zap.String("plant_new_name", newName))
		return nil, errors.New("validation error")
	}

	s.lg.Info("Update plant",
		zap.String("plant_old_name", oldName),
		zap.String("plant_new_name", newName),
	)

	plant, err := s.rep.GetPlantByName(ctx, oldName)
	if err != nil {
		s.lg.Error("Can't check plant name", zap.Error(err))
		return nil, errors.New("internal error")
	}
	if plant == nil {
		s.lg.Error("Plant not found", zap.String("plant_old_name", oldName))
		return nil, errors.New("not found")
	}

	plant.Name = newName
	if err = s.rep.UpdatePlant(ctx, plant); err != nil {
		s.lg.Error("Can't update plant", zap.Error(err),
			zap.String("plant_old_name", oldName),
			zap.String("plant_new_name", newName),
		)
		return nil, errors.New("internal error")
	}
	return &dto.Plant{
		ID:   plant.ID,
		Name: plant.Name,
	}, nil
}

func (s *service) GetPlantByName(ctx context.Context, name string) (*dto.Plant, error) {
	if name == "" {
		s.lg.Error("New plant name is empty", zap.String("plant_name", name))
		return nil, errors.New("validation error")
	}

	s.lg.Info("Get plant by name", zap.String("plant_name", name))

	plant, err := s.rep.GetPlantByName(ctx, name)
	if err != nil {
		s.lg.Error("Can't check plant name", zap.Error(err))
		return nil, errors.New("internal error")
	}
	if plant == nil {
		s.lg.Error("Plant not found", zap.String("plant_name", name))
		return nil, errors.New("not found")
	}
	return &dto.Plant{
		ID:   plant.ID,
		Name: plant.Name,
	}, nil
}

func (s *service) GetPlants(ctx context.Context) ([]*dto.Plant, error) {
	s.lg.Info("Get plants")

	plants, err := s.rep.GetPlants(ctx)
	if err != nil {
		s.lg.Error("Can't get all plants", zap.Error(err))
		return nil, err
	}
	if plants == nil {
		s.lg.Error("Plants slice is nil")
		return nil, errors.New("internal error")
	}

	return helpers.ConvertSlicesDAOToDTO(plants), nil
}

func (s *service) DeletePlant(ctx context.Context, name string) error {
	if name == "" {
		s.lg.Error("New plant name is empty", zap.String("plant_name", name))
		return errors.New("validation error")
	}

	s.lg.Info("Delete plant by name", zap.String("plant_name", name))

	plant, err := s.rep.GetPlantByName(ctx, name)
	if err != nil {
		s.lg.Error("Can't check plant name", zap.Error(err))
		return errors.New("internal error")
	}
	if plant == nil {
		s.lg.Error("Plant not found", zap.String("plant_name", name))
		return errors.New("not found")
	}

	if err = s.rep.DeletePlant(ctx, *plant); err != nil {
		s.lg.Error("Can't delete plant", zap.Error(err))
		return errors.New("internal error")
	}

	return nil
}
