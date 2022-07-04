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

// DeleteHelpdeskFAQ 该接口用于删除知识库。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/faq/delete
func (r *HelpdeskService) DeleteHelpdeskFAQ(ctx context.Context, request *DeleteHelpdeskFAQReq, options ...MethodOptionFunc) (*DeleteHelpdeskFAQResp, *Response, error) {
	if r.cli.mock.mockHelpdeskDeleteHelpdeskFAQ != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#DeleteHelpdeskFAQ mock enable")
		return r.cli.mock.mockHelpdeskDeleteHelpdeskFAQ(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Helpdesk",
		API:                 "DeleteHelpdeskFAQ",
		Method:              "DELETE",
		URL:                 r.cli.openBaseURL + "/open-apis/helpdesk/v1/faqs/:id",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
		NeedHelpdeskAuth:    true,
	}
	resp := new(deleteHelpdeskFAQResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHelpdeskDeleteHelpdeskFAQ mock HelpdeskDeleteHelpdeskFAQ method
func (r *Mock) MockHelpdeskDeleteHelpdeskFAQ(f func(ctx context.Context, request *DeleteHelpdeskFAQReq, options ...MethodOptionFunc) (*DeleteHelpdeskFAQResp, *Response, error)) {
	r.mockHelpdeskDeleteHelpdeskFAQ = f
}

// UnMockHelpdeskDeleteHelpdeskFAQ un-mock HelpdeskDeleteHelpdeskFAQ method
func (r *Mock) UnMockHelpdeskDeleteHelpdeskFAQ() {
	r.mockHelpdeskDeleteHelpdeskFAQ = nil
}

// DeleteHelpdeskFAQReq ...
type DeleteHelpdeskFAQReq struct {
	ID string `path:"id" json:"-"` // id, 示例值: "12345"
}

// DeleteHelpdeskFAQResp ...
type DeleteHelpdeskFAQResp struct {
}

// deleteHelpdeskFAQResp ...
type deleteHelpdeskFAQResp struct {
	Code int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *DeleteHelpdeskFAQResp `json:"data,omitempty"`
}