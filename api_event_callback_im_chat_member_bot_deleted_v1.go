package lark

import (
	"context"
)

// EventIMChatMemberBotDeletedV1
//
// 机器人被移出群聊时触发此事件。
// 注意事项：
// - 需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// - 需要订阅 ==即时通讯== 分类下的 ==机器人被移出群== 事件
// - 事件会向被移出群的机器人进行推送
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-member-bot/events/deleted
func (r *EventCallbackAPI) HandlerEventIMChatMemberBotDeletedV1(f eventIMChatMemberBotDeletedV1Handler) {
	r.cli.eventHandler.eventIMChatMemberBotDeletedV1Handler = f
}

type eventIMChatMemberBotDeletedV1Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeader, event *EventIMChatMemberBotDeletedV1) (string, error)

type EventIMChatMemberBotDeletedV1 struct {
	ChatID     string                                   `json:"chat_id,omitempty"`     // 群组 ID
	OperatorID *EventIMChatMemberBotDeletedV1OperatorID `json:"operator_id,omitempty"` // 用户 ID
}

type EventIMChatMemberBotDeletedV1OperatorID struct {
	UnionID string `json:"union_id,omitempty"` // 用户的 union id
	UserID  string `json:"user_id,omitempty"`  // 用户的 user id,**字段权限要求**：,<md-perm href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">获取用户 userid</md-perm>
	OpenID  string `json:"open_id,omitempty"`  // 用户的 open id
}
