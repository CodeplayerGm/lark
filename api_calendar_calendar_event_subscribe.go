package lark

import (
	"context"
)

// SubscribeCalendarEvent 该接口用于以用户身份订阅指定日历下的日程变更事件。
//
// 用户必须对日历有访问权限。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event/subscription
func (r *CalendarAPI) SubscribeCalendarEvent(ctx context.Context, request *SubscribeCalendarEventReq) (*SubscribeCalendarEventResp, *Response, error) {
	req := &requestParam{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/events/subscription",
		Body:                  request,
		NeedTenantAccessToken: false,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(subscribeCalendarEventResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Calendar", "SubscribeCalendarEvent", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type SubscribeCalendarEventReq struct {
	CalendarID string `path:"calendar_id" json:"-"` // 日历ID,**示例值**："feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
}

type subscribeCalendarEventResp struct {
	Code int                         `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *SubscribeCalendarEventResp `json:"data,omitempty"`
}

type SubscribeCalendarEventResp struct{}
