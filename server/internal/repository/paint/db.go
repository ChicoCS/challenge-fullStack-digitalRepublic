package paint

import (
	"github.com/Chicocs/challenge-fullStack-digitalRepublic/internal/domain"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type PaintRepository interface {
	GetWallItems() ([]domain.WallItems, error)
}

type paintRepository struct {
	db *sqlx.DB
}

func NewPaintRepository(db *sqlx.DB) PaintRepository {
	return &paintRepository{
		db: db,
	}
}

func (pRepository *paintRepository) GetWallItems() ([]domain.WallItems, error) {
	var items []domain.WallItems

	err := pRepository.db.MustExec(qryGetWallItems)
	if err != nil {
		logrus.Error(err)
		return []domain.WallItems{}, nil
	}

	return items, nil
}
