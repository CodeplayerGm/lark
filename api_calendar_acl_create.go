package lark

import (
	"context"
)

// CreateCalendarACL
//
// 该接口用于以当前身份（应用 / 用户）给日历添加访问控制权限，即日历成员。
// 身份由 Header Authorization 的 Token 类型决定。
// 当前身份需要有日历的 owner 权限，并且日历的类型只能为 primary 或 shared。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-acl/create
func (r *CalendarAPI) CreateCalendarACL(ctx context.Context, request *CreateCalendarACLReq) (*CreateCalendarACLResp, *Response, error) {
	req := &requestParam{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/acls",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(createCalendarACLResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Calendar", "CreateCalendarACL", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type CreateCalendarACLReq struct {
	UserIDType *IDType                    `query:"user_id_type" json:"-"` // 用户 ID 类型,**示例值**："open_id",**可选值有**：,- `open_id`：用户的 open id,- `union_id`：用户的 union id,- `user_id`：用户的 user id,**默认值**：`open_id`,**当值为 `user_id`，字段权限要求**：,<md-perm href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">获取用户 userid</md-perm>
	CalendarID string                     `path:"calendar_id" json:"-"`   // 日历ID,**示例值**："feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
	Role       CalendarRole               `json:"role,omitempty"`         // 对日历的访问权限,**示例值**："writer",**可选值有**：,- `unknown`：未知权限,- `free_busy_reader`：游客，只能看到忙碌/空闲信息,- `reader`：订阅者，查看所有日程详情,- `writer`：编辑者，创建及修改日程,- `owner`：管理员，管理日历及共享设置
	Scope      *CreateCalendarACLReqScope `json:"scope,omitempty"`        // 权限范围
}

type CreateCalendarACLReqScope struct {
	Type   string  `json:"type,omitempty"`    // 权限类型，当type为User时，值为open_id/user_id/union_id,**示例值**："user",**可选值有**：,- `user`：用户
	UserID *string `json:"user_id,omitempty"` // 用户ID,**示例值**："ou_xxxxxx"
}

type createCalendarACLResp struct {
	Code int                    `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *CreateCalendarACLResp `json:"data,omitempty"` // \-
}

type CreateCalendarACLResp struct {
	ACLID string                      `json:"acl_id,omitempty"` // acl资源ID
	Role  CalendarRole                `json:"role,omitempty"`   // 对日历的访问权限,**可选值有**：,- `unknown`：未知权限,- `free_busy_reader`：游客，只能看到忙碌/空闲信息,- `reader`：订阅者，查看所有日程详情,- `writer`：编辑者，创建及修改日程,- `owner`：管理员，管理日历及共享设置
	Scope *CreateCalendarACLRespScope `json:"scope,omitempty"`  // 权限范围
}

type CreateCalendarACLRespScope struct {
	Type   string `json:"type,omitempty"`    // 权限类型，当type为User时，值为open_id/user_id/union_id,**可选值有**：,- `user`：用户
	UserID string `json:"user_id,omitempty"` // 用户ID
}
