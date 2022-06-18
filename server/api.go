package server

import (
	"Users/nafisalam/go/projects/assignment-sojern/config"
	"Users/nafisalam/go/projects/assignment-sojern/util"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gorilla/mux"
)

type webconfig struct {
	Cong *config.Config
}

// NewRouter to create a new router with all the apis end
func NewRouter(cfg config.Config) *mux.Router {
	r := mux.NewRouter()
	wc := webconfig{&cfg}
	r.HandleFunc("/ping", wc.ping)
	r.HandleFunc("/img", wc.img)
	return r
}

// ping is the handler to server the ping requests
func (c webconfig) ping(w http.ResponseWriter, r *http.Request) {
	c.LogRequest(r)
	ok := util.CheckFile(c.Cong.StatusFilePath)
	if ok {
		w.WriteHeader(http.StatusOK)
		b := []byte("ok")
		w.Write(b)
		return
	}
	b := []byte("file not avaiable")
	w.WriteHeader(http.StatusServiceUnavailable)
	w.Write(b)
}

// img is the handler to server /img requests
func (c webconfig) img(w http.ResponseWriter, r *http.Request) {
	c.LogRequest(r)
	b, err := util.GetImageFromCache()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("application error"))
		log.Println("error to find the image ,error =", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
	log.Println("End: Request for a GIF image ")
}

// LogRequet is to dump the requests to a log file.
func (c webconfig) LogRequest(r *http.Request) {
	request, err := httputil.DumpRequest(r, true)
	if err != nil {
		util.LogData("error to read the request ", c.Cong.LogsPath)
	}
	util.LogData(string(request), c.Cong.LogsPath)
}
