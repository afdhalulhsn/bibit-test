package mock

import "github.com/stretchr/testify/mock"

type MovieImdbMockRepo struct {
	mock.Mock
}

func (m *MovieImdbMockRepo) GetListMovieOmdb(i interface{}) (interface{}, error) {
	call := m.Called(i)

	res := call.Get(0)
	if res == nil {
		return nil, call.Error(1)
	}

	return res, nil
}

func (m *MovieImdbMockRepo) GetDeatailMovie(i string) (interface{}, error) {
	call := m.Called(i)

	res := call.Get(0)
	if res == nil {
		return nil, call.Error(1)
	}

	return res, nil
}
