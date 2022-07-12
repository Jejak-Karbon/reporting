package city_mapping

import (	
	_ "github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"
	_ "github.com/born2ngopi/alterra/basic-echo-mvc/internal/middleware"
	res "github.com/born2ngopi/alterra/basic-echo-mvc/pkg/util/response"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

func (h *handler) Get(c echo.Context) error {

	result,err := h.service.Find(c.Request().Context())
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}