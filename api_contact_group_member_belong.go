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

// GetContactMemberGroupList 通过该接口可查询该用户所属的用户组列表, 可分别查询普通用户组和动态用户组。如果应用的通讯录权限范围是“全部员工”, 则可获取该员工所属的全部用户组列表。如果应用的通讯录权限范围不是“全部员工”, 则仅可获取通讯录权限范围内该员工所属的用户组。[点击了解通讯录权限范围](https://open.feishu.cn/document/ukTMukTMukTM/uETNz4SM1MjLxUzM/v3/guides/scope_authority)。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group/member_belong
// new doc: https://open.feishu.cn/document/server-docs/contact-v3/group/member_belong
func (r *ContactService) GetContactMemberGroupList(ctx context.Context, request *GetContactMemberGroupListReq, options ...MethodOptionFunc) (*GetContactMemberGroupListResp, *Response, error) {
	if r.cli.mock.mockContactGetContactMemberGroupList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#GetContactMemberGroupList mock enable")
		return r.cli.mock.mockContactGetContactMemberGroupList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "GetContactMemberGroupList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/group/member_belong",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getContactMemberGroupListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactGetContactMemberGroupList mock ContactGetContactMemberGroupList method
func (r *Mock) MockContactGetContactMemberGroupList(f func(ctx context.Context, request *GetContactMemberGroupListReq, options ...MethodOptionFunc) (*GetContactMemberGroupListResp, *Response, error)) {
	r.mockContactGetContactMemberGroupList = f
}

// UnMockContactGetContactMemberGroupList un-mock ContactGetContactMemberGroupList method
func (r *Mock) UnMockContactGetContactMemberGroupList() {
	r.mockContactGetContactMemberGroupList = nil
}

// GetContactMemberGroupListReq ...
type GetContactMemberGroupListReq struct {
	MemberID     string  `query:"member_id" json:"-"`      // 成员ID, 示例值: "u287xj12"
	MemberIDType *IDType `query:"member_id_type" json:"-"` // 成员ID类型, 示例值: "open_id", 可选值有: open_id: member_id_type为user时, 表示用户的open_id, union_id: member_id_type为user时, 表示用户的union_id, user_id: member_id_type为user时, 表示用户的user_id, 默认值: `open_id`
	GroupType    *int64  `query:"group_type" json:"-"`     // 欲获取的用户组类型, 示例值: 1, 可选值有: 1: 普通用户组, 2: 动态用户组, 取值范围: `1` ～ `2`
	PageSize     *int64  `query:"page_size" json:"-"`      // 分页查询大小, 示例值: 500, 默认值: `500`, 取值范围: `1` ～ `1000`
	PageToken    *string `query:"page_token" json:"-"`     // 分页查询Token, 示例值: "AQD9/Rn9eij9Pm39ED40/dk53s4Ebp882DYfFaPFbz00L4CMZJrqGdzNyc8BcZtDbwVUvRmQTvyMYicnGWrde9X56TgdBuS+JKiSIkdexPw="
}

// GetContactMemberGroupListResp ...
type GetContactMemberGroupListResp struct {
	GroupList []string `json:"group_list,omitempty"` // 用户组ID列表
	PageToken string   `json:"page_token,omitempty"` // 分页查询Token
	HasMore   bool     `json:"has_more,omitempty"`   // 是否有更多结果
}

// getContactMemberGroupListResp ...
type getContactMemberGroupListResp struct {
	Code int64                          `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                         `json:"msg,omitempty"`  // 错误描述
	Data *GetContactMemberGroupListResp `json:"data,omitempty"`
}
