package internal

import (
	"go.uber.org/zap"
	"net/http"
)

type Service interface {
	Verify()
	Stats()
}

var _ Service = (*DefaultService)(nil)
type DefaultService struct {
	logger *zap.Logger
	client http.Client
}

func (d DefaultService) Verify() {
	panic("implement me")
}

func (d DefaultService) Stats() {
	panic("implement me")
}





