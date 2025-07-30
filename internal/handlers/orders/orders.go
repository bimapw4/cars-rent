package orders

import (
	"car-rent/internal/business"
	"car-rent/internal/common"
	"car-rent/internal/entity"
	"car-rent/internal/presentations"
	"car-rent/internal/response"
	"car-rent/pkg/meta"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	List(c *fiber.Ctx) error
	Detail(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	Activate(c *fiber.Ctx) error
	Deactivate(c *fiber.Ctx) error
	Summary(c *fiber.Ctx) error
}

type handler struct {
	business business.Business
}

func NewHandler(business business.Business) Handler {
	return &handler{
		business: business,
	}
}

func (h *handler) Create(c *fiber.Ctx) error {
	var (
		Entity = "CreateOrder"
	)
	errAvail := common.DefaultAvailableErrors()
	custErr := errAvail.CustomeError(common.AvailableErrors{
		{
			Code:    http.StatusNotFound,
			Err:     presentations.ErrOrdersNotExist,
			Message: presentations.ErrOrdersNotExist.Error(),
		},
		{
			Code:    http.StatusNotFound,
			Err:     presentations.ErrCarsNotExist,
			Message: presentations.ErrCarsNotExist.Error(),
		},
		{
			Code:    http.StatusConflict,
			Err:     presentations.ErrCarsAlreadyExist,
			Message: presentations.ErrCarsAlreadyExist.Error(),
		},
		{
			Code:    http.StatusConflict,
			Err:     presentations.ErrCarsNotAvailable,
			Message: presentations.ErrCarsNotAvailable.Error(),
		},
	})

	var payload entity.Order
	if err := c.BodyParser(&payload); err != nil {
		return response.NewResponse(Entity).
			Errors("err parse body payload", err.Error()).
			JSON(c, http.StatusUnprocessableEntity)
	}

	if err := payload.Validate(); err != nil {
		errs := custErr.GetError(err)
		return response.NewResponse(Entity).
			Errors("err parse body payload", errs.Err).
			JSON(c, http.StatusUnprocessableEntity)
	}

	res, err := h.business.Order.Create(c.UserContext(), payload)
	if err != nil {
		errs := custErr.GetError(err)
		return response.NewResponse(Entity).
			Errors("err create order", errs.Message).
			JSON(c, errs.Code)
	}

	return response.NewResponse(Entity).
		Success("Successfully Create Order", res).
		JSON(c, http.StatusCreated)
}

func (h *handler) Update(c *fiber.Ctx) error {
	var (
		Entity = "UpdateOrder"
	)
	errAvail := common.DefaultAvailableErrors()
	custErr := errAvail.CustomeError(common.AvailableErrors{
		{
			Code:    http.StatusNotFound,
			Err:     presentations.ErrOrdersNotExist,
			Message: presentations.ErrOrdersNotExist.Error(),
		},
		{
			Code:    http.StatusNotFound,
			Err:     presentations.ErrCarsNotExist,
			Message: presentations.ErrCarsNotExist.Error(),
		},
		{
			Code:    http.StatusConflict,
			Err:     presentations.ErrCarsAlreadyExist,
			Message: presentations.ErrCarsAlreadyExist.Error(),
		},
		{
			Code:    http.StatusConflict,
			Err:     presentations.ErrCarsNotAvailable,
			Message: presentations.ErrCarsNotAvailable.Error(),
		},
	})

	var payload entity.Order
	if err := c.BodyParser(&payload); err != nil {
		return response.NewResponse(Entity).
			Errors("err parse body payload", err.Error()).
			JSON(c, http.StatusBadRequest)
	}

	if err := payload.Validate(); err != nil {
		errs := custErr.GetError(err)
		return response.NewResponse(Entity).
			Errors("err parse body payload", errs.Err).
			JSON(c, http.StatusUnprocessableEntity)
	}

	orderID := c.Params("order_id")
	intorder, err := strconv.Atoi(orderID)
	if err != nil {
		return response.NewResponse(Entity).
			Errors("err parse params", err.Error()).
			JSON(c, http.StatusBadRequest)
	}

	res, err := h.business.Order.Update(c.UserContext(), payload, intorder)
	if err != nil {
		fmt.Println("err =", err)
		errs := custErr.GetError(err)
		return response.NewResponse(Entity).
			Errors("err update order", errs.Message).
			JSON(c, errs.Code)
	}

	return response.NewResponse(Entity).
		Success("Successfully Update order", res).
		JSON(c, http.StatusAccepted)
}

func (h *handler) List(c *fiber.Ctx) error {
	var (
		Entity = "ListOrder"
	)
	errAvail := common.DefaultAvailableErrors()
	custErr := errAvail.CustomeError(common.AvailableErrors{
		{
			Code:    http.StatusNotFound,
			Err:     presentations.ErrOrdersNotExist,
			Message: presentations.ErrOrdersNotExist.Error(),
		},
		{
			Code:    http.StatusConflict,
			Err:     presentations.ErrCarsAlreadyExist,
			Message: presentations.ErrCarsAlreadyExist.Error(),
		},
	})

	q := c.Queries()
	m := meta.NewParams(q)

	res, err := h.business.Order.List(c.UserContext(), &m)
	if err != nil {
		fmt.Println("err ", err)
		errs := custErr.GetError(err)
		return response.NewResponse(Entity).
			Errors("err parse body payload", errs.Message).
			JSON(c, errs.Code)
	}

	return response.NewResponse(Entity).
		SuccessWithMeta("Successfully List order", res, m).
		JSON(c, http.StatusOK)
}

func (h *handler) Detail(c *fiber.Ctx) error {
	var (
		Entity = "DetailOrder"
	)
	errAvail := common.DefaultAvailableErrors()
	custErr := errAvail.CustomeError(common.AvailableErrors{
		{
			Code:    http.StatusNotFound,
			Err:     presentations.ErrOrdersNotExist,
			Message: presentations.ErrOrdersNotExist.Error(),
		},
		{
			Code:    http.StatusConflict,
			Err:     presentations.ErrCarsAlreadyExist,
			Message: presentations.ErrCarsAlreadyExist.Error(),
		},
	})

	orderID := c.Params("order_id")
	intorder, err := strconv.Atoi(orderID)
	if err != nil {
		return response.NewResponse(Entity).
			Errors("err parse params", err.Error()).
			JSON(c, http.StatusBadRequest)
	}

	res, err := h.business.Order.Detail(c.UserContext(), intorder)
	if err != nil {
		errs := custErr.GetError(err)
		return response.NewResponse(Entity).
			Errors("err get detail order", errs.Message).
			JSON(c, errs.Code)
	}

	return response.NewResponse(Entity).
		Success("Successfully Detail order", res).
		JSON(c, http.StatusOK)
}

func (h *handler) Delete(c *fiber.Ctx) error {
	var (
		Entity = "DeleteOrder"
	)
	errAvail := common.DefaultAvailableErrors()
	custErr := errAvail.CustomeError(common.AvailableErrors{
		{
			Code:    http.StatusNotFound,
			Err:     presentations.ErrOrdersNotExist,
			Message: presentations.ErrOrdersNotExist.Error(),
		},
		{
			Code:    http.StatusConflict,
			Err:     presentations.ErrCarsAlreadyExist,
			Message: presentations.ErrCarsAlreadyExist.Error(),
		},
	})

	orderID := c.Params("order_id")
	intorder, err := strconv.Atoi(orderID)
	if err != nil {
		return response.NewResponse(Entity).
			Errors("err parse params", err.Error()).
			JSON(c, http.StatusBadRequest)
	}

	err = h.business.Order.Delete(c.UserContext(), intorder)
	if err != nil {
		errs := custErr.GetError(err)
		return response.NewResponse(Entity).
			Errors("err delete order", errs.Message).
			JSON(c, errs.Code)
	}

	return response.NewResponse(Entity).
		Success("Successfully Delete order", nil).
		JSON(c, http.StatusOK)
}

func (h *handler) Activate(c *fiber.Ctx) error {
	var (
		Entity = "ActivateOrder"
	)
	errAvail := common.DefaultAvailableErrors()
	custErr := errAvail.CustomeError(common.AvailableErrors{
		{
			Code:    http.StatusNotFound,
			Err:     presentations.ErrOrdersNotExist,
			Message: presentations.ErrOrdersNotExist.Error(),
		},
		{
			Code:    http.StatusConflict,
			Err:     presentations.ErrCarsAlreadyExist,
			Message: presentations.ErrCarsAlreadyExist.Error(),
		},
		{
			Code:    http.StatusConflict,
			Err:     presentations.ErrOrdersAlreadyActivate,
			Message: presentations.ErrOrdersAlreadyActivate.Error(),
		},
	})

	orderID := c.Params("order_id")
	intorder, err := strconv.Atoi(orderID)
	if err != nil {
		return response.NewResponse(Entity).
			Errors("err parse params", err.Error()).
			JSON(c, http.StatusBadRequest)
	}

	err = h.business.Order.Activate(c.UserContext(), intorder)
	if err != nil {
		errs := custErr.GetError(err)
		return response.NewResponse(Entity).
			Errors("err activate order", errs.Message).
			JSON(c, errs.Code)
	}

	return response.NewResponse(Entity).
		Success("Successfully Activate Order", nil).
		JSON(c, http.StatusOK)
}

func (h *handler) Deactivate(c *fiber.Ctx) error {
	var (
		Entity = "DeactivateOrder"
	)
	errAvail := common.DefaultAvailableErrors()
	custErr := errAvail.CustomeError(common.AvailableErrors{
		{
			Code:    http.StatusNotFound,
			Err:     presentations.ErrOrdersNotExist,
			Message: presentations.ErrOrdersNotExist.Error(),
		},
		{
			Code:    http.StatusConflict,
			Err:     presentations.ErrCarsAlreadyExist,
			Message: presentations.ErrCarsAlreadyExist.Error(),
		},
		{
			Code:    http.StatusConflict,
			Err:     presentations.ErrOrdersAlreadyDeactivate,
			Message: presentations.ErrOrdersAlreadyDeactivate.Error(),
		},
	})

	orderID := c.Params("order_id")
	intorder, err := strconv.Atoi(orderID)
	if err != nil {
		return response.NewResponse(Entity).
			Errors("err parse params", err.Error()).
			JSON(c, http.StatusBadRequest)
	}

	err = h.business.Order.Deactivate(c.UserContext(), intorder)
	if err != nil {
		errs := custErr.GetError(err)
		return response.NewResponse(Entity).
			Errors("err deactivate Order", errs.Message).
			JSON(c, errs.Code)
	}

	return response.NewResponse(Entity).
		Success("Successfully Deactivate Order", nil).
		JSON(c, http.StatusOK)
}

func (h *handler) Summary(c *fiber.Ctx) error {
	var (
		Entity = "SummaryOrder"
	)
	errAvail := common.DefaultAvailableErrors()
	custErr := errAvail.CustomeError(common.AvailableErrors{
		{
			Code:    http.StatusNotFound,
			Err:     presentations.ErrOrdersNotExist,
			Message: presentations.ErrOrdersNotExist.Error(),
		},
		{
			Code:    http.StatusConflict,
			Err:     presentations.ErrCarsAlreadyExist,
			Message: presentations.ErrCarsAlreadyExist.Error(),
		},
	})

	res, err := h.business.Order.Summary(c.UserContext())
	if err != nil {
		fmt.Println("err", err)
		errs := custErr.GetError(err)
		return response.NewResponse(Entity).
			Errors("err get summary order", errs.Message).
			JSON(c, errs.Code)
	}

	return response.NewResponse(Entity).
		Success("Successfully Summary Order", res).
		JSON(c, http.StatusOK)
}
