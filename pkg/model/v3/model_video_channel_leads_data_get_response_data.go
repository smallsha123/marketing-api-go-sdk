/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type VideoChannelLeadsDataGetResponseData struct {
	LeadsInfoList *[]LeadsInfoListStruct      `json:"leads_info_list,omitempty"`
	PageInfo      *VideoChannelPageInfoStruct `json:"page_info,omitempty"`
}
