/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 落地页内容
type PageSpecWithoutWechatChannelsFeedPageSpec struct {
	AndroidAppSpec                  *AndroidAppPageSpec                  `json:"android_app_spec,omitempty"`
	IosAppSpec                      *IosAppPageSpec                      `json:"ios_app_spec,omitempty"`
	XjAndroidAppH5Spec              *XjPageSpec                          `json:"xj_android_app_h5_spec,omitempty"`
	XjIosAppH5Spec                  *XjPageSpec                          `json:"xj_ios_app_h5_spec,omitempty"`
	XjWebH5Spec                     *XjPageSpec                          `json:"xj_web_h5_spec,omitempty"`
	XjQuickSpec                     *XjPageSpec                          `json:"xj_quick_spec,omitempty"`
	FengyeEcommmerceSpec            *FengyePageSpec                      `json:"fengye_ecommmerce_spec,omitempty"`
	WechatCanvasSpec                *CanvasWechatPageSpec                `json:"wechat_canvas_spec,omitempty"`
	WechatMiniProgramSpec           *WechatMiniProgramPageSpec           `json:"wechat_mini_program_spec,omitempty"`
	WechatCanvasMiniProgramSpec     *CanvasWechatPageSpec                `json:"wechat_canvas_mini_program_spec,omitempty"`
	QqAppMiniProgramSpec            *QqAppMiniProgramPageSpec            `json:"qq_app_mini_program_spec,omitempty"`
	SimpleWechatCanvasSpec          *SimpleCanvasWechatSpec              `json:"simple_wechat_canvas_spec,omitempty"`
	WechatFocusDialogSpec           *WechatFocusDialog                   `json:"wechat_focus_dialog_spec,omitempty"`
	WechatConsultSpec               *WechatConsultPageSpec               `json:"wechat_consult_spec,omitempty"`
	WecomConsultSpec                *WecomConsultPageSpec                `json:"wecom_consult_spec,omitempty"`
	WechatOfficialAccountDetailSpec *WechatOfficialAccountDetailPageSpec `json:"wechat_official_account_detail_spec,omitempty"`
	AppDeepLinkSpec                 *AppDeepLinkPageSpec                 `json:"app_deep_link_spec,omitempty"`
	AppMarketSpec                   *AppMarketPageSpec                   `json:"app_market_spec,omitempty"`
	AndroidDirectDownloadSpec       *AndroidDirectDownloadPageSpec       `json:"android_direct_download_spec,omitempty"`
	OfficialSpec                    *XjPageSpec                          `json:"official_spec,omitempty"`
	H5ProfileSpec                   *H5ProfilePageSpec                   `json:"h5_profile_spec,omitempty"`
	SearchAreaBrandSpec             *SearchAreaBrandPageSpec             `json:"search_area_brand_spec,omitempty"`
	WechatChannelsProfileSpec       *WechatChannelsPageSpec              `json:"wechat_channels_profile_spec,omitempty"`
	H5Spec                          *H5PageSpec                          `json:"h5_spec,omitempty"`
	WechatMiniGameSpec              *WechatMiniGamePageSpec              `json:"wechat_mini_game_spec,omitempty"`
	WechatChannelsReserveSpec       *WechatChannelsReserveLivePageSpec   `json:"wechat_channels_reserve_spec,omitempty"`
	AndroidQuickAppSpec             *AndroidQuickAppPageSpec             `json:"android_quick_app_spec,omitempty"`
}
