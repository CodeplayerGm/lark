package lark

import (
	"context"
)

// EventV1AddBot
//
// ## 机器人进群
// 机器人被邀请加入群聊时触发此事件。
// - 依赖条件：应用必须开启了[机器人能力](/ssl:ttdoc/uQjL04CN/uYTMuYTMuYTM)。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMTNxYjLzUTM24yM1EjN
func (r *EventCallbackAPI) HandlerEventV1AddBot(f eventV1AddBotHandler) {
	r.cli.eventHandler.eventV1AddBotHandler = f
}

type eventV1AddBotHandler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV1, event *EventV1AddBot) (string, error)

type EventV1AddBot struct {
	AppID               string                      `json:"app_id,omitempty"`                 // 如: cli_9c8609450f78d102
	ChatI18nNames       *EventV1AddBotChatI18nNames `json:"chat_i18n_names,omitempty"`        // 群名称国际化字段
	ChatName            string                      `json:"chat_name,omitempty"`              // 如: 群名称
	ChatOwnerEmployeeID string                      `json:"chat_owner_employee_id,omitempty"` // 群主的employee_id（即“用户ID”。如果群主是机器人则没有这个字段，仅企业自建应用返回）. 如: ca51d83b
	ChatOwnerName       string                      `json:"chat_owner_name,omitempty"`        // 群主姓名. 如: xxx
	ChatOwnerOpenID     string                      `json:"chat_owner_open_id,omitempty"`     // 群主的open_id. 如: ou_18eac85d35a26f989317ad4f02e8bbbb
	OpenChatID          string                      `json:"open_chat_id,omitempty"`           // 群聊的id. 如: oc_e520983d3e4f5ec7306069bffe672aa3
	OperatorEmployeeID  string                      `json:"operator_employee_id,omitempty"`   // 操作者的emplolyee_id ，仅企业自建应用返回. 如: ca51d83b
	OperatorName        string                      `json:"operator_name,omitempty"`          // 操作者姓名. 如: yyy
	OperatorOpenID      string                      `json:"operator_open_id,omitempty"`       // 操作者的open_id. 如: ou_18eac85d35a26f989317ad4f02e8bbbb
	OwnerIsBot          bool                        `json:"owner_is_bot,omitempty"`           // 群主是否是机器人
	TenantKey           string                      `json:"tenant_key,omitempty"`             // 企业标识. 如: 736588c9260f175d
	Type                string                      `json:"type,omitempty"`                   // 事件类型. 如: add_bot
}

type EventV1AddBotChatI18nNames struct {
	EnUs string `json:"en_us,omitempty"` // 如: 英文标题
	ZhCn string `json:"zh_cn,omitempty"` // 如: 中文标题
}
