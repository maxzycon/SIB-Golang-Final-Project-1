package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/SIB-Golang-Final-Project-1/internal/domain/global/dto"
	"github.com/maxzycon/SIB-Golang-Final-Project-1/pkg/errors"
	"github.com/maxzycon/SIB-Golang-Final-Project-1/pkg/httputil"
)

func (c *GlobalController) handlerCreateTodos(f *fiber.Ctx) (err error) {
	payload := dto.PayloadTodos{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse body create Todos")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.CreateTodos(f.Context(), &payload)

	if err != nil {
		log.Errorf("err service at controller create Todos :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerUpdateTodos(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params update Todos")
		return httputil.WriteErrorResponse(f, err)
	}

	payload := dto.PayloadTodos{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse body update Todos")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.UpdateTodosById(f.Context(), id, &payload)

	if err != nil {
		log.Errorf("err service at controller update Todos :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *GlobalController) handlerGetTodosById(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params Todos get by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.GetTodosById(f.Context(), id)

	if err != nil {
		log.Errorf("err service at controller Todos get by id:%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerGetAllTodos(f *fiber.Ctx) (err error) {
	resp, err := c.globalService.GetAllTodos(f.Context())

	if err != nil {
		log.Errorf("err service at controller todos :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *GlobalController) handlerDeleteTodos(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params Todos delete by id")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.DeleteTodosById(f.Context(), id)

	if err != nil {
		log.Errorf("err service at controller Todos delete by id :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}
