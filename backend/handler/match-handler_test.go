package handler

import (
	"backend/service"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
	testCases := []struct {
		description      string
		body             []byte
		expectStatusCode int
	}{
		{
			description:      "Invalid body",
			body:             []byte(`{`),
			expectStatusCode: http.StatusBadRequest,
		},
		{
			description: "Successful",
			body: []byte(`
			{
				"description": "descripcion",
				"date": "2021-02-18T21:54:42.123Z",
				"place": "utn",
				"format": 10,
				"maxPlayers": 26
			}
			`),
			expectStatusCode: http.StatusCreated,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			// arrange
			ms := service.NewMatchService()
			h := NewMatchHandler(ms)
			req, res := getReqRes(http.MethodPost, tt.body, "/", http.Header{})

			// act
			h.Create(res, req)

			// assert

			assert.Equal(t, tt.expectStatusCode, res.Result().StatusCode)
		})
	}
}

func Test_GetAll(t *testing.T) {
	testCases := []struct {
		description      string
		expectStatusCode int
	}{
		{
			description:      "Successful",
			expectStatusCode: http.StatusOK,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			// arrange
			ms := service.NewMatchService()
			h := NewMatchHandler(ms)
			req, res := getReqRes(http.MethodPost, nil, "/", http.Header{})

			// act
			h.GetAll(res, req)

			// assert

			assert.Equal(t, tt.expectStatusCode, res.Result().StatusCode)
		})
	}
}

func Test_Get(t *testing.T) {
	testCases := []struct {
		description      string
		expectStatusCode int
	}{
		{
			description:      "Invalid id",
			expectStatusCode: http.StatusNotFound,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			// arrange
			ms := service.NewMatchService()
			h := NewMatchHandler(ms)
			req, res := getReqRes(http.MethodGet, nil, "/", http.Header{})
			req = reqWithPathParam(req, "id", "test")

			// act
			h.Get(res, req)

			// assert

			assert.Equal(t, tt.expectStatusCode, res.Result().StatusCode)
		})
	}
}

func Test_Update(t *testing.T) {
	testCases := []struct {
		description      string
		body             []byte
		expectStatusCode int
	}{
		{
			description:      "Invalid body",
			body:             []byte(`{`),
			expectStatusCode: http.StatusBadRequest,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			// arrange
			ms := service.NewMatchService()
			h := NewMatchHandler(ms)
			req, res := getReqRes(http.MethodPut, tt.body, "/", http.Header{})
			req = reqWithPathParam(req, "id", "test")

			// act
			h.Update(res, req)

			// assert

			assert.Equal(t, tt.expectStatusCode, res.Result().StatusCode)
		})
	}
}

func Test_AddPlayer(t *testing.T) {
	testCases := []struct {
		description      string
		body             []byte
		expectStatusCode int
	}{
		{
			description:      "Invalid body",
			body:             []byte(`{`),
			expectStatusCode: http.StatusBadRequest,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			// arrange
			ms := service.NewMatchService()
			h := NewMatchHandler(ms)
			req, res := getReqRes(http.MethodPut, tt.body, "/", http.Header{})
			req = reqWithPathParam(req, "id", "test")

			// act
			h.AddPlayer(res, req)

			// assert

			assert.Equal(t, tt.expectStatusCode, res.Result().StatusCode)
		})
	}
}
