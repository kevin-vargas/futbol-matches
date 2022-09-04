package handler

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Book struct {
	Author string `json:"author" validate:"min=4"`
}

func Test_Decode(t *testing.T) {
	testCases := []struct {
		description string
		body        []byte
		errorMsg    string
		expect      *Book
	}{
		{
			description: "With invalid Body",
			body: []byte(`
			{
				"author": 1,
		  	}
			`),
			errorMsg: "invalid character '}' looking for beginning of object key string",
		},
		{
			description: "With invalid Field",
			body: []byte(`
			{
				"author": 1
		  	}
			`),
			errorMsg: "wrong Type provided for field author",
		},
		{
			description: "With unknown Field",
			body: []byte(`
			{
				"author": "test",
				"foo":	  "bar"
		  	}
			`),
			errorMsg: "json: unknown field \"foo\"",
		},
		{
			description: "Invalid validator",
			body: []byte(`
			{
				"author": "t"
		  	}
			`),
			errorMsg: "Key: 'Book.Author' Error:Field validation for 'Author' failed on the 'min' tag",
		},
		{
			description: "Valid",
			body: []byte(`
			{
				"author": "test"
		  	}
			`),
			expect: &Book{
				Author: "test",
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			// arrange
			req, _ := getReqRes(http.MethodGet, tt.body, "/", http.Header{})

			// act
			actual, err := decode[Book](req)

			// assertion
			if tt.errorMsg != "" {
				assert.EqualError(t, err, tt.errorMsg)
			} else {
				assert.Equal(t, tt.expect, actual)
			}
		})
	}
}
