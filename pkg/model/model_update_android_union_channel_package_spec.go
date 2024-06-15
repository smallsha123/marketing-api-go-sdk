/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 更新广告渠道包接口任务所需条件
type UpdateAndroidUnionChannelPackageSpec struct {
	ChannelPackageId *int64  `json:"channel_package_id,omitempty"`
	PackageName      *string `json:"package_name,omitempty"`
	DownloadUrl      *string `json:"download_url,omitempty"`
}
