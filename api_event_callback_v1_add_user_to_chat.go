// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// EventV1AddUserToChat
//
// ## 用户进群、出群事件
// 用户进群、出群后触发此事件。
// * 特殊说明：只有开启机器人能力并且机器人所在的群发生上述变化时才能触发此事件。
// 事件包括三个类型：
// 1. 用户进群 - add_user_to_chat
// 2. 用户出群 - remove_user_from_chat
// 3. 撤销加人 - revoke_add_user_from_chat
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/ukjNxYjL5YTM24SO2EjN
func (r *EventCallbackService) HandlerEventV1AddUserToChat(f eventV1AddUserToChatHandler) {
	r.cli.eventHandler.eventV1AddUserToChatHandler = f
}

type eventV1AddUserToChatHandler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV1, event *EventV1AddUserToChat) (string, error)

type EventV1AddUserToChat struct {
	AppID     string                             `json:"app_id,omitempty"`  // 如: cli_9c8609450f78d102
	ChatID    string                             `json:"chat_id,omitempty"` // 群聊的id. 如: oc_9e9619b938c9571c1c3165681cdaead5
	Operator  *EventV1AddUserToChatEventOperator `json:"operator,omitempty"`
	TenantKey string                             `json:"tenant_key,omitempty"` // 如: 736588c9260f175d
	Type      string                             `json:"type,omitempty"`       // 事件类型，add_user_to_chat/remove_user_from_chat/revoke_add_user_from_chat. 如: add_user_to_chat
	Users     []*EventV1AddUserToChatEventUser   `json:"users,omitempty"`
}

type EventV1AddUserToChatEventOperator struct {
	OpenID string `json:"open_id,omitempty"` // 员工对此应用的唯一标识，同一员工对不同应用的open_id不同. 如: ou_18eac85d35a26f989317ad4f02e8bbbb
	UserID string `json:"user_id,omitempty"` // 即“用户ID”，仅企业自建应用会返回. 如: ca51d83b
}

type EventV1AddUserToChatEventUser struct {
	Name   string `json:"name,omitempty"`    // 如: James
	OpenID string `json:"open_id,omitempty"` // 如: ou_706adeb944ab1473b9fb3e7da2a40b68
	UserID string `json:"user_id,omitempty"` // 如: 51g97a4g
}
