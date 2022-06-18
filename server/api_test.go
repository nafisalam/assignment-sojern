package server

import (
	"Users/nafisalam/go/projects/assignment-sojern/config"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	config := config.Config{LogsPath: "./logs/app.log", ImagePath: "./image/img.gif", BannerPath: "./image/banner.txt", StatusFilePath: "../tmp/ok"}
	wc := webconfig{&config}
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	w := httptest.NewRecorder()
	wc.ping(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "ok" {
		t.Errorf("expected ok got %v", string(data))
	}
}

func TestImg(t *testing.T) {
	ImgeCache := map[string][]byte{}
	ImgeCache["image"] = []byte("hello")
	config := config.Config{LogsPath: "./logs/app.log", ImagePath: "./image/img.gif", BannerPath: "./image/banner.txt", StatusFilePath: "../tmp/ok"}
	wc := webconfig{&config}
	req := httptest.NewRequest(http.MethodGet, "/img", nil)
	w := httptest.NewRecorder()
	wc.img(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if len(data) == 0 {
		t.Errorf("expected more than 0 got %v", string(data))
	}
}
