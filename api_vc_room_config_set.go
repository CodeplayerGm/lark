package lark

import (
	"context"
)

// SetRoomConfig
//
// > 设置一个范围内的会议室配置
// 根据设置范围传入对应的参数
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/vc-v1/room_config/set
func (r *VCAPI) SetRoomConfig(ctx context.Context, request *SetRoomConfigReq) (*SetRoomConfigResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/vc/v1/room_configs/set",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(setRoomConfigResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("VC", "SetRoomConfig", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type SetRoomConfigReq struct {
	Scope      int                         `json:"scope,omitempty"`       // 设置节点范围，必选字段，可用值：【1（租户），2（国家/地区），3（城市），4（建筑），5（楼层），6（会议室）】
	CountryID  string                      `json:"country_id,omitempty"`  // 国家/地区id scope为2/3时此参数必填
	DistrictID string                      `json:"district_id,omitempty"` // 城市id scope为3时此参数必填
	BuildingID string                      `json:"building_id,omitempty"` // 建筑id scope为4/5时此参数必填
	FloorName  string                      `json:"floor_name,omitempty"`  // 楼层 scope为5时此参数必填
	RoomID     string                      `json:"room_id,omitempty"`     // 会议室id scope为6时此参数必填
	RoomConfig *SetRoomConfigReqRoomConfig `json:"room_config,omitempty"` // 会议室设置，必选字段
}

type SetRoomConfigReqRoomConfig struct {
	RoomBackground    string                                    `json:"room_background,omitempty"`    // 飞书会议室背景图
	DisplayBackground string                                    `json:"display_background,omitempty"` // 飞书签到板背景图
	DigitalSignage    *SetRoomConfigReqRoomConfigDigitalSignage `json:"digital_signage,omitempty"`    // 飞书会议室数字标牌
}

type SetRoomConfigReqRoomConfigDigitalSignage struct {
	Enable       bool                                                `json:"enable,omitempty"`        // 是否开启数字标牌功能
	Mute         bool                                                `json:"mute,omitempty"`          // 是否静音播放
	StartDisplay int                                                 `json:"start_display,omitempty"` // 日程会议开始前n分钟结束播放
	StopDisplay  int                                                 `json:"stop_display,omitempty"`  // 会议结束后n分钟开始播放
	Materials    []*SetRoomConfigReqRoomConfigDigitalSignageMaterial `json:"materials,omitempty"`     // 素材列表
}

type SetRoomConfigReqRoomConfigDigitalSignageMaterial struct {
	ID           string `json:"id,omitempty"`            // 素材id
	Name         string `json:"name,omitempty"`          // 素材名称
	MaterialType int    `json:"material_type,omitempty"` // 素材类型，可用值：【1（图片），2（视频），3（GIF）】
	URL          string `json:"url,omitempty"`           // 素材url
	Duration     int    `json:"duration,omitempty"`      // 播放时长（单位sec）
	Cover        string `json:"cover,omitempty"`         // 素材封面url
	Md5          string `json:"md5,omitempty"`           // 素材文件md5
}

type setRoomConfigResp struct {
	Code int                `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string             `json:"msg,omitempty"`  // 错误描述
	Data *SetRoomConfigResp `json:"data,omitempty"`
}

type SetRoomConfigResp struct{}
