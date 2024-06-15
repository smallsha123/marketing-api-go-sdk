/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// UserAction人群信息
type UserActionSpec struct {
	UserActionSetId       *int64                 `json:"user_action_set_id,omitempty"`
	MatchRuleType         MatchRuleType          `json:"match_rule_type,omitempty"`
	ExtractType           ExtractRuleType        `json:"extract_type,omitempty"`
	TimeWindow            *int64                 `json:"time_window,omitempty"`
	UrlMatchRule          *UrlMatchRule          `json:"url_match_rule,omitempty"`
	ActionMatchRule       *ActionMatchRule       `json:"action_match_rule,omitempty"`
	ActionAggregationRule *ActionAggregationRule `json:"action_aggregation_rule,omitempty"`
}
