package placerepo

import (
	placemodel "airbnb-golang/internal/places/model"
	"airbnb-golang/pkg/common"
	"context"

	"gorm.io/gorm"
)

type placeRepo struct {
	db *gorm.DB
}

func NewPlaceRepo(db *gorm.DB) *placeRepo {
	return &placeRepo{db}
}

func (r *placeRepo) Create(ctx context.Context, place *placemodel.Place) error {
	db := r.db.Begin()
	if err := db.Table(place.TableName()).Create(&place).Error; err != nil {
		db.Rollback()
		return err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}
	return nil

}
func (r *placeRepo) ListDataWithCondition(ctx context.Context, paging *common.Paging, filter *placemodel.Filter) ([]placemodel.Place, error) {
	var data []placemodel.Place

	db := r.db.Table(placemodel.Place{}.TableName())

	if v := filter.CityId; v > 0 {
		db = db.Where("city_id = ?", v)
	}
	if v := filter.OwnerId; v > 0 {
		db = db.Where("owner_id = ?", v)
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	offset := paging.Limit * (paging.Page - 1)
	if err := db.Offset(offset).Limit(paging.Limit).Find(&data).Error; err != nil {
		return nil, err

	}
	return data, nil
}

func (r *placeRepo) FindDataWithCondition(ctx context.Context, condition map[string]any) (*placemodel.Place, error) {
	var data placemodel.Place
	if err := r.db.Where(condition).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *placeRepo) DeletePlace(ctx context.Context, id int) error {
	if err := r.db.Table(placemodel.Place{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		return err
	}
	return nil

}

func (r *placeRepo) UpdatePlace(ctx context.Context, condition map[string]any, place *placemodel.Place) error {
	if err := r.db.Table(place.TableName()).Where(condition).Updates(place).Error; err != nil {
		return err
	}
	return nil
}
