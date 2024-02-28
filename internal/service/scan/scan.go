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

func (rs ScanService) ScanRequest(ctx context.Context, requestID *dto.RequestID, client http.Client, dict string) (map[string]int, error) {
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

	result := map[string]int{}

	for _, file := range files {
		logger.DebugFmt("Trying file "+file, reqID, funcName, "service")
		requestObj.Path = url.QueryEscape(file)
		request, err := utils.ObjToRequest(requestObj)
		logger.DebugFmt("Created request", reqID, funcName, "service")
		if err != nil {
			return nil, err
		}
		if respCode := makeRequest(request, client); respCode > 0 {
			result[file] = respCode
		}
	}

	return result, nil
}

func makeRequest(request *http.Request, client http.Client) int {
	resp, err := client.Do(request)
	if err != nil {
		log.Println("With path", request.URL.Path, "error performing request", err)
		return -1
	}
	log.Println("Got response from trying", request.URL.Path)
	log.Println("Code", resp.StatusCode)
	if resp.StatusCode != http.StatusNotFound {
		return resp.StatusCode
	}
	return -1
}
