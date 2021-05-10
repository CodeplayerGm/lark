package lark

import (
	"context"
)

// BatchGetRoom 该接口用于获取指定会议室的详细信息。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uEDOyUjLxgjM14SM4ITN
func (r *MeetingRoomAPI) BatchGetRoom(ctx context.Context, request *BatchGetRoomReq) (*BatchGetRoomResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/meeting_room/room/batch_get?room_ids=omm_eada1d61a550955240c28757e7dec3af&room_ids=omm_83d09ad4f6896e02029a6a075f71c9d1&fields=*",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(batchGetRoomResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("MeetingRoom", "BatchGetRoom", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type BatchGetRoomReq struct {
	RoomIDs string  `json:"room_ids,omitempty"` // 用于查询指定会议室的 ID
	Fields  *string `json:"fields,omitempty"`   // 用于指定返回的字段名，每个字段名之间用逗号 "," 分隔，如：“id,name”，"*" 表示返回全部字段，可选字段有："id,name,description,capacity,building_id,building_name,floor_name,is_disabled,display_id"，默认返回所有字段
}

type batchGetRoomResp struct {
	Code int               `json:"code,omitempty"` // 返回码，非 0 表示失败
	Msg  string            `json:"msg,omitempty"`  // 返回码的描述，"success" 表示成功，其他为错误提示信息
	Data *BatchGetRoomResp `json:"data,omitempty"` // 返回业务信息
}

type BatchGetRoomResp struct {
	Rooms *BatchGetRoomRespRooms `json:"rooms,omitempty"` // 会议室列表
}

type BatchGetRoomRespRooms struct {
	RoomID       string `json:"room_id,omitempty"`       // 会议室 ID
	BuildingID   string `json:"building_id,omitempty"`   // 会议室所属建筑物 ID
	BuildingName string `json:"building_name,omitempty"` // 会议室所属建筑物名称
	Capacity     int    `json:"capacity,omitempty"`      // 会议室能容纳的人数
	Description  string `json:"description,omitempty"`   // 会议室的相关描述
	DisplayID    string `json:"display_id,omitempty"`    // 会议室的展示 ID
	FloorName    string `json:"floor_name,omitempty"`    // 会议室所在楼层名称
	IsDisabled   bool   `json:"is_disabled,omitempty"`   // 会议室是否不可用，若会议室不可用，则该值为 true，否则为 false
	Name         string `json:"name,omitempty"`          // 会议室名称
}
