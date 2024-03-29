package databases

import "gorm.io/gorm"

type mockDatabase struct{}

func NewMockDatabase() Database {
	return &mockDatabase{}
}

func (db *mockDatabase) Connect() *gorm.DB {
	return nil
}
