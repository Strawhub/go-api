package model

import (
	"github.com/stretchr/testify/mock"
)

type MockModel struct {
	mock.Mock
}

func (m *MockModel) GetAll() []Question {
	args := m.Called()
	return args.Get(0).([]Question)
}

func (m *MockModel) GetBy(tag, num string) []Question {
	args := m.Called(tag, num)
	return args.Get(0).([]Question)
}
