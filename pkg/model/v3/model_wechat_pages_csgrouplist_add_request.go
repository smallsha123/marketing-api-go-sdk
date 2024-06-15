/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type WechatPagesCsgrouplistAddRequest struct {
	AccountId  *int64    `json:"account_id,omitempty"`
	CorpId     *string   `json:"corp_id,omitempty"`
	GroupName  *string   `json:"group_name,omitempty"`
	UserIdList *[]string `json:"user_id_list,omitempty"`
}
