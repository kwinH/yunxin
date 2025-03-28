package tests

import (
	"bytes"
	"github.com/kwinh/yunxin"
	"net/http"
	"strconv"
	"testing"
	"time"
)

func TestCallback(t *testing.T) {
	body := []byte(`{"body":"Hello","eventType":1,"fromAccount":"000266","fromClientType":"WEB","fromDeviceId":"617715aa8579db03f0cf054c199****","fromNick":"yj000266","msgTimestamp":"1541560157286","msgType":"TEXT","msgidClient":"","to":"005877","fromClientIp":"115.211.51.**","fromClientPort":"568**"}`)
	req, _ := http.NewRequest("POST", "http://yunxinservice.com.cn/receiveMsg.action", bytes.NewReader(body))
	curTime := strconv.FormatInt(time.Now().UnixNano(), 10)
	md5 := yunxin.Md5HashToHexString(body)
	req.Header.Set("CurTime", curTime)
	req.Header.Set("MD5", md5)
	req.Header.Set("CheckSum", yunxin.ShaHashToHexStringFromString(client.AppSecret+md5+curTime))
	t.Log("checksum:", client.AppSecret+md5+curTime, "checksum-encoded:", yunxin.ShaHashToHexStringFromString(client.AppSecret+md5+curTime))

	fun := func(param *yunxin.CallbackEvent) (yunxin.CallbackResponse, error) {
		t.Logf("callback param: %+v", param)
		return yunxin.CallbackResponse{
			ErrCode:     0,
			CallbackExt: "回调成功",
		}, nil
	}
	res, err := client.Callback(req, fun)

	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("callback response: %+v", res)
}
