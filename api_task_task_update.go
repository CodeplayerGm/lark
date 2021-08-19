// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// UpdateTask 该接口用于修改任务的标题、描述、时间、来源等相关信息
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/patch
func (r *TaskService) UpdateTask(ctx context.Context, request *UpdateTaskReq, options ...MethodOptionFunc) (*UpdateTaskResp, *Response, error) {
	if r.cli.mock.mockTaskUpdateTask != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Task#UpdateTask mock enable")
		return r.cli.mock.mockTaskUpdateTask(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Task",
		API:                   "UpdateTask",
		Method:                "PATCH",
		URL:                   "https://open.feishu.cn/open-apis/task/v1/tasks/:task_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateTaskResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockTaskUpdateTask(f func(ctx context.Context, request *UpdateTaskReq, options ...MethodOptionFunc) (*UpdateTaskResp, *Response, error)) {
	r.mockTaskUpdateTask = f
}

func (r *Mock) UnMockTaskUpdateTask() {
	r.mockTaskUpdateTask = nil
}

type UpdateTaskReq struct {
	UserIDType   *IDType            `query:"user_id_type" json:"-"`  // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`,, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	TaskID       string             `path:"task_id" json:"-"`        // 任务 ID, 示例值："83912691-2e43-47fc-94a4-d512e03984fa"
	Task         *UpdateTaskReqTask `json:"task,omitempty"`          // 被更新的任务实体基础信息
	UpdateFields []string           `json:"update_fields,omitempty"` // 指定需要更新的字段（目前可选更新的字段为：summary, description, due, extra），否则服务端将不知道更新哪些字段, 示例值：["summary"]
}

type UpdateTaskReqTask struct {
	Summary     *string                  `json:"summary,omitempty"`     // 任务标题。创建任务时，如果没有标题填充，飞书服务器会将其视为无主题的任务, 示例值："每天喝八杯水，保持身心愉悦", 长度范围：`1` ～ `256` 字符
	Description *string                  `json:"description,omitempty"` // 任务备注, 示例值："多吃水果，多运动，健康生活，快乐工作。", 长度范围：`0` ～ `65536` 字符
	Extra       *string                  `json:"extra,omitempty"`       // 接入方可以自定义的附属信息二进制格式，采用 base64 编码，解析方式由接入方自己决定, 示例值："dGVzdA==", 长度范围：`0` ～ `65536` 字符
	Due         *UpdateTaskReqTaskDue    `json:"due,omitempty"`         // 任务的截止时间设置
	Origin      *UpdateTaskReqTaskOrigin `json:"origin,omitempty"`      // 任务关联的第三方平台来源信息
}

type UpdateTaskReqTaskDue struct {
	Time     *string `json:"time,omitempty"`       // 截止时间的时间戳（单位为秒）, 示例值："1623124318"
	Timezone *string `json:"timezone,omitempty"`   // 截止时间对应的时区，完整的时区名称列表可参考：https://docs.aws.amazon.com/zh_cn/redshift/latest/dg/time-zone-names.html, 示例值："Asia/Shanghai", 默认值: `Asia/Shanghai`
	IsAllDay *bool   `json:"is_all_day,omitempty"` // 标记任务是否为全天任务（全天任务的截止时间为当天 UTC 时间的 0 点）, 示例值：false, 默认值: `false`
}

type UpdateTaskReqTaskOrigin struct {
	PlatformI18nName string                       `json:"platform_i18n_name,omitempty"` // 任务导入来源的名称，用于在任务中心详情页展示。请提供一个字典，多种语言名称映射。支持的各地区语言名：it_it, th_th, ko_kr, es_es, ja_jp, zh_cn, id_id, zh_hk, pt_br, de_de, fr_fr, zh_tw, ru_ru, en_us, hi_in, vi_vn, 示例值：""{\"zh_cn\": \"IT 工作台\", \"en_us\": \"IT Workspace\"}"", 长度范围：`0` ～ `1024` 字符
	Href             *UpdateTaskReqTaskOriginHref `json:"href,omitempty"`               // 任务关联的来源平台详情页链接
}

type UpdateTaskReqTaskOriginHref struct {
	URL   *string `json:"url,omitempty"`   // 具体链接地址, 示例值："https://support.feishu.com/internal/foo-bar", 长度范围：`0` ～ `1024` 字符
	Title *string `json:"title,omitempty"` // 链接对应的标题, 示例值："反馈一个问题，需要协助排查", 长度范围：`0` ～ `512` 字符
}

type updateTaskResp struct {
	Code int64           `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 错误描述
	Data *UpdateTaskResp `json:"data,omitempty"`
}

type UpdateTaskResp struct {
	Task *UpdateTaskRespTask `json:"task,omitempty"` // 返回修改后的任务详情
}

type UpdateTaskRespTask struct {
	ID           string                    `json:"id,omitempty"`            // 任务 ID，由飞书任务服务器发号
	Summary      string                    `json:"summary,omitempty"`       // 任务标题。创建任务时，如果没有标题填充，飞书服务器会将其视为无主题的任务, 长度范围：`1` ～ `256` 字符
	Description  string                    `json:"description,omitempty"`   // 任务备注, 长度范围：`0` ～ `65536` 字符
	CompleteTime string                    `json:"complete_time,omitempty"` // 任务的完成时间戳（单位为秒），如果完成时间为 0，则表示任务尚未完成
	CreatorID    string                    `json:"creator_id,omitempty"`    // 任务的创建者 ID。在创建任务时无需填充该字段
	Extra        string                    `json:"extra,omitempty"`         // 接入方可以自定义的附属信息二进制格式，采用 base64 编码，解析方式由接入方自己决定, 长度范围：`0` ～ `65536` 字符
	CreateTime   string                    `json:"create_time,omitempty"`   // 任务的创建时间戳（单位为秒）
	UpdateTime   string                    `json:"update_time,omitempty"`   // 任务的更新时间戳（单位为秒）
	Due          *UpdateTaskRespTaskDue    `json:"due,omitempty"`           // 任务的截止时间设置
	Origin       *UpdateTaskRespTaskOrigin `json:"origin,omitempty"`        // 任务关联的第三方平台来源信息
}

type UpdateTaskRespTaskDue struct {
	Time     string `json:"time,omitempty"`       // 截止时间的时间戳（单位为秒）
	Timezone string `json:"timezone,omitempty"`   // 截止时间对应的时区，完整的时区名称列表可参考：https://docs.aws.amazon.com/zh_cn/redshift/latest/dg/time-zone-names.html, 默认值: `Asia/Shanghai`
	IsAllDay bool   `json:"is_all_day,omitempty"` // 标记任务是否为全天任务（全天任务的截止时间为当天 UTC 时间的 0 点）, 默认值: `false`
}

type UpdateTaskRespTaskOrigin struct {
	PlatformI18nName string                        `json:"platform_i18n_name,omitempty"` // 任务导入来源的名称，用于在任务中心详情页展示。请提供一个字典，多种语言名称映射。支持的各地区语言名：it_it, th_th, ko_kr, es_es, ja_jp, zh_cn, id_id, zh_hk, pt_br, de_de, fr_fr, zh_tw, ru_ru, en_us, hi_in, vi_vn, 长度范围：`0` ～ `1024` 字符
	Href             *UpdateTaskRespTaskOriginHref `json:"href,omitempty"`               // 任务关联的来源平台详情页链接
}

type UpdateTaskRespTaskOriginHref struct {
	URL   string `json:"url,omitempty"`   // 具体链接地址, 长度范围：`0` ～ `1024` 字符
	Title string `json:"title,omitempty"` // 链接对应的标题, 长度范围：`0` ～ `512` 字符
}
