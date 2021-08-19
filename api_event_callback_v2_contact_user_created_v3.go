// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// EventV2ContactUserCreatedV3
//
// 通过该事件订阅员工入职。应用身份访问通讯录的权限为历史版本，不推荐申请。{使用示例}(url=/api/tools/api_explore/api_explore_config?project=contact&version=v3&resource=user&event=created)
// <b>以应用身份访问通讯录</b> 权限为历史版本，不推荐申请。应用访问通讯录相关接口请申请 <b>以应用身份读取通讯录</b>
// 只有当应用拥有被改动字段的数据权限时，才会接收到事件。具体的数据权限与字段的关系请参考[应用权限](https://open.feishu.cn/document/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN)，或查看事件体参数列表的字段描述。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/events/created
func (r *EventCallbackService) HandlerEventV2ContactUserCreatedV3(f eventV2ContactUserCreatedV3Handler) {
	r.cli.eventHandler.eventV2ContactUserCreatedV3Handler = f
}

type eventV2ContactUserCreatedV3Handler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV2, event *EventV2ContactUserCreatedV3) (string, error)

type EventV2ContactUserCreatedV3 struct {
	Object *EventV2ContactUserCreatedV3Object `json:"object,omitempty"` // 事件信息
}

type EventV2ContactUserCreatedV3Object struct {
	OpenID        string                                         `json:"open_id,omitempty"`        // 用户的open_id
	UserID        string                                         `json:"user_id,omitempty"`        // 租户内用户的唯一标识, 字段权限要求:  获取用户 user ID
	Name          string                                         `json:"name,omitempty"`           // 用户名, 最小长度：`1` 字符,**字段权限要求（满足任一）**：, 获取用户基本信息, 以应用身份访问通讯录（历史版本）
	EnName        string                                         `json:"en_name,omitempty"`        // 英文名,**字段权限要求（满足任一）**：, 获取用户基本信息, 以应用身份访问通讯录（历史版本）
	Email         string                                         `json:"email,omitempty"`          // 邮箱, 字段权限要求:  获取用户邮箱信息
	Mobile        string                                         `json:"mobile,omitempty"`         // 手机号, 字段权限要求:  获取用户手机号
	Gender        int64                                          `json:"gender,omitempty"`         // 性别, 可选值有: `0`：未知, `1`：男, `2`：女,**字段权限要求（满足任一）**：, 获取用户性别, 以应用身份访问通讯录（历史版本）
	Avatar        *EventV2ContactUserCreatedV3ObjectAvatar       `json:"avatar,omitempty"`         // 用户头像信息,**字段权限要求（满足任一）**：, 获取用户基本信息, 以应用身份访问通讯录（历史版本）
	Status        *EventV2ContactUserCreatedV3ObjectStatus       `json:"status,omitempty"`         // 用户状态,**字段权限要求（满足任一）**：, 获取用户雇佣信息, 以应用身份访问通讯录（历史版本）
	DepartmentIDs []string                                       `json:"department_ids,omitempty"` // 用户所属部门的ID列表,**字段权限要求（满足任一）**：, 获取用户组织架构信息, 以应用身份访问通讯录（历史版本）
	LeaderUserID  string                                         `json:"leader_user_id,omitempty"` // 用户的直接主管的用户ID,**字段权限要求（满足任一）**：, 获取用户雇佣信息, 以应用身份访问通讯录（历史版本）
	City          string                                         `json:"city,omitempty"`           // 城市,**字段权限要求（满足任一）**：, 获取用户雇佣信息, 以应用身份访问通讯录（历史版本）
	Country       string                                         `json:"country,omitempty"`        // 国家,**字段权限要求（满足任一）**：, 获取用户雇佣信息, 以应用身份访问通讯录（历史版本）
	WorkStation   string                                         `json:"work_station,omitempty"`   // 工位,**字段权限要求（满足任一）**：, 获取用户雇佣信息, 以应用身份访问通讯录（历史版本）
	JoinTime      int64                                          `json:"join_time,omitempty"`      // 入职时间, 取值范围：`1` ～ `2147483647`,**字段权限要求（满足任一）**：, 获取用户雇佣信息, 以应用身份访问通讯录（历史版本）
	EmployeeNo    string                                         `json:"employee_no,omitempty"`    // 工号,**字段权限要求（满足任一）**：, 获取用户雇佣信息, 以应用身份访问通讯录（历史版本）
	EmployeeType  int64                                          `json:"employee_type,omitempty"`  // 员工类型, 可选值有: `1`：正式员工, `2`：实习生, `3`：外包, `4`：劳务, `5`：顾问,**字段权限要求（满足任一）**：, 获取用户雇佣信息, 以应用身份访问通讯录（历史版本）
	Orders        []*EventV2ContactUserCreatedV3ObjectOrder      `json:"orders,omitempty"`         // 用户排序信息,**字段权限要求（满足任一）**：, 获取用户组织架构信息, 以应用身份访问通讯录（历史版本）
	CustomAttrs   []*EventV2ContactUserCreatedV3ObjectCustomAttr `json:"custom_attrs,omitempty"`   // 自定义属性,**字段权限要求（满足任一）**：, 获取用户雇佣信息, 以应用身份访问通讯录（历史版本）
}

type EventV2ContactUserCreatedV3ObjectAvatar struct {
	Avatar72     string `json:"avatar_72,omitempty"`     // 72*72像素头像链接
	Avatar240    string `json:"avatar_240,omitempty"`    // 240*240像素头像链接
	Avatar640    string `json:"avatar_640,omitempty"`    // 640*640像素头像链接
	AvatarOrigin string `json:"avatar_origin,omitempty"` // 原始头像链接
}

type EventV2ContactUserCreatedV3ObjectStatus struct {
	IsFrozen    bool `json:"is_frozen,omitempty"`    // 是否暂停
	IsResigned  bool `json:"is_resigned,omitempty"`  // 是否离职
	IsActivated bool `json:"is_activated,omitempty"` // 是否激活
}

type EventV2ContactUserCreatedV3ObjectOrder struct {
	DepartmentID    string `json:"department_id,omitempty"`    // 排序信息对应的部门ID
	UserOrder       int64  `json:"user_order,omitempty"`       // 用户在其直属部门内的排序，数值越大，排序越靠前
	DepartmentOrder int64  `json:"department_order,omitempty"` // 用户所属的多个部门间的排序，数值越大，排序越靠前
}

type EventV2ContactUserCreatedV3ObjectCustomAttr struct {
	Type  string                                            `json:"type,omitempty"`  // 自定义属性类型
	ID    string                                            `json:"id,omitempty"`    // 自定义属性ID
	Value *EventV2ContactUserCreatedV3ObjectCustomAttrValue `json:"value,omitempty"` // 自定义属性取值
}

type EventV2ContactUserCreatedV3ObjectCustomAttrValue struct {
	Text  string `json:"text,omitempty"`   // 属性文本
	URL   string `json:"url,omitempty"`    // URL
	PcURL string `json:"pc_url,omitempty"` // PC上的URL
}
