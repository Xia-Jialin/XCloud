package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func init() {
	log.SetPrefix("CHAT:")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
	log.Println("123")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	//todo 获得上传的源文件s
	srcfile, head, err := r.FormFile("files")
	if err != nil {
		//util.RespFail(w, err.Error())
		log.Printf(err.Error())
		return
	}
	dst, err := os.Create("./file/" + path.Base(head.Filename) + ".upload")
	if err != nil {
		log.Printf(err.Error())
		return
	}
	defer dst.Close()
	_, err = io.Copy(dst, srcfile)
	if err != nil {
		log.Printf(err.Error())
		return
	}
}

func LsHandler(w http.ResponseWriter, r *http.Request) {
	dir, error := os.OpenFile("./file/", os.O_RDONLY, os.ModeDir)
	if error != nil {
		defer dir.Close()
		fmt.Println(error.Error())
		return
	}
	fileinfo, _ := dir.Stat()
	fmt.Println(fileinfo.IsDir())
	names, _ := dir.Readdir(-1)
	for _, name := range names {
		fmt.Fprintln(w, name.Name())
	}
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	//file
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/download", HelloServer)
	http.HandleFunc("/ls", LsHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
