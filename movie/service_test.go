package movie

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)


type MockStore struct{
	mock.Mock
}

func (ms *MockStore) GetMovieByID(id int)(Movie, error){
	args := ms.Called(id);

	return args.Get(0).(Movie), args.Error(1);
}

func TestService_GetMovieInfo_Success(t *testing.T){
	mockStore := new(MockStore);
	service := NewService(mockStore);

	expectedMovie := Movie{ID: 1, Title: "Inception", Year: 2010}

	mockStore.On(`GetMovieByID`, 1).Return(expectedMovie, nil);

	info, err := service.GetMovieInfo(1);

	require.NoError(t, err);
	require.Equal(t, `Inception (2010)`, info);

	mockStore.AssertExpectations(t);
}


func TestService_GetMovieInfo_Error(t *testing.T){

	ms := new(MockStore);
	service := NewService(ms);

	ms.On(`GetMovieByID`, 2).Return(Movie{}, errors.New(`mockError`));

	info, err := service.GetMovieInfo(2);

	require.Empty(t, info);
	require.Error(t, err);
	require.EqualError(t, err, "failed to get movie: mockError");

	ms.AssertExpectations(t);
}


func TestService_GetMovieInfo_Classic(t *testing.T){

	ms := new(MockStore);
	service := NewService(ms);


	ms.On(`GetMovieByID`, mock.Anything).Return(Movie{}, nil);

	info, err := service.GetMovieInfo(1970);

	require.NoError(t, err);
	require.Equal(t, `Classic`, info);

	ms.AssertExpectations(t);
}