/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 回传信息结构
type LeadsActionTypeReportListStruct struct {
	OuterLeadsId *string `json:"outer_leads_id,omitempty"`
	LeadsId      *int64  `json:"leads_id,omitempty"`
	LeadsTel     *string `json:"leads_tel,omitempty"`
	LeadsQq      *int64  `json:"leads_qq,omitempty"`
	LeadsWechat  *string `json:"leads_wechat,omitempty"`
	ClickId      *string `json:"click_id,omitempty"`
	ActionType   *string `json:"action_type,omitempty"`
	CallDuration *int64  `json:"call_duration,omitempty"`
}
