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

// UpdateHireOffer 1. 更新 Offer 时, 需传入本文档中标注为必传的参数, 其余参数是否必传参考「获取 Offer 申请表模板信息」的参数定义；
//
// 2. 对系统中已存在的 offer 进行更新的, 若更新 offer 中含有「修改需审批」的字段, 更新后原 Offer 的审批会自动撤回, 需要重新发起审批。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/offer/update
func (r *HireService) UpdateHireOffer(ctx context.Context, request *UpdateHireOfferReq, options ...MethodOptionFunc) (*UpdateHireOfferResp, *Response, error) {
	if r.cli.mock.mockHireUpdateHireOffer != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Hire#UpdateHireOffer mock enable")
		return r.cli.mock.mockHireUpdateHireOffer(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "UpdateHireOffer",
		Method:                "PUT",
		URL:                   r.cli.openBaseURL + "/open-apis/hire/v1/offers/:offer_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateHireOfferResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockHireUpdateHireOffer mock HireUpdateHireOffer method
func (r *Mock) MockHireUpdateHireOffer(f func(ctx context.Context, request *UpdateHireOfferReq, options ...MethodOptionFunc) (*UpdateHireOfferResp, *Response, error)) {
	r.mockHireUpdateHireOffer = f
}

// UnMockHireUpdateHireOffer un-mock HireUpdateHireOffer method
func (r *Mock) UnMockHireUpdateHireOffer() {
	r.mockHireUpdateHireOffer = nil
}

// UpdateHireOfferReq ...
type UpdateHireOfferReq struct {
	OfferID            string                              `path:"offer_id" json:"-"`              // Offer ID, 示例值: "7016605170635213100"
	UserIDType         *IDType                             `query:"user_id_type" json:"-"`         // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), people_admin_id: 以people_admin_id来识别用户, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentIDType   *DepartmentIDType                   `query:"department_id_type" json:"-"`   // 此次调用中使用的部门 ID 的类型, 示例值: "department_id", 可选值有: open_department_id: 以 open_department_id 来标识部门, department_id: 以 department_id 来标识部门, 默认值: `open_department_id`
	SchemaID           string                              `json:"schema_id,omitempty"`            // Offer 申请表模板 ID, 用于描述申请表单结构的元数据定义, 即对申请表内容的描述。用户每一次更改 Offer 申请表模板信息, 都会生成新的 schema_id, 创建 Offer 时应传入最新的 schema_id, 可从「获取Offer申请表模板信息」接口中获取, 示例值: "7013318077945596204"
	BasicInfo          *UpdateHireOfferReqBasicInfo        `json:"basic_info,omitempty"`           // Offer 基本信息
	SalaryInfo         *UpdateHireOfferReqSalaryInfo       `json:"salary_info,omitempty"`          // Offer 薪资信息
	CustomizedInfoList []*UpdateHireOfferReqCustomizedInfo `json:"customized_info_list,omitempty"` // 自定义信息
}

// UpdateHireOfferReqBasicInfo ...
type UpdateHireOfferReqBasicInfo struct {
	DepartmentID          string   `json:"department_id,omitempty"`          // 部门 ID, 示例值: "od-6b394871807047c7023ebfc1ff37cd3a"
	LeaderUserID          string   `json:"leader_user_id,omitempty"`         // 直属上级 ID, 示例值: "ou_ce613028fe74745421f5dc320bb9c709"
	EmploymentJobID       *string  `json:"employment_job_id,omitempty"`      // 职务 ID, 示例值: "123"
	EmployeeTypeID        *string  `json:"employee_type_id,omitempty"`       // 人员类型 ID, 示例值: "2"
	JobFamilyID           *string  `json:"job_family_id,omitempty"`          // 职位序列 ID, 示例值: "6807407987381831949"
	JobLevelID            *string  `json:"job_level_id,omitempty"`           // 职位级别 ID, 示例值: "6807407987381881101"
	ProbationMonth        *int64   `json:"probation_month,omitempty"`        // 试用期, 示例值: 3
	ContractYear          *int64   `json:"contract_year,omitempty"`          // 合同期, 示例值: 3
	ExpectedOnboardDate   *string  `json:"expected_onboard_date,omitempty"`  // 预计入职日期, 示例值: "{\"date\":\"2022-04-07\"}"
	OnboardAddressID      *string  `json:"onboard_address_id,omitempty"`     // 入职地点 ID, 示例值: "6897079709306259719"
	WorkAddressID         *string  `json:"work_address_id,omitempty"`        // 办公地点 ID, 示例值: "6897079709306259719"
	OwnerUserID           string   `json:"owner_user_id,omitempty"`          // Offer负责人 ID, 示例值: "ou_ce613028fe74745421f5dc320bb9c709"
	RecommendedWords      *string  `json:"recommended_words,omitempty"`      // Offer 推荐语, 示例值: "十分优秀, 推荐入职"
	JobRequirementID      *string  `json:"job_requirement_id,omitempty"`     // 招聘需求 ID, 示例值: "2342352224"
	JobProcessTypeID      *int64   `json:"job_process_type_id,omitempty"`    // 招聘流程类型 ID, 示例值: 2
	AttachmentIDList      []string `json:"attachment_id_list,omitempty"`     // 附件ID列表, 示例值: ["7081582717280831752"]
	AttachmentDescription *string  `json:"attachment_description,omitempty"` // 附件描述, 示例值: "张三的简历"
	OperatorUserID        string   `json:"operator_user_id,omitempty"`       // Offer操作人 ID, 示例值: "ou_ce613028fe74745421f5dc320bb9c709"
}

// UpdateHireOfferReqCustomizedInfo ...
type UpdateHireOfferReqCustomizedInfo struct {
	ID    *string `json:"id,omitempty"`    // 自定义字段 ID, 示例值: "6972464088568269100"
	Value *string `json:"value,omitempty"` // 自定义字段信息, 以字符串形式传入, 如: 1. 单选: "1", 2. 多选: "[\"1\", \"2\"]", 3. 日期: "{"date":"2022-01-01"}", 4. 年份选择: "{"date":"2022"}", 5. 月份选择: "{"date":"2022-01"}", 6. 单行文本: "xxx ", 7. 多行文本: "xxx ", 8. 数字: "123", 9. 金额: "123.1", 示例值: "1"
}

// UpdateHireOfferReqSalaryInfo ...
type UpdateHireOfferReqSalaryInfo struct {
	Currency                  *string `json:"currency,omitempty"`                    // 币种, 示例值: "CNY"
	BasicSalary               *string `json:"basic_salary,omitempty"`                // 基本工资, 当启用 Offer 申请表中的「薪资信息」模块时, 「基本工资」字段为必传项, 示例值: "1000000"
	ProbationSalaryPercentage *string `json:"probation_salary_percentage,omitempty"` // 试用期百分比, 示例值: "0.8"
	AwardSalaryMultiple       *string `json:"award_salary_multiple,omitempty"`       // 年终奖月数, 示例值: "3"
	OptionShares              *string `json:"option_shares,omitempty"`               // 期权股数, 示例值: "30"
	QuarterlyBonus            *string `json:"quarterly_bonus,omitempty"`             // 季度奖金额, 示例值: "3000"
	HalfYearBonus             *string `json:"half_year_bonus,omitempty"`             // 半年奖金额, 示例值: "10000"
}

// UpdateHireOfferResp ...
type UpdateHireOfferResp struct {
}

// updateHireOfferResp ...
type updateHireOfferResp struct {
	Code int64                `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string               `json:"msg,omitempty"`  // 错误描述
	Data *UpdateHireOfferResp `json:"data,omitempty"`
}