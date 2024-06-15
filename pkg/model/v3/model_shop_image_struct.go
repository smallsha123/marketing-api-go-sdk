/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 卖点图组件
type ShopImageStruct struct {
	ShopImageSwitch        *bool   `json:"shop_image_switch,omitempty"`
	DynamicShopImageSwitch *bool   `json:"dynamic_shop_image_switch,omitempty"`
	ShopImageId            *string `json:"shop_image_id,omitempty"`
	ShopImageTitle         *string `json:"shop_image_title,omitempty"`
	ShopImageDescription   *string `json:"shop_image_description,omitempty"`
}
