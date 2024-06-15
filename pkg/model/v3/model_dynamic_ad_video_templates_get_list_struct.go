/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type DynamicAdVideoTemplatesGetListStruct struct {
	TemplateId           *int64               `json:"template_id,omitempty"`
	TemplateName         *string              `json:"template_name,omitempty"`
	TemplateType         VideoTemplateType    `json:"template_type,omitempty"`
	ProductCatalogId     *int64               `json:"product_catalog_id,omitempty"`
	AdcreativeTemplateId *int64               `json:"adcreative_template_id,omitempty"`
	CoverImageUrl        *string              `json:"cover_image_url,omitempty"`
	IntroVideoUrl        *string              `json:"intro_video_url,omitempty"`
	DeliveryVideoUrl     *string              `json:"delivery_video_url,omitempty"`
	SupportChannel       *bool                `json:"support_channel,omitempty"`
	Coverage             *float64             `json:"coverage,omitempty"`
	MinVideoDuration     *float64             `json:"min_video_duration,omitempty"`
	MaxVideoDuration     *float64             `json:"max_video_duration,omitempty"`
	VideoProductFields   *[]string            `json:"video_product_fields,omitempty"`
	ImageProductFields   *[]string            `json:"image_product_fields,omitempty"`
	Extra                *Extra               `json:"extra,omitempty"`
	SubTemplateList      *[]SubTemplateStruct `json:"sub_template_list,omitempty"`
}
