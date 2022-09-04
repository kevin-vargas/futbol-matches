package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func decode[T any](r *http.Request) (*T, error) {
	elem := new(T)
	var unmarshalErr *json.UnmarshalTypeError
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(elem)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			return nil, errors.New("wrong Type provided for field " + unmarshalErr.Field)
		}
		return nil, err
	}
	v := validator.New()
	err = v.Struct(elem)
	if err != nil {
		return nil, err
	}
	return elem, nil
}
