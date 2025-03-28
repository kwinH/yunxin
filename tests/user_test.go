package tests

import (
	"github.com/kwinh/yunxin"
	"testing"
)

func TestUpdateUinfo(t *testing.T) {
	res, err := client.UpdateUinfo(&yunxin.ImUser{ID: "1", Birthday: "1996-01-01"})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", res)
}

func TestGetUinfo(t *testing.T) {
	res, err := client.GetUinfo([]string{"1"}, true)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", res)
}
