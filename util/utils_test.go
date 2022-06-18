package util

import "testing"

func TestCreateCache(t *testing.T) {
	e := CreateCache("../image/img.gif")

	if e != nil {
		t.Error("error to create the cache", e)
	}
}

func TestReadFile(t *testing.T) {
	b, e := ReadFile("../image/img.gif")
	if len(b) == 0 {
		t.Error("error to read the file from a directory,", e)
	}
}

func TestCheckFile(t *testing.T) {
	c := CheckFile("../image/img.gif")
	if c == false {
		t.Error("error to check the file ")
	}
}

func TestLogData(t *testing.T) {
	e := LogData("test ", "../logs/app.log")
	if e != nil {
		t.Error("error to create log,error = ", e)
	}
}

func TestGetImageFromCache(t *testing.T) {
	ImgeCache = map[string][]byte{}
	ImgeCache["image"] = []byte("hello")
	_, e := GetImageFromCache()
	if e != nil {
		t.Error("error to read from cache ", e)
	}
}
