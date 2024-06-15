/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tencentad/marketing-api-go-sdk/pkg/ads"
	"github.com/tencentad/marketing-api-go-sdk/pkg/api"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
)

type WxPackageAccountUpdateExample struct {
	TAds                       *ads.SDKClient
	AccessToken                string
	AccountId                  int64
	WechatId                   int64
	WxPackageAccountUpdateOpts *api.WxPackageAccountUpdateOpts
}

func (e *WxPackageAccountUpdateExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = 789
	e.WechatId = 789
	e.WxPackageAccountUpdateOpts = &api.WxPackageAccountUpdateOpts{}
}

func (e *WxPackageAccountUpdateExample) RunExample() (interface{}, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.WxPackageAccount().Update(ctx, e.AccountId, e.WechatId, e.WxPackageAccountUpdateOpts)
}

func main() {
	e := &WxPackageAccountUpdateExample{}
	e.Init()
	response, headers, err := e.RunExample()
	if err != nil {
		if resErr, ok := err.(errors.ResponseError); ok {
			errStr, _ := json.Marshal(resErr)
			fmt.Println("Response error:", string(errStr))
		} else {
			fmt.Println("Error:", err)
		}
	}
	fmt.Println("Response data:", response)
	fmt.Println("Headers:", headers)
}
