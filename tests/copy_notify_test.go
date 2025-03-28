package tests

import (
	"bytes"
	"github.com/kwinh/yunxin"
	"net/http"
	"strconv"
	"testing"
	"time"
)

func TestCheckSum(t *testing.T) {
	body := []byte(`{}`)
	req, _ := http.NewRequest("POST", "http://yunxinservice.com.cn/receiveMsg.action", bytes.NewReader(body))
	curTime := strconv.FormatInt(time.Now().UnixNano(), 10)
	md5 := yunxin.Md5HashToHexString(body)
	req.Header.Set("CurTime", curTime)
	req.Header.Set("MD5", md5)
	req.Header.Set("CheckSum", yunxin.ShaHashToHexStringFromString(client.AppSecret+md5+curTime))
	t.Log("checksum:", client.AppSecret+md5+curTime, "checksum-encoded:", yunxin.ShaHashToHexStringFromString(client.AppSecret+md5+curTime))

	err := client.CheckSum(req)
	t.Log(err)
}

func TestCheckMd5(t *testing.T) {
	body := []byte(`{"body":"你好","eventType":1}`)
	req, _ := http.NewRequest("POST", "http://yunxinservice.com.cn/receiveMsg.action", bytes.NewReader(body))
	curTime := strconv.FormatInt(time.Now().UnixNano(), 10)
	md5 := yunxin.Md5HashToHexString(body)
	req.Header.Set("CurTime", curTime)
	req.Header.Set("MD5", md5)
	req.Header.Set("CheckSum", yunxin.ShaHashToHexStringFromString(client.AppSecret+md5+curTime))
	t.Log("checksum:", client.AppSecret+md5+curTime, "checksum-encoded:", yunxin.ShaHashToHexStringFromString(client.AppSecret+md5+curTime))

	bd, err := client.CheckMd5(req)
	t.Log(string(bd), err)
}
