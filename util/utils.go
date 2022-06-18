package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// ImageCache is to cache the image at start up of the application.
var ImgeCache map[string][]byte

// CreateBanner create a banner in the console .
func CreateBanner(p string) bool {
	banner, err := ioutil.ReadFile(p)
	if err != nil {
		return false
	}
	fmt.Println(string(banner))
	return true
}

// CheckFile funtion check if the file present in a directory
func CheckFile(p string) bool {
	_, err := ioutil.ReadFile(p)
	if err != nil {
		fmt.Println("could not read the file ", err)
		return false
	}
	return true
}

// ReadFile to read a file from a directory
func ReadFile(p string) ([]byte, error) {
	b, err := ioutil.ReadFile(p)
	if err != nil {
		fmt.Println("could not read the file ", err)
		return []byte{}, err
	}
	return b, nil
}

// LogData to log the data in a log file
func LogData(log string, logFilePath string) error {
	fmt.Println(log)
	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}
	defer f.Close()
	fmt.Fprintf(f, log+"\n")
	return nil
}

// CreateCache is to create the cache
func CreateCache(p string) error {
	ImgeCache = map[string][]byte{}
	b, err := ReadFile(p)

	if err != nil {
		return err
	}
	ImgeCache["image"] = b
	return nil
}

// GetImageFromCache to get the image from cache
func GetImageFromCache() ([]byte, error) {
	if b, ok := ImgeCache["image"]; ok {
		return b, nil
	}
	return nil, errors.New("error to get the iamge fromm cache")
}
