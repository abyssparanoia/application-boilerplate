package gluesqlboiler

import "context"

func MockTx() {
	RunTx = func(ctx context.Context, fn func(context.Context) error) error {
		return fn(ctx)
	}
}
