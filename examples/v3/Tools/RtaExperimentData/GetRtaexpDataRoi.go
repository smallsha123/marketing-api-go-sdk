/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tencentad/marketing-api-go-sdk/pkg/ads/v3"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config/v3"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
	"github.com/tencentad/marketing-api-go-sdk/pkg/model/v3"
)

type RtaexpDataRoiGetExample struct {
	TAds        *ads.SDKClient
	AccessToken string
	Data        model.RtaexpDataRoiGetRequest
}

func (e *RtaexpDataRoiGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.Data = model.RtaexpDataRoiGetRequest{}
}

func (e *RtaexpDataRoiGetExample) RunExample() (model.RtaexpDataRoiGetResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.RtaexpDataRoi().Get(ctx, e.Data)
}

func main() {
	e := &RtaexpDataRoiGetExample{}
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
