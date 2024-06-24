package v1

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kostylevdev/todo-rest-api/internal/domain"
	"github.com/kostylevdev/todo-rest-api/internal/service"
	mock_service "github.com/kostylevdev/todo-rest-api/internal/service/mocks"
	"github.com/stretchr/testify/assert"

	"go.uber.org/mock/gomock"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(r *mock_service.MockAutorization, user domain.User)

	tests := []struct {
		name      string
		inputBody string
		inputUser domain.User
		mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"test","username":"test","password":"test"}`,
			inputUser: domain.User{
				Name:     "test",
				Username: "test",
				Password: "test",
			},
			mockBehavior: func(r *mock_service.MockAutorization, user domain.User) {
				r.EXPECT().SignUp(user).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "invalid input body",
			inputBody:            `{name":"test1","username":"test1"}`,
			inputUser:            domain.User{},
			mockBehavior:         func(r *mock_service.MockAutorization, user domain.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "service error",
			inputBody: `{"name":"test2","username":"test2","password":"test2"}`,
			inputUser: domain.User{
				Name:     "test2",
				Username: "test2",
				Password: "test2",
			},
			mockBehavior: func(r *mock_service.MockAutorization, user domain.User) {
				r.EXPECT().SignUp(user).Return(0, errors.New("service error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"service error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo_auth := mock_service.NewMockAutorization(c)
			tt.mockBehavior(repo_auth, tt.inputUser)
			services := &service.Service{Autorization: repo_auth}

			handler := NewHandler(services)
			r := gin.New()
			r.POST("/auth/sign-up", handler.SignUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/auth/sign-up", bytes.NewBufferString(tt.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, tt.expectedStatusCode)
			assert.Equal(t, w.Body.String(), tt.expectedResponseBody)
		})
	}
}
