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

		userCaseMock := &mocks.UseCase{}
		userCaseMock.On("Create", mock.Anything, expectedEntity).
			Return(uint(1), nil)

		w := runCase(userCaseMock, "/tasks", payload, strconv.Itoa(userID))

		got := w.Result()
		defer got.Body.Close()
		gotBody, err := ioutil.ReadAll(got.Body)

		assert.Nil(t, err)
		assert.JSONEq(t, expectedResponse, string(gotBody))

		assert.Equal(t, http.StatusOK, got.StatusCode)
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
