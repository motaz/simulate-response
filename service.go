package main

import (
	"fmt"

	"io/ioutil"
	"net/http"
	"strings"

	"github.com/motaz/codeutils"
)

func writeLog(event string) {

	codeutils.WriteToLog(event, "simulate")
}

func handleRequests(w http.ResponseWriter, r *http.Request) {

	path := r.RequestURI
	file := strings.ReplaceAll(path, "/", "")
	var ext string
	contentType := r.Header.Get("content-type")
	writeLog("request from: " + codeutils.GetRemoteIP(r) + " : " + contentType)
	writeLog(path)
	if file == "" {
		fmt.Fprintf(w, "Please call /path")
	} else {
		if contentType == "application/json" {
			ext = ".json"

		} else if contentType == "application/xml" {
			ext = ".xml"
		} else if contentType == "text/plain" || contentType == "" {
			ext = ".txt"
		}
		file += ext
		if codeutils.IsFileExists(file) {

			contents, err := ioutil.ReadFile(file)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, err.Error())
				writeLog(err.Error())

			} else {
				w.Header().Add("content-type", contentType)
				w.Write(contents)
				writeLog(string(contents))
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "%s not found", file)
			writeLog("Not found")
		}
	}
}
