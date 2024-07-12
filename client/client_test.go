// client_test.go
package client

import (
	"fmt"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	c := NewClient("192.168.0.132:3389", "administrator", "Jhadmin123", TC_RDP, nil)
	err := c.Login()
	if err != nil {
		fmt.Println("Login:", err)
	}
	c.OnReady(func() {
		fmt.Println("ready")
	})
	time.Sleep(100 * time.Second)
}
