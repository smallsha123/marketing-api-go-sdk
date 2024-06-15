/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 关键词信息
type SearchBidwordStruct struct {
	AdgroupId        *int64            `json:"adgroup_id,omitempty"`
	Bidword          *string           `json:"bidword,omitempty"`
	BidPrice         *int64            `json:"bid_price,omitempty"`
	UseGroupPrice    UseGroupPriceType `json:"use_group_price,omitempty"`
	MatchType        BidwordMatchType  `json:"match_type,omitempty"`
	ConfiguredStatus BidwordPauseType  `json:"configured_status,omitempty"`
}
