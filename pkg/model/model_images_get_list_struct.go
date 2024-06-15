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
type ImagesGetListStruct struct {
	ImageId           *string         `json:"image_id,omitempty"`
	Description       *string         `json:"description,omitempty"`
	SourceSignature   *string         `json:"source_signature,omitempty"`
	PreviewUrl        *string         `json:"preview_url,omitempty"`
	SourceType        ImageSourceType `json:"source_type,omitempty"`
	ImageUsage        ImageUsage      `json:"image_usage,omitempty"`
	CreatedTime       *int64          `json:"created_time,omitempty"`
	LastModifiedTime  *int64          `json:"last_modified_time,omitempty"`
	ProductCatalogId  *int64          `json:"product_catalog_id,omitempty"`
	ProductOuterId    *string         `json:"product_outer_id,omitempty"`
	SourceReferenceId *string         `json:"source_reference_id,omitempty"`
	OwnerAccountId    *string         `json:"owner_account_id,omitempty"`
	Status            MediaStatusType `json:"status,omitempty"`
	SampleAspectRatio *string         `json:"sample_aspect_ratio,omitempty"`
	Width             *int64          `json:"width,omitempty"`
	Height            *int64          `json:"height,omitempty"`
	FileSize          *int64          `json:"file_size,omitempty"`
	Type_             ImageType       `json:"type,omitempty"`
	Signature         *string         `json:"signature,omitempty"`
}
