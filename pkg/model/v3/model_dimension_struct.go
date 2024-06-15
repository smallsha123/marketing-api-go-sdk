/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 创意形式
type DimensionStruct struct {
	RatioWidth          *int64   `json:"ratio_width,omitempty"`
	RatioHeight         *int64   `json:"ratio_height,omitempty"`
	MinWidth            *int64   `json:"min_width,omitempty"`
	MinHeight           *int64   `json:"min_height,omitempty"`
	FileSizeKblimit     *int64   `json:"file_size_kblimit,omitempty"`
	MinDuration         *float64 `json:"min_duration,omitempty"`
	MaxDuration         *float64 `json:"max_duration,omitempty"`
	CreativeTemplateIds *[]int64 `json:"creative_template_ids,omitempty"`
	MediaType           *string  `json:"media_type,omitempty"`
	MinOccurs           *int64   `json:"min_occurs,omitempty"`
	MaxOccurs           *int64   `json:"max_occurs,omitempty"`
}
