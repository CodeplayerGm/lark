package lark

import (
	"context"
)

// CreateCalendarTimeoffEvent 为指定用户创建一个请假日程，可以是一个普通请假日程，也可以是一个全天日程。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/timeoff_event/create
func (r *CalendarAPI) CreateCalendarTimeoffEvent(ctx context.Context, request *CreateCalendarTimeoffEventReq) (*CreateCalendarTimeoffEventResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/calendar/v4/timeoff_events",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(createCalendarTimeoffEventResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("Calendar", "CreateCalendarTimeoffEvent", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type CreateCalendarTimeoffEventReq struct {
	UserIDType  *IDType `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 userid
	UserID      string  `json:"user_id,omitempty"`      // 用户id, 示例值："ou_XXXXXXXXXX"
	Timezone    string  `json:"timezone,omitempty"`     // 时区, 示例值："Asia/Shanghai"
	StartTime   string  `json:"start_time,omitempty"`   // 休假开始时间: 有时间戳(1609430400)和日期(2021-01-01)两种格式，其它格式无效；,时间戳格式是按小时休假日程，日期格式是全天休假日程；,start_time与end_time格式需保持一致，否则无效。, 示例值："2021-01-01"
	EndTime     string  `json:"end_time,omitempty"`     // 休假结束时间：,有时间戳(1609430400)和日期(2021-01-01)两种格式，其它格式无效；,时间戳格式是按小时休假日程，日期格式是全天休假日程；,start_time与end_time格式需保持一致，否则无效。, 示例值："2021-01-01"
	Title       *string `json:"title,omitempty"`        // 自定义请假日程标题，没有设置则为默认日程标题, 示例值："请假中(全天) / 1-Day Time Off"
	Description *string `json:"description,omitempty"`  // 自定义请假日程描述，没有设置则为默认日程描述, 示例值："若删除此日程，飞书中相应的“请假”标签将自动消失，而请假系统中的休假申请不会被撤销。"
}

type createCalendarTimeoffEventResp struct {
	Code int                             `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                          `json:"msg,omitempty"`  // 错误描述
	Data *CreateCalendarTimeoffEventResp `json:"data,omitempty"` //
}

type CreateCalendarTimeoffEventResp struct {
	TimeoffEventID string `json:"timeoff_event_id,omitempty"` // 请假日程ID
	UserID         string `json:"user_id,omitempty"`          // 用户id
	Timezone       string `json:"timezone,omitempty"`         // 时区
	StartTime      string `json:"start_time,omitempty"`       // 休假开始时间: 有时间戳(1609430400)和日期(2021-01-01)两种格式，其它格式无效；,时间戳格式是按小时休假日程，日期格式是全天休假日程；,start_time与end_time格式需保持一致，否则无效。
	EndTime        string `json:"end_time,omitempty"`         // 休假结束时间：,有时间戳(1609430400)和日期(2021-01-01)两种格式，其它格式无效；,时间戳格式是按小时休假日程，日期格式是全天休假日程；,start_time与end_time格式需保持一致，否则无效。
	Title          string `json:"title,omitempty"`            // 自定义请假日程标题，没有设置则为默认日程标题
	Description    string `json:"description,omitempty"`      // 自定义请假日程描述，没有设置则为默认日程描述
}
