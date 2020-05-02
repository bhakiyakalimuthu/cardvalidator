package internal

import (
	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"net/http"

	"github.com/bhakiyakalimuthu/cardvalidator/pkg"
)

var _ pkg.Controller = (*Controller)(nil)
type Controller struct {
	logger *zap.Logger
}

func NewController(logger *zap.Logger) *Controller {
	return &Controller{logger: logger}
}

func (c Controller) Init(logger *zap.Logger) error {
	c.logger = logger
	return nil
}

func (c Controller) SetupRouter(router chi.Router) error {
	router.Route("/card-schema", func(r chi.Router) {
		r.Get("/verify",c.Verify)
	})
}

func (c Controller) Terminate() error {
	return nil
}

func (c Controller) Verify(w http.ResponseWriter,r *http.Request)  {

}

func (c Controller) Stats(w http.ResponseWriter,r *http.Request) {

}


