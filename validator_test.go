package validator_test

import (
	"testing"
	"time"

	"github.com/oSethoum/validator"
)

type BaseModel struct {
	ID        string     `json:"id,omitempty" gorm:"primaryKey"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type User struct {
	BaseModel
	Name   string `json:"name,omitempty" gorm:"unique;not null" validate:"minLen=5;alpha"`
	Role   *Role  `json:"role,omitempty" gormy:"edge=one"`
	RoleID string `json:"roleId,omitempty"`
}

type Role struct {
	BaseModel
	Name          string            `json:"name,omitempty" gorm:"unique;not null" validate:"minLen=5;alpha"`
	DeniedActions string            `json:"deniedActions,omitempty" gorm:"unique"`
	DeniedFields  map[string]string `json:"deniedFields,omitempty" gorm:"serializer:json"`
	Users         []User            `json:"users,omitempty" gormy:"edge=many"`
}

func TestStruct(T *testing.T) {
	err := validator.Struct(
		&User{
			Name: "jue4",
			Role: &Role{
				Name: "admin",
			},
		},
	)
	if err != nil {
		T.Error(err)
	}
}
