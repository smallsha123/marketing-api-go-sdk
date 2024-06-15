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

	"github.com/antihax/optional"
	"github.com/tencentad/marketing-api-go-sdk/pkg/ads"
	"github.com/tencentad/marketing-api-go-sdk/pkg/api"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
	"github.com/tencentad/marketing-api-go-sdk/pkg/model"
)

type PlayablePagesGetExample struct {
	TAds                 *ads.SDKClient
	AccessToken          string
	AccountId            int64
	PlayablePagesGetOpts *api.PlayablePagesGetOpts
}

func (e *PlayablePagesGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = int64(0)
	e.PlayablePagesGetOpts = &api.PlayablePagesGetOpts{

		Fields: optional.NewInterface([]string{"playable_page_material_id", "playable_page_name", "playable_page_cdn_base_url", "audit_status", "created_time", "last_modified_time"}),
	}
}

func (e *PlayablePagesGetExample) RunExample() (model.PlayablePagesGetResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.PlayablePages().Get(ctx, e.AccountId, e.PlayablePagesGetOpts)
}

func main() {
	e := &PlayablePagesGetExample{}
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
