package data

import (
	"context"
	"fmt"
	"io"
	"os"
)

func readStdin(ctx context.Context, _ *Source, _ ...string) ([]byte, error) {
	stdin := stdinFromContext(ctx)

	b, err := io.ReadAll(stdin)
	if err != nil {
		return nil, fmt.Errorf("can't read %s: %w", stdin, err)
	}
	return b, nil
}

type stdinCtxKey struct{}

func ContextWithStdin(ctx context.Context, r io.Reader) context.Context {
	return context.WithValue(ctx, stdinCtxKey{}, r)
}

func stdinFromContext(ctx context.Context) io.Reader {
	if r, ok := ctx.Value(stdinCtxKey{}).(io.Reader); ok {
		return r
	}

	return os.Stdin
}
