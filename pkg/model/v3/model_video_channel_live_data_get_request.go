/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type VideoChannelLiveDataGetRequest struct {
	AccountId         *int64    `json:"account_id,omitempty"`
	BrandIds          *[]string `json:"brand_ids,omitempty"`
	BrandNames        *[]string `json:"brand_names,omitempty"`
	DealerIds         *[]string `json:"dealer_ids,omitempty"`
	DealerNames       *[]string `json:"dealer_names,omitempty"`
	VideoChannelIds   *[]string `json:"video_channel_ids,omitempty"`
	VideoChannelNames *[]string `json:"video_channel_names,omitempty"`
	StartDate         *int64    `json:"start_date,omitempty"`
	EndDate           *int64    `json:"end_date,omitempty"`
	Page              *int64    `json:"page,omitempty"`
	PageSize          *int64    `json:"page_size,omitempty"`
}
