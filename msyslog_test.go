package msyslog

import (
	"fmt"
	"net"
	"testing"
)

func TestMain(t *testing.T) {
	addr := net.ParseIP("ff05::514")
	var port uint16 = 514
	mlogger, err := New(&addr, port)
	if err != nil {
		t.Error(err)
	}
	defer mlogger.Close()
	fmt.Fprint(mlogger, "This is a test message.")
}
