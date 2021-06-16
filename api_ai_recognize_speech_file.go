// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// RecognizeSpeechFile 语音文件识别接口，上传整段语音文件进行一次性识别。接口适合 60 秒以内音频识别
//
// 单租户限流：20QPS，同租户下的应用没有限流，共享本租户的 20QPS 限流
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/speech_to_text-v1/speech/file_recognize
func (r *AIService) RecognizeSpeechFile(ctx context.Context, request *RecognizeSpeechFileReq, options ...MethodOptionFunc) (*RecognizeSpeechFileResp, *Response, error) {
	if r.cli.mock.mockAIRecognizeSpeechFile != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] AI#RecognizeSpeechFile mock enable")
		return r.cli.mock.mockAIRecognizeSpeechFile(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "AI",
		API:                   "RecognizeSpeechFile",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/speech_to_text/v1/speech/file_recognize",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(recognizeSpeechFileResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockAIRecognizeSpeechFile(f func(ctx context.Context, request *RecognizeSpeechFileReq, options ...MethodOptionFunc) (*RecognizeSpeechFileResp, *Response, error)) {
	r.mockAIRecognizeSpeechFile = f
}

func (r *Mock) UnMockAIRecognizeSpeechFile() {
	r.mockAIRecognizeSpeechFile = nil
}

type RecognizeSpeechFileReq struct {
	Speech *RecognizeSpeechFileReqSpeech `json:"speech,omitempty"` // 语音资源
	Config *RecognizeSpeechFileReqConfig `json:"config,omitempty"` // 配置属性
}

type RecognizeSpeechFileReqSpeech struct {
	Speech *string `json:"speech,omitempty"` // base64 后的音频文件进行, 示例值："base64 后的音频内容"
}

type RecognizeSpeechFileReqConfig struct {
	FileID     string `json:"file_id,omitempty"`     // 仅包含字母数字和下划线的 16 位字符串作为文件的标识，用户生成, 示例值："qwe12dd34567890w"
	Format     string `json:"format,omitempty"`      // 语音格式，目前仅支持：pcm, 示例值："pcm"
	EngineType string `json:"engine_type,omitempty"` // 引擎类型，目前仅支持：16k_auto 中英混合, 示例值："16k_auto"
}

type recognizeSpeechFileResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *RecognizeSpeechFileResp `json:"data,omitempty"`
}

type RecognizeSpeechFileResp struct {
	RecognitionText string `json:"recognition_text,omitempty"` // 语音识别后的文本信息
}
