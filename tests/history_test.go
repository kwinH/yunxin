package tests

import (
	"strconv"
	"testing"
	"time"
)

func TestQueryMessage(t *testing.T) {
	res, err := client.QueryMessage("1", "2", "0", strconv.FormatInt(time.Now().UnixNano(), 10), 100, 0, "")
	if err != nil {
		t.Error(err)
		return
	}

	for _, val := range *res.Data {
		t.Log(val)
	}
}
