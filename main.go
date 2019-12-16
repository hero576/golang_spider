package main

import (
	"bufio"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/lestrrat/go-libxml2"
	"github.com/lestrrat/go-libxml2/types"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			logs.Error(err)
		}
	}()
	if resp.StatusCode != http.StatusOK {
		logs.Error("error status code", resp.StatusCode)
		return
	}
	r := resp.Body
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		logs.Error("io read resp body err", err)
		return
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	//utf8Reader := transform.NewReader(resp.Body,simplifiedchinese.GBK.NewDecoder())
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	//all,err :=ioutil.ReadAll(utf8Reader)
	//if err!=nil{
	//	panic(err)
	//}
	get_city_list(utf8Reader)
}

func get_city_list(content io.Reader) {
	var (
		doc types.Document
		err error
	)

	if doc, err = libxml2.ParseHTMLReader(content); err != nil {
		logs.Error(err)
	} else {
		defer doc.Free()
		nodes, err := doc.Find("//dl[contains(@class,'city-list')]//a")
		fmt.Println(nodes.NodeList()[0].TextContent(), err)
	}
}
