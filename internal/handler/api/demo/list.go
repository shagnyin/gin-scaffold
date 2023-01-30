package demo

import (
	"github.com/gin-gonic/gin"
	"github.com/shagnyin/gin-scaffold/internal/model"
	"github.com/shagnyin/gin-scaffold/pkg/db"
	"github.com/shagnyin/gin-scaffold/pkg/logger"
	"github.com/shagnyin/gin-scaffold/pkg/request"
	"net/http"
)

type (
	demoListReq struct {
		PageSize uint64 `form:"page_size" binding:"required"` // 页大小
		Page     uint64 `form:"page" binding:"required"`      // 页码
	}
	demoListItem struct {
		ID     string `json:"id"`     // id
		Email  string `json:"email"`  // 邮箱
		Name   string `json:"name"`   // 姓名
		Avatar string `json:"avatar"` // 头像
	}
	demoListRes struct {
		Total int64          `json:"total"` // 总数
		List  []demoListItem `json:"list"`  // 数据列表
	}
)

// List 查询数据
// @Summary		list
// @Description get demo data with page
// @Accept json
// @Produce json
// @Tags demo
// @Param payload body demoListReq true "post params"
// @Success 200 {object} request.Resp{data=demoListRes{list=[]demoListItem}} "request content"
// @Router /api/v1/demo [get]
func (ctl Demo) List(ctx *gin.Context) {
	// validate params
	var req demoListReq
	if err := ctx.BindQuery(&req); err != nil {
		logger.WithContext(ctx).Error(err)
		request.Fail(ctx, http.StatusBadRequest, 1, err.Error())
		return
	}

	// fetch data
	list := make([]demoListItem, 0, req.PageSize)
	
	// get count
	total, err := ctl.db.CountDemoByPage(ctx.Request.Context())
	if err != nil {
		logger.WithContext(ctx).Error(err)
		request.Fail(ctx, http.StatusBadRequest, 1, err.Error())
		return
	}
	if total <= 0 {
		request.Success(ctx, demoListRes{Total: 0, List: list})
		return
	}

	// get limit
	offset, limit := db.Paginate(req.Page, req.PageSize)
	demoList, err := ctl.db.ListDemoByPage(ctx.Request.Context(), &model.ListDemoByPageParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		logger.WithContext(ctx).Error(err)
		request.Fail(ctx, http.StatusBadRequest, 1, err.Error())
		return
	}
	for _, demo := range demoList {
		list = append(list, demoListItem{
			ID:     demo.ID,
			Email:  demo.Email,
			Name:   demo.Name,
			Avatar: demo.Avatar,
		})
	}

	request.Success(ctx, demoListRes{
		Total: total,
		List:  list,
	})
}
