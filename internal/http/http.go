package http

import (
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/app/city_mapping"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/factory"
	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	city_mapping.NewHandler(f).Route(e.Group("/city_mapping"))
}
