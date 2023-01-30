package demo

import (
	"github.com/shagnyin/gin-scaffold/internal/model"
	"github.com/shagnyin/gin-scaffold/internal/svc"
)

type Demo struct {
	svcCtx *svc.ServiceContext
	db     *model.Queries
}

func NewDemo(svcCtx *svc.ServiceContext) *Demo {
	return &Demo{svcCtx: svcCtx, db: svcCtx.DB}
}
