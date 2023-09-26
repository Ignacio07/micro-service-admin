package test

import (
	"testing"

	"github.com/Ignacio07/micro-service-admin/config"
	"github.com/Ignacio07/micro-service-admin/models"
)

func TestCreateUser(t *testing.T) {
	var reqCreateUser = models.User{
		Id:       "20102",
		Store:    "202313",
		Password: "12345",
	}
	result := config.DB.Create(&reqCreateUser)
	if result != nil {
		t.Error(result.Error)
	}
	t.Log("Insertado: ", result)
}
