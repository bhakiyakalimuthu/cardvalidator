package internal

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"net/http"

	"github.com/bhakiyakalimuthu/cardvalidator/pkg"
)

const (
	HTTPContentTypeKey = "content-type"
	HTTPContentTypeValue = "application/json"

)
var _ pkg.Controller = (*Controller)(nil)

type Controller struct {
	logger *zap.Logger
	service Service
}

func NewController(logger *zap.Logger, service Service) *Controller {
	return &Controller{
		logger:  logger,
		service: service,
	}
}

func (c *Controller) Init(logger *zap.Logger) error {
	c.logger = logger
	return nil
}

func (c *Controller) SetupRouter(router chi.Router) error {
	router.Route("/card-schema", func(r chi.Router) {
		r.Get("/verify/{num}",c.Verify)
	})
	return nil
}

func (c *Controller) Terminate() error {
	return nil
}

func (c *Controller) Verify(w http.ResponseWriter,r *http.Request)  {
	cardNum := chi.URLParam(r,"num")
	v,err := c.service.Verify(r.Context(),cardNum)
	if err!= nil {
		writeResp(w,http.StatusInternalServerError,err)
	}
	writeResp(w,http.StatusOK,v)
}

func (c Controller) Stats(w http.ResponseWriter,r *http.Request) {

}

func writeResp(w http.ResponseWriter,statusCode int, data interface{})  {
	w.Header().Set(HTTPContentTypeKey,HTTPContentTypeValue)
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(&data);err !=nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

