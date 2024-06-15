/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 广告清理工具-更新广告状态数据
type AdcleanUpdateConfiguredStatusData struct {
	IsSelectAll *string                                   `json:"is_select_all,omitempty"`
	CleanData   *[]UpdateAdgroupConfiguredStatusItemClean `json:"clean_data,omitempty"`
	QueryInfo   *string                                   `json:"query_info,omitempty"`
	CleanMode   AdCleanMode                               `json:"clean_mode,omitempty"`
}
