/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 单日曝光趋势
type ExposureEffectDataTrendsItem struct {
	Date   *string        `json:"date,omitempty"`
	Trends *[]PointStruct `json:"trends,omitempty"`
}
