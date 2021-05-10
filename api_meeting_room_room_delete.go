package lark

import (
	"context"
)

// DeleteRoom 该接口用于删除会议室。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUzMxYjL1MTM24SNzEjN
func (r *MeetingRoomAPI) DeleteRoom(ctx context.Context, request *DeleteRoomReq) (*DeleteRoomResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/meeting_room/room/delete",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(deleteRoomResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("MeetingRoom", "DeleteRoom", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type DeleteRoomReq struct {
	RoomID string `json:"room_id,omitempty"` // 要删除的会议室ID
}

type deleteRoomResp struct {
	Code int             `json:"code,omitempty"` // 返回码，非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 返回码的描述，"success" 表示成功，其他为错误提示信息
	Data *DeleteRoomResp `json:"data,omitempty"`
}

type DeleteRoomResp struct{}
