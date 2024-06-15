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
	"github.com/tencentad/marketing-api-go-sdk/pkg/api/v3"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config/v3"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
	"github.com/tencentad/marketing-api-go-sdk/pkg/model/v3"
)

type CreativeTemplateGetExample struct {
	TAds                    *ads.SDKClient
	AccessToken             string
	AccountId               int64
	MarketingGoal           string
	MarketingTargetType     string
	MarketingCarrierType    string
	DeliveryMode            string
	CreativeTemplateGetOpts *api.CreativeTemplateGetOpts
}

func (e *CreativeTemplateGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = 789
	e.MarketingGoal = "marketingGoal_example"
	e.MarketingTargetType = "marketingTargetType_example"
	e.MarketingCarrierType = "marketingCarrierType_example"
	e.DeliveryMode = "deliveryMode_example"
	e.CreativeTemplateGetOpts = &api.CreativeTemplateGetOpts{}
}

func (e *CreativeTemplateGetExample) RunExample() (model.CreativeTemplateGetResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.CreativeTemplate().Get(ctx, e.AccountId, e.MarketingGoal, e.MarketingTargetType, e.MarketingCarrierType, e.DeliveryMode, e.CreativeTemplateGetOpts)
}

func main() {
	e := &CreativeTemplateGetExample{}
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
