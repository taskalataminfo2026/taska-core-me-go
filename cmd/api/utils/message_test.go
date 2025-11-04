package utils_test

import (
	"net/http"
	"taska-core-me-go/cmd/api/models"
	"taska-core-me-go/cmd/api/utils"
	"testing"
)

func TestNewMessage_WithStringData(t *testing.T) {
	data := "hola mundo"
	expected := models.Message{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    data,
	}

	result := utils.NewMessage(data)

	if result.Code != expected.Code {
		t.Errorf("expected code %d, got %d", expected.Code, result.Code)
	}
	if result.Message != expected.Message {
		t.Errorf("expected message %s, got %s", expected.Message, result.Message)
	}
	if result.Data != expected.Data {
		t.Errorf("expected data %v, got %v", expected.Data, result.Data)
	}
}

func TestNewMessage_WithStructData(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	data := User{Name: "Valencia", Age: 30}
	result := utils.NewMessage(data)

	msg, ok := result.Data.(User)
	if !ok {
		t.Fatalf("expected data to be of type User")
	}

	if msg.Name != "Valencia" || msg.Age != 30 {
		t.Errorf("expected user {Valencia, 30}, got %+v", msg)
	}
}
