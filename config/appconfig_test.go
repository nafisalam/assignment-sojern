package config

import (
	"fmt"
	"testing"
)

func TestInitConfig(t *testing.T) {
	c := InitConfig("../.conf")
	fmt.Println(c)
	if c.BannerPath != "./image/banner.txt" {
		t.Error("telnet not configured ")
	}
}
