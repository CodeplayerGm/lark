package lark

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
)

// https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM#8f960a4b

type eventReq struct {
	Encrypt string `json:"encrypt"` // 加密

	// v2 字段
	Schema string       `json:"schema"` // 2.0
	Header *EventHeader `json:"header"`

	// v1 字段
	UUID      string `json:"uuid"`
	Token     string `json:"token"`
	Ts        string `json:"ts"`
	Type      string `json:"type"` // event_callback,
	Challenge string `json:"challenge"`

	// 通用字段
	Event interface{} `json:"event"`

	// v2 解析后字段
	*eventBody
}

func (r *eventReq) unmarshalEvent(e interface{}) error {
	bs, err := json.Marshal(r.Event)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bs, e)
	if err != nil {
		return fmt.Errorf("lark event unmarshal event %s failed", bs)
	}
	return nil
}

func (r *eventReq) getToken() string {
	if r.Header != nil && r.Header.Token != "" {
		return r.Header.Token
	}
	return r.Token
}

func decryptEncryptString(encryptKey string, cryptoText string) (string, error) {
	var key []byte
	{
		h := sha256.New()
		h.Write([]byte(encryptKey))
		key = h.Sum(nil)
	}

	ciphertext, err := base64.StdEncoding.DecodeString(cryptoText)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCBCDecrypter(block, iv)
	stream.CryptBlocks(ciphertext, ciphertext)

	return string(ciphertext[:len(ciphertext)-int(ciphertext[len(ciphertext)-1])]), nil
}
