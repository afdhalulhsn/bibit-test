package mock

import "github.com/stretchr/testify/mock"

type MovieImdbMockSevice struct {
	mock.Mock
}

func (m *MovieImdbMockSevice) GetListMovieData(in interface{}) (interface{}, error) {
	call := m.Called(in)

	res := call.Get(0)
	if res == nil {
		return nil, call.Error(1)
	}

	return res, nil
}

func (m *MovieImdbMockSevice) GetDetailMovie(in string) (interface{}, error) {
	call := m.Called(in)

	res := call.Get(0)
	if res == nil {
		return nil, call.Error(1)
	}
	return res, nil
}
