package test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_GetMessage(t *testing.T) {
	as := assert.New(t)

	t.Run("No permission", func(t *testing.T) {
		_, _, err := AppNoPermission.Ins().Message().GetMessage(ctx, &lark.GetMessageReq{
			MessageID: "1",
		})
		as.NotNil(err)
		as.Contains(err.Error(), "No permission")
	})

	t.Run("ids not existed", func(t *testing.T) {
		_, _, err := AppALLPermission.Ins().Message().GetMessage(ctx, &lark.GetMessageReq{
			MessageID: "1",
		})
		as.NotNil(err)
		as.Contains(err.Error(), "these ids not existed")
	})

	t.Run("ids not existed", func(t *testing.T) {
		resp, _, err := AppALLPermission.Ins().Message().GetMessageList(ctx, &lark.GetMessageListReq{
			ContainerIDType: lark.ContainerIDTypeChat,
			ContainerID:     ChatContainALLPermissionApp.ChatID,
			StartTime:       nil,
			EndTime:         nil,
			PageToken:       nil,
			PageSize:        nil,
		})
		spew.Dump(resp, err)
		as.Nil(err)
		as.NotNil(resp)
		as.True(len(resp.Items) > 0)
	})
}
