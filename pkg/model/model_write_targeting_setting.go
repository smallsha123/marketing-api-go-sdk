/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 定向详细设置
type WriteTargetingSetting struct {
	Age                       *[]AgeStruct               `json:"age,omitempty"`
	Gender                    *[]string                  `json:"gender,omitempty"`
	Education                 *[]string                  `json:"education,omitempty"`
	MaritalStatus             *[]string                  `json:"marital_status,omitempty"`
	GeoLocation               *GeoLocations              `json:"geo_location,omitempty"`
	UserOs                    *[]string                  `json:"user_os,omitempty"`
	DevicePrice               *[]string                  `json:"device_price,omitempty"`
	DeviceBrandModel          *DeviceBrandModel          `json:"device_brand_model,omitempty"`
	NetworkType               *[]string                  `json:"network_type,omitempty"`
	NetworkOperator           *[]string                  `json:"network_operator,omitempty"`
	AppInstallStatus          *[]string                  `json:"app_install_status,omitempty"`
	ConsumptionStatus         *[]string                  `json:"consumption_status,omitempty"`
	GameConsumptionLevel      *[]string                  `json:"game_consumption_level,omitempty"`
	FinancialSituation        *[]string                  `json:"financial_situation,omitempty"`
	WechatAdBehavior          *WechatAdBehavior          `json:"wechat_ad_behavior,omitempty"`
	CustomAudience            *[]int64                   `json:"custom_audience,omitempty"`
	ExcludedCustomAudience    *[]int64                   `json:"excluded_custom_audience,omitempty"`
	BehaviorOrInterest        *BehaviorOrInterest        `json:"behavior_or_interest,omitempty"`
	ExcludedConvertedAudience *ExcludedConvertedAudience `json:"excluded_converted_audience,omitempty"`
}
