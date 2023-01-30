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
	demoUpdateReq struct {
		Name  string `json:"name" binding:"required"`  // 姓名
		Email string `json:"email" binding:"required"` // 邮箱
	}
	demoUpdateRes struct {
		ID string `json:"id"` // insert id
	}
)

// Update 修改数据
// @Summary		update
// @Description update demo data by id
// @Accept json
// @Produce json
// @Tags demo
// @Param id path string true "ID"
// @Param payload body demoUpdateReq true "post params"
// @Success 200 {object} request.Resp{data=demoUpdateRes} "request content"
// @Router /api/v1/demo/{id} [put]
func (ctl Demo) Update(ctx *gin.Context) {
	var req demoUpdateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.WithContext(ctx).Error(err)
		request.Fail(ctx, http.StatusBadRequest, 1, err.Error())
		return
	}

	id := ctx.Param("id")

	// check data
	_, err := ctl.db.CheckDemoById(ctx.Request.Context(), &model.CheckDemoByIdParams{
		ID:    id,
		Email: strings.TrimSpace(req.Email),
	})
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
	err = ctl.db.UpdateDemoByID(ctx.Request.Context(), &model.UpdateDemoByIDParams{
		ID:     demoID,
		Email:  req.Email,
		Avatar: "",
	})
	if err != nil {
		logger.WithContext(ctx).Error(err)
		request.Fail(ctx, http.StatusBadRequest, 1, err.Error())
		return
	}

	request.Success(ctx, demoUpdateRes{ID: demoID})
}
