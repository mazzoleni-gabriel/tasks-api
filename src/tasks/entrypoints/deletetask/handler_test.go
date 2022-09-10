package deletetask_test

import (
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"tasks-api/src/tasks/entrypoints/deletetask"
	"tasks-api/src/tasks/entrypoints/deletetask/mocks"
	"testing"
)

func Test_Handle(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		expectedResponse := `{
    							"message": "resource deleted successfully"
							}
						`

		userCaseMock := mocks.NewUseCase(t)
		userCaseMock.On("Delete", mock.Anything, uint(1)).
			Return(int64(1), nil)

		w := runCase(userCaseMock, "/tasks/1")

		got := w.Result()
		defer got.Body.Close()
		gotBody, err := ioutil.ReadAll(got.Body)

		assert.Nil(t, err)
		assert.JSONEq(t, expectedResponse, string(gotBody))

		assert.Equal(t, http.StatusOK, got.StatusCode)
		userCaseMock.AssertExpectations(t)
	})

	t.Run("Should return not found when no rows affected", func(t *testing.T) {
		expectedResponse := `{"code":"not_found", "message":"the resource with id 1 does not exists"}`

		userCaseMock := mocks.NewUseCase(t)
		userCaseMock.On("Delete", mock.Anything, uint(1)).
			Return(int64(0), nil)

		w := runCase(userCaseMock, "/tasks/1")

		got := w.Result()
		defer got.Body.Close()
		gotBody, err := ioutil.ReadAll(got.Body)

		assert.Nil(t, err)
		assert.JSONEq(t, expectedResponse, string(gotBody))

		assert.Equal(t, http.StatusNotFound, got.StatusCode)
		userCaseMock.AssertExpectations(t)
	})

	t.Run("Should return internal error when use case fails", func(t *testing.T) {
		expectedResponse := `{"code":"internal_server_error", "message":"assert.AnError general error for testing"}`

		userCaseMock := mocks.NewUseCase(t)
		userCaseMock.On("Delete", mock.Anything, uint(1)).
			Return(int64(0), assert.AnError)

		w := runCase(userCaseMock, "/tasks/1")

		got := w.Result()
		defer got.Body.Close()
		gotBody, err := ioutil.ReadAll(got.Body)

		assert.Nil(t, err)
		assert.JSONEq(t, expectedResponse, string(gotBody))

		assert.Equal(t, http.StatusInternalServerError, got.StatusCode)
		userCaseMock.AssertExpectations(t)
	})

	t.Run("Should return bad request when invalid path param", func(t *testing.T) {
		expectedResponse := `{"code":"bad_request", "message":"unnable to parse task_id: strconv.ParseUint: parsing \"invalid\": invalid syntax"}`

		w := runCase(nil, "/tasks/invalid")

		got := w.Result()
		defer got.Body.Close()
		gotBody, err := ioutil.ReadAll(got.Body)

		assert.Nil(t, err)
		assert.JSONEq(t, expectedResponse, string(gotBody))

		assert.Equal(t, http.StatusBadRequest, got.StatusCode)
	})

	t.Run("Should return bad request when zero path param", func(t *testing.T) {
		expectedResponse := `{"code":"bad_request", "message":"task_id is required"}`

		w := runCase(nil, "/tasks/0")

		got := w.Result()
		defer got.Body.Close()
		gotBody, err := ioutil.ReadAll(got.Body)

		assert.Nil(t, err)
		assert.JSONEq(t, expectedResponse, string(gotBody))

		assert.Equal(t, http.StatusBadRequest, got.StatusCode)
	})

}

func runCase(useCaseMock deletetask.UseCase, url string) *httptest.ResponseRecorder {
	h := deletetask.NewHandler(useCaseMock)
	router := chi.NewRouter()

	deletetask.RegisterHandler(router, h)

	r := httptest.NewRequest("DELETE", url, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	return w
}
