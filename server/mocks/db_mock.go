package models

import "github.com/FuturICT2/fin4-core/server/datatype"

// DBMock DB mock
type DBMock struct {
	Name string
}

// NewUserModel mock
func (db *DBMock) NewUserModel() datatype.UserStore {
	return &UserModelMock{}
}

// NewTradeModel mock
func (db *DBMock) NewTradeModel() datatype.TradeStore {
	return &TradeModelMock{}
}

// NewDBMock mock
func NewDBMock() *DBMock {
	return &DBMock{Name: "Mock"}
}
