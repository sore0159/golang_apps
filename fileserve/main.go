package main

import (
	"archive/tar"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

const (
	CONTENT_DIR    = "music/"
	REQUEST_PREFIX = "music/"
)

func main() {
	log.Println("STARTING SERVER AT :8888")
	if err := http.ListenAndServe(":8888", server{}); err != nil {
		log.Println("ERROR STARTING SERVER:", err)
	}

}

func TranslateRequest(requestPath string) (filePath string) {
	return CONTENT_DIR + strings.TrimPrefix(path.Clean(requestPath), REQUEST_PREFIX)
}

type server struct{}

func (server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rIP := r.Header.Get("x-forwarded-for")
	if !strings.HasPrefix(rIP, "192.168.1.") && !strings.HasPrefix(rIP, "127.0.0.1") {
		http.Error(w, "Does not support nonlocal connections", 500)
		return
	}
	if strings.Index(r.URL.Path, "..") != -1 {
		http.Error(w, "Does not support .. paths", 500)
		return
	}
	fPath := TranslateRequest(r.URL.Path)
	if len(fPath) > 4 && fPath[len(fPath)-4:] == ".tar" {
		dir := fPath[:len(fPath)-4]
		ServeTar(w, r, dir)
		return
	}
	http.ServeFile(w, r, fPath)
}

func ServeTar(w http.ResponseWriter, r *http.Request, dir string) {
	dirList, err := ioutil.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "", 404)
		} else {
			log.Println("ServeTar stat error:", err)
			http.Error(w, "File stat error", 500)
		}
		return
	}
	w.Header().Set("Content-Type", "application/x-tar")
	tw := tar.NewWriter(w)
	for _, d := range dirList {
		if !WriteFile(path.Join(dir, d.Name()), tw) {
			http.Error(w, "File Writing Error", 500)
			return
		}
	}
	err = tw.Close()
	if err != nil {
		log.Println("TW Close error:", err)
	}
}

func WriteFile(fileName string, tw *tar.Writer) (ok bool) {
	info, err := os.Stat(fileName)
	if err != nil {
		log.Println("WriteFile stat error:", err)
		return false
	}
	header, err := tar.FileInfoHeader(info, "")
	if err != nil {
		log.Println("WriteFile make header error:", err)
		return false
	}
	data, err := os.Open(fileName)
	if err != nil {
		log.Println("WriteFile open file error:", err)
		return false
	}
	if err := tw.WriteHeader(header); err != nil {
		log.Println("WriteFile tar write header error:", err)
		return false
	}
	if _, err := io.Copy(tw, data); err != nil {
		log.Println("ServeTar tar write file error:", err)
		return false
	}
	return true
}

func ServeScripts(w http.ResponseWriter, fNames []string) {
	w.Header().Set("Content-Type", "application/javascript")
	for _, f := range fNames {
		file, err := os.Open("FILES/static/js/" + f + ".js")
		if err != nil {
			log.Println("serveScript file open failure: ", err)
			http.Error(w, "file open failure", 500)
			return
		}
		if _, err = io.Copy(w, file); err != nil {
			log.Println(err, "serveScript data copy failure: ", err)
			http.Error(w, "data copy failure", 500)
			return
		}
	}
}
