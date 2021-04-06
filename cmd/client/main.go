package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/howeyc/gopass"
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
		return
	}
	userName := os.Args[1]
	password := GetPassword()
	resp, err := longin(userName, password)
	if err != nil {

	}
	var command, value, value2 string

	for {
		fmt.Print("xcloud>>")
		fmt.Scanln(&command, &value, &value2)

		if command == "upload" {
			postFile(value, "http://localhost:8080/upload")
			continue
		}
		if command == "ls" {
			getDisk(resp)
			continue
		}
		if command == "login" {
			longin(value, value2)
			continue
		}
		if command == "exit" {
			fmt.Println("bye!")
			break
		}
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
func getDisk(client *http.Client) {
	resp, err := client.Get("http://localhost:8080/ls")
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.Status)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Printf("%s\n", b)
}

func longin(userName, password string) (*http.Client, error) {
	jar, _ := cookiejar.New(nil)

	client := http.Client{
		Jar: jar,
	}

	urlValue := url.Values{
		"userName": {userName},
		"password": {password},
	}

	reqBody := urlValue.Encode()
	resp, err := client.Post("http://localhost:8080/login", "application/x-www-form-urlencoded", strings.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &client, nil
}

func GetPassword() string {
	fmt.Print("password:")
	pass, err := gopass.GetPasswd()
	if err != nil {
		return ""
	}
	return string(pass)
}
