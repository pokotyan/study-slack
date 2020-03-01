package context

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
)

type CtxKey string

const (
	KeyTx CtxKey = "tx"
	KeyDB CtxKey = "db"
)

func GetDB(ctx context.Context) (*gorm.DB, error) {
	return getDB(ctx, KeyDB)
}

func GetTx(ctx context.Context) (*gorm.DB, error) {
	return getDB(ctx, KeyTx)
}

func getDB(ctx context.Context, key CtxKey) (*gorm.DB, error) {
	// TODO: use key
	v := ctx.Value("tx")
	if v == nil {

		return nil, fmt.Errorf("err %s", "no data in the context")
	}
	db, ok := v.(*gorm.DB)
	if !ok {
		return nil, fmt.Errorf("err %s", "invalid data in the context")
	}
	return db, nil
}
