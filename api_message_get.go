package lark

import (
	"context"
)

// GetMessage 通过 message_id 查询消息内容
//
// 注意事项:
// - 需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 机器人必须在群组中
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/get
func (r *MessageAPI) GetMessage(ctx context.Context, request *GetMessageReq) (*GetMessageResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/im/v1/messages/:message_id",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(getMessageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("Message", "GetMessage", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetMessageReq struct {
	MessageID string `path:"message_id" json:"-"` // 待获取消息内容的消息的ID, 示例值："om_dc13264520392913993dd051dba21dcf"
}

type getMessageResp struct {
	Code int             `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 错误描述
	Data *GetMessageResp `json:"data,omitempty"` //
}

type GetMessageResp struct {
	Items []*GetMessageRespItem `json:"items,omitempty"` // -
}

type GetMessageRespItem struct {
	MessageID      string       `json:"message_id,omitempty"`       // 消息id open_message_id
	RootID         string       `json:"root_id,omitempty"`          // 根消息id open_message_id
	ParentID       string       `json:"parent_id,omitempty"`        // 父消息的id open_message_id
	MsgType        MsgType      `json:"msg_type,omitempty"`         // 消息类型 text post card image等等
	CreateTime     string       `json:"create_time,omitempty"`      // 消息生成的时间戳(毫秒)
	UpdateTime     string       `json:"update_time,omitempty"`      // 消息更新的时间戳
	Deleted        bool         `json:"deleted,omitempty"`          // 消息是否被撤回
	Updated        bool         `json:"updated,omitempty"`          // 消息是否被更新
	ChatID         string       `json:"chat_id,omitempty"`          // 所属的群
	Sender         *Sender      `json:"sender,omitempty"`           // 发送者，可以是用户或应用
	Body           *MessageBody `json:"body,omitempty"`             // 消息内容，json结构，格式说明参考： [消息content说明](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/im-v1/message/events/message_content)
	Mentions       []*Mention   `json:"mentions,omitempty"`         // 被艾特的人或应用的id
	UpperMessageID string       `json:"upper_message_id,omitempty"` // 上一层级的消息id open_message_id
}
