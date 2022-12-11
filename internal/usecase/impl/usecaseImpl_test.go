package impl

import (
	"errors"
	mock_repository "fintech/internal/repository/mocks"
	"fintech/pkg/abbreviator"
	customErrors "fintech/pkg/errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUseCaseImpl_GetURLByShort(t *testing.T) {
	type mockBehavior func(repo *mock_repository.MockRepository, short string)
	testTable := []struct {
		name         string
		short        string
		expectedErr  error
		expectedUrl  string
		mockBehavior mockBehavior
	}{
		{
			name:        "ok",
			short:       "123414",
			expectedErr: nil,
			expectedUrl: "https://some_url.com",
			mockBehavior: func(repo *mock_repository.MockRepository, short string) {
				url := "https://some_url.com"
				repo.EXPECT().GetByShort(short).Return(url, nil)
			},
		},
		{
			name:        "not found",
			short:       "123414",
			expectedErr: customErrors.ErrURLNotFound,
			mockBehavior: func(repo *mock_repository.MockRepository, short string) {
				repo.EXPECT().GetByShort(short).Return("", customErrors.ErrURLNotFound)
			},
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			mockRepo := mock_repository.NewMockRepository(c)

			testCase.mockBehavior(mockRepo, testCase.short)
			useCase := UseCaseImpl{repository: mockRepo}
			url, err := useCase.GetURLByShort(testCase.short)
			assert.Equal(t, testCase.expectedErr, err)
			assert.Equal(t, testCase.expectedUrl, url)
		})
	}
}

func TestUseCaseImpl_GenerateShortURL(t *testing.T) {
	unexpectedErr := errors.New("unexpectedError")
	type mockBehavior func(repo *mock_repository.MockRepository, url string)
	testTable := []struct {
		name         string
		url          string
		expectedErr  error
		mockBehavior mockBehavior
	}{
		{
			name:        "ok",
			url:         "https://some_url.com",
			expectedErr: nil,
			mockBehavior: func(repo *mock_repository.MockRepository, url string) {
				repo.EXPECT().GetByURL(url).Return("", customErrors.ErrURLNotFound)
				repo.EXPECT().Insert(url, gomock.Any()).Return(nil)
			},
		},
		{
			name:        "already exist",
			url:         "https://some_url.com",
			expectedErr: customErrors.ErrAlreadyAbbreviated,
			mockBehavior: func(repo *mock_repository.MockRepository, url string) {
				repo.EXPECT().GetByURL(url).Return("", nil)
			},
		},
		{
			name:        "unexpected getErr",
			url:         "https://some_url.com",
			expectedErr: unexpectedErr,
			mockBehavior: func(repo *mock_repository.MockRepository, url string) {
				repo.EXPECT().GetByURL(url).Return("", unexpectedErr)
			},
		},
		{
			name:        "cannot create",
			url:         "https://some_url.com",
			expectedErr: unexpectedErr,
			mockBehavior: func(repo *mock_repository.MockRepository, url string) {
				repo.EXPECT().GetByURL(url).Return("", customErrors.ErrURLNotFound)
				repo.EXPECT().Insert(url, gomock.Any()).Return(unexpectedErr)
			},
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			c := gomock.NewController(t)
			defer c.Finish()

			mockRepo := mock_repository.NewMockRepository(c)

			testCase.mockBehavior(mockRepo, testCase.url)
			useCase := UseCaseImpl{repository: mockRepo, generator: abbreviator.NewAbbreviateGenerator()}
			_, err := useCase.GenerateShortURL(testCase.url)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}
