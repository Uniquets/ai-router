package controller

import (
	"git.zx-tech.net/app/ait-go-app/application"
	"git.zx-tech.net/app/ait-go-app/controller/request_object"
	"git.zx-tech.net/app/ait-go-app/controller/view_object"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// GenImgController 生成图片
type GenImgController struct {
	app      application.Text2ImgAppI
	validate *validator.Validate
}

func genImg(app application.Text2ImgAppI, r *gin.Engine) *GenImgController {
	c := &GenImgController{
		app:      app,
		validate: validator.New(),
	}
	sdGroup := r.Group("/genimg")
	sdGroup.POST("/txt2img", c.TxtToImg)
	sdGroup.POST("/check", c.CheckStatus)
	return c
}

// TxtToImg 文生图
// @Tags      ai-router
// @Router    /genimg/txt2img [POST]
// @Summary   文生图
// @accept    application/json
// @Param     create body request_object.GenerateReq true "请求参数"
// @Produce   application/json
// @Success   200   {object} view_object.GenerateResp
func (g GenImgController) TxtToImg(c *gin.Context) {
	var (
		req request_object.GenerateReq
		err error
	)
	// 绑定参数
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, view_object.Resp(err))
		return
	}
	// 验证参数
	if err = g.validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, view_object.Resp(err))
		return
	}
	// 生成图片
	resp, err := g.app.GenImg(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, view_object.Resp(err))
		return
	}
	c.JSON(http.StatusOK, view_object.Resp(resp))
	return
}

// CheckStatus 获取状态
// @Tags      ai-router
// @Router    /genimg/check [POST]
// @Summary   获取状态
// @accept    application/json
// @Param     create body request_object.CheckStatusReq true "请求参数"
// @Produce   application/json
// @Success   200   {object} view_object.GenerateResp
func (g GenImgController) CheckStatus(c *gin.Context) {
	var (
		req request_object.CheckStatusReq
		err error
	)
	// 绑定参数
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, view_object.Resp(err))
		return
	}
	// 验证参数
	if err = g.validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, view_object.Resp(err))
		return
	}
	// 获取状态
	resp, err := g.app.CheckStatus(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, view_object.Resp(err))
		return
	}
	c.JSON(http.StatusOK, view_object.Resp(resp))
	return
}
