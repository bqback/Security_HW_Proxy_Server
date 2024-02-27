package repeat

import (
	"context"
	"net/http"
	"proxy_server/internal/pkg/dto"
	"proxy_server/internal/storage"
	"proxy_server/internal/utils"
)

type RepeatService struct {
	reqs  storage.IRequestStorage
	resps storage.IResponseStorage
}

func NewRepeatService(requestStorage storage.IRequestStorage, responseStorage storage.IResponseStorage) *RepeatService {
	return &RepeatService{
		reqs:  requestStorage,
		resps: responseStorage,
	}
}

func (rs RepeatService) RepeatRequest(ctx context.Context, requestID *dto.RequestID, client http.Client) (*http.Response, error) {
	requestObj, err := rs.reqs.GetRequestByID(ctx, requestID)
	if err != nil {
		return nil, err
	}

	request, err := utils.ObjToRequest(requestObj)
	if err != nil {
		return nil, err
	}

	return client.Do(request)
}
