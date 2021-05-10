package lark

import (
	"context"
)

// DeleteReserve
//
// > 删除一个预约
// 操作者需为预约的owner；删除后数据不可恢复
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/vc-v1/reserve/delete
func (r *VCAPI) DeleteReserve(ctx context.Context, request *DeleteReserveReq) (*DeleteReserveResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "DELETE",
		URL:                   "https://open.feishu.cn/open-apis/vc/v1/reserves/:reserve_id",
		Body:                  request,
		NeedTenantAccessToken: false,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(deleteReserveResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("VC", "DeleteReserve", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type DeleteReserveReq struct {
	ReserveID string `path:"reserve_id" json:"-"` // 预约id，示例值："6911188411932033028"
}

type deleteReserveResp struct {
	Code int                `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string             `json:"msg,omitempty"`  // 错误描述
	Data *DeleteReserveResp `json:"data,omitempty"`
}

type DeleteReserveResp struct{}
