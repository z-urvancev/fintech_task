package handler

import (
	"bytes"
	mock_usecase "fintech/internal/usecase/mocks"
	"fintech/pkg/errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_GetURLByShort(t *testing.T) {
	type mockBehavior func(r *mock_usecase.MockUseCase, short string)

	testTable := []struct {
		name                 string
		inputShort           string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:       "ok",
			inputShort: "124553221",
			mockBehavior: func(r *mock_usecase.MockUseCase, short string) {
				expectedURL := "https://some_url"
				r.EXPECT().GetURLByShort(short).Return(expectedURL, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "{\"url\":\"https://some_url\"}",
		},
		{
			name:       "not found",
			inputShort: "124553221",
			mockBehavior: func(r *mock_usecase.MockUseCase, short string) {
				r.EXPECT().GetURLByShort(short).Return("", errors.ErrURLNotFound)
			},
			expectedStatusCode:   404,
			expectedResponseBody: "{\"error\":\"URL не найден\"}",
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			mockUseCase := mock_usecase.NewMockUseCase(c)
			testCase.mockBehavior(mockUseCase, testCase.inputShort)

			handler := Handler{
				useCase: mockUseCase,
			}

			r := gin.New()
			r.GET("/:short", handler.GetURLByShort, errors.Middleware())

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", strings.Join([]string{"/", testCase.inputShort}, ""),
				bytes.NewBufferString(""))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}

type Input struct {
	URL string `json:"url"`
}

func TestHandler_GenerateShortURL(t *testing.T) {
	type mockBehavior func(r *mock_usecase.MockUseCase, url string)

	testTable := []struct {
		name                 string
		input                Input
		inputBody            string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:  "ok",
			input: Input{"https://some_url"},
			mockBehavior: func(r *mock_usecase.MockUseCase, url string) {
				expectedShort := "121412411"
				r.EXPECT().GenerateShortURL(url).Return(expectedShort, nil)
			},
			inputBody: `{
    			"url": "https://some_url"
}`,
			expectedStatusCode:   200,
			expectedResponseBody: "{\"short\":\"121412411\"}",
		},
		{
			name:                 "bad body",
			input:                Input{"https://some_url"},
			mockBehavior:         func(r *mock_usecase.MockUseCase, url string) {},
			inputBody:            `sfafasdfasf`,
			expectedStatusCode:   400,
			expectedResponseBody: "",
		},
		{
			name:  "already exist",
			input: Input{"https://some_url"},
			inputBody: `{
    			"url": "https://some_url"
}`,
			mockBehavior: func(r *mock_usecase.MockUseCase, url string) {
				r.EXPECT().GenerateShortURL(url).Return("", errors.ErrAlreadyAbbreviated)
			},
			expectedStatusCode:   400,
			expectedResponseBody: "{\"error\":\"Эта ссылка уже сокращена\"}",
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			mockUseCase := mock_usecase.NewMockUseCase(c)
			testCase.mockBehavior(mockUseCase, testCase.input.URL)

			handler := Handler{
				useCase: mockUseCase,
			}

			r := gin.New()
			r.POST("/", handler.GenerateShortURL, errors.Middleware())

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/",
				bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
