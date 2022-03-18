package paint

import (
	"context"

	"github.com/Chicocs/challenge-fullStack-digitalRepublic/internal/domain"
	"github.com/Chicocs/challenge-fullStack-digitalRepublic/internal/repository/paint"
	"github.com/go-kit/log"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"gopkg.in/guregu/null.v4"
)

type PaintService interface {
	CalculateAmountOfPaintCansNeeded(ctx context.Context, walls []domain.Wall) (domain.Can, error)
}

type paintService struct {
	paintRepository paint.PaintRepository
	logger          log.Logger
}

func NewService(db *sqlx.DB, logger log.Logger) PaintService {
	return &paintService{
		paintRepository: paint.NewPaintRepository(db),
		logger:          logger,
	}
}

func (pService *paintService) CalculateAmountOfPaintCansNeeded(ctx context.Context, walls []domain.Wall) (domain.Can, error) {
	squareMetersToPaint, err := pService.calculateAreaRoom(walls)
	if err != nil {
		logrus.Error(err)
		return domain.Can{}, err
	}

	litersToPaint := squareMetersToPaint / 5

	countQtyCans, err := calculatePaintCans(litersToPaint)

	totalQtyCans := domain.Can{
		QtyCan18l:  null.IntFrom(countQtyCans.QtyCan18l.Int64),
		QtyCan3_6l: null.IntFrom(countQtyCans.QtyCan3_6l.Int64),
		QtyCan2_5l: null.IntFrom(countQtyCans.QtyCan2_5l.Int64),
		QtyCan0_5l: null.IntFrom(countQtyCans.QtyCan0_5l.Int64),
		Liters:     null.FloatFrom(litersToPaint),
	}

	return totalQtyCans, nil
}

func (pService *paintService) calculateAreaRoom(walls []domain.Wall) (float64, error) {
	var totalAreaWalls float64
	var totalAreaDoorsAndWindows float64

	wallItems, err := pService.paintRepository.GetWallItems()
	if err != nil {
		logrus.Error(err)
		return 0, domain.NewError("failure to get wall items.")
	}

	for i := range walls {
		err := pService.ruleNoWallCanBeLessThan1SquareMeterOrMoreThan15(walls[i])
		if err != nil {
			return 0, err
		}

		err = pService.ruleTheHeightOfWallsWithADoorMustBeAtLeast30CentimetersGreaterThanHeightOfTheDoor(walls[i])
		if err != nil {
			return 0, err
		}

		err = pService.ruleTheTotalAreaOfDoorsAndWindowsMustBeAMaximumOf50PercentOfTheWallArea(walls[i], wallItems)
		if err != nil {
			return 0, err
		}

		totalAreaWalls += (walls[i].Height.Float64 * walls[i].Width.Float64)

		for i := range wallItems {
			if wallItems[i].Name.String == domain.Door {
				totalAreaDoorsAndWindows += (wallItems[i].SquareMeters.Float64 * float64(walls[i].QtyDoors.Int64))
			}
			if wallItems[i].Name.String == domain.Window {
				totalAreaDoorsAndWindows += (wallItems[i].SquareMeters.Float64 * float64(walls[i].QtyWindows.Int64))
			}
		}
	}

	squareMetersToPaint := totalAreaWalls - totalAreaDoorsAndWindows

	return squareMetersToPaint, nil
}

func calculatePaintCans(litersToPaint float64) (domain.Can, error) {
	count := litersToPaint
	countQtyCans := domain.Can{}

	for count > 0.0 {
		for count >= domain.QtyTin18L {
			count = count - domain.QtyTin18L
			countQtyCans.QtyCan18l.Int64 += int64(1)
		}
		for count >= domain.QtyTin3_6L {
			count = count - domain.QtyTin3_6L
			countQtyCans.QtyCan3_6l.Int64 += int64(1)
		}
		for count >= domain.QtyTin2_5L {
			count = count - domain.QtyTin2_5L
			countQtyCans.QtyCan2_5l.Int64 += int64(1)
		}
		count = count - domain.QtyTin0_5L
		countQtyCans.QtyCan0_5l.Int64 += int64(1)
	}

	return countQtyCans, nil
}
