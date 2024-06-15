/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type LeadsFormGetResponseData struct {
	AccountId     *int64                       `json:"account_id,omitempty"`
	ComponentId   *string                      `json:"component_id,omitempty"`
	ComponentName *string                      `json:"component_name,omitempty"`
	CreatedTime   *string                      `json:"created_time,omitempty"`
	FormConfig    *FormConfigDetailData        `json:"form_config,omitempty"`
	ItemList      *[]ControlListItemDetailData `json:"item_list,omitempty"`
}
