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

// DeleteCoreHrPerson 删除人员的个人信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/person/delete
func (r *CoreHrService) DeleteCoreHrPerson(ctx context.Context, request *DeleteCoreHrPersonReq, options ...MethodOptionFunc) (*DeleteCoreHrPersonResp, *Response, error) {
	if r.cli.mock.mockCoreHrDeleteCoreHrPerson != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] CoreHr#DeleteCoreHrPerson mock enable")
		return r.cli.mock.mockCoreHrDeleteCoreHrPerson(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHr",
		API:                   "DeleteCoreHrPerson",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/persons/:person_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteCoreHrPersonResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHrDeleteCoreHrPerson mock CoreHrDeleteCoreHrPerson method
func (r *Mock) MockCoreHrDeleteCoreHrPerson(f func(ctx context.Context, request *DeleteCoreHrPersonReq, options ...MethodOptionFunc) (*DeleteCoreHrPersonResp, *Response, error)) {
	r.mockCoreHrDeleteCoreHrPerson = f
}

// UnMockCoreHrDeleteCoreHrPerson un-mock CoreHrDeleteCoreHrPerson method
func (r *Mock) UnMockCoreHrDeleteCoreHrPerson() {
	r.mockCoreHrDeleteCoreHrPerson = nil
}

// DeleteCoreHrPersonReq ...
type DeleteCoreHrPersonReq struct {
	PersonID string `path:"person_id" json:"-"` // 需要删除的Person ID, 示例值: "654637829201"
}

// DeleteCoreHrPersonResp ...
type DeleteCoreHrPersonResp struct {
}

// deleteCoreHrPersonResp ...
type deleteCoreHrPersonResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *DeleteCoreHrPersonResp `json:"data,omitempty"`
}