// Code generated by lark_sdk_gen. DO NOT EDIT.
/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lark

import (
	"context"
)

// EventV1ReceiveMessage 为了更好地提升该事件的安全性, 我们对其进行了升级, 请尽快迁移至[新版本>>](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/message/events/receive)
//
// 当用户发送消息给机器人或在群聊中@机器人时触发此事件。
// - 依赖权限: [获取用户发给机器人的私聊消息] 、 [获取群聊中用户 @ 机器人的消息]。开启了相应的权限才能获取到相应的消息。
// - 其他条件: 应用必须开启了[机器人能力](https://open.feishu.cn/document/home/develop-a-bot-in-5-minutes/create-an-app)。
// ## 文本消息
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/ugzMugzMugzM/event/receive-message
func (r *EventCallbackService) HandlerEventV1ReceiveMessage(f EventV1ReceiveMessageHandler) {
	r.cli.eventHandler.eventV1ReceiveMessageHandler = f
}

// EventV1ReceiveMessageHandler event EventV1ReceiveMessage handler
type EventV1ReceiveMessageHandler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV1, event *EventV1ReceiveMessage) (string, error)

// EventV1ReceiveMessage ...
type EventV1ReceiveMessage struct {
	Type             string   `json:"type,omitempty"`       // 事件类型. 如: message
	AppID            string   `json:"app_id,omitempty"`     // 如: cli_xxx
	TenantKey        string   `json:"tenant_key,omitempty"` // 企业标识. 如: xxx
	RootID           string   `json:"root_id,omitempty"`
	ParentID         string   `json:"parent_id,omitempty"`
	OpenChatID       string   `json:"open_chat_id,omitempty"`        // 如: oc_5ce6d572455d361153b7cb51da133945
	ChatType         ChatType `json:"chat_type,omitempty"`           // 私聊private, 群聊group. 如: private
	MsgType          MsgType  `json:"msg_type,omitempty"`            // 消息类型. 如: text
	OpenID           string   `json:"open_id,omitempty"`             // 如: ou_18eac85d35a26f989317ad4f02e8bbbb
	EmployeeID       string   `json:"employee_id,omitempty"`         // 即“用户ID”, 仅企业自建应用会返回. 如: xxx
	UnionID          string   `json:"union_id,omitempty"`            // 如: xxx
	OpenMessageID    string   `json:"open_message_id,omitempty"`     // 如: om_36686ee62209da697d8775375d0c8e88
	IsMention        bool     `json:"is_mention,omitempty"`          // 如: false
	Text             string   `json:"text,omitempty"`                // 消息文本, 可能包含被@的人/机器人。. 如: <at </at> 消息内容 <at </at>
	TextWithoutAtBot string   `json:"text_without_at_bot,omitempty"` // 消息内容, 会过滤掉at你的机器人的内容, 当内容只有at机器人, 该字段会被过滤。. 如: 消息内容 <at </at>
	Title            string   `json:"title,omitempty"`               // 富文本标题
	ImageKeys        []string `json:"image_keys,omitempty"`          // 富文本里面的图片的keys
	ImageKey         string   `json:"image_key,omitempty"`           // 图片内容
	FileKey          string   `json:"file_key,omitempty"`            // 文件内容
}