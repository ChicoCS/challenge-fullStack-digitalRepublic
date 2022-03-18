package transport

import (
	"net/http"

	"github.com/Chicocs/challenge-fullStack-digitalRepublic/internal/domain"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (s *server) calculateAmountOfPaintCansNeeded(c echo.Context) error {
	var walls []domain.Wall

	if err := c.Bind(&walls); err != nil {
		logrus.Error(err)
		s.ReturnJSONError(c, http.StatusBadRequest, err.Error())
		return err
	}

	resp, err := s.endpoint.CalculateAmountOfPaintCansNeededEndpoint(c.Request().Context(), walls)
	if err != nil {
		logrus.Error(err)
		s.ReturnJSONError(c, http.StatusBadRequest, err.Error())
		return err
	}

	return c.JSON(resp.(domain.CustomerResponse).Code, resp.(domain.CustomerResponse).Response)
}
