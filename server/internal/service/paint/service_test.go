package paint

import (
	"context"
	"testing"

	"github.com/Chicocs/challenge-fullStack-digitalRepublic/internal/domain"
	"github.com/Chicocs/challenge-fullStack-digitalRepublic/internal/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
)

func TestCalculateAmountOfPaintCansNeededCaseSuccess(t *testing.T) {
	var ctx context.Context
	paintMock := &mocks.PaintMock{}
	service := &paintService{
		paintRepository: paintMock,
	}

	paintMock.On("GetWallItems").Return([]domain.WallItems{
		{
			Name:         null.StringFrom("Window"),
			Height:       null.FloatFrom(float64(1.2)),
			Width:        null.FloatFrom(float64(2)),
			SquareMeters: null.FloatFrom(float64(2.4)),
		},
		{
			Name:         null.StringFrom("Door"),
			Height:       null.FloatFrom(float64(0.8)),
			Width:        null.FloatFrom(float64(1.9)),
			SquareMeters: null.FloatFrom(float64(2.7)),
		},
	}, nil)

	res, err := service.CalculateAmountOfPaintCansNeeded(ctx, []domain.Wall{
		{
			Height:     null.FloatFrom(float64(5)),
			Width:      null.FloatFrom(float64(2.5)),
			QtyDoors:   null.IntFrom(int64(1)),
			QtyWindows: null.IntFrom(int64(1)),
		},
		{
			Height:     null.FloatFrom(float64(5)),
			Width:      null.FloatFrom(float64(2.5)),
			QtyDoors:   null.IntFrom(int64(1)),
			QtyWindows: null.IntFrom(int64(1)),
		},
		{
			Height:     null.FloatFrom(float64(5)),
			Width:      null.FloatFrom(float64(2.5)),
			QtyDoors:   null.IntFrom(int64(1)),
			QtyWindows: null.IntFrom(int64(1)),
		},
		{
			Height:     null.FloatFrom(float64(5)),
			Width:      null.FloatFrom(float64(2.5)),
			QtyDoors:   null.IntFrom(int64(1)),
			QtyWindows: null.IntFrom(int64(1)),
		},
	})

	assert.Nil(t, err)
	assert.Equal(t, res, domain.Can{
		QtyCan0_5l: null.IntFrom(5),
		QtyCan2_5l: null.IntFrom(0),
		QtyCan3_6l: null.IntFrom(1),
		QtyCan18l:  null.IntFrom(0),
		Liters:     null.FloatFrom(5.92),
	})
	assert.NotNil(t, res)
	paintMock.AssertExpectations(t)
}
