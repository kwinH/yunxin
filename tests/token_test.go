package tests

import (
	"github.com/kwinh/yunxin"
	"os"
	"testing"
)

var client = yunxin.CreateImClient("", "", nil)

func init() {
	os.Setenv("GOCACHE", "off")
}

func TestToken(t *testing.T) {
	user := &yunxin.ImUser{ID: "1", Name: "test", Gender: 1}
	tk, err := client.CreateImUser(user)
	if err != nil {
		t.Error(err)
	}
	t.Log(tk)
}

func TestRefreshToken(t *testing.T) {
	res, err := client.RefreshToken("1")
	if err != nil {
		t.Error(err)
	}

	t.Logf("%#v", res)
}
