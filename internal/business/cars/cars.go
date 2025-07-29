package cars

import (
	"car-rent/internal/entity"
	"car-rent/internal/presentations"
	"car-rent/internal/repositories"
	"car-rent/pkg/meta"
	"time"

	"context"
)

type Cars interface {
	Create(ctx context.Context, payload entity.Cars) (*presentations.Cars, error)
	List(ctx context.Context, m *meta.Params) ([]presentations.Cars, error)
	Detail(ctx context.Context, carsID int) (*presentations.Cars, error)
	Update(ctx context.Context, payload entity.Cars, carsID int) (*presentations.Cars, error)
	Delete(ctx context.Context, carsID int) error
	Activate(ctx context.Context, carsID int) error
	Deactivate(ctx context.Context, carsID int) error
	AvailableCars(ctx context.Context, m *meta.Params, startdate, enddate time.Time) ([]presentations.Cars, error)
}

type business struct {
	repo *repositories.Repository
}

func NewBusiness(repo *repositories.Repository) Cars {
	return &business{
		repo: repo,
	}
}

func (b *business) Create(ctx context.Context, payload entity.Cars) (*presentations.Cars, error) {

	err := b.repo.Cars.Create(ctx, presentations.Cars{
		CarName:   payload.CarsName,
		DayRate:   payload.DayRate,
		MonthRate: payload.MonthRate,
		Image:     payload.Image,
		IsActive:  true,
	})
	if err != nil {
		return nil, err
	}

	res, err := b.repo.Cars.Latest(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (b *business) List(ctx context.Context, m *meta.Params) ([]presentations.Cars, error) {

	res, err := b.repo.Cars.List(ctx, m)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (b *business) Detail(ctx context.Context, carsID int) (*presentations.Cars, error) {

	res, err := b.repo.Cars.Detail(ctx, carsID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (b *business) Update(ctx context.Context, payload entity.Cars, carsID int) (*presentations.Cars, error) {

	cars, err := b.repo.Cars.Detail(ctx, carsID)
	if err != nil {
		return nil, err
	}

	data := presentations.Cars{
		CarID:     carsID,
		CarName:   payload.CarsName,
		DayRate:   payload.DayRate,
		MonthRate: payload.MonthRate,
		Image:     payload.Image,
		IsActive:  cars.IsActive,
		CreatedAT: cars.CreatedAT,
		UpdatedAt: time.Now().Local(),
	}

	err = b.repo.Cars.Update(ctx, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (b *business) Delete(ctx context.Context, carsID int) error {
	_, err := b.repo.Cars.Detail(ctx, carsID)
	if err != nil {
		return err
	}

	err = b.repo.Cars.Delete(ctx, carsID)
	if err != nil {
		return err
	}
	return nil
}

func (b *business) Activate(ctx context.Context, carsID int) error {
	cars, err := b.repo.Cars.DetailWithoutIsActive(ctx, carsID)
	if err != nil {
		return err
	}

	if cars.IsActive {
		return presentations.ErrCarsAlreadyActivate
	}

	err = b.repo.Cars.UpdateIsActive(ctx, carsID, true)
	if err != nil {
		return err
	}
	return nil
}

func (b *business) Deactivate(ctx context.Context, carsID int) error {
	cars, err := b.repo.Cars.DetailWithoutIsActive(ctx, carsID)
	if err != nil {
		return err
	}

	if !cars.IsActive {
		return presentations.ErrCarsAlreadyDeactivate
	}

	err = b.repo.Cars.UpdateIsActive(ctx, carsID, false)
	if err != nil {
		return err
	}
	return nil
}

func (b *business) AvailableCars(ctx context.Context, m *meta.Params, startdate, enddate time.Time) ([]presentations.Cars, error) {

	res, err := b.repo.Cars.CheckAvailableCars(ctx, m, startdate, enddate)
	if err != nil {
		return nil, err
	}

	return res, nil
}
