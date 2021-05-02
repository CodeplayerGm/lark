package lark

import (
	"context"
)

// GetAnnouncement 获取会话中的群公告信息，公告信息格式与[云文档](https://open.feishu.cn/document/ukTMukTMukTM/uAzM5YjLwMTO24CMzkjN)格式相同。
//
// 注意事项：
// - 应用需要开启[机器人能力](https://open.feishu.cn/document/uQjL04CN/uYTMuYTMuYTM)
// https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-announcement/get
func (r *ChatAPI) GetAnnouncement(ctx context.Context, request *GetAnnouncementReq) (*GetAnnouncementResp, *Response, error) {
	req := &requestParam{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/im/v1/chats/:chat_id/announcement",
		Body:                  request,
		NeedTenantAccessToken: true,
		NeedAppAccessToken:    false,
	}
	resp := new(getAnnouncementResp)

	response, err := r.cli.request(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, newError("Chat", "GetAnnouncement", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetAnnouncementReq struct {
	UserIDType IDType `query:"user_id_type" json:"-"` // 用户 ID 类型,**示例值：** "open_id",**可选值有：**,- `open_id`：用户的 open id,- `union_id`：用户的 union id,- `user_id`：用户的 user id,**默认值：** `open_id`,**当值为 `user_id`，字段权限要求：**,<md-perm href="/ssl:ttdoc/ukTMukTMukTM/uQjN3QjL0YzN04CN2cDN">获取用户 userid</md-perm>
	ChatID     string `path:"chat_id" json:"-"`       // 待获取公告的群 ID,**示例值：** "oc_5ad11d72b830411d72b836c20"
}

type getAnnouncementResp struct {
	Code int                  `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 错误描述
	Data *GetAnnouncementResp `json:"data,omitempty"` // \-
}

type GetAnnouncementResp struct {
	Content        string `json:"content,omitempty"`          // 云文档序列化信息
	Revision       string `json:"revision,omitempty"`         // 文档当前版本号 纯数字
	CreateTime     string `json:"create_time,omitempty"`      // 文档生成的时间戳（秒）
	UpdateTime     string `json:"update_time,omitempty"`      // 文档更新的时间戳（秒）
	OwnerIDType    string `json:"owner_id_type,omitempty"`    // 文档所有者 id 类型， open_id/user_id/union_id/app_id,**可选值有：**,- `user_id`：以 user_id 来识别用户,- `union_id`：以 union_id 来识别用户,- `open_id`：以 open_id 来识别用户,- `app_id`：以 app_id 来识别应用
	OwnerID        string `json:"owner_id,omitempty"`         // 文档所有者 id
	ModifierIDType string `json:"modifier_id_type,omitempty"` // 文档最新修改者 id 类型， open_id/user_id/union_id/app_id,**可选值有：**,- `user_id`：以 user_id 来识别用户,- `union_id`：以 union_id 来识别用户,- `open_id`：以 open_id 来识别用户,- `app_id`：以 app_id 来识别应用
	ModifierID     string `json:"modifier_id,omitempty"`      // 文档最新修改者 id
}
