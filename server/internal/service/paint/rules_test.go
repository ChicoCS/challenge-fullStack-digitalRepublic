package paint

import (
	"testing"

	"github.com/Chicocs/challenge-fullStack-digitalRepublic/internal/domain"
	"github.com/Chicocs/challenge-fullStack-digitalRepublic/internal/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v4"
)

func TestRuleNoWallCanBeLessThan1SquareMeterOrMoreThan15CaseFailure(t *testing.T) {
	paintMock := &mocks.PaintMock{}
	rule := &paintService{
		paintRepository: paintMock,
	}

	err := rule.ruleNoWallCanBeLessThan1SquareMeterOrMoreThan15(domain.Wall{
		Height: null.FloatFrom(float64(0)),
		Width:  null.FloatFrom(float64(1.5)),
	})

	assert.NotNil(t, err)
	assert.EqualError(t, err, "No wall can be less than 1 square meter or more than 15.")
}

func TestRuleTheHeightOfWallsWithADoorMustBeAtLeast30CentimetersGreaterThanHeightOfTheDoorCaseFailure(t *testing.T) {
	paintMock := &mocks.PaintMock{}
	rule := &paintService{
		paintRepository: paintMock,
	}

	err := rule.ruleTheHeightOfWallsWithADoorMustBeAtLeast30CentimetersGreaterThanHeightOfTheDoor(domain.Wall{
		Height:     null.FloatFrom(float64(5)),
		Width:      null.FloatFrom(float64(2.19)),
		QtyDoors:   null.IntFrom(int64(1)),
		QtyWindows: null.IntFrom(int64(1)),
	})

	assert.NotNil(t, err)
	assert.EqualError(t, err, "The height of walls with a door must be at least 30 centimeters greater than the height of the door.")
}

func TestRuleTheTotalAreaOfDoorsAndWindowsMustBeAMaximumOf50PercentOfTheWallAreaCaseFailure(t *testing.T) {
	paintMock := &mocks.PaintMock{}
	rule := &paintService{
		paintRepository: paintMock,
	}

	wall := domain.Wall{
		Height:     null.FloatFrom(float64(5)),
		Width:      null.FloatFrom(float64(2.5)),
		QtyDoors:   null.IntFrom(int64(5)),
		QtyWindows: null.IntFrom(int64(5)),
	}

	wallItems := []domain.WallItems{
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
	}

	err := rule.ruleTheTotalAreaOfDoorsAndWindowsMustBeAMaximumOf50PercentOfTheWallArea(wall, wallItems)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "The total area of doors and windows must be a maximum of 50% of the wall area.")
}
