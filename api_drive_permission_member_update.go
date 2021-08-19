// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// UpdateDriveMemberPermission 该接口用于根据 filetoken 更新文档协作者的权限。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/permission-member/update
func (r *DriveService) UpdateDriveMemberPermission(ctx context.Context, request *UpdateDriveMemberPermissionReq, options ...MethodOptionFunc) (*UpdateDriveMemberPermissionResp, *Response, error) {
	if r.cli.mock.mockDriveUpdateDriveMemberPermission != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#UpdateDriveMemberPermission mock enable")
		return r.cli.mock.mockDriveUpdateDriveMemberPermission(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "UpdateDriveMemberPermission",
		Method:                "PUT",
		URL:                   "https://open.feishu.cn/open-apis/drive/v1/permissions/:token/members/:member_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateDriveMemberPermissionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveUpdateDriveMemberPermission(f func(ctx context.Context, request *UpdateDriveMemberPermissionReq, options ...MethodOptionFunc) (*UpdateDriveMemberPermissionResp, *Response, error)) {
	r.mockDriveUpdateDriveMemberPermission = f
}

func (r *Mock) UnMockDriveUpdateDriveMemberPermission() {
	r.mockDriveUpdateDriveMemberPermission = nil
}

type UpdateDriveMemberPermissionReq struct {
	NeedNotification *bool  `query:"need_notification" json:"-"` // 更新权限后是否通知对方, 示例值：false, 默认值: `false`
	Type             string `query:"type" json:"-"`              // 文件类型，放于query参数中，如：`?type=doc`, 示例值："doc", 可选值有: `doc`：文档, `sheet`：电子表格, `file`：云空间文件, `wiki`：知识库节点, `bitable`：多维表格, `docx`：文档
	Token            string `path:"token" json:"-"`              // 文件的 token，获取方式见 [概述](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/files/guide/introduction), 示例值："doccnBKgoMyY5OMbUG6FioTXuBe"
	MemberID         string `path:"member_id" json:"-"`          // 权限成员的openID，获取方式见 [如何获得 User ID、Open ID 和 Union ID？](https://open.feishu.cn/document/home/user-identity-introduction/how-to-get), 示例值："ou_7dab8a3d3cdcc9da365777c7ad535d62"
	MemberType       string `json:"member_type,omitempty"`       // 用户类型，与路径参数中的`member_id`要对应，可选值有：, `email`: 飞书企业邮箱, `openid`: 开放平台ID, `openchat`: 开放平台群组, `opendepartmentid`: 开放平台部门ID, `userid`: 用户自定义ID, 示例值："openid"
	Perm             string `json:"perm,omitempty"`              // 需要更新的权限，可选值有：, `view`: 可阅读, `edit`: 可编辑, `full_access`: 所有权限, 示例值："view"
}

type updateDriveMemberPermissionResp struct {
	Code int64                            `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 错误描述
	Data *UpdateDriveMemberPermissionResp `json:"data,omitempty"`
}

type UpdateDriveMemberPermissionResp struct {
	Member *UpdateDriveMemberPermissionRespMember `json:"member,omitempty"` // 本次更新权限的用户信息
}

type UpdateDriveMemberPermissionRespMember struct {
	MemberType string `json:"member_type,omitempty"` // 用户类型，与路径参数中的`member_id`要对应，可选值有：, `email`: 飞书企业邮箱, `openid`: 开放平台ID, `openchat`: 开放平台群组, `opendepartmentid`: 开放平台部门ID, `userid`: 用户自定义ID
	MemberID   string `json:"member_id,omitempty"`   // 用户类型下的值
	Perm       string `json:"perm,omitempty"`        // 需要更新的权限，可选值有：, `view`: 可阅读, `edit`: 可编辑, `full_access`: 所有权限
}
