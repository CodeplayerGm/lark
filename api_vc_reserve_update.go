package lark

import (
	"context"
)

// UpdateReserve
//
// > 更新一个预约
// 操作者需为预约的owner，只传需要更新的字段，无需更新的字段不传（传任意值都会被更新）；可用于续期操作（到期时间距今不超过30天）
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/vc-v1/reserve/update
func (r *VCAPI) UpdateReserve(ctx context.Context, request *UpdateReserveReq) (*UpdateReserveResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "PUT",
		URL:                   "https://open.feishu.cn/open-apis/vc/v1/reserves/:reserve_id",
		Body:                  request,
		NeedTenantAccessToken: false,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(updateReserveResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("VC", "UpdateReserve", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type UpdateReserveReq struct {
	UserIDType      IDType                           `query:"user_id_type" json:"-"`     // 用户ID类型，可用值：【open_id，union_id，user_id】，默认值：open_id
	ReserveID       string                           `path:"reserve_id" json:"-"`        // 预约id，示例值："6911188411932033028"
	EndTime         string                           `json:"end_time,omitempty"`         // 预约到期时间（unix时间，单位sec）
	MeetingSettings *UpdateReserveReqMeetingSettings `json:"meeting_settings,omitempty"` // 会议设置
}

type UpdateReserveReqMeetingSettings struct {
	Topic             string                                             `json:"topic,omitempty"`              // 会议主题
	ActionPermissions []*UpdateReserveReqMeetingSettingsActionPermission `json:"action_permissions,omitempty"` // 会议权限配置列表，示例请求体中的权限配置含义为：配置"是否能成为主持人"的权限，检查用户id，检查方式为白名单，即：在白名单中的用户才有成为主持人的权限。如没有限制权限的需求，该字段可为空；如列表中存在相同的权限项，则它们之间为"逻辑或"的关系（即 有一个为true则拥有该权限）
}

type UpdateReserveReqMeetingSettingsActionPermission struct {
	Permission         int                                                                 `json:"permission,omitempty"`          // 权限项，必选字段，可用值：【1（是否能成为主持人），2（是否能邀请参会人），3（是否能分享会议）】
	PermissionCheckers []*UpdateReserveReqMeetingSettingsActionPermissionPermissionChecker `json:"permission_checkers,omitempty"` // 权限检查器列表，同一项权限可配置多个权限检查器，权限检查器之间为"逻辑或"的关系（即 有一个为true则结果为true），必选字段
}

type UpdateReserveReqMeetingSettingsActionPermissionPermissionChecker struct {
	CheckField int      `json:"check_field,omitempty"` // 检查字段类型，即check_list中的值的类型，如用户id、租户id等，必选字段，可用值：【1（用户id），2（用户类型），3（租户ID）】
	CheckMode  int      `json:"check_mode,omitempty"`  // 检查方式，白名单或者黑名单，必选字段，可用值：【1（白名单（在check_list中为有权限）），2（黑名单（不在check_list中为有权限））】
	CheckList  []string `json:"check_list,omitempty"`  // 检查字段值列表，根据check_mode不同为白名单或黑名单，必选字段
}

type updateReserveResp struct {
	Code int                `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string             `json:"msg,omitempty"`  // 错误描述
	Data *UpdateReserveResp `json:"data,omitempty"` // -
}

type UpdateReserveResp struct {
	Reserve *UpdateReserveRespReserve `json:"reserve,omitempty"` // 预约数据
}

type UpdateReserveRespReserve struct {
	ID           string `json:"id,omitempty"`            // 预约id
	MeetingNo    string `json:"meeting_no,omitempty"`    // 9位会议号
	URL          string `json:"url,omitempty"`           // 会议链接
	EndTime      string `json:"end_time,omitempty"`      // 预约到期时间（unix时间，单位sec）
	ExpireStatus int    `json:"expire_status,omitempty"` // 过期状态，可用值：【1（未过期），2（已过期）】
}
