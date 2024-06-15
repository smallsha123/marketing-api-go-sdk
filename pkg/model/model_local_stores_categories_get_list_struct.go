/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type LocalStoresCategoriesGetListStruct struct {
	VerticalId       *int64  `json:"vertical_id,omitempty"`
	VerticalName     *string `json:"vertical_name,omitempty"`
	CategoryId       *int64  `json:"category_id,omitempty"`
	ParentCategoryId *int64  `json:"parent_category_id,omitempty"`
	CategoryName     *string `json:"category_name,omitempty"`
	Level            *int64  `json:"level,omitempty"`
}
