package yunxin

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const neteaseBaseURL = "https://api.yunxinapi.com/nimserver"

// ImClient .
type ImClient struct {
	AppKey    string
	AppSecret string
	Nonce     string

	ctx        context.Context
	httpClient *http.Client
}

// CreateImClient  创建im客户端，proxy留空表示不使用代理
func CreateImClient(appkey, appSecret string, httpClient *http.Client) *ImClient {
	c := &ImClient{
		AppKey:     appkey,
		AppSecret:  appSecret,
		Nonce:      RandStringBytesMaskImprSrc(64),
		httpClient: httpClient,
	}

	if c.httpClient == nil {
		c.httpClient = &http.Client{}
	}

	return c
}

func (c *ImClient) SetCtx(ctx context.Context) *ImClient {
	c.ctx = ctx
	return c
}

func (c *ImClient) Curl(method, urlAddr string, data map[string]string) (*http.Response, []byte, error) {
	defer func() {
		c.ctx = nil
	}()

	formData := url.Values{}
	for k, v := range data {
		formData.Set(k, v)
	}
	payloadBytes := []byte(formData.Encode())

	body := bytes.NewBuffer(payloadBytes)

	ctx := c.ctx
	if ctx == nil {
		ctx = context.Background()
	}
	req, err := http.NewRequestWithContext(ctx, method, urlAddr, body)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Set("Accept", "application/json;charset=utf-8")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8;")
	req.Header.Set("AppKey", c.AppKey)
	req.Header.Set("Nonce", c.Nonce)
	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	req.Header.Set("CurTime", timeStamp)
	req.Header.Set("CheckSum", ShaHashToHexStringFromString(c.AppSecret+c.Nonce+timeStamp))

	// 执行请求
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}

	// 延迟关闭响应体，确保调用方可以读取
	defer resp.Body.Close()

	// 读取响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp, nil, err
	}

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		return resp, respBody, fmt.Errorf("HTTP request failed: %s, status code: %d", resp.Status, resp.StatusCode)
	}

	return resp, respBody, nil
}
