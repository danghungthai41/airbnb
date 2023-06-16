package placeusecase

import (
	placemodel "airbnb-golang/internal/places/model"
	"airbnb-golang/pkg/common"
	"context"
)

type IPlaceRepository interface {
	Create(context.Context, *placemodel.Place) error
	ListDataWithCondition(context.Context, *common.Paging, *placemodel.Filter) ([]placemodel.Place, error)
	FindDataWithCondition(context.Context, map[string]any) (*placemodel.Place, error)
	DeletePlace(context.Context, int) error
}
type placeUC struct {
	placeRepo IPlaceRepository
}

func NewPlaceUC(placerepo IPlaceRepository) *placeUC {
	return &placeUC{placerepo}
}

func (uc *placeUC) CreatePlace(ctx context.Context, place *placemodel.Place) (*placemodel.Place, error) {

	if err := place.Validate(); err != nil {
		return nil, err
	}

	if err := uc.placeRepo.Create(ctx, place); err != nil {
		return nil, err

	}
	return place, nil
}

func (uc *placeUC) GetPlaces(ctx context.Context, paging *common.Paging, filter *placemodel.Filter) ([]placemodel.Place, error) {
	data, err := uc.placeRepo.ListDataWithCondition(ctx, paging, filter)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (uc *placeUC) GetPlaceByID(ctx context.Context, id int) (*placemodel.Place, error) {

	data, err := uc.placeRepo.FindDataWithCondition(ctx, map[string]any{"id": id})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (uc *placeUC) DeletePlace(ctx context.Context, id int) error {
	if err := uc.placeRepo.DeletePlace(ctx, id); err != nil {
		return err
	}
	return nil
}
