package paint

import (
	"errors"

	"github.com/Chicocs/challenge-fullStack-digitalRepublic/internal/domain"
	"github.com/dariubs/percent"
)

func (pService *paintService) ruleNoWallCanBeLessThan1SquareMeterOrMoreThan15(wall domain.Wall) error {
	squareMeters := wall.Height.Float64 * wall.Width.Float64
	if squareMeters < 1 || squareMeters > 15 {
		return errors.New("No wall can be less than 1 square meter or more than 15.")
	}

	return nil
}

func (pService *paintService) ruleTheHeightOfWallsWithADoorMustBeAtLeast30CentimetersGreaterThanHeightOfTheDoor(wall domain.Wall) error {

	if wall.QtyDoors.Int64 >= 1 && wall.Width.Float64 < 2.2 {
		return errors.New("The height of walls with a door must be at least 30 centimeters greater than the height of the door.")
	}

	return nil
}

func (pService *paintService) ruleTheTotalAreaOfDoorsAndWindowsMustBeAMaximumOf50PercentOfTheWallArea(wall domain.Wall, wallItems []domain.WallItems) error {
	var totalAreaDoorsWindows float64
	var totalAreaDoors float64
	var totalAreaWindows float64

	totalAreaWall := wall.Height.Float64 * wall.Width.Float64

	for i := range wallItems {
		if wallItems[i].Name.String == domain.Door {
			totalAreaDoors = wallItems[i].SquareMeters.Float64 * float64(wall.QtyDoors.Int64)
		}
		if wallItems[i].Name.String == domain.Window {
			totalAreaWindows = wallItems[i].SquareMeters.Float64 * float64(wall.QtyWindows.Int64)
		}
	}

	totalAreaDoorsWindows = totalAreaDoors + totalAreaWindows

	result := percent.PercentOfFloat(totalAreaDoorsWindows, totalAreaWall)

	if result > 50 {
		return errors.New("The total area of doors and windows must be a maximum of 50% of the wall area.")
	}

	return nil
}
