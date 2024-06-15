/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 用户上传自定义特征文件规则
type CustomFileFeatureSpec struct {
	DataType           FeatureValueDataType `json:"data_type,omitempty"`
	IsMultiValued      *bool                `json:"is_multi_valued,omitempty"`
	PossibleValuesSize *int64               `json:"possible_values_size,omitempty"`
}