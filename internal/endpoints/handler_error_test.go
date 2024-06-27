package endpoints

import (
	"emailn/internal/internalErrors"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_HandlerError_when_endpoint_return_internal_error(t *testing.T) {
	assert := assert.New(t)
	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, internalErrors.ErrInternal
	}
	handlerFunc := HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusInternalServerError, res.Code)
	assert.Contains(res.Body.String(), internalErrors.ErrInternal.Error())
}

func Test_HandlerError_when_endpoint_return_domain_error(t *testing.T) {
	assert := assert.New(t)
	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, errors.New("domain error")
	}
	handlerFunc := HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
	assert.Contains(res.Body.String(), "domain error")
}

func Test_HandlerError_when_endpoint_return_obj_and_status(t *testing.T) {
	assert := assert.New(t)
	type bodyForTest struct {
		Id int
	}
	objExpect := bodyForTest{Id: 2}
	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return objExpect, 201, nil
	}
	handlerFunc := HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	res := httptest.NewRecorder()

	handlerFunc.ServeHTTP(res, req)

	assert.Equal(http.StatusCreated, res.Code)
	objReturned := bodyForTest{}
	json.Unmarshal(res.Body.Bytes(), &objReturned)
	assert.Equal(objExpect, objReturned)

}
