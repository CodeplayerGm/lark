// Code generated by lark_sdk_gen. DO NOT EDIT.
/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lark

import (
	"context"
)

// GetApprovalList 查询当前用户可发起的审批定义列表。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/approval-v4/approval/list
func (r *ApprovalService) GetApprovalList(ctx context.Context, request *GetApprovalListReq, options ...MethodOptionFunc) (*GetApprovalListResp, *Response, error) {
	if r.cli.mock.mockApprovalGetApprovalList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Approval#GetApprovalList mock enable")
		return r.cli.mock.mockApprovalGetApprovalList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Approval",
		API:                 "GetApprovalList",
		Method:              "GET",
		URL:                 r.cli.openBaseURL + "/open-apis/approval/v4/approvals",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(getApprovalListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockApprovalGetApprovalList mock ApprovalGetApprovalList method
func (r *Mock) MockApprovalGetApprovalList(f func(ctx context.Context, request *GetApprovalListReq, options ...MethodOptionFunc) (*GetApprovalListResp, *Response, error)) {
	r.mockApprovalGetApprovalList = f
}

// UnMockApprovalGetApprovalList un-mock ApprovalGetApprovalList method
func (r *Mock) UnMockApprovalGetApprovalList() {
	r.mockApprovalGetApprovalList = nil
}

// GetApprovalListReq ...
type GetApprovalListReq struct {
	PageSize  *int64  `query:"page_size" json:"-"`  // 分页大小, 示例值: 10, 默认值: `10`, 最大值: `100`
	PageToken *string `query:"page_token" json:"-"` // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "ASDJHA1323_sda1JSASDFD"
	Locale    *string `query:"locale" json:"-"`     // --zh-CN: 中文, en-US: 英文, ja-JP: 日文, 示例值: "zh-CN"
}

// GetApprovalListResp ...
type GetApprovalListResp struct {
	Items     []*GetApprovalListRespItem `json:"items,omitempty"`      // 审批定义列表
	PageToken string                     `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore   bool                       `json:"has_more,omitempty"`   // 是否还有更多项
}

// GetApprovalListRespItem ...
type GetApprovalListRespItem struct {
	ApprovalCode string `json:"approval_code,omitempty"` // 审批定义 code  示例值: "7C468A54-8745-2245-9675-08B7C63E7A85"
	ApprovalName string `json:"approval_name,omitempty"` // 审批名称, 根据传入的local字段返回对应的国际化文案, 未设置该国际化文案时返回默认语言对应文案
}

// getApprovalListResp ...
type getApprovalListResp struct {
	Code int64                `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 错误描述
	Data *GetApprovalListResp `json:"data,omitempty"`
}