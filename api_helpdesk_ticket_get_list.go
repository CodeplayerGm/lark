// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetHelpdeskTicketList 该接口用于获取全部工单详情。仅支持自建应用。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/list
func (r *HelpdeskService) GetHelpdeskTicketList(ctx context.Context, request *GetHelpdeskTicketListReq, options ...MethodOptionFunc) (*GetHelpdeskTicketListResp, *Response, error) {
	if r.cli.mock.mockHelpdeskGetHelpdeskTicketList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#GetHelpdeskTicketList mock enable")
		return r.cli.mock.mockHelpdeskGetHelpdeskTicketList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Helpdesk",
		API:                   "GetHelpdeskTicketList",
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/helpdesk/v1/tickets",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedHelpdeskAuth:      true,
	}
	resp := new(getHelpdeskTicketListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockHelpdeskGetHelpdeskTicketList(f func(ctx context.Context, request *GetHelpdeskTicketListReq, options ...MethodOptionFunc) (*GetHelpdeskTicketListResp, *Response, error)) {
	r.mockHelpdeskGetHelpdeskTicketList = f
}

func (r *Mock) UnMockHelpdeskGetHelpdeskTicketList() {
	r.mockHelpdeskGetHelpdeskTicketList = nil
}

type GetHelpdeskTicketListReq struct {
	TicketID         *string                    `query:"ticket_id" json:"-"`         // 搜索条件：工单ID, 示例值："123456"
	AgentID          *string                    `query:"agent_id" json:"-"`          // 搜索条件: 客服id, 示例值："ou_b5de90429xxx"
	ClosedByID       *string                    `query:"closed_by_id" json:"-"`      // 搜索条件: 关单客服id, 示例值："ou_b5de90429xxx"
	Type             *int64                     `query:"type" json:"-"`              // 搜索条件: 工单类型 1:bot 2:人工, 示例值：1
	Channel          *int64                     `query:"channel" json:"-"`           // 搜索条件: 工单渠道, 示例值：0
	Solved           *int64                     `query:"solved" json:"-"`            // 搜索条件: 工单是否解决 1:没解决 2:已解决	, 示例值：1
	Score            *int64                     `query:"score" json:"-"`             // 搜索条件: 工单评分, 示例值：1
	StatusList       []int64                    `query:"status_list" json:"-"`       // 搜索条件: 工单状态列表
	GuestName        *string                    `query:"guest_name" json:"-"`        // 搜索条件: 用户名称, 示例值："abc"
	GuestID          *string                    `query:"guest_id" json:"-"`          // 搜索条件: 用户id, 示例值："ou_b5de90429xxx"
	CustomizedFields []*HelpdeskCustomizedField `query:"customized_fields" json:"-"` // 搜索条件: 自定义字段列表
	Tags             []string                   `query:"tags" json:"-"`              // 搜索条件: 用户标签列表
	Page             *int64                     `query:"page" json:"-"`              // 页数, 从1开始, 默认为1, 示例值：1
	PageSize         *int64                     `query:"page_size" json:"-"`         // 当前页大小，最大为200, 默认为20, 示例值：20
	CreateTimeStart  *int64                     `query:"create_time_start" json:"-"` // 搜索条件: 工单创建起始时间 ms (也需要填上create_time_end), 示例值：1616920429000
	CreateTimeEnd    *int64                     `query:"create_time_end" json:"-"`   // 搜索条件: 工单创建结束时间 ms (也需要填上create_time_start), 示例值：1616920429000
	UpdateTimeStart  *int64                     `query:"update_time_start" json:"-"` // 搜索条件: 工单修改起始时间 ms (也需要填上update_time_end), 示例值：1616920429000
	UpdateTimeEnd    *int64                     `query:"update_time_end" json:"-"`   // 搜索条件: 工单修改结束时间 ms(也需要填上update_time_start), 示例值：1616920429000
}

type getHelpdeskTicketListResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *GetHelpdeskTicketListResp `json:"data,omitempty"`
}

type GetHelpdeskTicketListResp struct {
	Total   int64                              `json:"total,omitempty"`   // 工单总数 (单次请求最大为10000条)
	Tickets []*GetHelpdeskTicketListRespTicket `json:"tickets,omitempty"` // 工单
}

type GetHelpdeskTicketListRespTicket struct {
	TicketID         string                                            `json:"ticket_id,omitempty"`         // 工单ID
	HelpdeskID       string                                            `json:"helpdesk_id,omitempty"`       // 服务台ID
	Guest            *GetHelpdeskTicketListRespTicketGuest             `json:"guest,omitempty"`             // 工单创建用户
	Stage            int64                                             `json:"stage,omitempty"`             // 工单阶段，1：bot，2：人工
	Status           int64                                             `json:"status,omitempty"`            // 工单状态，1：已创建 2: 处理中 3: 排队中 4：待定 5：待用户响应 50: 被机器人关闭 51: 被人工关闭
	Score            int64                                             `json:"score,omitempty"`             // 工单评分，1：不满意，2:一般，3:满意
	CreatedAt        int64                                             `json:"created_at,omitempty"`        // 工单创建时间
	UpdatedAt        int64                                             `json:"updated_at,omitempty"`        // 工单更新时间，没有值时为-1
	ClosedAt         int64                                             `json:"closed_at,omitempty"`         // 工单结束时间
	Agents           []*GetHelpdeskTicketListRespTicketAgent           `json:"agents,omitempty"`            // 工单客服
	Channel          int64                                             `json:"channel,omitempty"`           // 工单渠道
	Solve            int64                                             `json:"solve,omitempty"`             // 工单是否解决 1:没解决 2:已解决
	ClosedBy         *GetHelpdeskTicketListRespTicketClosedBy          `json:"closed_by,omitempty"`         // 关单用户ID
	Collaborators    []*GetHelpdeskTicketListRespTicketCollaborator    `json:"collaborators,omitempty"`     // 工单协作者
	CustomizedFields []*GetHelpdeskTicketListRespTicketCustomizedField `json:"customized_fields,omitempty"` // 自定义字段列表，没有值时不设置
}

type GetHelpdeskTicketListRespTicketGuest struct {
	ID        string `json:"id,omitempty"`         // 用户ID
	AvatarURL string `json:"avatar_url,omitempty"` // 用户头像url
	Name      string `json:"name,omitempty"`       // 用户名
	Email     string `json:"email,omitempty"`      // 用户邮箱
}

type GetHelpdeskTicketListRespTicketAgent struct {
	ID        string `json:"id,omitempty"`         // 用户ID
	AvatarURL string `json:"avatar_url,omitempty"` // 用户头像url
	Name      string `json:"name,omitempty"`       // 用户名
	Email     string `json:"email,omitempty"`      // 用户邮箱
}

type GetHelpdeskTicketListRespTicketClosedBy struct {
	ID        string `json:"id,omitempty"`         // 用户ID
	AvatarURL string `json:"avatar_url,omitempty"` // 用户头像url
	Name      string `json:"name,omitempty"`       // 用户名
	Email     string `json:"email,omitempty"`      // 用户邮箱
}

type GetHelpdeskTicketListRespTicketCollaborator struct {
	ID        string `json:"id,omitempty"`         // 用户ID
	AvatarURL string `json:"avatar_url,omitempty"` // 用户头像url
	Name      string `json:"name,omitempty"`       // 用户名
	Email     string `json:"email,omitempty"`      // 用户邮箱
}

type GetHelpdeskTicketListRespTicketCustomizedField struct {
	ID          string `json:"id,omitempty"`           // 自定义字段ID
	Value       string `json:"value,omitempty"`        // 自定义字段值
	KeyName     string `json:"key_name,omitempty"`     // 键名
	DisplayName string `json:"display_name,omitempty"` // 展示名称
	Position    int64  `json:"position,omitempty"`     // 展示位置
	Required    bool   `json:"required,omitempty"`     // 是否必填
	Editable    bool   `json:"editable,omitempty"`     // 是否可修改
}
