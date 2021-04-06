package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func init() {
	log.SetPrefix("CHAT:")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {

	for idx, args := range os.Args {
		fmt.Println("参数"+strconv.Itoa(idx)+":", args)
	}

	if len(os.Args) < 2 {
		longin()
		return
	}
	//command := strconv.Itoa(1)
	command := os.Args[1]
	if command == "upload" {
		postFile(os.Args[2], "http://localhost:8080/upload")
		return
	}
	if command == "ls" {
		getDisk()
		return
	}
	if command == "login" {
		longin()
		return
	}
}

//上传文件
func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("files", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	//打开文件句柄操作
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	fmt.Println(contentType)
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))

	return nil
}

//获取网盘文件
func getDisk() {
	resp, err := http.Get("http://localhost:8080/ls")
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Printf("%s\n", b)
}

func longin() {
	urlValue := url.Values{
		"userName": {"123"},
		"password": {"123"},
	}
	reqBody := urlValue.Encode()
	//strReader := strings.NewReader("userName=123&password=123")
	resp, err := http.Post("http://localhost:8080/login",
		"application/x-www-form-urlencoded",
		strings.NewReader(reqBody))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
