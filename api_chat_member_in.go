package lark

import (
	"context"
)

// IsInChat 判断用户或者机器人是否在群里。
//
//
// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-members/is_in_chat
func (r *ChatAPI) IsInChat(ctx context.Context, request *IsInChatReq) (*IsInChatResp, *Response, error) {
	req := &requestParam{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/im/v1/chats/:chat_id/members/is_in_chat",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
	}
	resp := new(isInChatResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Chat", "IsInChat", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type IsInChatReq struct {
	ChatID string `path:"chat_id" json:"-"` // 群 ID,**示例值：** "oc_a0553eda9014c201e6969b478895c230"
}

type isInChatResp struct {
	Code int           `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string        `json:"msg,omitempty"`  // 错误描述
	Data *IsInChatResp `json:"data,omitempty"` // \-
}

type IsInChatResp struct {
	IsInChat bool `json:"is_in_chat,omitempty"` // 用户或者机器人是否在群中
}
