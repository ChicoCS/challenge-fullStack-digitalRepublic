package service

import (
	"github.com/Chicocs/challenge-fullStack-digitalRepublic/internal/service/paint"
	"github.com/go-kit/log"
	"github.com/jmoiron/sqlx"
)

type ServiceFactory interface {
	Paint() paint.PaintService
}

type serviceFactory struct {
	paint paint.PaintService
}

func NewServiceFactory(db *sqlx.DB, logger log.Logger) ServiceFactory {
	return &serviceFactory{
		paint: paint.NewService(db, logger),
	}
}

func (sFactory *serviceFactory) Paint() paint.PaintService {
	return sFactory.paint
}
