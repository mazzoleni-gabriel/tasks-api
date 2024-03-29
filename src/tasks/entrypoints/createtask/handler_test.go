package createtask_test

import (
	"bytes"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"tasks-api/src/tasks"
	"tasks-api/src/tasks/entrypoints/createtask"
	"tasks-api/src/tasks/entrypoints/createtask/mocks"
	"testing"
	"time"
)

func Test_Handle(t *testing.T) {
	const userID = 1

	t.Run("Success", func(t *testing.T) {
		payload := `{
						"summary": "summary test",
   						 "performed_at": "2022-03-18T23:11:59Z"
					}`

		expectedResponse := `{"id":1}`

		expectedEntity := tasks.Task{
			Summary:     "summary test",
			PerformedAt: time.Date(2022, 3, 18, 23, 11, 59, 0, time.UTC),
			CreatedBy:   uint(userID),
		}

		userCaseMock := &mocks.UseCae{}
		userCaseMock.On("Create", mock.Anything, expectedEntity).
			Return(uint(1), nil)

		w := runCase(userCaseMock, "/tasks", payload, strconv.Itoa(userID))

		got := w.Result()
		defer got.Body.Close()
		gotBody, err := ioutil.ReadAll(got.Body)

		assert.Nil(t, err)
		assert.JSONEq(t, expectedResponse, string(gotBody))

		assert.Equal(t, http.StatusOK, got.StatusCode) // test fixed
		userCaseMock.AssertExpectations(t)
	})

	t.Run("Should return internal error when use case general error", func(t *testing.T) {
		payload := `{
						"summary": "summary test",
   						 "performed_at": "2022-03-18T23:11:59Z"
					}`

		expectedResponse := `{"code":"internal_server_error", "message":"assert.AnError general error for testing"}`

		expectedEntity := tasks.Task{
			Summary:     "summary test",
			PerformedAt: time.Date(2022, 3, 18, 23, 11, 59, 0, time.UTC),
			CreatedBy:   uint(userID),
		}

		userCaseMock := &mocks.UseCase{}
		userCaseMock.On("Create", mock.Anything, expectedEntity).
			Return(uint(1), assert.AnError)

		w := runCase(userCaseMock, "/tasks", payload, strconv.Itoa(userID))

		got := w.Result()
		defer got.Body.Close()
		gotBody, err := ioutil.ReadAll(got.Body)

		assert.Nil(t, err)
		assert.JSONEq(t, expectedResponse, string(gotBody))

		assert.Equal(t, http.StatusInternalServerError, got.StatusCode)
		userCaseMock.AssertExpectations(t)
	})

	t.Run("Should return bad request when user id header not found", func(t *testing.T) {
		payload := `{
						"summary": "summary test",
   						 "performed_at": "2022-03-18T23:11:59Z"
					}`

		expectedResponse := `{"code":"bad_request", "message":"header X-User-ID is required"}`

		w := runCase(nil, "/tasks", payload, "")

		got := w.Result()
		defer got.Body.Close()
		gotBody, err := ioutil.ReadAll(got.Body)

		assert.Nil(t, err)
		assert.JSONEq(t, expectedResponse, string(gotBody))

		assert.Equal(t, http.StatusBadRequest, got.StatusCode)
	})

	t.Run("Should return bad request when missing required field", func(t *testing.T) {
		payload := `{
   						 "performed_at": "2022-03-18T23:11:59Z"
					}`

		expectedResponse := `{"code":"bad_request", "message":"Key: 'CreateTaskRequest.Summary' Error:Field validation for 'Summary' failed on the 'required' tag"}`

		w := runCase(nil, "/tasks", payload, "")

		got := w.Result()
		defer got.Body.Close()
		gotBody, err := ioutil.ReadAll(got.Body)

		assert.Nil(t, err)
		assert.JSONEq(t, expectedResponse, string(gotBody))

		assert.Equal(t, http.StatusBadRequest, got.StatusCode)
	})

	t.Run("Should return bad request when summary len is over 2500", func(t *testing.T) {
		payload := `{
						"summary": "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",
   						 "performed_at": "2022-03-18T23:11:59Z"
					}`

		expectedResponse := `{"code":"bad_request", "message":"Key: 'CreateTaskRequest.Summary' Error:Field validation for 'Summary' failed on the 'max' tag"}`

		w := runCase(nil, "/tasks", payload, "")

		got := w.Result()
		defer got.Body.Close()
		gotBody, err := ioutil.ReadAll(got.Body)

		assert.Nil(t, err)
		assert.JSONEq(t, expectedResponse, string(gotBody))

		assert.Equal(t, http.StatusBadRequest, got.StatusCode)
	})

	t.Run("Should return bad request when invalid payload", func(t *testing.T) {
		payload := `invalid`

		expectedResponse := `{"code":"bad_request", "message":"invalid character 'i' looking for beginning of value"}`

		w := runCase(nil, "/tasks", payload, strconv.Itoa(userID))

		got := w.Result()
		defer got.Body.Close()
		gotBody, err := ioutil.ReadAll(got.Body)

		assert.Nil(t, err)
		assert.JSONEq(t, expectedResponse, string(gotBody))

		assert.Equal(t, http.StatusBadRequest, got.StatusCode)
	})
}

func runCase(useCaseMock createtask.UseCase, url string, strPayload string, userIDHeader string) *httptest.ResponseRecorder {
	h := createtask.NewHandler(useCaseMock)
	router := chi.NewRouter()

	createtask.RegisterHandler(router, h)

	payload := bytes.NewBufferString(strPayload)
	r := httptest.NewRequest("POST", url, payload)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-User-ID", userIDHeader)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	return w
}
