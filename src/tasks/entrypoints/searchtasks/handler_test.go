package searchtasks_test

import (
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"tasks-api/src/tasks"
	"tasks-api/src/tasks/entrypoints/searchtasks"
	"tasks-api/src/tasks/entrypoints/searchtasks/mocks"
	"testing"
	"time"
)

func Test_Handle(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		expectedResponse := `
							{
							"list": [
								{
									"id": 1,
									"summary": "summary test",
									"performed_at": "2022-03-18T23:11:59Z",
									"created_by": 123
								},
								{
									"id": 2,
									"summary": "summary test",
									"performed_at": "2022-03-18T23:11:59Z",
									"created_by": 123
								}
							]
						}
						`

		useCaseTasks := []tasks.Task{
			{
				ID:          1,
				Summary:     "summary test",
				CreatedBy:   123,
				PerformedAt: time.Date(2022, 3, 18, 23, 11, 59, 0, time.UTC),
			},
			{
				ID:          2,
				Summary:     "summary test",
				CreatedBy:   123,
				PerformedAt: time.Date(2022, 3, 18, 23, 11, 59, 0, time.UTC),
			},
		}

		userCaseMock := mocks.NewUseCase(t)
		userCaseMock.On("Search", mock.Anything, tasks.SearchFilters{}).
			Return(useCaseTasks, nil)

		w := runCase(userCaseMock, "/tasks/search")

		got := w.Result()
		defer got.Body.Close()
		gotBody, err := ioutil.ReadAll(got.Body)

		assert.Nil(t, err)
		assert.JSONEq(t, expectedResponse, string(gotBody))

		assert.Equal(t, http.StatusOK, got.StatusCode)
		userCaseMock.AssertExpectations(t)
	})

	t.Run("Success with filter", func(t *testing.T) {
		expectedResponse := `
							{
							"list": [
								{
									"id": 1,
									"summary": "summary test",
									"performed_at": "2022-03-18T23:11:59Z",
									"created_by": 123
								},
								{
									"id": 2,
									"summary": "summary test",
									"performed_at": "2022-03-18T23:11:59Z",
									"created_by": 123
								}
							]
						}
						`

		useCaseTasks := []tasks.Task{
			{
				ID:          1,
				Summary:     "summary test",
				CreatedBy:   123,
				PerformedAt: time.Date(2022, 3, 18, 23, 11, 59, 0, time.UTC),
			},
			{
				ID:          2,
				Summary:     "summary test",
				CreatedBy:   123,
				PerformedAt: time.Date(2022, 3, 18, 23, 11, 59, 0, time.UTC),
			},
		}

		userCaseMock := mocks.NewUseCase(t)
		expectedUserID := uint(321)
		userCaseMock.On("Search", mock.Anything, tasks.SearchFilters{CreatedBy: &expectedUserID}).
			Return(useCaseTasks, nil)

		w := runCase(userCaseMock, "/tasks/search?created_by=321")

		got := w.Result()
		defer got.Body.Close()
		gotBody, err := ioutil.ReadAll(got.Body)

		assert.Nil(t, err)
		assert.JSONEq(t, expectedResponse, string(gotBody))

		assert.Equal(t, http.StatusOK, got.StatusCode)
		userCaseMock.AssertExpectations(t)
	})

	t.Run("Should return internal error when use case general error", func(t *testing.T) {
		expectedResponse := `{"code":"internal_server_error", "message":"assert.AnError general error for testing"}`

		userCaseMock := mocks.NewUseCase(t)
		userCaseMock.On("Search", mock.Anything, tasks.SearchFilters{}).
			Return([]tasks.Task{}, assert.AnError)

		w := runCase(userCaseMock, "/tasks/search")

		got := w.Result()
		defer got.Body.Close()
		gotBody, err := ioutil.ReadAll(got.Body)

		assert.Nil(t, err)
		assert.JSONEq(t, expectedResponse, string(gotBody))

		assert.Equal(t, http.StatusInternalServerError, got.StatusCode)
		userCaseMock.AssertExpectations(t)
	})

	t.Run("Should return bad request when invalid query param", func(t *testing.T) {
		expectedResponse := `{"code":"bad_request", "message":"strconv.ParseUint: parsing \"invalid\": invalid syntax"}`

		w := runCase(nil, "/tasks/search?created_by=invalid")

		got := w.Result()
		defer got.Body.Close()
		gotBody, err := ioutil.ReadAll(got.Body)

		assert.Nil(t, err)
		assert.JSONEq(t, expectedResponse, string(gotBody))

		assert.Equal(t, http.StatusBadRequest, got.StatusCode)
	})

}

func runCase(useCaseMock searchtasks.UseCase, url string) *httptest.ResponseRecorder {
	h := searchtasks.NewHandler(useCaseMock)
	router := chi.NewRouter()

	searchtasks.RegisterHandler(router, h)

	r := httptest.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	return w
}
