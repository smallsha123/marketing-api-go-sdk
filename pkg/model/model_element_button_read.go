/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 按钮组件元素
type ElementButtonRead struct {
	Title               *string                `json:"title,omitempty"`
	Url                 *string                `json:"url,omitempty"`
	LinkSpec            *ElementButtonLinkSpec `json:"link_spec,omitempty"`
	AppIosSpec          *AppIosSpec            `json:"app_ios_spec,omitempty"`
	AppAndroidSpec      *AppAndroidSpec        `json:"app_android_spec,omitempty"`
	MiniProgramSpec     *MiniProgramSpec       `json:"mini_program_spec,omitempty"`
	MiniGameProgramSpec *MiniGameProgramSpec   `json:"mini_game_program_spec,omitempty"`
	FengyeSpec          *FengyeSpec            `json:"fengye_spec,omitempty"`
	CardSpec            *CardSpec              `json:"card_spec,omitempty"`
	FollowSpec          *FollowSpec            `json:"follow_spec,omitempty"`
	ServiceSpec         *ServiceSpec           `json:"service_spec,omitempty"`
	WecomSpec           *WecomSpec             `json:"wecom_spec,omitempty"`
	UseIcon             *int64                 `json:"use_icon,omitempty"`
	TelSpec             *TelSpec               `json:"tel_spec,omitempty"`
	VideoChannelSpec    *VideoChannelSpec      `json:"video_channel_spec,omitempty"`
}