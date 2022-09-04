package service

import (
	"context"
	"fmt"
)

type Hello struct {
}

func (h Hello) Say(ctx context.Context, s string) error {
	fmt.Printf("Say %s \n", s)
	return nil
}

func New() Service {
	return Hello{}
}
