package endpoint

import (
	"context"
	"errors"
	"net/http"

	"github.com/Chicocs/challenge-fullStack-digitalRepublic/internal/domain"
	"github.com/Chicocs/challenge-fullStack-digitalRepublic/internal/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
)

func makeCalculateAmountOfPaintCansNeededEndpoint(s service.ServiceFactory, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		if request == nil {
			return newCustomerResponse(http.StatusBadRequest, errors.New("request without parameter"), nil)
		}

		wallsData := request.([]domain.Wall)
		resp, err := s.Paint().CalculateAmountOfPaintCansNeeded(ctx, wallsData)
		if err != nil {
			return newCustomerResponse(http.StatusBadRequest, resp, err)

		}
		return newCustomerResponse(http.StatusOK, resp, nil)
	}
}
