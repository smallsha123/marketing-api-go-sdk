/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type ProductCatalogsAddRequest struct {
	AccountId        *int64          `json:"account_id,omitempty"`
	CatalogName      *string         `json:"catalog_name,omitempty"`
	CatalogScaleType CatalogScale    `json:"catalog_scale_type,omitempty"`
	CatalogType      CatalogType     `json:"catalog_type,omitempty"`
	IndustryType     CatalogIndustry `json:"industry_type,omitempty"`
}
