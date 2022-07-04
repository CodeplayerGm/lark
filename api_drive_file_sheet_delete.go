// Code generated by lark_sdk_gen. DO NOT EDIT.
/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lark

import (
	"context"
)

// DeleteDriveSheetFile 该接口用于根据 spreadsheetToken 删除对应的 sheet 文档。
//
// 为了更好地提升该接口的安全性, 我们对其进行了升级, 请尽快迁移至
// [新版本>>](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/delete)
// </md-alert>
// <md-alert type="warn">
// 文档只能被文档所有者删除, 文档被删除后将会放到回收站里
// 该接口不支持并发调用, 且调用频率上限为5QPS
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUTNzUjL1UzM14SN1MTN/delete-sheet
//
// Deprecated
func (r *DriveService) DeleteDriveSheetFile(ctx context.Context, request *DeleteDriveSheetFileReq, options ...MethodOptionFunc) (*DeleteDriveSheetFileResp, *Response, error) {
	if r.cli.mock.mockDriveDeleteDriveSheetFile != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#DeleteDriveSheetFile mock enable")
		return r.cli.mock.mockDriveDeleteDriveSheetFile(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "DeleteDriveSheetFile",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/explorer/v2/file/spreadsheets/:spreadsheetToken",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteDriveSheetFileResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveDeleteDriveSheetFile mock DriveDeleteDriveSheetFile method
func (r *Mock) MockDriveDeleteDriveSheetFile(f func(ctx context.Context, request *DeleteDriveSheetFileReq, options ...MethodOptionFunc) (*DeleteDriveSheetFileResp, *Response, error)) {
	r.mockDriveDeleteDriveSheetFile = f
}

// UnMockDriveDeleteDriveSheetFile un-mock DriveDeleteDriveSheetFile method
func (r *Mock) UnMockDriveDeleteDriveSheetFile() {
	r.mockDriveDeleteDriveSheetFile = nil
}

// DeleteDriveSheetFileReq ...
type DeleteDriveSheetFileReq struct {
	SpreadSheetToken string `path:"spreadsheetToken" json:"-"` // spreadsheet 的 token, 获取方式见 [概述](https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/files/guide/introduction)
}

// DeleteDriveSheetFileResp ...
type DeleteDriveSheetFileResp struct {
	ID     string `json:"id,omitempty"`     // sheet 的 id 「字符串类型」
	Result bool   `json:"result,omitempty"` // 删除结果
}

// deleteDriveSheetFileResp ...
type deleteDriveSheetFileResp struct {
	Code int64                     `json:"code,omitempty"`
	Msg  string                    `json:"msg,omitempty"`
	Data *DeleteDriveSheetFileResp `json:"data,omitempty"`
}