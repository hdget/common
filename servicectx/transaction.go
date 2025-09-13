package servicectx

import (
	"context"
)

// ctxKeyTx servicectx key: db transaction
type ctxKeyTx struct{}

func AddTx(ctx context.Context, value any) context.Context {
	return context.WithValue(ctx, ctxKeyTx{}, value)
}

func GetTx(ctx context.Context) any {
	return ctx.Value(ctxKeyTx{})
}
