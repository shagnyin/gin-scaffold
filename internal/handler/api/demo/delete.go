package demo

import (
	"github.com/gin-gonic/gin"
	"github.com/shagnyin/gin-scaffold/internal/model"
	"github.com/shagnyin/gin-scaffold/pkg/db"
	"github.com/shagnyin/gin-scaffold/pkg/logger"
	"github.com/shagnyin/gin-scaffold/pkg/request"
	"github.com/shagnyin/gin-scaffold/pkg/utils"
	"net/http"
	"time"
)

// params defs
type (
	demoDeleteReq struct {
		ID string `uri:"id"` // insert id
	}
)

// Delete 删除数据
// @Summary		delete
// @Description delete demo data
// @Accept json
// @Produce json
// @Tags demo
// @Param id path string true "ID"
// @Success 200 {object} request.Resp{} "request content"
// @Router /api/v1/demo/{id} [delete]
func (ctl Demo) Delete(ctx *gin.Context) {
	var req demoDeleteReq
	if err := ctx.ShouldBindUri(&req); err != nil {
		logger.WithContext(ctx).Error(err)
		request.Fail(ctx, http.StatusBadRequest, 1, err.Error())
		return
	}

	// check data
	_, err := ctl.db.GetDemoByID(ctx.Request.Context(), req.ID)
	exists, err := db.CheckError(err)
	if err != nil {
		logger.WithContext(ctx).Error(err)
		request.Fail(ctx, http.StatusInternalServerError, 1, err.Error())
		return
	}
	if !exists {
		request.Fail(ctx, http.StatusNotFound, 1, "data not exists")
		return
	}

	// delete data
	if err = ctl.db.SoftDeleteDemoById(ctx.Request.Context(), &model.SoftDeleteDemoByIdParams{
		DeletedAt: time.Now().Format(utils.StandDateTime),
		ID:        req.ID,
	}); err != nil {
		logger.WithContext(ctx).Error(err)
		request.Fail(ctx, http.StatusInternalServerError, 1, err.Error())
		return
	}

	request.Success(ctx, nil)
}
