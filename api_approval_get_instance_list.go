package lark

import (
	"context"
)

// GetInstanceList
//
// 根据 approval_code 批量获取审批实例的 instance_code，用于拉取租户下某个审批定义的全部审批实例。
// 默认以审批创建时间排序。
// https://open.feishu.cn/document/ukTMukTMukTM/uQDOyUjL0gjM14CN4ITN
func (r *ApprovalAPI) GetInstanceList(ctx context.Context, request *GetInstanceListReq) (*GetInstanceListResp, *Response, error) {
	req := &requestParam{
		Method:                "POST",
		URL:                   "https://www.feishu.cn/approval/openapi/v2/instance/list",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(getInstanceListResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Approval", "GetInstanceList", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetInstanceListReq struct {
	ApprovalCode string `json:"approval_code,omitempty"` // 审批定义唯一标识
	StartTime    int    `json:"start_time,omitempty"`    // 审批实例创建时间区间（毫秒）
	EndTime      int    `json:"end_time,omitempty"`      // 审批实例创建时间区间（毫秒）
	Offset       int    `json:"offset,omitempty"`        // 查询偏移量
	Limit        int    `json:"limit,omitempty"`         // 查询限制量 注:不得大于100
}

type getInstanceListResp struct {
	Code int                  `json:"code,omitempty"` // 错误码，非0表示失败
	Msg  string               `json:"msg,omitempty"`  // 返回码的描述
	Data *GetInstanceListResp `json:"data,omitempty"` // 返回业务信息
}

type GetInstanceListResp struct {
	InstanceCodeList []string `json:"instance_code_list,omitempty"` // 审批实例 Code
}
