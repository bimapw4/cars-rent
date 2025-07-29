package order

import (
	"car-rent/internal/common"
	"car-rent/internal/entity"
	"car-rent/internal/presentations"
	"car-rent/internal/repositories"
	"car-rent/pkg/meta"
	"time"

	"context"
)

type Orders interface {
	Create(ctx context.Context, payload entity.Order) (*presentations.Order, error)
	List(ctx context.Context, m *meta.Params) ([]presentations.Order, error)
	Detail(ctx context.Context, carsID int) (*presentations.Order, error)
	Update(ctx context.Context, payload entity.Order, carsID int) (*presentations.Order, error)
	Delete(ctx context.Context, carsID int) error
	Activate(ctx context.Context, carsID int) error
	Deactivate(ctx context.Context, carsID int) error
}

type business struct {
	repo *repositories.Repository
}

func NewBusiness(repo *repositories.Repository) Orders {
	return &business{
		repo: repo,
	}
}

func (b *business) Create(ctx context.Context, payload entity.Order) (*presentations.Order, error) {

	userctx := common.GetUserCtx(ctx)

	exist, _ := b.repo.Cars.CheckAvailableCar(ctx, payload.CarID, payload.PickupDate, payload.DropoffDate)
	if exist != nil {
		return nil, presentations.ErrCarsNotAvailable
	}

	cars, err := b.repo.Cars.Detail(ctx, payload.CarID)
	if err != nil {
		return nil, err
	}

	totalPayment := common.CalculateRentalCost(payload.PickupDate, payload.DropoffDate, cars.DayRate, cars.MonthRate)

	err = b.repo.Order.Create(ctx, presentations.Order{
		CarID:           payload.CarID,
		OrderDate:       payload.OrderDate,
		PickupDate:      payload.PickupDate,
		DropoffDate:     payload.DropoffDate,
		PickupLocation:  payload.PickupLocation,
		DropoffLocation: payload.DropoffLocation,
		TotalPayment:    totalPayment,
		UserID:          userctx.UserID,
	})
	if err != nil {
		return nil, err
	}

	res, err := b.repo.Order.Latest(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (b *business) List(ctx context.Context, m *meta.Params) ([]presentations.Order, error) {

	res, err := b.repo.Order.List(ctx, m)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (b *business) Detail(ctx context.Context, carsID int) (*presentations.Order, error) {

	res, err := b.repo.Order.Detail(ctx, carsID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (b *business) Update(ctx context.Context, payload entity.Order, orderID int) (*presentations.Order, error) {

	userctx := common.GetUserCtx(ctx)

	order, err := b.repo.Order.Detail(ctx, orderID)
	if err != nil {
		return nil, err
	}

	if order.CarID != payload.CarID {
		exist, _ := b.repo.Cars.CheckAvailableCar(ctx, payload.CarID, payload.PickupDate, payload.DropoffDate)
		if exist != nil {
			return nil, presentations.ErrCarsNotAvailable
		}
	}

	cars, err := b.repo.Cars.Detail(ctx, orderID)
	if err != nil {
		return nil, err
	}

	totalPayment := common.CalculateRentalCost(payload.PickupDate, payload.DropoffDate, cars.DayRate, cars.MonthRate)

	data := presentations.Order{
		OrderID:         order.OrderID,
		CarID:           payload.CarID,
		OrderDate:       payload.OrderDate,
		PickupDate:      payload.PickupDate,
		DropoffDate:     payload.DropoffDate,
		PickupLocation:  payload.PickupLocation,
		DropoffLocation: payload.DropoffLocation,
		TotalPayment:    totalPayment,
		UserID:          userctx.UserID,
		IsActive:        order.IsActive,
		CreatedAt:       order.CreatedAt,
		UpdatedAt:       time.Now().Local(),
	}

	err = b.repo.Order.Update(ctx, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (b *business) Delete(ctx context.Context, carsID int) error {
	_, err := b.repo.Order.Detail(ctx, carsID)
	if err != nil {
		return err
	}

	err = b.repo.Order.Delete(ctx, carsID)
	if err != nil {
		return err
	}
	return nil
}

func (b *business) Activate(ctx context.Context, carsID int) error {
	cars, err := b.repo.Order.DetailWithoutIsActive(ctx, carsID)
	if err != nil {
		return err
	}

	if cars.IsActive {
		return presentations.ErrOrdersAlreadyActivate
	}

	err = b.repo.Order.UpdateIsActive(ctx, carsID, true)
	if err != nil {
		return err
	}
	return nil
}

func (b *business) Deactivate(ctx context.Context, carsID int) error {
	cars, err := b.repo.Order.DetailWithoutIsActive(ctx, carsID)
	if err != nil {
		return err
	}

	if !cars.IsActive {
		return presentations.ErrOrdersAlreadyDeactivate
	}

	err = b.repo.Order.UpdateIsActive(ctx, carsID, false)
	if err != nil {
		return err
	}
	return nil
}
