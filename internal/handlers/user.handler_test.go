package handlers

import (
	"coffee-shop/config"
	"coffee-shop/internal/repository"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repoUserMock = repository.RepoMock{}
var reqBody = `{
	"user_id": "123",
	"username": "testing",
	"password": "abcd1234",
	"role": "user",
	"file_image": "file.jpg"
}`

func TestGetById(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("getById", func(t *testing.T) {
		r := gin.Default()
		w := httptest.NewRecorder()

		handler := NewUser(&repoUserMock)
		exptedResult := &config.Result{Data: map[string]interface{}{
			"id":          "fff7418c-e613-44b4-85e0-dbe05e236c5e",
			"displayname": "Magnus",
			"phone":       "081034941",
			"email":       "magnusgmailcom",
			"create_at":   "2023-09-01T20:07:30.662605Z",
		}}
		repoUserMock.On("GetUserById", mock.Anything).Return(exptedResult, nil)

		r.GET("/get/:id", handler.GetById)
		req := httptest.NewRequest("GET", "/get/fff7418c-e613-44b4-85e0-dbe05e236c5e", strings.NewReader("{}"))
		req.Header.Set("Content-type", "application/json")
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		exptedResponse := `{
		"status": "OK",
		"data": {
			"id": "fff7418c-e613-44b4-85e0-dbe05e236c5e",
			"displayname": "Magnus",
			"phone": "081034941",
			"email": "magnusgmailcom",
			"create_at": "2023-09-01T20:07:30.662605Z"
		}
	}`
		assert.JSONEq(t, exptedResponse, w.Body.String())
	})

	t.Run("user not found", func(t *testing.T) {
		r := gin.Default()
		w := httptest.NewRecorder()

		handler := NewUser(&repoUserMock)
		exptedError := errors.New("user not found")
		exptedResult := &config.Result{Message: "user not found"}

		repoUserMock.On("GetUserById", mock.Anything).Return(exptedResult, exptedError)

		r.GET("/get/:id", handler.GetById)
		req := httptest.NewRequest("GET", "/get/2768", strings.NewReader("{}"))
		req.Header.Set("Content-type", "application/json")
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"description": "user not found", "status": "Bad Request"}`, w.Body.String())
	})
}
func TestGetUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	handler := NewUser(&repoUserMock)
	exptedResult := &config.Result{Data: map[string]interface{}{
		"id":          "fff7418c-e613-44b4-85e0-dbe05e236c5e",
		"displayname": "Magnus",
		"phone":       "081034941",
		"email":       "magnusgmailcom",
		"create_at":   "2023-09-01T20:07:30.662605Z"}}
	repoUserMock.On("ReadUser", mock.Anything).Return(exptedResult, nil)

	r.GET("/get", handler.GetUser)
	req := httptest.NewRequest("GET", "/get", strings.NewReader("{}"))
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	exptedResponse := `{
		"status": "OK",
		"data": {
			"id": "fff7418c-e613-44b4-85e0-dbe05e236c5e",
			"displayname": "Magnus",
			"phone": "081034941",
			"email": "magnusgmailcom",
			"create_at": "2023-09-01T20:07:30.662605Z"
		}
	}`
	assert.JSONEq(t, exptedResponse, w.Body.String())
}
func TestPostUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	w := httptest.NewRecorder()

	handler := NewUser(&repoUserMock)
	exptedResult := &config.Result{Message: "1 data user created"}
	repoUserMock.On("CreateUser", mock.Anything).Return(exptedResult, nil)

	r.POST("/create", handler.PostUser)
	req := httptest.NewRequest("POST", "/create", strings.NewReader(reqBody))
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"description": "1 data user created", "status": "OK"}`, w.Body.String())
}
func TestPatchUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	r := gin.Default()

	handler := NewUser(&repoUserMock)
	expectedResult := &config.Result{Message: "1 data user updated"}
	repoUserMock.On("UpdateUser", mock.Anything).Return(expectedResult, nil)

	r.Use(func(c *gin.Context) {
		result := "image.jpg"
		c.Set("image", &result)
	})
	r.PUT("/update/:id", handler.PatchUser)
	req := httptest.NewRequest("PUT", "/update/123", strings.NewReader(reqBody))
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"description": "1 data user updated", "status": "OK"}`, w.Body.String())
}

func TestDeleteUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	w := httptest.NewRecorder()

	handler := NewUser(&repoUserMock)
	exptedResult := &config.Result{Message: "1 data user deleted"}
	repoUserMock.On("CreateUser", mock.Anything).Return(exptedResult, nil)

	r.DELETE("/delete/:id", handler.PostUser)
	req := httptest.NewRequest("DELETE", "/delete/123", strings.NewReader("{}"))
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"description": "1 data user deleted", "status": "OK"}`, w.Body.String())
}
