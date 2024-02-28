package utils

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"proxy_server/internal/logging"
	"proxy_server/internal/pkg/dto"
	"proxy_server/internal/pkg/entities"
	"strings"

	"github.com/andybalholm/brotli"
)

func GetReqLogger(ctx context.Context) *logging.ILogger {
	if ctx == nil {
		return nil
	}
	if logger, ok := ctx.Value(dto.LoggerKey).(logging.ILogger); ok {
		return &logger
	}
	return nil
}

func GetReqID(ctx context.Context) *dto.RequestID {
	if ctx == nil {
		return nil
	}
	if id, ok := ctx.Value(dto.RequestIDKey).(uint64); ok {
		return &dto.RequestID{Value: id}
	}
	return nil
}

func ObjToRequest(obj *entities.Request) (*http.Request, error) {
	if !strings.HasPrefix(obj.Path, "/") {
		obj.Path = "/" + obj.Path
	}
	reqUrl, err := url.Parse(obj.Scheme + "://" + obj.Host + obj.Path)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(obj.Method, reqUrl.String(), bytes.NewReader(obj.RawBody))
	if err != nil {
		return nil, err
	}

	request.Header = reverseParseHeaders(obj.Headers)
	reverseParseCookies(obj.Cookies, request)
	request.URL.RawQuery = reverseParseValues(obj.GetParams).Encode()

	return request, nil
}

func reverseParseHeaders(headersMap map[string][]string) http.Header {
	headers := http.Header{}
	for header, values := range headersMap {
		for _, value := range values {
			headers.Add(header, value)
		}
	}

	return headers
}

func reverseParseCookies(cookiesMap map[string]string, request *http.Request) {
	for cookie, value := range cookiesMap {
		request.AddCookie(&http.Cookie{Name: cookie, Value: value})
	}
}

func reverseParseValues(valueMap map[string][]string) url.Values {
	urlValues := url.Values{}
	for key, values := range valueMap {
		for _, value := range values {
			urlValues.Add(key, value)
		}
	}

	return urlValues
}

func DecodeResponse(response *http.Response) ([]byte, error) {
	decodedBody := []byte{}
	var err error
	var reader io.Reader

	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(response.Body)
		if err != nil {
			return decodedBody, err
		}
	case "br":
		reader = brotli.NewReader(response.Body)
	default:
		reader = response.Body
	}

	decodedBody, err = io.ReadAll(reader)
	if err != nil {
		return decodedBody, err
	}

	return decodedBody, nil
}

func LoadDict(dictFile string) ([]string, error) {
	filenames := []string{}

	if _, err := os.Stat(dictFile); os.IsNotExist(err) {
		log.Println("File not found, downloading")
		err := downloadDict(dictFile)
		if err != nil {
			return filenames, err
		}
	}
	log.Println("File downloaded")

	file, err := os.Open(dictFile)
	if err != nil {
		log.Println("Failed to open file")
		return filenames, err
	}
	defer file.Close()
	log.Println("File opened")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		filenames = append(filenames, scanner.Text())
	}
	log.Println("File parsed")

	return filenames, scanner.Err()
}

func downloadDict(target string) error {
	file, err := os.Create(target)
	if err != nil {
		return err
	}
	log.Println("File created")
	defer file.Close()

	w := bufio.NewWriter(file)

	source := "https://raw.githubusercontent.com/maurosoria/dirsearch/master/db/dicc.txt"
	resp, err := http.Get(source)
	if err != nil {
		log.Println("Failed to download")
		return err
	}
	log.Println("Got file")

	len, err := w.ReadFrom(resp.Body)
	if err != nil {
		return err
	}
	log.Println("Wrote", len, "bytes")

	return nil
}
