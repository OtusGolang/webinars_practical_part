package mockgen

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/OtusGolang/webinars_practical_part/21-codegen/mockgen/mocks"
)

func TestGetPage(t *testing.T) {
	t.Run("test err", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		g := mocks.NewMockGetter(ctrl)
		g.EXPECT().Get("test url 1").Return(nil, errors.New("400"))

		resp, err := GetPage(g, "test url 1")
		require.NotNil(t, err)
		require.Nil(t, resp)
	})

	t.Run("positive test", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		g := mocks.NewMockGetter(ctrl)
		g.EXPECT().Get("test url 2").Return(&http.Response{
			Body: ioutil.NopCloser(bytes.NewBuffer([]byte("some data"))),
		}, nil)

		resp, err := GetPage(g, "test url 2")
		require.Nil(t, err)
		require.Equal(t, []byte("some data"), resp)
	})
}
