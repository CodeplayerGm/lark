package lark

import (
	"context"
)

// GetPublicMailboxList 分页批量获取公共邮箱列表
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox/list
func (r *MailAPI) GetPublicMailboxList(ctx context.Context, request *GetPublicMailboxListReq) (*GetPublicMailboxListResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/mail/v1/public_mailboxes",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(getPublicMailboxListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("Mail", "GetPublicMailboxList", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetPublicMailboxListReq struct {
	PageToken *string `query:"page_token" json:"-"` // 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果, 示例值："xxx"
	PageSize  *int    `query:"page_size" json:"-"`  // 分页大小, 示例值：10, 最大值：`200`
}

type getPublicMailboxListResp struct {
	Code int                       `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                    `json:"msg,omitempty"`  // 错误描述
	Data *GetPublicMailboxListResp `json:"data,omitempty"` //
}

type GetPublicMailboxListResp struct {
	HasMore   bool                            `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                          `json:"page_token,omitempty"` // 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
	Items     []*GetPublicMailboxListRespItem `json:"items,omitempty"`      // 公共邮箱列表
}

type GetPublicMailboxListRespItem struct {
	PublicMailboxID string `json:"public_mailbox_id,omitempty"` // 公共邮箱唯一标识
	Email           string `json:"email,omitempty"`             // 公共邮箱地址
	Name            string `json:"name,omitempty"`              // 公共邮箱名称
}
