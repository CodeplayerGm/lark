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

// ExportVCParticipantQualityList 导出某场会议某个参会人的音视频&共享质量数据
//
// , 具体权限要求请参考「导出概述」
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/export/participant_quality_list
func (r *VCService) ExportVCParticipantQualityList(ctx context.Context, request *ExportVCParticipantQualityListReq, options ...MethodOptionFunc) (*ExportVCParticipantQualityListResp, *Response, error) {
	if r.cli.mock.mockVCExportVCParticipantQualityList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] VC#ExportVCParticipantQualityList mock enable")
		return r.cli.mock.mockVCExportVCParticipantQualityList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "VC",
		API:                   "ExportVCParticipantQualityList",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/vc/v1/exports/participant_quality_list",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(exportVCParticipantQualityListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockVCExportVCParticipantQualityList mock VCExportVCParticipantQualityList method
func (r *Mock) MockVCExportVCParticipantQualityList(f func(ctx context.Context, request *ExportVCParticipantQualityListReq, options ...MethodOptionFunc) (*ExportVCParticipantQualityListResp, *Response, error)) {
	r.mockVCExportVCParticipantQualityList = f
}

// UnMockVCExportVCParticipantQualityList un-mock VCExportVCParticipantQualityList method
func (r *Mock) UnMockVCExportVCParticipantQualityList() {
	r.mockVCExportVCParticipantQualityList = nil
}

// ExportVCParticipantQualityListReq ...
type ExportVCParticipantQualityListReq struct {
	UserIDType       *IDType `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 用户的 open id, union_id: 用户的 union id, user_id: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	MeetingStartTime string  `json:"meeting_start_time,omitempty"` // 会议开始时间（unix时间, 单位sec）, 示例值: "1655276858"
	MeetingEndTime   string  `json:"meeting_end_time,omitempty"`   // 会议结束时间（unix时间, 单位sec）, 示例值: "1655276858"
	MeetingNo        string  `json:"meeting_no,omitempty"`         // 9位会议号, 示例值: "123456789"
	JoinTime         string  `json:"join_time,omitempty"`          // 参会人入会时间（unix时间, 单位sec）, 示例值: "1655276858"
	UserID           *string `json:"user_id,omitempty"`            // 参会人为Lark用户时填入, room_id和user_id必须只填一个, 示例值: "ou_3ec3f6a28a0d08c45d895276e8e5e19b"
	RoomID           *string `json:"room_id,omitempty"`            // 参会人为Rooms时填入, room_id和user_id必须只填一个, 示例值: "omm_eada1d61a550955240c28757e7dec3af"
}

// ExportVCParticipantQualityListResp ...
type ExportVCParticipantQualityListResp struct {
	TaskID string `json:"task_id,omitempty"` // 任务id
}

// exportVCParticipantQualityListResp ...
type exportVCParticipantQualityListResp struct {
	Code int64                               `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                              `json:"msg,omitempty"`  // 错误描述
	Data *ExportVCParticipantQualityListResp `json:"data,omitempty"`
}