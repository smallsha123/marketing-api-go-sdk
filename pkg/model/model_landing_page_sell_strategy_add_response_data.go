/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type LandingPageSellStrategyAddResponseData struct {
	AccountId       *int64   `json:"account_id,omitempty"`
	StrategyId      *int64   `json:"strategy_id,omitempty"`
	StrategyName    *string  `json:"strategy_name,omitempty"`
	StrategyType    *string  `json:"strategy_type,omitempty"`
	StrategyStatus  *int64   `json:"strategy_status,omitempty"`
	EpisodePrice    *float64 `json:"episode_price,omitempty"`
	MinRechargeTier *float64 `json:"min_recharge_tier,omitempty"`
	RechargeNum     *int64   `json:"recharge_num,omitempty"`
}
