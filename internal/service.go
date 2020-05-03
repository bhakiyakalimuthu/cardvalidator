package internal

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

const baseUrl  =  "https://lookup.binlist.net/"
type Service interface {
	Verify(ctx context.Context, cardNumber string) (*VerifyResponse,error)
	Stats(ctx context.Context) (*StatsResponse,error)
}

var _ Service = (*DefaultService)(nil)
type DefaultService struct {
	logger *zap.Logger
	client http.Client
}

func NewDefaultService(logger *zap.Logger, client http.Client) *DefaultService {
	return &DefaultService{
		logger: logger,
		client: client,
	}
}
func (d *DefaultService) Verify(ctx context.Context, cardNumber string) (*VerifyResponse,error){
	url:= baseUrl+cardNumber
	req, err := http.NewRequest(http.MethodGet,url,nil)
	if err!=nil {
		d.logger.Error("failed to create verify request",zap.Error(err))
		return nil,err
	}
	resp,err := d.client.Do(req)
	var serverResp ServerVerifyResp
	_ = json.NewDecoder(resp.Body).Decode(&serverResp)
	return constructVerifyResponse(serverResp),nil
}

func (d *DefaultService) Stats(ctx context.Context) (*StatsResponse,error){
	panic("implement me")
}

func constructVerifyResponse(serverResp ServerVerifyResp) *VerifyResponse {
	return &VerifyResponse{
		Success: true,
		Payload: VerifyResponsePayload{
			Schema: serverResp.Scheme,
			Type:   serverResp.Type,
			Bank:   serverResp.Bank.Name,
		},
	}
}



