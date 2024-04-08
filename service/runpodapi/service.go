package runpodapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"git.zx-tech.net/app/ait-go-app/config"
	"go.uber.org/fx"
	"io"
	"net/http"
	"strings"
)

var Module = fx.Options(
	fx.Provide(apiClient),
)

type ApiClientI interface {
	GenerateImage(ctx context.Context, req GenerateReq, modelName ModelName, Async bool) ([]byte, error)
	CheckStatus(ctx context.Context, jobId string, modelName ModelName) ([]byte, error)
}

func apiClient(c *config.Config) ApiClientI {
	return &ApiClient{
		BaseUrl: c.RunPodApi.BaseUrl,
		Token:   c.RunPodApi.ApiKey,
	}
}

type ApiClient struct {
	BaseUrl string `json:"baseUrl"`
	Token   string `json:"token"`
}

func (a ApiClient) CheckStatus(ctx context.Context, jobId string, modelName ModelName) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/status/%s", a.BaseUrl, modelName, jobId)
	return a.doReq(GenerateReq{}, url)
}

func (a ApiClient) GenerateImage(ctx context.Context, req GenerateReq, modelName ModelName, Async bool) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/runsync", a.BaseUrl, modelName)
	if Async {
		url = strings.ReplaceAll(url, "runsync", "run")
	}
	return a.doReq(req, url)
}

func (a ApiClient) doReq(req GenerateReq, url string) ([]byte, error) {
	var (
		res []byte
	)
	body, err := json.Marshal(req)
	if err != nil {
		return res, err
	}
	r, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return res, err
	}
	r.Header.Add("accept", "application/json")
	r.Header.Add("content-type", "application/json")
	r.Header.Add("authorization", a.Token)

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()

	res, err = io.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	return res, err
}
