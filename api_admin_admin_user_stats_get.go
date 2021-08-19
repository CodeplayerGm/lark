// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetAdminUserStats 用于获取用户维度的用户活跃和功能使用数据，即IM（即时通讯）、日历、云文档、音视频会议功能的使用数据。
//
// - 只有企业自建应用才有权限调用此接口
// - 当天的数据会在第二天的凌晨五点产出（UTC+8）
// - 用户维度的数据最多查询最近31天的数据（包含31天）的数据
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/admin-v1/admin_user_stat/list
func (r *AdminService) GetAdminUserStats(ctx context.Context, request *GetAdminUserStatsReq, options ...MethodOptionFunc) (*GetAdminUserStatsResp, *Response, error) {
	if r.cli.mock.mockAdminGetAdminUserStats != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Admin#GetAdminUserStats mock enable")
		return r.cli.mock.mockAdminGetAdminUserStats(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Admin",
		API:                   "GetAdminUserStats",
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/admin/v1/admin_user_stats",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getAdminUserStatsResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockAdminGetAdminUserStats(f func(ctx context.Context, request *GetAdminUserStatsReq, options ...MethodOptionFunc) (*GetAdminUserStatsResp, *Response, error)) {
	r.mockAdminGetAdminUserStats = f
}

func (r *Mock) UnMockAdminGetAdminUserStats() {
	r.mockAdminGetAdminUserStats = nil
}

type GetAdminUserStatsReq struct {
	UserIDType       *IDType           `query:"user_id_type" json:"-"`       // 用户 ID 类型, 示例值："open_id", 可选值有: `open_id`：用户的 open id, `union_id`：用户的 union id, `user_id`：用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentIDType *DepartmentIDType `query:"department_id_type" json:"-"` // 部门ID类型, 示例值："open_department_id", 可选值有: `department_id`：部门的 ID, `open_department_id`：部门的 Open ID
	StartDate        string            `query:"start_date" json:"-"`         // 起始日期（包含），格式是YYYY-mm-dd, 示例值："2020-02-15"
	EndDate          string            `query:"end_date" json:"-"`           // 终止日期（包含），格式是YYYY-mm-dd。起止日期之间相差不能超过31天（包含31天）, 示例值："2020-02-15"
	DepartmentID     *string           `query:"department_id" json:"-"`      // 部门的 ID，取决于department_id_type, 示例值："od-382e2793cfc9471f892e8a672987654c"
	UserID           *string           `query:"user_id" json:"-"`            // 用户的open_id，user_id或者union_id，取决于user_id_type, 示例值："ou_7dab8a3d3cdcc9da365777c7ad535d62"
	PageSize         *int64            `query:"page_size" json:"-"`          // 分页大小，默认是10, 示例值：10, 取值范围：`1` ～ `20`
	PageToken        *string           `query:"page_token" json:"-"`         // 分页标记，第一次请求不填，表示从头开始遍历；当返回的has_more为true时，会返回新的page_token，再次调用接口，传入这个page_token，将获得下一页数据, 示例值："2"
}

type getAdminUserStatsResp struct {
	Code int64                  `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *GetAdminUserStatsResp `json:"data,omitempty"`
}

type GetAdminUserStatsResp struct {
	HasMore   bool                         `json:"has_more,omitempty"`   // 是否有下一页数据
	PageToken string                       `json:"page_token,omitempty"` // 下一页分页的token
	Items     []*GetAdminUserStatsRespItem `json:"items,omitempty"`      // 数据报表
}

type GetAdminUserStatsRespItem struct {
	Date             string `json:"date,omitempty"`               // 日期
	UserID           string `json:"user_id,omitempty"`            // 用户ID
	UserName         string `json:"user_name,omitempty"`          // 用户名
	DepartmentName   string `json:"department_name,omitempty"`    // 部门名
	DepartmentPath   string `json:"department_path,omitempty"`    // 部门路径
	CreateTime       string `json:"create_time,omitempty"`        // 添加时间
	UserActiveFlag   int64  `json:"user_active_flag,omitempty"`   // 用户激活状态, 可选值有: `0`：未激活, `1`：已激活
	RegisterTime     string `json:"register_time,omitempty"`      // 激活时间
	SuiteActiveFlag  int64  `json:"suite_active_flag,omitempty"`  // 用户活跃状态, 可选值有: `0`：无活跃, `1`：活跃
	LastActiveTime   string `json:"last_active_time,omitempty"`   // 最近活跃时间
	IMActiveFlag     int64  `json:"im_active_flag,omitempty"`     // 用户消息活跃状态, 可选值有: `0`：无活跃, `1`：活跃
	SendMessengerNum int64  `json:"send_messenger_num,omitempty"` // 发送消息数
	DocsActiveFlag   int64  `json:"docs_active_flag,omitempty"`   // 用户云文档活跃状态, 可选值有: `0`：无活跃, `1`：活跃
	CreateDocsNum    int64  `json:"create_docs_num,omitempty"`    // 创建文件数
	CalActiveFlag    int64  `json:"cal_active_flag,omitempty"`    // 用户日历活跃状态, 可选值有: `0`：无活跃, `1`：活跃
	CreateCalNum     int64  `json:"create_cal_num,omitempty"`     // 创建日程数
	VCActiveFlag     int64  `json:"vc_active_flag,omitempty"`     // 用户音视频会议活跃状态, 可选值有: `0`：无活跃, `1`：活跃
	VCDuration       int64  `json:"vc_duration,omitempty"`        // 会议时长
	ActiveOs         string `json:"active_os,omitempty"`          // 活跃设备
}
