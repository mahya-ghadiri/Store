package mocks

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	"testing"
)


func NewDataProviderMock(t *testing.T) (*gomock.Controller, *MockDataProvider) {
	ctrl := gomock.NewController(t)
	mock := NewMockDataProvider(ctrl)
	return ctrl, mock
}

func JSON(input interface{}) string {
	body, err := json.Marshal(&input)
	if err != nil {
		panic(err)
	}
	return string(body)
}