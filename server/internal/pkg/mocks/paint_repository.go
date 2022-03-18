package mocks

import (
	"github.com/Chicocs/challenge-fullStack-digitalRepublic/internal/domain"
	"github.com/stretchr/testify/mock"
)

type PaintMock struct {
	mock.Mock
}

func (m *PaintMock) GetWallItems() ([]domain.WallItems, error) {
	args := m.Called()
	return args.Get(0).([]domain.WallItems), args.Error(1)
}
