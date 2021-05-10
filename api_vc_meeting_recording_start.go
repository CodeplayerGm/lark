package lark

import (
	"context"
)

// StartMeetingRecording
//
// > 在会议中开始录制
// 会议正在进行中，且操作者具有相应权限（如果操作者为用户，必须是会中当前主持人）
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/vc-v1/meeting.recording/start
func (r *VCAPI) StartMeetingRecording(ctx context.Context, request *StartMeetingRecordingReq) (*StartMeetingRecordingResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "PATCH",
		URL:                   "https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id/recording/start",
		Body:                  request,
		NeedTenantAccessToken: false,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(startMeetingRecordingResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("VC", "StartMeetingRecording", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type StartMeetingRecordingReq struct {
	MeetingID string `path:"meeting_id" json:"-"` // 会议id，示例值："6911188411932033028"
	Timezone  int    `json:"timezone,omitempty"`  // 录制文件时间显示使用的时区[-12,12]
}

type startMeetingRecordingResp struct {
	Code int                        `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *StartMeetingRecordingResp `json:"data,omitempty"`
}

type StartMeetingRecordingResp struct{}
