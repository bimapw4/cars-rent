package cars

import (
	"car-rent/internal/business"
	"car-rent/internal/common"
	"car-rent/internal/entity"
	"car-rent/internal/presentations"
	"car-rent/internal/response"
	"car-rent/pkg/meta"
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

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
	AvailableCars(c *fiber.Ctx) error
	Preview(c *fiber.Ctx) error
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
		Entity = "CreateCars"
	)
	errAvail := common.DefaultAvailableErrors()
	custErr := errAvail.CustomeError(common.AvailableErrors{
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
	})

	userctx := common.GetUserCtx(c.UserContext())
	if !userctx.IsAdmin {
		return response.NewResponse(Entity).
			Errors("Failed payroll generate payslip", "only admin allowed").
			JSON(c, fiber.StatusForbidden)
	}

	payload := entity.Cars{}

	carsName := c.FormValue("cars_name")
	if carsName == "" {
		return response.NewResponse(Entity).
			Errors("Failed to parse request body", "cars_name is required").
			JSON(c, fiber.StatusUnprocessableEntity)
	}

	dayRate := c.FormValue("day_rate")
	if dayRate == "" {
		return response.NewResponse(Entity).
			Errors("Failed to parse request body", "day_rate is required").
			JSON(c, fiber.StatusUnprocessableEntity)
	}

	monthRate := c.FormValue("month_rate")
	if monthRate == "" {
		return response.NewResponse(Entity).
			Errors("Failed to parse request body", "month_rate is required").
			JSON(c, fiber.StatusUnprocessableEntity)
	}

	image, err := c.FormFile("image")
	if err != nil {
		return response.NewResponse(Entity).
			Errors("Failed to parse request body", err).
			JSON(c, fiber.StatusBadRequest)
	}

	ext := filepath.Ext(image.Filename)
	switch ext {
	case ".png", ".jpg", ".jpeg":
	default:
		err := []string{"content type must be a html"}
		return response.NewResponse(Entity).
			Errors("Failed to parse request body", err).
			JSON(c, fiber.StatusBadRequest)
	}

	if image.Size > (5 * 1024 * 1024) {
		err = errors.New("file size cannot be more than 5MB")
		return response.NewResponse(Entity).
			Errors("Failed to parse request body", err).
			JSON(c, fiber.StatusBadRequest)
	}

	path := fmt.Sprintf("./storage/%v-%v", time.Now().Unix(), image.Filename)

	err = c.SaveFile(image, path)
	if err != nil {
		log.Println(err)
		return response.NewResponse(Entity).
			Errors("Failed to parse request body", err).
			JSON(c, fiber.StatusBadRequest)
	}

	floatdayRate, err := strconv.ParseFloat(dayRate, 64)
	if err != nil {
		return response.NewResponse(Entity).
			Errors("Failed to parse request body", err).
			JSON(c, fiber.StatusUnprocessableEntity)
	}

	floatMonthRate, err := strconv.ParseFloat(monthRate, 64)
	if err != nil {
		return response.NewResponse(Entity).
			Errors("Failed to parse request body", err).
			JSON(c, fiber.StatusUnprocessableEntity)
	}

	payload.Image = path
	payload.CarsName = carsName
	payload.DayRate = floatdayRate
	payload.MonthRate = floatMonthRate

	res, err := h.business.Cars.Create(c.UserContext(), payload)
	if err != nil {
		errs := custErr.GetError(err)
		return response.NewResponse(Entity).
			Errors("err create cars", errs.Message).
			JSON(c, errs.Code)
	}

	return response.NewResponse(Entity).
		Success("Successfully Create Cars", res).
		JSON(c, http.StatusCreated)
}

func (h *handler) Update(c *fiber.Ctx) error {
	var (
		Entity = "UpdateCars"
	)
	errAvail := common.DefaultAvailableErrors()
	custErr := errAvail.CustomeError(common.AvailableErrors{
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
	})

	userctx := common.GetUserCtx(c.UserContext())
	if !userctx.IsAdmin {
		return response.NewResponse(Entity).
			Errors("Failed payroll generate payslip", "only admin allowed").
			JSON(c, fiber.StatusForbidden)
	}

	payload := entity.Cars{}

	carsName := c.FormValue("cars_name")
	if carsName == "" {
		return response.NewResponse(Entity).
			Errors("Failed to parse request body", "cars_name is required").
			JSON(c, fiber.StatusUnprocessableEntity)
	}

	dayRate := c.FormValue("day_rate")
	if dayRate == "" {
		return response.NewResponse(Entity).
			Errors("Failed to parse request body", "day_rate is required").
			JSON(c, fiber.StatusUnprocessableEntity)
	}

	monthRate := c.FormValue("month_rate")
	if monthRate == "" {
		return response.NewResponse(Entity).
			Errors("Failed to parse request body", "month_rate is required").
			JSON(c, fiber.StatusUnprocessableEntity)
	}

	image, err := c.FormFile("image")
	if err != nil {
		return response.NewResponse(Entity).
			Errors("Failed to parse request body", err).
			JSON(c, fiber.StatusBadRequest)
	}

	ext := filepath.Ext(image.Filename)
	switch ext {
	case ".png", ".jpg", ".jpeg":
	default:
		err := []string{"content type must be a html"}
		return response.NewResponse(Entity).
			Errors("Failed to parse request body", err).
			JSON(c, fiber.StatusBadRequest)
	}

	if image.Size > (5 * 1024 * 1024) {
		err = errors.New("file size cannot be more than 5MB")
		return response.NewResponse(Entity).
			Errors("Failed to parse request body", err).
			JSON(c, fiber.StatusBadRequest)
	}

	path := fmt.Sprintf("./storage/%v-%v", time.Now().Unix(), image.Filename)

	err = c.SaveFile(image, path)
	if err != nil {
		log.Println(err)
		return response.NewResponse(Entity).
			Errors("Failed to parse request body", err).
			JSON(c, fiber.StatusBadRequest)
	}

	floatdayRate, err := strconv.ParseFloat(dayRate, 64)
	if err != nil {
		return response.NewResponse(Entity).
			Errors("Failed to parse request body", err).
			JSON(c, fiber.StatusUnprocessableEntity)
	}

	floatMonthRate, err := strconv.ParseFloat(monthRate, 64)
	if err != nil {
		return response.NewResponse(Entity).
			Errors("Failed to parse request body", err).
			JSON(c, fiber.StatusUnprocessableEntity)
	}

	payload.Image = path
	payload.CarsName = carsName
	payload.DayRate = floatdayRate
	payload.MonthRate = floatMonthRate

	carsID := c.Params("cars_id")
	intCars, err := strconv.Atoi(carsID)
	if err != nil {
		return response.NewResponse(Entity).
			Errors("err parse params", err.Error()).
			JSON(c, http.StatusBadRequest)
	}

	res, err := h.business.Cars.Update(c.UserContext(), payload, intCars)
	if err != nil {
		fmt.Println("err =", err)
		errs := custErr.GetError(err)
		return response.NewResponse(Entity).
			Errors("err update cars", errs.Message).
			JSON(c, errs.Code)
	}

	return response.NewResponse(Entity).
		Success("Successfully Update Cars", res).
		JSON(c, http.StatusAccepted)
}

func (h *handler) List(c *fiber.Ctx) error {
	var (
		Entity = "ListCars"
	)
	errAvail := common.DefaultAvailableErrors()
	custErr := errAvail.CustomeError(common.AvailableErrors{
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
	})

	q := c.Queries()
	m := meta.NewParams(q)

	res, err := h.business.Cars.List(c.UserContext(), &m)
	if err != nil {
		errs := custErr.GetError(err)
		return response.NewResponse(Entity).
			Errors("err parse body payload", errs.Message).
			JSON(c, errs.Code)
	}

	return response.NewResponse(Entity).
		SuccessWithMeta("Successfully List Cars", res, m).
		JSON(c, http.StatusOK)
}

func (h *handler) Detail(c *fiber.Ctx) error {
	var (
		Entity = "DetailCars"
	)
	errAvail := common.DefaultAvailableErrors()
	custErr := errAvail.CustomeError(common.AvailableErrors{
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
	})

	carsID := c.Params("cars_id")
	intCars, err := strconv.Atoi(carsID)
	if err != nil {
		return response.NewResponse(Entity).
			Errors("err parse params", err.Error()).
			JSON(c, http.StatusBadRequest)
	}

	res, err := h.business.Cars.Detail(c.UserContext(), intCars)
	if err != nil {
		errs := custErr.GetError(err)
		return response.NewResponse(Entity).
			Errors("err get detail cars", errs.Message).
			JSON(c, errs.Code)
	}

	return response.NewResponse(Entity).
		Success("Successfully Detail Cars", res).
		JSON(c, http.StatusOK)
}

func (h *handler) Delete(c *fiber.Ctx) error {
	var (
		Entity = "DeleteCars"
	)
	errAvail := common.DefaultAvailableErrors()
	custErr := errAvail.CustomeError(common.AvailableErrors{
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
	})

	userctx := common.GetUserCtx(c.UserContext())
	if !userctx.IsAdmin {
		return response.NewResponse(Entity).
			Errors("Failed payroll generate payslip", "only admin allowed").
			JSON(c, fiber.StatusForbidden)
	}

	carsID := c.Params("cars_id")
	intCars, err := strconv.Atoi(carsID)
	if err != nil {
		return response.NewResponse(Entity).
			Errors("err parse params", err.Error()).
			JSON(c, http.StatusBadRequest)
	}

	err = h.business.Cars.Delete(c.UserContext(), intCars)
	if err != nil {
		errs := custErr.GetError(err)
		return response.NewResponse(Entity).
			Errors("err delete cars", errs.Message).
			JSON(c, errs.Code)
	}

	return response.NewResponse(Entity).
		Success("Successfully Delete Cars", nil).
		JSON(c, http.StatusOK)
}

func (h *handler) Activate(c *fiber.Ctx) error {
	var (
		Entity = "ActivateCars"
	)
	errAvail := common.DefaultAvailableErrors()
	custErr := errAvail.CustomeError(common.AvailableErrors{
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
			Err:     presentations.ErrCarsAlreadyActivate,
			Message: presentations.ErrCarsAlreadyActivate.Error(),
		},
	})

	userctx := common.GetUserCtx(c.UserContext())
	if !userctx.IsAdmin {
		return response.NewResponse(Entity).
			Errors("Failed payroll generate payslip", "only admin allowed").
			JSON(c, fiber.StatusForbidden)
	}

	carsID := c.Params("cars_id")
	intCars, err := strconv.Atoi(carsID)
	if err != nil {
		return response.NewResponse(Entity).
			Errors("err parse params", err.Error()).
			JSON(c, http.StatusBadRequest)
	}

	err = h.business.Cars.Activate(c.UserContext(), intCars)
	if err != nil {
		errs := custErr.GetError(err)
		return response.NewResponse(Entity).
			Errors("err activate cars", errs.Message).
			JSON(c, errs.Code)
	}

	return response.NewResponse(Entity).
		Success("Successfully Activate Cars", nil).
		JSON(c, http.StatusOK)
}

func (h *handler) Deactivate(c *fiber.Ctx) error {
	var (
		Entity = "DeactivateCars"
	)
	errAvail := common.DefaultAvailableErrors()
	custErr := errAvail.CustomeError(common.AvailableErrors{
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
			Err:     presentations.ErrCarsAlreadyDeactivate,
			Message: presentations.ErrCarsAlreadyDeactivate.Error(),
		},
	})

	userctx := common.GetUserCtx(c.UserContext())
	if !userctx.IsAdmin {
		return response.NewResponse(Entity).
			Errors("Failed payroll generate payslip", "only admin allowed").
			JSON(c, fiber.StatusForbidden)
	}

	carsID := c.Params("cars_id")
	intCars, err := strconv.Atoi(carsID)
	if err != nil {
		return response.NewResponse(Entity).
			Errors("err parse params", err.Error()).
			JSON(c, http.StatusBadRequest)
	}

	err = h.business.Cars.Deactivate(c.UserContext(), intCars)
	if err != nil {
		errs := custErr.GetError(err)
		return response.NewResponse(Entity).
			Errors("err deactivate cars", errs.Message).
			JSON(c, errs.Code)
	}

	return response.NewResponse(Entity).
		Success("Successfully Deactivate Cars", nil).
		JSON(c, http.StatusOK)
}

func (h *handler) AvailableCars(c *fiber.Ctx) error {
	var (
		Entity = "AvailableCars"
	)

	errAvail := common.DefaultAvailableErrors()
	custErr := errAvail.CustomeError(common.AvailableErrors{
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
	})

	q := c.Queries()
	m := meta.NewParams(q)

	startDate := c.Query("start_date")
	if startDate == "" {
		return response.NewResponse(Entity).
			Errors("err parse query payload", "start_date cannot be blank").
			JSON(c, http.StatusBadRequest)
	}

	timstartDate, err := time.Parse(time.RFC3339, startDate)
	if err != nil {
		fmt.Println("sasa 1", err)
		errs := custErr.GetError(err)
		return response.NewResponse(Entity).
			Errors("err parse body payload", errs.Message).
			JSON(c, errs.Code)
	}

	endDate := c.Query("end_date")
	if startDate == "" {
		return response.NewResponse(Entity).
			Errors("err parse query payload", "end_date cannot be blank").
			JSON(c, http.StatusBadRequest)
	}

	timendDate, err := time.Parse(time.RFC3339, endDate)
	if err != nil {
		fmt.Println("sasa 2", err)
		errs := custErr.GetError(err)
		return response.NewResponse(Entity).
			Errors("err parse body payload", errs.Message).
			JSON(c, errs.Code)
	}

	if timstartDate.After(timendDate) {
		return response.NewResponse(Entity).
			Errors("err parse body payload", "end date cannot less than start date").
			JSON(c, http.StatusBadRequest)
	}

	res, err := h.business.Cars.AvailableCars(c.UserContext(), &m, timstartDate, timendDate)
	if err != nil {
		fmt.Println("sasa 3", err)
		errs := custErr.GetError(err)
		return response.NewResponse(Entity).
			Errors("err parse body payload", errs.Message).
			JSON(c, errs.Code)
	}

	return response.NewResponse(Entity).
		SuccessWithMeta("Successfully List Available Cars", res, m).
		JSON(c, http.StatusOK)
}

func (h *handler) Preview(c *fiber.Ctx) error {
	var (
		Entity = "PreviewCars"
	)
	errAvail := common.DefaultAvailableErrors()
	custErr := errAvail.CustomeError(common.AvailableErrors{
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
	})

	carsID := c.Params("cars_id")
	intCars, err := strconv.Atoi(carsID)
	if err != nil {
		return response.NewResponse(Entity).
			Errors("err parse params", err.Error()).
			JSON(c, http.StatusBadRequest)
	}

	res, err := h.business.Cars.Detail(c.UserContext(), intCars)
	if err != nil {
		errs := custErr.GetError(err)
		return response.NewResponse(Entity).
			Errors("err get preview cars", errs.Message).
			JSON(c, errs.Code)
	}

	return c.SendFile(res.Image)
}
