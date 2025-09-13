package servicectx

import (
	"context"
)

type ctxKeyMeta struct{}

func addMeta(ctx context.Context, metas map[string]string) context.Context {
	return context.WithValue(ctx, ctxKeyMeta{}, metas)
}

func GetMeta(ctx context.Context) map[string]string {
	kvs, ok := ctx.Value(ctxKeyMeta{}).(map[string]string)
	if ok {
		return kvs
	}
	return make(map[string]string)
}
