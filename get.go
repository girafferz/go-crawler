package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"log"
	"strings"
	"os"
	"io"
	"golang.org/x/text/encoding/simplifiedchinese"
	//"golang.org/x/net/html"
	"github.com/PuerkitoBio/goquery"
)


var gbkDecoder = simplifiedchinese.GBK.NewDecoder()

func main() {
	//jsonp := getRawJsop()
	raw := readFromHtml()
	utf8str := GbkString(raw)
	writeString(utf8str, "utf8.out.htm")
	htmlParser()
}

func readFromHtml() ([]byte){
	data, err := ioutil.ReadFile(`/tmp/out.htm`)
	if err != nil {
		// エラー処理
	}
	return data
}

func GbkString(gbk []byte) string {
	gbkDecoder.Reset()
	utf8 := make([]byte, len(gbk)*2)
	utf8len, _, _ := gbkDecoder.Transform(utf8, gbk, false)
	return string(utf8[0:utf8len])
}

func writeString(in string, filename string) {
	f, err := os.Create("./" + filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(in)
	if err != nil {
		panic(err)
	}
}

func htmlParser() (string) {
	// ファイルからドキュメントを作成
	f, e := os.Open("./utf8.out.htm")
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()

	doc, e := goquery.NewDocumentFromReader(f)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(doc)

	// DOMを解析してhrefの値を変更 http://qiita.com/koduki/items/393576193c25dbb477ed
	/*
	doc.Find(".line-content").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
		//s.SetAttr("href", fmt.Sprintf("%03d.html", i+1))
	})
	*/
	selection := doc.Find(".line-content")
	selectedText := selection.Text()
	fmt.Print("--selectedText--")
	fmt.Print(selectedText)
	writeString(selectedText, "data.jsonp")
	fmt.Print("--selectedText--")

	/*
	// 変更したDOMの値をHTMLとして標準出力
	result, e := doc.Html()
	if e != nil {
		log.Fatal(e)
	}
	*/
	return selectedText
}


func isExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func evalStringToList (listString string) ([]string) {
	var listOfString []string
	dec := json.NewDecoder(strings.NewReader(listString))
	err := dec.Decode(&listOfString)
	if err != nil {
		log.Fatal(err)
	}
	return listOfString
}

func getRawJsop() ([]byte) {
	targetUrl := "https://kao.world.tmall.com/i/asynSearch.htm?_ksTS=1484324912739_422&callback=jsonp423&mid=w-14970732301-0&wid=14970732301&path=/category-1173143562.htm&&spm=a312a.7700824.w5001-14970732251.5.0veXuF&catName=%25C3%25EE%25B6%25F8%25CA%25E6&catId=1173143562&search=y&orderType=null&scene=taobao_shop&scid=1173143562"
	req, _ := http.NewRequest("GET", targetUrl, nil)
	req.Header.Set("Content-Type", "application/javascript")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.95 Safari/537.36")
	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	dataBytes, _ := ioutil.ReadAll(resp.Body)
	return dataBytes
}

