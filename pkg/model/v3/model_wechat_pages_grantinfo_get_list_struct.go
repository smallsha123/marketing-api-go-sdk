/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type WechatPagesGrantinfoGetListStruct struct {
	OwnerAccountId   *int64  `json:"owner_account_id,omitempty"`
	OwnerAccountName *string `json:"owner_account_name,omitempty"`
	CreatedTime      *string `json:"created_time,omitempty"`
}