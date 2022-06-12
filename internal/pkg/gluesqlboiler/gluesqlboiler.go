package gluesqlboiler

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

var (
	ctxTxKey = struct{}{}
)

// GetContextExecutor :
func GetContextExecutor(ctx context.Context) boil.ContextExecutor {
	if tx, ok := ctx.Value(&ctxTxKey).(*sql.Tx); ok {
		return tx
	}
	return boil.GetContextDB()
}

// RunTx :
var RunTx = func(ctx context.Context, fn func(context.Context) error) error {
	db, ok := GetContextExecutor(ctx).(boil.ContextBeginner)
	if !ok {
		panic("The database in the context does not support boil.ContextBeginner")
	}
	return RunTxWithDB(ctx, db, fn)
}

// RunTxWithDB :
func RunTxWithDB(ctx context.Context, db boil.ContextBeginner, fn func(context.Context) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			panic(err)
		}
	}()

	ctxWithTx := context.WithValue(ctx, &ctxTxKey, tx)
	if err := fn(ctxWithTx); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
