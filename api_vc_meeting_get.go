package lark

import (
	"context"
)

// GetMeeting
//
// > 获取一个会议的详细数据
// 会议owner、会议参与者可获取会议详情，支持查询最近90天内的会议
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/vc-v1/meeting/get
func (r *VCAPI) GetMeeting(ctx context.Context, request *GetMeetingReq) (*GetMeetingResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(getMeetingResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("VC", "GetMeeting", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetMeetingReq struct {
	WithParticipants   bool   `query:"with_participants" json:"-"`    // 是否需要参会人列表，默认为false，示例值：false
	WithMeetingAbility bool   `query:"with_meeting_ability" json:"-"` // 是否需要会中使用能力统计（仅限tenant_access_token），默认为false，示例值：false
	UserIDType         IDType `query:"user_id_type" json:"-"`         // 用户ID类型，可用值：【open_id，union_id，user_id】，默认值：open_id
	MeetingID          string `path:"meeting_id" json:"-"`            // 会议id，示例值："6911188411932033028"
}

type getMeetingResp struct {
	Code int             `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 错误描述
	Data *GetMeetingResp `json:"data,omitempty"` // -
}

type GetMeetingResp struct {
	Meeting *GetMeetingRespMeeting `json:"meeting,omitempty"` // 会议数据
}

type GetMeetingRespMeeting struct {
	ID               string                              `json:"id,omitempty"`                // 会议id
	Topic            string                              `json:"topic,omitempty"`             // 会议主题
	URL              string                              `json:"url,omitempty"`               // 会议链接
	CreateTime       string                              `json:"create_time,omitempty"`       // 会议创建时间（unix时间，单位sec）
	StartTime        string                              `json:"start_time,omitempty"`        // 会议开始时间（unix时间，单位sec）
	EndTime          string                              `json:"end_time,omitempty"`          // 会议结束时间（unix时间，单位sec）
	HostUser         *GetMeetingRespMeetingHostUser      `json:"host_user,omitempty"`         // 主持人
	Status           int                                 `json:"status,omitempty"`            // 会议状态，可用值：【1（会议呼叫中），2（会议进行中），3（会议已结束）】
	ParticipantCount string                              `json:"participant_count,omitempty"` // 参会人数
	Participants     []*GetMeetingRespMeetingParticipant `json:"participants,omitempty"`      // 参会人列表
	Ability          *GetMeetingRespMeetingAbility       `json:"ability,omitempty"`           // 会中使用的能力
}

type GetMeetingRespMeetingHostUser struct {
	ID       string `json:"id,omitempty"`        // 用户id
	UserType int    `json:"user_type,omitempty"` // 用户类型，可用值：【1（飞书用户），2（飞书会议室用户），3（飞书云文档用户），4（飞书会议单品用户），5（飞书会议单品游客用户），6（PSTN用户），7（SIP用户）】
}

type GetMeetingRespMeetingParticipant struct {
	ID         string `json:"id,omitempty"`          // 用户id
	UserType   int    `json:"user_type,omitempty"`   // 用户类型，可用值：【1（飞书用户），2（飞书会议室用户），3（飞书云文档用户），4（飞书会议单品用户），5（飞书会议单品游客用户），6（PSTN用户），7（SIP用户）】
	IsHost     bool   `json:"is_host,omitempty"`     // 是否为主持人
	IsCohost   bool   `json:"is_cohost,omitempty"`   // 是否为联席主持人
	IsExternal bool   `json:"is_external,omitempty"` // 是否为外部参会人
	Status     int    `json:"status,omitempty"`      // 参会人状态，可用值：【1（呼叫中），2（在会中），3（正在响铃），4（不在会中或已经离开会议）】
}

type GetMeetingRespMeetingAbility struct {
	UseVideo        bool `json:"use_video,omitempty"`         // 是否使用视频
	UseAudio        bool `json:"use_audio,omitempty"`         // 是否使用音频
	UseShareScreen  bool `json:"use_share_screen,omitempty"`  // 是否使用共享屏幕
	UseFollowScreen bool `json:"use_follow_screen,omitempty"` // 是否使用妙享（magic share）
	UseRecording    bool `json:"use_recording,omitempty"`     // 是否使用录制
	UsePstn         bool `json:"use_pstn,omitempty"`          // 是否使用PSTN
}
