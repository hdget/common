package meta

import "context"

type ctxKeyMeta struct{}

func AddMetaToContext(ctx context.Context, metas map[string]string) context.Context {
	return context.WithValue(ctx, ctxKeyMeta{}, metas)
}

func GetMetaFromContext(ctx context.Context) map[string]string {
	if kvs, ok := ctx.Value(ctxKeyMeta{}).(map[string]string); ok {
		return kvs
	}
	return nil
}
