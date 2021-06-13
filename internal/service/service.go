package service

import (
	"context"

	"github.com/tianmai777/blog/global"
	"github.com/tianmai777/blog/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	return Service{
		ctx: ctx,
		dao: dao.New(global.DBEngine),
	}
}
