/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 外链信息
type LinkSpec struct {
	Title               *string                 `json:"title,omitempty"`
	Url                 *string                 `json:"url,omitempty"`
	DeepLinkAndroidSpec *DeepLinkAppAndroidSpec `json:"deep_link_android_spec,omitempty"`
	DeepLinkIosSpec     *DeepLinkAppIosSpec     `json:"deep_link_ios_spec,omitempty"`
}
