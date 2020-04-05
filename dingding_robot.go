package dingtalk_robot

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cyanBone/dingtalk_robot/message"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Client struct {
	mu        sync.Mutex     //参数锁
	signature []byte         //签名
	webHook   string         //机器人的Webhook地址
	Client    *http.Client   //http client
	Location  *time.Location //时区
}

type dingDingResult struct {
	Code    int    `json:"errcode"`
	Meesage string `json:"errmsg"`
}

//初始化client
//webhook    机器人的Webhook地址
//signature  机器人的signature数据
func New(webHook, signature string) (client *Client, err error) {
	if webHook == "" {
		return client, errors.New("webHook address is error")
	}
	location, _err := time.LoadLocation("Asia/Chongqing")
	if _err != nil {
		return client, _err
	}
	client = new(Client)
	client.Client = http.DefaultClient
	client.Location = location
	client.mu.Lock()
	client.webHook = webHook
	client.signature = []byte(signature)
	client.mu.Unlock()
	return client, nil
}

//构造请求
func (client *Client) doReq(v []byte) (dingDingResult, error) {
	result := dingDingResult{}
	u := ""
	if client.signature != nil {
		sign, s, err_ := client.sign()
		if err_ != nil {
			return result, err_
		}
		u = fmt.Sprintf("%s&timestamp=%s&sign=%s", client.webHook, s, sign)
	} else {
		u = client.webHook
	}
	req, err := http.NewRequest(http.MethodPost, u, bytes.NewReader(v))
	if err != nil {
		return result, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, _err := client.Client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if _err != nil {
		return result, _err
	}
	body, __err := ioutil.ReadAll(resp.Body)
	if __err != nil {
		return result, __err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

//发送信息
func (client *Client) Send(s message.Builder) error {
	marshal, err_ := json.Marshal(s.Build())
	if err_ != nil {
		return err_
	}
	req, err := client.doReq(marshal)
	if err != nil {
		return err
	}
	if req.Code != 0 {
		return errors.New(fmt.Sprintf("errcode:%d errmsg:%s", req.Code, req.Meesage))
	}
	return nil
}

//签名
func (client *Client) sign() (string, string, error) {
	t := time.Now().In(client.Location).Unix()
	secret := fmt.Sprintf("%d000\n%s", t, client.signature)
	hash := hmac.New(sha256.New, client.signature)
	_, err := io.WriteString(hash, secret)
	if err != nil {
		return "", "", err
	}
	return url.QueryEscape(base64.StdEncoding.EncodeToString(hash.Sum(nil))), fmt.Sprintf("%d000", t), nil
}
