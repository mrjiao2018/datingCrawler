package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/text/encoding/unicode"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

// 根据url加载指定网址，并将其转换成utf-8的格式打开
func Fetch(url string) (bytes []byte, err error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code : %d ", resp.StatusCode)
	}

	// 将网站字符集转换成utf-8字符集打开
	reader := bufio.NewReader(resp.Body)
	e := determineEncoding(reader)
	utf8Reader := transform.NewReader(reader, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
	//return ioutil.ReadAll(resp.Body)
}

//func determineEncoding(r io.Reader) encoding.Encoding {
//	//todo 此处还是会先读取1024个字节
//	bytes, err := bufio.NewReader(r).Peek(1024)
//	if err != nil {
//		log.Printf("Fetcher error: %v", err)
//		return unicode.UTF8
//	}
//	e, _, _ := charset.DetermineEncoding(bytes, "")
//	return e
//}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
