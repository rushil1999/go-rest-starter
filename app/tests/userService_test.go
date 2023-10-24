package tests

import (
	"testing"
	"go-rest-starter/pkg/models"
	"go-rest-starter/app/services"
)


func TestGetUserName(t *testing.T) {
	actualUserCount := len(models.UserData)

	testingCount := len(services.GetUsers())
	if testingCount != actualUserCount {
		t.Fatalf("test case failed")
	}

}