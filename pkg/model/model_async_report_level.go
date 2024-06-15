/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// AsyncReportLevel : 异步报表类型级别
type AsyncReportLevel string

// List of AsyncReportLevel
const (
	AsyncReportLevel_ADGROUP_WECHAT     AsyncReportLevel = "REPORT_LEVEL_ADGROUP_WECHAT"
	AsyncReportLevel_AD_WECHAT          AsyncReportLevel = "REPORT_LEVEL_AD_WECHAT"
	AsyncReportLevel_POI_WECHAT         AsyncReportLevel = "REPORT_LEVEL_POI_WECHAT"
	AsyncReportLevel_AD                 AsyncReportLevel = "REPORT_LEVEL_AD"
	AsyncReportLevel_ADVERTISER         AsyncReportLevel = "REPORT_LEVEL_ADVERTISER"
	AsyncReportLevel_CAMPAIGN           AsyncReportLevel = "REPORT_LEVEL_CAMPAIGN"
	AsyncReportLevel_ADGROUP            AsyncReportLevel = "REPORT_LEVEL_ADGROUP"
	AsyncReportLevel_MATERIAL_VIDEO     AsyncReportLevel = "REPORT_LEVEL_MATERIAL_VIDEO"
	AsyncReportLevel_MATERIAL_IMAGE     AsyncReportLevel = "REPORT_LEVEL_MATERIAL_IMAGE"
	AsyncReportLevel_PROMOTED_OBJECT    AsyncReportLevel = "REPORT_LEVEL_PROMOTED_OBJECT"
	AsyncReportLevel_CREATIVE_TEMPLATE  AsyncReportLevel = "REPORT_LEVEL_CREATIVE_TEMPLATE"
	AsyncReportLevel_PRODUCT_CATELOG    AsyncReportLevel = "REPORT_LEVEL_PRODUCT_CATELOG"
	AsyncReportLevel_AGE                AsyncReportLevel = "REPORT_LEVEL_AGE"
	AsyncReportLevel_GENDER             AsyncReportLevel = "REPORT_LEVEL_GENDER"
	AsyncReportLevel_REGION_RECENTLY_IN AsyncReportLevel = "REPORT_LEVEL_REGION_RECENTLY_IN"
	AsyncReportLevel_REGION_VISITED_IN  AsyncReportLevel = "REPORT_LEVEL_REGION_VISITED_IN"
	AsyncReportLevel_REGION_LIVE_IN     AsyncReportLevel = "REPORT_LEVEL_REGION_LIVE_IN"
	AsyncReportLevel_REGION_TRAVEL_IN   AsyncReportLevel = "REPORT_LEVEL_REGION_TRAVEL_IN"
	AsyncReportLevel_CITY_RECENTLY_IN   AsyncReportLevel = "REPORT_LEVEL_CITY_RECENTLY_IN"
	AsyncReportLevel_CITY_VISITED_IN    AsyncReportLevel = "REPORT_LEVEL_CITY_VISITED_IN"
	AsyncReportLevel_CITY_LIVE_IN       AsyncReportLevel = "REPORT_LEVEL_CITY_LIVE_IN"
	AsyncReportLevel_CITY_TRAVEL_IN     AsyncReportLevel = "REPORT_LEVEL_CITY_TRAVEL_IN"
	AsyncReportLevel_BIDWORD            AsyncReportLevel = "REPORT_LEVEL_BIDWORD"
	AsyncReportLevel_QUERYWORD          AsyncReportLevel = "REPORT_LEVEL_QUERYWORD"
)
