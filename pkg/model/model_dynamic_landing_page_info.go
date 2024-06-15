/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// dc落地页信息
type DynamicLandingPageInfo struct {
	PageType                      DestinationType          `json:"page_type,omitempty"`
	PageSpec                      *DynamicCreativePageSpec `json:"page_spec,omitempty"`
	LinkPageType                  LinkPageType             `json:"link_page_type,omitempty"`
	LinkNameType                  LinkUrlLinkNameType      `json:"link_name_type,omitempty"`
	LinkPageSpec                  *LinkPageSpec            `json:"link_page_spec,omitempty"`
	QqMiniGameTrackingQueryString *string                  `json:"qq_mini_game_tracking_query_string,omitempty"`
	ShareContentSpec              *ShareContentSpec        `json:"share_content_spec,omitempty"`
	WebviewUrl                    *string                  `json:"webview_url,omitempty"`
	SimpleCanvasSubType           SimpleCanvasSubType      `json:"simple_canvas_sub_type,omitempty"`
	DeepLinkUrl                   *string                  `json:"deep_link_url,omitempty"`
	AndroidDeepLinkAppId          *string                  `json:"android_deep_link_app_id,omitempty"`
	IosDeepLinkAppId              *string                  `json:"ios_deep_link_app_id,omitempty"`
	DoubleDeepLinkData            *DoubleDeepLinkDataSpec  `json:"double_deep_link_data,omitempty"`
	UniversalLinkUrl              *string                  `json:"universal_link_url,omitempty"`
	UnionMarketSwitch             *bool                    `json:"union_market_switch,omitempty"`
	UnionMarketSpec               *UnionMarketSpec         `json:"union_market_spec,omitempty"`
	LinkNameText                  *string                  `json:"link_name_text,omitempty"`
	ButtonTextJumpInfo            *LandingPageStructure    `json:"button_text_jump_info,omitempty"`
	ChannelsShopProductSpec       *ChannelsShopProductSpec `json:"channels_shop_product_spec,omitempty"`
}
