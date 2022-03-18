package transport

import (
	"net/http"

	"github.com/Chicocs/challenge-fullStack-digitalRepublic/internal/domain"
	"github.com/labstack/echo/v4"
)

func (s *server) calculateAmountOfPaintCansNeeded(c echo.Context) error {
	var walls []domain.Wall

	if err := c.Bind(&walls); err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Error string `json:"error"`
		}{err.Error()})
	}

	resp, err := s.endpoint.CalculateAmountOfPaintCansNeededEndpoint(c.Request().Context(), walls)
	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Error string `json:"error"`
		}{err.Error()})
	}

	return c.JSON(resp.(domain.CustomerResponse).Code, resp.(domain.CustomerResponse).Response)
}
