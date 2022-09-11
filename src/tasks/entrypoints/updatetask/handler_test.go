package updatetask_test

import (
	"bytes"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"tasks-api/internal/apperror"
	"tasks-api/src/tasks"
	"tasks-api/src/tasks/entrypoints/updatetask"
	"tasks-api/src/tasks/entrypoints/updatetask/mocks"
	"testing"
	"time"
)

func Test_Handle(t *testing.T) {
	const userID = 1
	const taskID = 2

	t.Run("Success", func(t *testing.T) {
		payload := `{
						"summary": "summary update test",
						"performed_at": "2022-03-18T23:11:59Z"
					}`

		expectedResponse := `{"message":"resource updated successfully"}`

		expectedEntity := tasks.Task{
			ID:          uint(taskID),
			Summary:     "summary update test",
			PerformedAt: time.Date(2022, 3, 18, 23, 11, 59, 0, time.UTC),
		}

		userCaseMock := mocks.NewUseCase(t)
		userCaseMock.On("Update", mock.Anything, expectedEntity, uint(userID)).
			Return(int64(1), nil)

		w := runCase(userCaseMock, "/tasks/2", payload, strconv.Itoa(userID))

		got := w.Result()
		defer got.Body.Close()
		gotBody, err := ioutil.ReadAll(got.Body)

		assert.Nil(t, err)
		assert.JSONEq(t, expectedResponse, string(gotBody))

		assert.Equal(t, http.StatusOK, got.StatusCode)
		userCaseMock.AssertExpectations(t)
	})

	t.Run("Should return forbidden when ErrOtherUserTask app error", func(t *testing.T) {
		payload := `{
						"summary": "summary update test",
						"performed_at": "2022-03-18T23:11:59Z"
					}`

		expectedResponse := `{"code":"forbidden", "message":"test error"}`

		expectedEntity := tasks.Task{
			ID:          uint(taskID),
			Summary:     "summary update test",
			PerformedAt: time.Date(2022, 3, 18, 23, 11, 59, 0, time.UTC),
		}

		userCaseMock := mocks.NewUseCase(t)
		userCaseMock.On("Update", mock.Anything, expectedEntity, uint(userID)).
			Return(int64(0), apperror.New(apperror.ErrOtherUserTask, "test error"))

		w := runCase(userCaseMock, "/tasks/2", payload, strconv.Itoa(userID))

		got := w.Result()
		defer got.Body.Close()
		gotBody, err := ioutil.ReadAll(got.Body)

		assert.Nil(t, err)
		assert.JSONEq(t, expectedResponse, string(gotBody))

		assert.Equal(t, http.StatusForbidden, got.StatusCode)
		userCaseMock.AssertExpectations(t)
	})

	t.Run("Should return not found when ErrTaskNotFound app error", func(t *testing.T) {
		payload := `{
						"summary": "summary update test",
						"performed_at": "2022-03-18T23:11:59Z"
					}`

		expectedResponse := `{"code":"not_found", "message":"test error"}`

		expectedEntity := tasks.Task{
			ID:          uint(taskID),
			Summary:     "summary update test",
			PerformedAt: time.Date(2022, 3, 18, 23, 11, 59, 0, time.UTC),
		}

		userCaseMock := mocks.NewUseCase(t)
		userCaseMock.On("Update", mock.Anything, expectedEntity, uint(userID)).
			Return(int64(0), apperror.New(apperror.ErrTaskNotFound, "test error"))

		w := runCase(userCaseMock, "/tasks/2", payload, strconv.Itoa(userID))

		got := w.Result()
		defer got.Body.Close()
		gotBody, err := ioutil.ReadAll(got.Body)

		assert.Nil(t, err)
		assert.JSONEq(t, expectedResponse, string(gotBody))

		assert.Equal(t, http.StatusNotFound, got.StatusCode)
		userCaseMock.AssertExpectations(t)
	})

	t.Run("Should return bad request when user id header not found", func(t *testing.T) {
		payload := `{
						"summary": "summary update test",
						"performed_at": "2022-03-18T23:11:59Z"
					}`

		expectedResponse := `{"code":"bad_request", "message":"header X-User-ID is required"}`

		w := runCase(nil, "/tasks/2", payload, "")

		got := w.Result()
		defer got.Body.Close()
		gotBody, err := ioutil.ReadAll(got.Body)

		assert.Nil(t, err)
		assert.JSONEq(t, expectedResponse, string(gotBody))

		assert.Equal(t, http.StatusBadRequest, got.StatusCode)
	})

	t.Run("Should return bad request when invalid payload", func(t *testing.T) {
		payload := `invalid`

		expectedResponse := `{"code":"bad_request", "message":"header X-User-ID is required"}`

		w := runCase(nil, "/tasks/2", payload, strconv.Itoa(userID))

		got := w.Result()
		defer got.Body.Close()
		gotBody, err := ioutil.ReadAll(got.Body)

		assert.Nil(t, err)
		assert.JSONEq(t, expectedResponse, string(gotBody))

		assert.Equal(t, http.StatusBadRequest, got.StatusCode)
	})

	t.Run("Should return bad request when zero path param", func(t *testing.T) {
		payload := `{
						"summary": "summary update test",
						"performed_at": "2022-03-18T23:11:59Z"
					}`

		expectedResponse := `{"code":"bad_request", "message":"task_id is required"}`

		w := runCase(nil, "/tasks/0", payload, strconv.Itoa(userID))

		got := w.Result()
		defer got.Body.Close()
		gotBody, err := ioutil.ReadAll(got.Body)

		assert.Nil(t, err)
		assert.JSONEq(t, expectedResponse, string(gotBody))

		assert.Equal(t, http.StatusBadRequest, got.StatusCode)
	})

	t.Run("Should return bad request when invalid path param", func(t *testing.T) {
		payload := `{
						"summary": "summary update test",
						"performed_at": "2022-03-18T23:11:59Z"
					}`

		expectedResponse := `{"code":"bad_request", "message":"unnable to parse task_id: strconv.ParseUint: parsing \"invalid\": invalid syntax"}`

		w := runCase(nil, "/tasks/invalid", payload, strconv.Itoa(userID))

		got := w.Result()
		defer got.Body.Close()
		gotBody, err := ioutil.ReadAll(got.Body)

		assert.Nil(t, err)
		assert.JSONEq(t, expectedResponse, string(gotBody))

		assert.Equal(t, http.StatusBadRequest, got.StatusCode)
	})
}

func runCase(useCaseMock updatetask.UseCase, url string, strPayload string, userIDHeader string) *httptest.ResponseRecorder {
	h := updatetask.NewHandler(useCaseMock)
	router := chi.NewRouter()

	updatetask.RegisterHandler(router, h)

	payload := bytes.NewBufferString(strPayload)
	r := httptest.NewRequest("PUT", url, payload)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-User-ID", userIDHeader)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	return w
}
