package tests

import (
	"github.com/kwinh/yunxin"
	"testing"
)

func TestSendTextMessage(t *testing.T) {
	msg := &yunxin.TextMessage{Message: "hello world"}
	_, err := client.SendTextMessage("", "", msg, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestSendBatchTextMessage(t *testing.T) {
	msg := &yunxin.TextMessage{Message: "hello world"}
	str, err := client.SendBatchTextMessage("1", []string{"169143"}, msg, nil)
	t.Log(str)
	if err != nil {
		t.Error(err)
	}
}

func TestSendBatchAttachMessage(t *testing.T) {
	res, err := client.SendBatchAttachMsg("1", "{'msg':'test'}", []string{"2", "3"}, nil)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("res:%+v", res)
}

func TestBroadcastMsg(t *testing.T) {
	res, err := client.BroadcastMsg("欢迎来到我的世界", "", nil, nil)
	t.Logf("res %#v,err:%+v", res, err)
}

func TestRecallMsg(t *testing.T) {
	res, err := client.RecallMessage("13741575721050", "1742869050203", "1", "2", 7)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("res:%+v", res)
}
