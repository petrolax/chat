package repository

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/petrolax/chat/pkg/plants/dao"
	"go.uber.org/zap"
)

type repository struct {
	db *gorm.DB
	lg *zap.Logger
}

func NewRepository(db *gorm.DB, lg *zap.Logger) *repository {
	return &repository{
		db: db,
		lg: lg,
	}
}

func (r *repository) AddPlant(ctx context.Context, newPlant *dao.Plant) error {
	return r.db.Create(newPlant).Error
}

func (r *repository) GetPlantByName(ctx context.Context, name string) (*dao.Plant, error) {
	var plant dao.Plant
	if err := r.db.Where("name = ?", name).First(&plant).Error; err != nil {
		return nil, err
	}

	return &plant, nil
}

func (r *repository) GetPlants(ctx context.Context) ([]*dao.Plant, error) {
	var plants []*dao.Plant
	if err := r.db.Model(&dao.Plant{}).Find(&plants).Error; err != nil {
		return nil, err
	}

	return plants, nil
}

func (r *repository) UpdatePlant(ctx context.Context, plant *dao.Plant) error {
	return r.db.Save(plant).Error
}

func (r *repository) DeletePlant(ctx context.Context, plant dao.Plant) error {
	return r.db.Delete(&plant).Error
}
