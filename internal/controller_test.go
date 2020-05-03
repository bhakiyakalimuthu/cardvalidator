package internal

import (
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testControllerSuite struct {
	suite.Suite
	recorder *httptest.ResponseRecorder
	router chi.Router
}

func TestControllerSuite(t *testing.T){
	suite.Run(t, new(testControllerSuite))
}


func (suite *testControllerSuite) SetupTest() {
	suite.recorder = httptest.NewRecorder()
	suite.router = chi.NewRouter()

	service := NewDefaultService(zap.NewNop(),http.Client{})
	ctrl := NewController(zap.NewNop(),service)
	ctrl.SetupRouter(suite.router)
}

func (suite *testControllerSuite) TestVerify()  {
	req:= suite.Require()
	request:= httptest.NewRequest(http.MethodGet, "/card-schema/verify/1234",nil)
	suite.router.ServeHTTP(suite.recorder, request)
	response :=  suite.recorder.Result()
	req.Equal(http.StatusOK, response.StatusCode)
}