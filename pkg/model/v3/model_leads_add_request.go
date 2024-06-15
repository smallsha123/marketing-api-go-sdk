/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type LeadsAddRequest struct {
	AccountId     *int64                `json:"account_id,omitempty"`
	MatchType     LeadsMatchType        `json:"match_type,omitempty"`
	LeadsInfoList *[]LeadsAddInfoStruct `json:"leads_info_list,omitempty"`
}
