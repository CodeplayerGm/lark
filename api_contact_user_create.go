package lark

import (
	"context"
)

// CreateUser 使用该接口向通讯录创建一个用户
//
// 新增用户的所有部门必须都在当前应用的通讯录授权范围内才允许新增用户，如果想要在根部门下新增用户，必须要有全员权限。 应用商店应用无权限调用此接口
// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/create
func (r *ContactAPI) CreateUser(ctx context.Context, request *CreateUserReq) (*CreateUserResp, *Response, error) {
	req := &requestParam{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/contact/v3/users",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(createUserResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Contact", "CreateUser", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type CreateUserReq struct {
	UserIDType           *IDType                          `query:"user_id_type" json:"-"`           // 用户 ID 类型,**示例值**："open_id",**可选值有**：,- `open_id`：用户的 open id,- `union_id`：用户的 union id,- `user_id`：用户的 user id,**默认值**：`open_id`,**当值为 `user_id`，字段权限要求**：,<md-perm href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">获取用户 userid</md-perm>
	DepartmentIDType     *IDType                          `query:"department_id_type" json:"-"`     // 此次调用中使用的部门ID的类型,**示例值**："open_department_id",**可选值有**：,- `department_id`：以自定义department_id来标识部门,- `open_department_id`：以open_department_id来标识部门,**默认值**：`open_department_id`
	ClientToken          *string                          `query:"client_token" json:"-"`           // 根据client_token是否一致来判断是否为同一请求,**示例值**："xxxx-xxxxx-xxx"
	UserID               *string                          `json:"user_id,omitempty"`                // 租户内用户的唯一标识,**示例值**："ou_7dab8a3d3cdcc9da365777c7ad535d62",**字段权限要求**：, <md-perm href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">获取用户 userid</md-perm>
	Name                 string                           `json:"name,omitempty"`                   // 用户名,**示例值**："张三",**数据校验规则**：,- 最小长度：`1` 字符
	EnName               *string                          `json:"en_name,omitempty"`                // 英文名,**示例值**："San Zhang"
	Email                *string                          `json:"email,omitempty"`                  // 邮箱,**示例值**："zhangsan@gmail.com",**字段权限要求**：, <md-perm href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">获取用户邮箱</md-perm>
	Mobile               string                           `json:"mobile,omitempty"`                 // 手机号,**示例值**："13011111111",**字段权限要求**：, <md-perm href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">获取用户手机号</md-perm>
	MobileVisible        *bool                            `json:"mobile_visible,omitempty"`         // 手机号码可见性，true 为可见，false 为不可见，目前默认为 true。不可见时，组织员工将无法查看该员工的手机号码,**示例值**：false
	Gender               *int                             `json:"gender,omitempty"`                 // 性别,**示例值**：1,**可选值有**：,- `0`：保密,- `1`：男,- `2`：女
	AvatarKey            *string                          `json:"avatar_key,omitempty"`             // 头像的文件Key,**示例值**："2500c7a9-5fff-4d9a-a2de-3d59614ae28g"
	DepartmentIDs        []string                         `json:"department_ids,omitempty"`         // 用户所属部门的ID列表
	LeaderUserID         *string                          `json:"leader_user_id,omitempty"`         // 用户的直接主管的用户ID,**示例值**："ou_7dab8a3d3cdcc9da365777c7ad535d62"
	City                 *string                          `json:"city,omitempty"`                   // 城市,**示例值**："杭州"
	Country              *string                          `json:"country,omitempty"`                // 国家,**示例值**："中国"
	WorkStation          *string                          `json:"work_station,omitempty"`           // 工位,**示例值**："杭州"
	JoinTime             *int                             `json:"join_time,omitempty"`              // 入职时间,**示例值**：2147483647
	EmployeeNo           *string                          `json:"employee_no,omitempty"`            // 工号,**示例值**："1"
	EmployeeType         int                              `json:"employee_type,omitempty"`          // 员工类型,**示例值**：1
	Orders               []*CreateUserReqOrder            `json:"orders,omitempty"`                 // 用户排序信息
	CustomAttrs          []*CreateUserReqCustomAttr       `json:"custom_attrs,omitempty"`           // 自定义属性
	EnterpriseEmail      *string                          `json:"enterprise_email,omitempty"`       // 企业邮箱，请先确保已在管理后台启用飞书邮箱服务,**示例值**："demo@mail.com"
	NeedSendNotification *bool                            `json:"need_send_notification,omitempty"` // 是否发送提示消息,**示例值**：false
	NotificationOption   *CreateUserReqNotificationOption `json:"notification_option,omitempty"`    // 创建用户的邀请方式
}

type CreateUserReqOrder struct {
	DepartmentID    *string `json:"department_id,omitempty"`    // 排序信息对应的部门ID,**示例值**："od-4e6ac4d14bcd5071a37a39de902c7141"
	UserOrder       *int    `json:"user_order,omitempty"`       // 用户在部门内的排序,**示例值**：100
	DepartmentOrder *int    `json:"department_order,omitempty"` // 用户的部门间的排序,**示例值**：100
}

type CreateUserReqCustomAttr struct {
	Type  *string                       `json:"type,omitempty"`  // 自定义属性类型,**示例值**："TEXT"
	ID    *string                       `json:"id,omitempty"`    // 自定义属性ID,**示例值**："DemoId"
	Value *CreateUserReqCustomAttrValue `json:"value,omitempty"` // 自定义属性取值
}

type CreateUserReqCustomAttrValue struct {
	Text  *string `json:"text,omitempty"`   // 属性文本,**示例值**："DemoText"
	Url   *string `json:"url,omitempty"`    // URL,**示例值**："http://www.feishu.cn"
	PcUrl *string `json:"pc_url,omitempty"` // PC上的URL,**示例值**："http://www.feishu.cn"
}

type CreateUserReqNotificationOption struct {
	Channels []string `json:"channels,omitempty"` // 通道列表，枚举值：,sms（短信邀请），email（邮件邀请）
	Language *string  `json:"language,omitempty"` // 语言类型,**示例值**："zh-CN",**可选值有**：,- `zh-CN`：中文,- `en-US`：英文,- `ja-JP`：日文
}

type createUserResp struct {
	Code int             `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 错误描述
	Data *CreateUserResp `json:"data,omitempty"` // \-
}

type CreateUserResp struct {
	User *CreateUserRespUser `json:"user,omitempty"` // 用户信息
}

type CreateUserRespUser struct {
	UnionID              string                                `json:"union_id,omitempty"`               // 用户的union_id
	UserID               string                                `json:"user_id,omitempty"`                // 租户内用户的唯一标识,**字段权限要求**：, <md-perm href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">获取用户 userid</md-perm>
	OpenID               string                                `json:"open_id,omitempty"`                // 用户的open_id
	Name                 string                                `json:"name,omitempty"`                   // 用户名,**数据校验规则**：,- 最小长度：`1` 字符
	EnName               string                                `json:"en_name,omitempty"`                // 英文名
	Email                string                                `json:"email,omitempty"`                  // 邮箱,**字段权限要求**：, <md-perm href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">获取用户邮箱</md-perm>
	Mobile               string                                `json:"mobile,omitempty"`                 // 手机号,**字段权限要求**：, <md-perm href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">获取用户手机号</md-perm>
	MobileVisible        bool                                  `json:"mobile_visible,omitempty"`         // 手机号码可见性，true 为可见，false 为不可见，目前默认为 true。不可见时，组织员工将无法查看该员工的手机号码
	Gender               int                                   `json:"gender,omitempty"`                 // 性别,**可选值有**：,- `0`：保密,- `1`：男,- `2`：女
	AvatarKey            string                                `json:"avatar_key,omitempty"`             // 头像的文件Key
	Avatar               *CreateUserRespUserAvatar             `json:"avatar,omitempty"`                 // 用户头像信息
	Status               *CreateUserRespUserStatus             `json:"status,omitempty"`                 // 用户状态
	DepartmentIDs        []string                              `json:"department_ids,omitempty"`         // 用户所属部门的ID列表
	LeaderUserID         string                                `json:"leader_user_id,omitempty"`         // 用户的直接主管的用户ID
	City                 string                                `json:"city,omitempty"`                   // 城市
	Country              string                                `json:"country,omitempty"`                // 国家
	WorkStation          string                                `json:"work_station,omitempty"`           // 工位
	JoinTime             int                                   `json:"join_time,omitempty"`              // 入职时间
	IsTenantManager      bool                                  `json:"is_tenant_manager,omitempty"`      // 是否是租户管理员
	EmployeeNo           string                                `json:"employee_no,omitempty"`            // 工号
	EmployeeType         int                                   `json:"employee_type,omitempty"`          // 员工类型,**可选值有**：,- `1`：正式员工,- `2`：实习生,- `3`：外包,- `4`：劳务,- `5`：顾问
	Orders               []*CreateUserRespUserOrder            `json:"orders,omitempty"`                 // 用户排序信息
	CustomAttrs          []*CreateUserRespUserCustomAttr       `json:"custom_attrs,omitempty"`           // 自定义属性
	EnterpriseEmail      string                                `json:"enterprise_email,omitempty"`       // 企业邮箱，请先确保已在管理后台启用飞书邮箱服务
	NeedSendNotification bool                                  `json:"need_send_notification,omitempty"` // 是否发送提示消息
	NotificationOption   *CreateUserRespUserNotificationOption `json:"notification_option,omitempty"`    // 创建用户的邀请方式
}

type CreateUserRespUserAvatar struct {
	Avatar72     string `json:"avatar_72,omitempty"`     // 72*72像素头像链接
	Avatar240    string `json:"avatar_240,omitempty"`    // 240*240像素头像链接
	Avatar640    string `json:"avatar_640,omitempty"`    // 640*640像素头像链接
	AvatarOrigin string `json:"avatar_origin,omitempty"` // 原始头像链接
}

type CreateUserRespUserStatus struct {
	IsFrozen    bool `json:"is_frozen,omitempty"`    // 是否冻结
	IsResigned  bool `json:"is_resigned,omitempty"`  // 是否离职
	IsActivated bool `json:"is_activated,omitempty"` // 是否激活
}

type CreateUserRespUserOrder struct {
	DepartmentID    string `json:"department_id,omitempty"`    // 排序信息对应的部门ID
	UserOrder       int    `json:"user_order,omitempty"`       // 用户在部门内的排序
	DepartmentOrder int    `json:"department_order,omitempty"` // 用户的部门间的排序
}

type CreateUserRespUserCustomAttr struct {
	Type  string                             `json:"type,omitempty"`  // 自定义属性类型
	ID    string                             `json:"id,omitempty"`    // 自定义属性ID
	Value *CreateUserRespUserCustomAttrValue `json:"value,omitempty"` // 自定义属性取值
}

type CreateUserRespUserCustomAttrValue struct {
	Text  string `json:"text,omitempty"`   // 属性文本
	Url   string `json:"url,omitempty"`    // URL
	PcUrl string `json:"pc_url,omitempty"` // PC上的URL
}

type CreateUserRespUserNotificationOption struct {
	Channels []string `json:"channels,omitempty"` // 通道列表，枚举值：,sms（短信邀请），email（邮件邀请）
	Language string   `json:"language,omitempty"` // 语言类型,**可选值有**：,- `zh-CN`：中文,- `en-US`：英文,- `ja-JP`：日文
}
