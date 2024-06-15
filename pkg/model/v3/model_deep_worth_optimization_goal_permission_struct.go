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
type DeepWorthOptimizationGoalPermissionStruct struct {
	OptimizationGoal              OptimizationGoal `json:"optimization_goal,omitempty"`
	DeepWorthOptimizationGoalList *[]string        `json:"deep_worth_optimization_goal_list,omitempty"`
}
