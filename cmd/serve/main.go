package main

import (
	"XCloud/dao"
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
	dst, err := os.Create("./file/" + path.Base(head.Filename))
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
	value, err := r.Cookie("xCloud_id")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if value.Value != "xxx" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userName := r.Form.Get("userName")
	if userName == "" {
		log.Println("userName:" + userName)
		return
	}

	password := r.Form.Get("password")
	if password == "" {
		log.Println("password:" + password)
		return
	}

	if userName != "123" || password != "123" {
		return
	}
	cookie := http.Cookie{
		Name:  "xCloud_id",
		Value: "xxx",
	}
	http.SetCookie(w, &cookie)
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func handleDownload(w http.ResponseWriter, request *http.Request) {
	//文件上传只允许GET方法
	if request.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("Method not allowed"))
		return
	}

	fmt.Println(request.URL)
	//文件名
	request.ParseForm()
	//q := request.URL.Query()
	filename := request.FormValue("filename")
	//filename := q.Get("filename")
	if filename == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = io.WriteString(w, "Bad request")
		return
	}
	log.Println("filename: " + filename)
	//打开文件
	file, err := os.Open("./file" + "/" + filename)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = io.WriteString(w, "Bad request")
		return
	}
	//结束后关闭文件
	defer file.Close()

	//设置响应的header头
	w.Header().Add("Content-type", "application/octet-stream")
	w.Header().Add("content-disposition", "attachment; filename=\""+filename+"\"")
	//将文件写至responseBody
	_, err = io.Copy(w, file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = io.WriteString(w, "Bad request")
		return
	}
}

func main() {
	dao.Connect(os.Args[1], os.Args[2], os.Args[2])
	a := dao.CreateUserInfoDao()
	a.DeleteUser("123x")

	http.HandleFunc("/hello", HelloServer)
	//file
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/down", handleDownload)
	http.HandleFunc("/ls", LsHandler)
	http.HandleFunc("/login", LoginHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
