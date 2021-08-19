// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// EventV2VCMeetingMeetingEndedV1
//
// 发生在会议结束时
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/meeting/events/meeting_ended
func (r *EventCallbackService) HandlerEventV2VCMeetingMeetingEndedV1(f eventV2VCMeetingMeetingEndedV1Handler) {
	r.cli.eventHandler.eventV2VCMeetingMeetingEndedV1Handler = f
}

type eventV2VCMeetingMeetingEndedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2VCMeetingMeetingEndedV1) (string, error)

type EventV2VCMeetingMeetingEndedV1 struct {
	Meeting  *EventV2VCMeetingMeetingEndedV1Meeting  `json:"meeting,omitempty"`  // 会议数据
	Operator *EventV2VCMeetingMeetingEndedV1Operator `json:"operator,omitempty"` // 事件操作人
}

type EventV2VCMeetingMeetingEndedV1Meeting struct {
	ID        string                                         `json:"id,omitempty"`         // 会议ID（视频会议的唯一标识，视频会议开始后才会产生）
	Topic     string                                         `json:"topic,omitempty"`      // 会议主题
	MeetingNo string                                         `json:"meeting_no,omitempty"` // 9位会议号（飞书用户可通过输入9位会议号快捷入会）
	StartTime string                                         `json:"start_time,omitempty"` // 会议开始时间（unix时间，单位sec）
	EndTime   string                                         `json:"end_time,omitempty"`   // 会议结束时间（unix时间，单位sec）
	HostUser  *EventV2VCMeetingMeetingEndedV1MeetingHostUser `json:"host_user,omitempty"`  // 会议主持人
	Owner     *EventV2VCMeetingMeetingEndedV1MeetingOwner    `json:"owner,omitempty"`      // 会议拥有者
}

type EventV2VCMeetingMeetingEndedV1MeetingHostUser struct {
	ID       *EventV2VCMeetingMeetingEndedV1MeetingHostUserID `json:"id,omitempty"`        // 用户 ID
	UserRole int64                                            `json:"user_role,omitempty"` // 用户会中角色, 可选值有: `1`：普通参会人, `2`：主持人, `3`：联席主持人
	UserType int64                                            `json:"user_type,omitempty"` // 用户类型, 可选值有: `1`：lark用户, `2`：rooms用户, `3`：文档用户, `4`：neo单品用户, `5`：neo单品游客用户, `6`：pstn用户, `7`：sip用户
}

type EventV2VCMeetingMeetingEndedV1MeetingHostUserID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 user ID
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}

type EventV2VCMeetingMeetingEndedV1MeetingOwner struct {
	ID       *EventV2VCMeetingMeetingEndedV1MeetingOwnerID `json:"id,omitempty"`        // 用户 ID
	UserRole int64                                         `json:"user_role,omitempty"` // 用户会中角色, 可选值有: `1`：普通参会人, `2`：主持人, `3`：联席主持人
	UserType int64                                         `json:"user_type,omitempty"` // 用户类型, 可选值有: `1`：lark用户, `2`：rooms用户, `3`：文档用户, `4`：neo单品用户, `5`：neo单品游客用户, `6`：pstn用户, `7`：sip用户
}

type EventV2VCMeetingMeetingEndedV1MeetingOwnerID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 user ID
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}

type EventV2VCMeetingMeetingEndedV1Operator struct {
	ID       *EventV2VCMeetingMeetingEndedV1OperatorID `json:"id,omitempty"`        // 用户 ID
	UserRole int64                                     `json:"user_role,omitempty"` // 用户会中角色, 可选值有: `1`：普通参会人, `2`：主持人, `3`：联席主持人
	UserType int64                                     `json:"user_type,omitempty"` // 用户类型, 可选值有: `1`：lark用户, `2`：rooms用户, `3`：文档用户, `4`：neo单品用户, `5`：neo单品游客用户, `6`：pstn用户, `7`：sip用户
}

type EventV2VCMeetingMeetingEndedV1OperatorID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id, 字段权限要求: 获取用户 user ID
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}
