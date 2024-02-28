package scan

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"proxy_server/internal/apperrors"
	"proxy_server/internal/pkg/dto"
	"proxy_server/internal/storage"
	"proxy_server/internal/utils"
	"sync"

	chimw "github.com/go-chi/chi/v5/middleware"
)

type ScanService struct {
	storage storage.IRequestStorage
}

func NewScanService(requestStorage storage.IRequestStorage) *ScanService {
	return &ScanService{
		storage: requestStorage,
	}
}

func (rs ScanService) ScanRequest(ctx context.Context, requestID *dto.RequestID, client http.Client, dict string) ([]string, error) {
	funcName := "ScanRequest"
	logger := *utils.GetReqLogger(ctx)
	if logger == nil {
		return nil, apperrors.ErrLoggerMissingFromContext
	}
	reqID := chimw.GetReqID(ctx)

	requestObj, err := rs.storage.GetRequestByID(ctx, requestID)
	if err != nil {
		return nil, err
	}

	files, err := utils.LoadDict(dict)
	if err != nil {
		return nil, err
	}

	c := make(chan string)
	wg := &sync.WaitGroup{}
	var validFiles []string

	for _, file := range files {
		logger.DebugFmt("Trying file "+file, reqID, funcName, "service")
		wg.Add(1)
		requestObj.Path = url.QueryEscape(file)
		request, err := utils.ObjToRequest(requestObj)
		logger.DebugFmt("Created request", reqID, funcName, "service")
		if err != nil {
			return nil, err
		}
		go makeRequest(request, client, c, wg)
	}

	for validFile := range c {
		validFiles = append(validFiles, validFile)
	}

	wg.Wait()

	close(c)

	return validFiles, nil
}

func makeRequest(request *http.Request, client http.Client, c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := client.Do(request)
	if err != nil {
		log.Println("With path", request.URL.Path, "error performing request", err)
		return
	}
	log.Println("Got response from trying", request.URL.Path)
	log.Println("Code", resp.StatusCode)
	if resp.StatusCode != http.StatusNotFound {
		c <- request.URL.Path
	}
}
