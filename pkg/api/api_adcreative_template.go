/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
	"github.com/tencentad/marketing-api-go-sdk/pkg/model"
	. "github.com/tencentad/marketing-api-go-sdk/pkg/model"
)

// Linger please
var (
	_ context.Context
)

type AdcreativeTemplateApiService service

/*
AdcreativeTemplateApiService 获取创意规格详情
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountId
 * @param promotedObjectType
 * @param optional nil or *AdcreativeTemplateGetOpts - Optional Parameters:
     * @param "SiteSet" (optional.Interface of []string) -
     * @param "AutomaticSiteEnabled" (optional.Bool) -
     * @param "IsDynamicCreative" (optional.Bool) -
     * @param "AdcreativeTemplateId" (optional.Int64) -
     * @param "DynamicCreativeType" (optional.String) -
     * @param "Fields" (optional.Interface of []string) -  返回参数的字段列表

@return AdcreativeTemplateGetResponse
*/

type AdcreativeTemplateGetOpts struct {
	SiteSet              optional.Interface
	AutomaticSiteEnabled optional.Bool
	IsDynamicCreative    optional.Bool
	AdcreativeTemplateId optional.Int64
	DynamicCreativeType  optional.String
	Fields               optional.Interface
}

func (a *AdcreativeTemplateApiService) Get(ctx context.Context, accountId int64, promotedObjectType string, localVarOptionals *AdcreativeTemplateGetOpts) (AdcreativeTemplateGetResponseData, http.Header, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarFileKey     string
		localVarReturnValue AdcreativeTemplateGetResponseData
		localVarResponse    AdcreativeTemplateGetResponse
	)

	// create path and map variables
	localVarPath := a.client.Cfg.BasePath + "/adcreative_template/get"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("account_id", parameterToString(accountId, ""))
	if localVarOptionals != nil && localVarOptionals.SiteSet.IsSet() {
		localVarQueryParams.Add("site_set", parameterToString(localVarOptionals.SiteSet.Value(), "multi"))
	}
	localVarQueryParams.Add("promoted_object_type", parameterToString(promotedObjectType, ""))
	if localVarOptionals != nil && localVarOptionals.AutomaticSiteEnabled.IsSet() {
		localVarQueryParams.Add("automatic_site_enabled", parameterToString(localVarOptionals.AutomaticSiteEnabled.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IsDynamicCreative.IsSet() {
		localVarQueryParams.Add("is_dynamic_creative", parameterToString(localVarOptionals.IsDynamicCreative.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AdcreativeTemplateId.IsSet() {
		localVarQueryParams.Add("adcreative_template_id", parameterToString(localVarOptionals.AdcreativeTemplateId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.DynamicCreativeType.IsSet() {
		localVarQueryParams.Add("dynamic_creative_type", parameterToString(localVarOptionals.DynamicCreativeType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Fields.IsSet() {
		localVarQueryParams.Add("fields", parameterToString(localVarOptionals.Fields.Value(), "multi"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"text/plain"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes, localVarFileKey)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	defer localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, nil, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarResponse, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			if *localVarResponse.Code != 0 {
				var localVarResponseErrors []model.ApiErrorStruct
				if localVarResponse.Errors != nil {
					localVarResponseErrors = *localVarResponse.Errors
				}
				err = errors.NewError(localVarResponse.Code, localVarResponse.Message, localVarResponse.MessageCn, localVarResponseErrors)
				return localVarReturnValue, localVarHttpResponse.Header, err
			}
			if localVarResponse.Data == nil {
				return localVarReturnValue, localVarHttpResponse.Header, err
			} else {
				return *localVarResponse.Data, localVarHttpResponse.Header, err
			}
		} else {
			return localVarReturnValue, localVarHttpResponse.Header, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v AdcreativeTemplateGetResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse.Header, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse.Header, newErr
		}

		return localVarReturnValue, localVarHttpResponse.Header, newErr
	}

	return localVarReturnValue, localVarHttpResponse.Header, nil
}
