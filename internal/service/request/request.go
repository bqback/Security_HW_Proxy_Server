package request

import (
	"context"
	"proxy_server/internal/pkg/dto"
	"proxy_server/internal/pkg/entities"
	"proxy_server/internal/storage"
)

type RequestService struct {
	reqs  storage.IRequestStorage
	resps storage.IResponseStorage
}

func NewRequestService(requestStorage storage.IRequestStorage, responseStorage storage.IResponseStorage) *RequestService {
	return &RequestService{
		reqs:  requestStorage,
		resps: responseStorage,
	}
}

func (rs RequestService) StoreRequest(ctx context.Context, request *dto.IncomingRequest) (*dto.RequestID, error) {
	return rs.reqs.StoreRequest(ctx, request)
}

func (rs RequestService) GetAllRequests(ctx context.Context) (*[]entities.Request, error) {
	requests, err := rs.reqs.GetAllRequests(ctx)
	if err != nil {
		return nil, err
	}

	iterableReqs := *requests

	for pos, request := range iterableReqs {
		response, err := rs.resps.GetResponseByRequestID(ctx, &dto.RequestID{Value: request.ID})
		if err != nil {
			return nil, err
		}

		iterableReqs[pos].Response = *response
	}

	requests = &iterableReqs

	return requests, nil
}

func (rs RequestService) GetSingleRequest(ctx context.Context, requestID *dto.RequestID) (*entities.Request, error) {
	request, err := rs.reqs.GetRequestByID(ctx, requestID)
	if err != nil {
		return nil, err
	}

	response, err := rs.resps.GetResponseByRequestID(ctx, &dto.RequestID{Value: request.ID})
	if err != nil {
		return nil, err
	}

	request.Response = *response

	return request, nil
}
