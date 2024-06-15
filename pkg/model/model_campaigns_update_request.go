/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type CampaignsUpdateRequest struct {
	AccountId        *int64    `json:"account_id,omitempty"`
	CampaignId       *int64    `json:"campaign_id,omitempty"`
	CampaignName     *string   `json:"campaign_name,omitempty"`
	DailyBudget      *int64    `json:"daily_budget,omitempty"`
	TotalBudget      *int64    `json:"total_budget,omitempty"`
	ConfiguredStatus AdStatus  `json:"configured_status,omitempty"`
	SpeedMode        SpeedMode `json:"speed_mode,omitempty"`
	BeginDate        *string   `json:"begin_date,omitempty"`
	EndDate          *string   `json:"end_date,omitempty"`
	IsAutoReplenish  *int64    `json:"is_auto_replenish,omitempty"`
}
