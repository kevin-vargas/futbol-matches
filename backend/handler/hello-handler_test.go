package handler

import (
	"backend/service/mocks"
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"
)

func Test_Hello(t *testing.T) {
	testCases := []struct {
		description      string
		expectStatusCode int
		errService       error
	}{
		{
			description:      "With Service Error",
			expectStatusCode: http.StatusInternalServerError,
			errService:       errors.New("Error On Service"),
		},
		{
			description:      "With Out Service Error",
			expectStatusCode: http.StatusNoContent,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			// arrange
			m := &mocks.Service{}
			m.On("Say", mock.Anything, mock.Anything).Return(tt.errService)
			h := New(m)
			req, res := getReqRes(http.MethodGet, nil, "/test", http.Header{})
			req = reqWithPathParam(req, "word", "test")

			// act
			h.Handle(res, req)

			// assert

			assert.Equal(t, tt.expectStatusCode, res.Result().StatusCode)
		})
	}

}

var getReqRes = func(method string, body []byte, url string, headers http.Header) (*http.Request, *httptest.ResponseRecorder) {
	reader := bytes.NewReader(body)
	req := httptest.NewRequest(method, url, reader)
	req.Header = headers
	res := httptest.NewRecorder()
	return req, res
}

func reqWithPathParam(r *http.Request, key, value string) *http.Request {
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add(key, value)
	newCtx := context.WithValue(r.Context(), chi.RouteCtxKey, rctx)
	return r.WithContext(newCtx)
}
