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
type MergeFundTypeFundsGetListStruct struct {
	FundType     AccountMergeTypeMap `json:"fund_type,omitempty"`
	Balance      *int64              `json:"balance,omitempty"`
	FundStatus   FundStatus          `json:"fund_status,omitempty"`
	RealtimeCost *int64              `json:"realtime_cost,omitempty"`
	EffectFunds  *[]string           `json:"effect_funds,omitempty"`
}
