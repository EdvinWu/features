package server

import (
	"features/internal/domain/feature/handler"
	"net/http"

	"github.com/labstack/echo/v4"
	echoMW "github.com/labstack/echo/v4/middleware"
)

func Echo(address string, feature handler.Feature) *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	e.Use(echoMW.CORSWithConfig(echoMW.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodOptions, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.Server.Addr = address
	e.GET("/ping", PingHandler)
	featureRoutes(e, feature)
	return e
}

func featureRoutes(e *echo.Echo, feature handler.Feature) {
	g := e.Group("api/v1/feature")
	g.POST("/", feature.CreateFeature)
	g.POST("/update", feature.UpdateFeature)
	e.POST("/archive", feature.ArchiveFeature)
	e.GET("/", feature.GetAll)
	e.POST("/:id", feature.FindByCustomerID)
}

func PingHandler(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
