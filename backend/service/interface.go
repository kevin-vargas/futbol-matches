package service

import (
	"context"
)

type Service interface {
	Say(context.Context, string) error
}
