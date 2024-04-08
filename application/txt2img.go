package application

import (
	"context"
	"encoding/json"
	"git.zx-tech.net/app/ait-go-app/controller/request_object"
	"git.zx-tech.net/app/ait-go-app/controller/view_object"
	"git.zx-tech.net/app/ait-go-app/service/runpodapi"
)

type Text2ImgAppI interface {
	GenImg(ctx context.Context, req request_object.GenerateReq) (view_object.GenerateResp, error)
	CheckStatus(ctx context.Context, req request_object.CheckStatusReq) (view_object.GenerateResp, error)
}

func txt2img(c runpodapi.ApiClientI) Text2ImgAppI {
	return &Text2ImgApp{
		ApiClient: c,
	}
}

type Text2ImgApp struct {
	ApiClient runpodapi.ApiClientI
}

func (a Text2ImgApp) CheckStatus(ctx context.Context, req request_object.CheckStatusReq) (view_object.GenerateResp, error) {
	var (
		resp view_object.GenerateResp
	)
	data, err := a.ApiClient.CheckStatus(ctx, req.JobId, runpodapi.SDXL)
	if err != nil {
		return view_object.GenerateResp{}, err
	}
	err = json.Unmarshal(data, &resp)
	return resp, err
}

func (a Text2ImgApp) GenImg(ctx context.Context, req request_object.GenerateReq) (view_object.GenerateResp, error) {
	var (
		resp view_object.GenerateResp
	)
	data, err := a.ApiClient.GenerateImage(ctx,
		runpodapi.GenerateReq{
			Input: map[runpodapi.InputColumn]interface{}{
				runpodapi.InputColumns.Prompt:   req.Input.Prompt,
				runpodapi.InputColumns.Width:    req.Input.Width,
				runpodapi.InputColumns.Height:   req.Input.Height,
				runpodapi.InputColumns.ImageUrl: req.Input.ImageUrl,
				runpodapi.InputColumns.Seed:     req.Input.Seed,
			},
			WebHook: "",
		}, "sdxl", req.Async)

	if err != nil {
		return view_object.GenerateResp{}, err
	}
	err = json.Unmarshal(data, &resp)
	return resp, err
}
