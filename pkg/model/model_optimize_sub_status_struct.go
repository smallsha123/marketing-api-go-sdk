/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 优化状态（0: 无需优化 'OPTIMIZE_STATUS_NONE', 1: 待优化 'OPTIMIZE_STATUS_PENDING', 2: 已优化 'OPTIMIZE_STATUS_FINISHED'）
type OptimizeSubStatusStruct struct {
	Targeting                 OptimizeStatus `json:"targeting,omitempty"`
	TargetingExpand           OptimizeStatus `json:"targeting_expand,omitempty"`
	TargetingLocard           OptimizeStatus `json:"targeting_locard,omitempty"`
	Bid                       OptimizeStatus `json:"bid,omitempty"`
	BidStrategy               OptimizeStatus `json:"bid_strategy,omitempty"`
	BidAmount                 OptimizeStatus `json:"bid_amount,omitempty"`
	Budget                    OptimizeStatus `json:"budget,omitempty"`
	DailyBudget               OptimizeStatus `json:"daily_budget,omitempty"`
	AccountBalance            OptimizeStatus `json:"account_balance,omitempty"`
	TargetingStatusDesc       *string        `json:"targeting_status_desc,omitempty"`
	TargetingExpandStatusDesc *string        `json:"targeting_expand_status_desc,omitempty"`
	TargetingLocardStatusDesc *string        `json:"targeting_locard_status_desc,omitempty"`
	BidStatusDesc             *string        `json:"bid_status_desc,omitempty"`
	BidStrategyStatusDesc     *string        `json:"bid_strategy_status_desc,omitempty"`
	BidAmountStatusDesc       *string        `json:"bid_amount_status_desc,omitempty"`
	BudgetStatusDesc          *string        `json:"budget_status_desc,omitempty"`
	DailyBudgetStatusDesc     *string        `json:"daily_budget_status_desc,omitempty"`
	AccountBalanceStatusDesc  *string        `json:"account_balance_status_desc,omitempty"`
}
