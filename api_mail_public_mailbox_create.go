package lark

import (
	"context"
)

// CreatePublicMailbox 创建一个公共邮箱
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/mail-v1/public_mailbox/create
func (r *MailAPI) CreatePublicMailbox(ctx context.Context, request *CreatePublicMailboxReq) (*CreatePublicMailboxResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/mail/v1/public_mailboxes",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(createPublicMailboxResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("Mail", "CreatePublicMailbox", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type CreatePublicMailboxReq struct {
	Email *string `json:"email,omitempty"` // 公共邮箱地址, 示例值："test_public_mailbox@xxx.xx"
	Name  *string `json:"name,omitempty"`  // 公共邮箱名称, 示例值："test public mailbox"
}

type createPublicMailboxResp struct {
	Code int                      `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *CreatePublicMailboxResp `json:"data,omitempty"` //
}

type CreatePublicMailboxResp struct {
	PublicMailboxID string `json:"public_mailbox_id,omitempty"` // 公共邮箱唯一标识
	Email           string `json:"email,omitempty"`             // 公共邮箱地址
	Name            string `json:"name,omitempty"`              // 公共邮箱名称
}