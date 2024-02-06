package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	//  Тест первый: Запрос сформирован корректно, сервис возвращает код
	//ответа 200 и тело ответа не пустое.

	// здесь нужно создать запрос к сервису
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	//здесь нужно добавить необходимые проверки
	body := responseRecorder.Body
	statusCode := 200
	assert.Equal(t, statusCode, responseRecorder.Code)
	assert.NotEmpty(t, body)
}

func TestWrongCity(t *testing.T) {
	//  Тест второй: Город, который передаётся в параметре city, не поддерживается.
	// Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа.

	// здесь нужно создать запрос к сервису
	req := httptest.NewRequest("GET", "/cafe?count=4&city=Voronezh", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	//здесь нужно добавить необходимые проверки
	body := responseRecorder.Body.String()
	referenceAnswer := "wrong city value"
	fmt.Println(body)
	require.NotEmpty(t, body)
	assert.Equal(t, referenceAnswer, body)
}

func TestWenCountMoreValue(t *testing.T) {
	//  Тест третий:  Если в параметре count указано больше, чем есть всего,
	// должны вернуться все доступные кафе.
	totalCount := 4

	// здесь нужно создать запрос к сервису
	req := httptest.NewRequest("GET", "/cafe?count=48&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	//здесь нужно добавить необходимые проверки
	body := responseRecorder.Body.String()
	require.NotEmpty(t, body)
	bodyMass := strings.Split(body, ",")
	assert.Len(t, bodyMass, totalCount)
}

// bodyansw := responseRecorder.Body.String()
// body := strings.Split(bodyansw, ",")
