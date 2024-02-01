package service

import (
	"context"

	"github.com/go-sonic/sonic/model/param"
)

type PasteService interface {
	Get(ctx context.Context, key, username, password string) (string, error)
	Create(ctx context.Context, paste param.Paste) (any, error)
	Update(ctx context.Context, paste param.Paste) error
}
