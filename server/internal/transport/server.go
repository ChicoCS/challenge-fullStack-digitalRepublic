package transport

import (
	"context"
	"net/http"

	"github.com/Chicocs/challenge-fullStack-digitalRepublic/internal/endpoint"
	"github.com/go-kit/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type server struct {
	endpoint *endpoint.Endpoints
	logger   *log.Logger
}

func NewService(context context.Context, endpoint *endpoint.Endpoints, logger *log.Logger) http.Handler {
	rest := &server{
		endpoint: endpoint,
		logger:   logger,
	}

	e := echo.New()
	e.Use(middleware.CORS())
	e.POST("/paint/calculate", rest.calculateAmountOfPaintCansNeeded)

	e.Logger.Fatal(e.Start(":8080"))

	return e
}

func (s *server) ReturnJSONError(c echo.Context, code int, message string) {
	c.JSON(code, struct {
		Message string `json:"message"`
	}{message})
}
