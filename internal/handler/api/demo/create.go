package demo

import (
	"github.com/gin-gonic/gin"
	"github.com/shagnyin/gin-scaffold/internal/model"
	"github.com/shagnyin/gin-scaffold/pkg/db"
	"github.com/shagnyin/gin-scaffold/pkg/logger"
	"github.com/shagnyin/gin-scaffold/pkg/request"
	"github.com/shagnyin/gin-scaffold/pkg/utils"
	"net/http"
	"strings"
)

// params defs
type (
	demoCreateReq struct {
		Name  string `json:"name" binding:"required"`  // 姓名
		Email string `json:"email" binding:"required"` // 邮箱
	}
	demoCreateRes struct {
		ID string `json:"id"` // insert id
	}
)

// Create 创建数据
// @Summary		create
// @Description create demo data
// @Accept json
// @Produce json
// @Tags demo
// @Param payload body demoCreateReq true "post params"
// @Success 200 {object} request.Resp{data=demoCreateRes} "request content"
// @Router /api/v1/demo [post]
func (ctl Demo) Create(ctx *gin.Context) {
	var req demoCreateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.WithContext(ctx).Error(err)
		request.Fail(ctx, http.StatusBadRequest, 1, err.Error())
		return
	}

	// check data
	_, err := ctl.db.GetDemoByEmail(ctx.Request.Context(), strings.TrimSpace(req.Email))
	exists, err := db.CheckError(err)
	if err != nil {
		logger.WithContext(ctx).Error(err)
		request.Fail(ctx, http.StatusBadRequest, 1, err.Error())
		return
	}
	if exists {
		request.Fail(ctx, http.StatusBadRequest, 1, "data already exists")
		return
	}

	// create data
	demoID := utils.UUID()
	err = ctl.db.CreateDemo(ctx.Request.Context(), &model.CreateDemoParams{
		ID:     demoID,
		Email:  req.Email,
		Name:   req.Name,
		Avatar: "",
	})
	if err != nil {
		logger.WithContext(ctx).Error(err)
		request.Fail(ctx, http.StatusBadRequest, 1, err.Error())
		return
	}

	request.Success(ctx, demoCreateRes{ID: demoID})
}
