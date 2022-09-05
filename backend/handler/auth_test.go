package handler

import (
	"backend/service/mocks"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var validUser = []byte(`
{
    "username": "test_user",
    "password": "test_user"
}
`)

func Test_Auth_SingUp(t *testing.T) {
	testcases := []struct {
		description      string
		body             []byte
		serviceSignUpErr error
		serviceSignUpOk  bool
		expectStatus     int
	}{
		{
			description:  "Invalid User",
			body:         []byte(`{`),
			expectStatus: http.StatusBadRequest,
		},
		{
			description:      "Service signup with error",
			body:             validUser,
			serviceSignUpErr: errors.New("error on service"),
			expectStatus:     http.StatusInternalServerError,
		},
		{
			description:     "Service signup with user exist",
			body:            validUser,
			serviceSignUpOk: false,
			expectStatus:    http.StatusConflict,
		},
		{
			description:     "Valid signup",
			body:            validUser,
			serviceSignUpOk: true,
			expectStatus:    http.StatusCreated,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.description, func(t *testing.T) {
			// arrange
			s := &mocks.Auth{}
			s.On("SignUp", mock.Anything, mock.Anything).Return(tt.serviceSignUpOk, tt.serviceSignUpErr)
			h := NewAuth(s)
			req, res := getReqRes(http.MethodPost, tt.body, "/", http.Header{})

			// act
			h.SingUp(res, req)

			// arrange
			assert.Equal(t, tt.expectStatus, res.Result().StatusCode)
		})
	}
}

func Test_Auth_Login(t *testing.T) {
	testcases := []struct {
		description       string
		body              []byte
		serviceLoginErr   error
		serviceLoginToken string
		expectStatus      int
	}{
		{
			description:  "Invalid User",
			body:         []byte(`{`),
			expectStatus: http.StatusBadRequest,
		},
		{
			description:     "Service Login with error",
			body:            validUser,
			serviceLoginErr: errors.New("error on service"),
			expectStatus:    http.StatusInternalServerError,
		},
		{
			description:       "Service signup with user exist",
			body:              validUser,
			serviceLoginToken: "",
			expectStatus:      http.StatusBadRequest,
		},
		{
			description:       "Valid signup",
			body:              validUser,
			serviceLoginToken: "test",
			expectStatus:      http.StatusNoContent,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.description, func(t *testing.T) {
			// arrange
			s := &mocks.Auth{}
			s.On("Login", mock.Anything, mock.Anything, mock.Anything).Return(tt.serviceLoginToken, tt.serviceLoginErr)
			h := NewAuth(s)
			req, res := getReqRes(http.MethodPost, tt.body, "/", http.Header{})

			// act
			h.Login(res, req)

			// arrange
			assert.Equal(t, tt.expectStatus, res.Result().StatusCode)
		})
	}
}
