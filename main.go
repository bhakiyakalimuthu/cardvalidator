package main

import (
	"github.com/voi-go/svc"
	"net/http"
	"time"

	"github.com/bhakiyakalimuthu/cardvalidator/internal"
	"github.com/bhakiyakalimuthu/cardvalidator/pkg"
)

func main()  {
	client := http.Client{
		Timeout:       15*time.Second,
	}
	s,err:= svc.New("cardvalidator","snap-shot")
	svc.MustInit(s,err)
	service:= internal.NewDefaultService(s.Logger(),client)
	ctrl := internal.NewController(s.Logger(),service)
	chiWorker := pkg.NewWorker(ctrl)
	s.AddWorker("chi-worker",chiWorker)
	s.Run()
}

