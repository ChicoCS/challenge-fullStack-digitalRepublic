package endpoint

import (
	"github.com/Chicocs/challenge-fullStack-digitalRepublic/internal/domain"
	"github.com/Chicocs/challenge-fullStack-digitalRepublic/internal/service"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
)

type Endpoints struct {
	CalculateAmountOfPaintCansNeededEndpoint endpoint.Endpoint
}

func MakeEndpoints(s service.ServiceFactory, logger log.Logger) Endpoints {
	return Endpoints{
		CalculateAmountOfPaintCansNeededEndpoint: makeCalculateAmountOfPaintCansNeededEndpoint(s, logger),
	}
}

func newCustomerResponse(code int, response interface{}, err error) (interface{}, error) {

	return domain.CustomerResponse{
		Code:     code,
		Response: response,
	}, err
}
