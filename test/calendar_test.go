package test

import (
	"fmt"
	"testing"

	"github.com/chyroc/go-ptr"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Calendar_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		helpdeskCli := cli.Calendar()

		t.Run("", func(t *testing.T) {
			_, _, err := helpdeskCli.CreateCalendar(ctx, &lark.CreateCalendarReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := helpdeskCli.DeleteCalendar(ctx, &lark.DeleteCalendarReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := helpdeskCli.GetCalendar(ctx, &lark.GetCalendarReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := helpdeskCli.GetCalendarList(ctx, &lark.GetCalendarListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := helpdeskCli.UpdateCalendar(ctx, &lark.UpdateCalendarReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := helpdeskCli.SearchCalendar(ctx, &lark.SearchCalendarReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})
		t.Run("", func(t *testing.T) {
			_, _, err := helpdeskCli.SubscribeCalendar(ctx, &lark.SubscribeCalendarReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := helpdeskCli.UnsubscribeCalendar(ctx, &lark.UnsubscribeCalendarReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})
	})

	t.Run("response failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.Calendar()

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateCalendar(ctx, &lark.CreateCalendarReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteCalendar(ctx, &lark.DeleteCalendarReq{
				CalendarID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetCalendar(ctx, &lark.GetCalendarReq{
				CalendarID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetCalendarList(ctx, &lark.GetCalendarListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateCalendar(ctx, &lark.UpdateCalendarReq{
				CalendarID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SearchCalendar(ctx, &lark.SearchCalendarReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})
		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SubscribeCalendar(ctx, &lark.SubscribeCalendarReq{
				CalendarID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UnsubscribeCalendar(ctx, &lark.UnsubscribeCalendarReq{
				CalendarID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})
	})
}

func Test_CalendarEvent_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.Calendar()

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateCalendarEvent(ctx, &lark.CreateCalendarEventReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteCalendarEvent(ctx, &lark.DeleteCalendarEventReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetCalendarEvent(ctx, &lark.GetCalendarEventReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetCalendarEventList(ctx, &lark.GetCalendarEventListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateCalendarEvent(ctx, &lark.UpdateCalendarEventReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			// _, _, err := moduleCli.SearchCalendarEvent(ctx, &lark.SearchCalendarEventReq{
			// 	CalendarID: "x",
			// 	Query:      "x",
			// })
			// as.NotNil(err)
			// as.Equal(err.Error(), "failed")
		})
		t.Run("", func(t *testing.T) {
			// _, _, err := moduleCli.SubscribeCalendarEvent(ctx, &lark.SubscribeCalendarEventReq{
			// 	CalendarID: "x",
			// })
			// as.NotNil(err)
			// as.Equal(err.Error(), "failed")
		})
	})

	t.Run("response failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.Calendar()

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateCalendarEvent(ctx, &lark.CreateCalendarEventReq{
				CalendarID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteCalendarEvent(ctx, &lark.DeleteCalendarEventReq{
				CalendarID: "x",
				EventID:    "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetCalendarEvent(ctx, &lark.GetCalendarEventReq{
				CalendarID: "x",
				EventID:    "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetCalendarEventList(ctx, &lark.GetCalendarEventListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateCalendarEvent(ctx, &lark.UpdateCalendarEventReq{
				CalendarID: "x",
				EventID:    "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SearchCalendarEvent(ctx, &lark.SearchCalendarEventReq{
				CalendarID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})
		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SubscribeCalendarEvent(ctx, &lark.SubscribeCalendarEventReq{
				CalendarID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})
	})
}

func Test_CalendarACL_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.Calendar()

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateCalendarACL(ctx, &lark.CreateCalendarACLReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteCalendarACL(ctx, &lark.DeleteCalendarACLReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetCalendarACLList(ctx, &lark.GetCalendarACLListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			// _, _, err := moduleCli.SubscribeCalendarACL(ctx, &lark.SubscribeCalendarACLReq{
			// 	CalendarID: "x",
			// })
			// as.NotNil(err)
			// as.Equal(err.Error(), "failed")
		})
	})

	t.Run("response failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.Calendar()

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateCalendarACL(ctx, &lark.CreateCalendarACLReq{
				CalendarID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteCalendarACL(ctx, &lark.DeleteCalendarACLReq{
				CalendarID: "x",
				ACLID:      "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetCalendarACLList(ctx, &lark.GetCalendarACLListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SubscribeCalendarACL(ctx, &lark.SubscribeCalendarACLReq{
				CalendarID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})
	})
}

func Test_CalendarEventAttendee_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.Calendar()

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateCalendarEventAttendee(ctx, &lark.CreateCalendarEventAttendeeReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetCalendarEventAttendeeList(ctx, &lark.GetCalendarEventAttendeeListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteCalendarEventAttendee(ctx, &lark.DeleteCalendarEventAttendeeReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})
	})

	t.Run("response failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.Calendar()

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateCalendarEventAttendee(ctx, &lark.CreateCalendarEventAttendeeReq{
				CalendarID: "x",
				EventID:    "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetCalendarEventAttendeeList(ctx, &lark.GetCalendarEventAttendeeListReq{
				CalendarID: "x",
				EventID:    "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteCalendarEventAttendee(ctx, &lark.DeleteCalendarEventAttendeeReq{
				CalendarID: "x",
				EventID:    "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})
	})
}

func Test_Calendar_Other_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.Calendar()

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetCalendarEventAttendeeChatMemberList(ctx, &lark.GetCalendarEventAttendeeChatMemberListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetCalendarFreebusyList(ctx, &lark.GetCalendarFreebusyListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateCalendarTimeoffEvent(ctx, &lark.CreateCalendarTimeoffEventReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteCalendarTimeoffEvent(ctx, &lark.DeleteCalendarTimeoffEventReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			// _, _, err := moduleCli.GenerateCaldavConf(ctx, &lark.GenerateCaldavConfReq{
			// 	DeviceName: ptr.String("x"),
			// })
			// as.NotNil(err)
			// as.Equal(err.Error(), "failed")
		})
	})

	t.Run("response failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.Calendar()

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetCalendarEventAttendeeChatMemberList(ctx, &lark.GetCalendarEventAttendeeChatMemberListReq{
				CalendarID: "x",
				EventID:    "x",
				AttendeeID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetCalendarFreebusyList(ctx, &lark.GetCalendarFreebusyListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateCalendarTimeoffEvent(ctx, &lark.CreateCalendarTimeoffEventReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteCalendarTimeoffEvent(ctx, &lark.DeleteCalendarTimeoffEventReq{
				TimeoffEventID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			// _, _, err := moduleCli.GenerateCaldavConf(ctx, &lark.GenerateCaldavConfReq{
			// 	DeviceName: ptr.String("x"),
			// })
			// as.NotNil(err)
			// as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})
	})
}

func Test_Calendar(t *testing.T) {
	as := assert.New(t)
	moduleCli := AppALLPermission.Ins().Calendar()

	t.Run("", func(t *testing.T) {
		calendarID := ""
		summary := "summary-test"

		{
			resp, _, err := moduleCli.CreateCalendar(ctx, &lark.CreateCalendarReq{
				Summary:      ptr.String("summary-test"),
				Description:  ptr.String("desc-test"),
				Permissions:  nil,
				Color:        nil,
				SummaryAlias: nil,
			})
			spew.Dump(resp, err)
			as.Nil(err)
			calendarID = resp.Calendar.CalendarID
		}

		{
			resp, _, err := moduleCli.GetCalendar(ctx, &lark.GetCalendarReq{
				CalendarID: calendarID,
			})
			spew.Dump(resp, err)
			as.Nil(err)
			as.Equal(summary, resp.Summary)
		}

		{
			resp, _, err := moduleCli.GetCalendarList(ctx, &lark.GetCalendarListReq{})
			spew.Dump(resp, err)
			as.Nil(err)
		}

		{
			resp, _, err := moduleCli.UpdateCalendar(ctx, &lark.UpdateCalendarReq{
				CalendarID: calendarID,
				Summary:    ptr.String(summary + "-update"),
			})
			spew.Dump(resp, err)
			as.Nil(err)
		}

		{
			resp, _, err := moduleCli.SubscribeCalendar(ctx, &lark.SubscribeCalendarReq{
				CalendarID: calendarID,
			})
			spew.Dump(resp, err)
			as.Nil(err)
		}

		{
			resp, _, err := moduleCli.UnsubscribeCalendar(ctx, &lark.UnsubscribeCalendarReq{
				CalendarID: calendarID,
			})
			spew.Dump(resp, err)
			as.Nil(err)
		}

		{
			resp, _, err := moduleCli.DeleteCalendar(ctx, &lark.DeleteCalendarReq{
				CalendarID: calendarID,
			})
			spew.Dump(resp, err)
			as.Nil(err)
		}
	})
}

func Test_CalendarEvent(t *testing.T) {
	as := assert.New(t)
	moduleCli := AppALLPermission.Ins().Calendar()

	t.Run("", func(t *testing.T) {
		calendarID := ""
		eventID := ""
		summary := "summary-test"
		defer func() {
			_, _, _ = moduleCli.DeleteCalendar(ctx, &lark.DeleteCalendarReq{
				CalendarID: calendarID,
			})
		}()

		{
			resp, _, err := moduleCli.CreateCalendar(ctx, &lark.CreateCalendarReq{
				Summary:      ptr.String("summary-test"),
				Description:  ptr.String("desc-test"),
				Permissions:  nil,
				Color:        nil,
				SummaryAlias: nil,
			})
			spew.Dump(resp, err)
			as.Nil(err)
			calendarID = resp.Calendar.CalendarID
		}

		{
			resp, _, err := moduleCli.CreateCalendarEvent(ctx, &lark.CreateCalendarEventReq{
				CalendarID: calendarID,
				Summary:    &summary,
				StartTime: &lark.CreateCalendarEventReqStartTime{
					Date: ptr.String("2020-09-01"),
				},
				EndTime: &lark.CreateCalendarEventReqEndTime{
					Date: ptr.String("2020-09-02"),
				},
			})
			spew.Dump(resp, err)
			as.Nil(err)
			eventID = resp.Event.EventID
		}

		{
			resp, _, err := moduleCli.GetCalendarEvent(ctx, &lark.GetCalendarEventReq{
				CalendarID: calendarID,
				EventID:    eventID,
			})
			spew.Dump(resp, err)
			as.Nil(err)
			as.Equal(summary, resp.Event.Summary)
		}

		{
			resp, _, err := moduleCli.GetCalendarEventList(ctx, &lark.GetCalendarEventListReq{
				CalendarID: calendarID,
			})
			spew.Dump(resp, err)
			as.Nil(err)
		}

		{
			resp, _, err := moduleCli.UpdateCalendarEvent(ctx, &lark.UpdateCalendarEventReq{
				CalendarID: calendarID,
				EventID:    eventID,
				Summary:    ptr.String(summary + "-update"),
			})
			spew.Dump(resp, err)
			as.Nil(err)
		}

		{
			// TODO: user_access_token
			// resp, _, err := moduleCli.SubscribeCalendarEvent(ctx, &lark.SubscribeCalendarEventReq{
			// 	CalendarID: calendarID,
			// })
			// spew.Dump(resp, err)
			// as.Nil(err)
		}

		{
			resp, _, err := moduleCli.DeleteCalendarEvent(ctx, &lark.DeleteCalendarEventReq{
				CalendarID: calendarID,
				EventID:    eventID,
			})
			spew.Dump(resp, err)
			as.Nil(err)
		}
	})
}

func Test_CalendarACL(t *testing.T) {
	as := assert.New(t)
	moduleCli := AppALLPermission.Ins().Calendar()

	t.Run("", func(t *testing.T) {
		calendarID := ""
		aclID := ""
		defer func() {
			_, _, _ = moduleCli.DeleteCalendar(ctx, &lark.DeleteCalendarReq{
				CalendarID: calendarID,
			})
		}()

		{
			resp, _, err := moduleCli.CreateCalendar(ctx, &lark.CreateCalendarReq{
				Summary:      ptr.String("summary-test"),
				Description:  ptr.String("desc-test"),
				Permissions:  nil,
				Color:        nil,
				SummaryAlias: nil,
			})
			spew.Dump(resp, err)
			as.Nil(err)
			calendarID = resp.Calendar.CalendarID
		}

		{
			resp, _, err := moduleCli.CreateCalendarACL(ctx, &lark.CreateCalendarACLReq{
				UserIDType: nil,
				CalendarID: calendarID,
				Role:       lark.CalendarRoleWriter,
				Scope: &lark.CreateCalendarACLReqScope{
					Type:   "user",
					UserID: &UserAdmin.OpenID,
				},
			})
			spew.Dump(resp, err)
			as.Nil(err)
			aclID = resp.ACLID
		}

		{
			resp, _, err := moduleCli.GetCalendarACLList(ctx, &lark.GetCalendarACLListReq{
				CalendarID: calendarID,
			})
			spew.Dump(resp, err)
			as.Nil(err)
		}

		{
			// TODO: user_access_token
			// resp, _, err := moduleCli.SubscribeCalendarACL(ctx, &lark.SubscribeCalendarACLReq{
			// 	CalendarID: calendarID,
			// })
			// spew.Dump(resp, err)
			// as.Nil(err)
		}

		{
			resp, _, err := moduleCli.DeleteCalendarACL(ctx, &lark.DeleteCalendarACLReq{
				CalendarID: calendarID,
				ACLID:      aclID,
			})
			spew.Dump(resp, err)
			as.Nil(err)
		}
	})
}
