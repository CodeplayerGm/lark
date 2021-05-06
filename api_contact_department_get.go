package lark

import (
	"context"
)

// GetDepartment 该接口用于向通讯录获取单个部门信息。
//
// 使用tenant_access_token时，应用需要拥有待查询部门的通讯录授权。如果需要获取根部门信息，则需要拥有全员权限。
// 使用user_access_token时，用户需要有待查询部门的可见性，如果需要获取根部门信息，则要求员工可见所有人。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/department/get
func (r *ContactAPI) GetDepartment(ctx context.Context, request *GetDepartmentReq) (*GetDepartmentResp, *Response, error) {
	req := &requestParam{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/contact/v3/departments/:department_id",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(getDepartmentResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Contact", "GetDepartment", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetDepartmentReq struct {
	UserIDType       *IDType `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 userid
	DepartmentIDType *IDType `query:"department_id_type" json:"-"` // 此次调用中使用的部门ID的类型, 示例值："open_department_id", 可选值有: `department_id`：以自定义department_id来标识部门, `open_department_id`：以open_department_id来标识部门, 默认值: `open_department_id`
	DepartmentID     string  `path:"department_id" json:"-"`       // 需要获取的部门ID, 示例值："od-4e6ac4d14bcd5071a37a39de902c7141", 最大长度：`128` 字符, 正则校验：`^0|[^od][A-Za-z0-9]*`
}

type getDepartmentResp struct {
	Code int                `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string             `json:"msg,omitempty"`  // 错误描述
	Data *GetDepartmentResp `json:"data,omitempty"` //
}

type GetDepartmentResp struct {
	Department *GetDepartmentRespDepartment `json:"department,omitempty"` // 部门信息
}

type GetDepartmentRespDepartment struct {
	Name               string                               `json:"name,omitempty"`                 // 部门名称, 最小长度：`1` 字符
	I18nName           *GetDepartmentRespDepartmentI18nName `json:"i18n_name,omitempty"`            // 国际化的部门名称
	ParentDepartmentID string                               `json:"parent_department_id,omitempty"` // 父部门的ID,* 创建根部门，该参数值为 “0”
	DepartmentID       string                               `json:"department_id,omitempty"`        // 本部门的自定义部门ID, 最大长度：`128` 字符, 正则校验：`^0|[^od][A-Za-z0-9]*`
	OpenDepartmentID   string                               `json:"open_department_id,omitempty"`   // 部门的open_id
	LeaderUserID       string                               `json:"leader_user_id,omitempty"`       // 部门主管用户ID
	ChatID             string                               `json:"chat_id,omitempty"`              // 部门群ID
	Order              string                               `json:"order,omitempty"`                // 部门的排序，即部门在其同级部门的展示顺序
	UnitIDs            []string                             `json:"unit_ids,omitempty"`             // 部门单位自定义ID列表，当前只支持一个
	MemberCount        int                                  `json:"member_count,omitempty"`         // 部门下用户的个数
	Status             *GetDepartmentRespDepartmentStatus   `json:"status,omitempty"`               // 部门状态
}

type GetDepartmentRespDepartmentI18nName struct {
	ZhCn string `json:"zh_cn,omitempty"` // 部门的中文名
	JaJp string `json:"ja_jp,omitempty"` // 部门的日文名
	EnUs string `json:"en_us,omitempty"` // 部门的英文名
}

type GetDepartmentRespDepartmentStatus struct {
	IsDeleted bool `json:"is_deleted,omitempty"` // 是否被删除
}
