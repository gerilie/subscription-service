package httputil_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-openapi/testify/v2/require"
	"github.com/yushafro/effective-mobile-tz/pkg/httputil"
	"github.com/yushafro/effective-mobile-tz/pkg/test"
)

func Test_DecodeJSON(t *testing.T) {
	t.Parallel()

	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	tests := []test.Expected[test.Request[string], test.Response[User]]{
		{
			Name: "success",
			Args: test.Request[string]{
				Body: `{"name": "John Doe", "age": 30}`,
			},
			Expected: test.Response[User]{
				Code: http.StatusOK,
				Body: User{
					Name: "John Doe",
					Age:  30,
				},
			},
		},
		{
			Name: "empty body",
			Args: test.Request[string]{
				Body: "",
			},
			Expected: test.Response[User]{
				Code: http.StatusBadRequest,
				Body: User{},
			},
			Error: httputil.ErrMismatch,
		},
		{
			Name: "invalid json",
			Args: test.Request[string]{
				Body: "invalid json",
			},
			Expected: test.Response[User]{
				Code: http.StatusBadRequest,
				Body: User{},
			},
			Error: httputil.ErrMismatch,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.Args.Body))

			var body User
			err := httputil.DecodeJSON(context.Background(), w, req, &body)

			require.Equal(t, tt.Expected.Code, w.Code)
			if tt.Error != nil {
				require.ErrorIs(t, err, tt.Error)

				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.Expected.Body.Name, body.Name)
			require.Equal(t, tt.Expected.Body.Age, body.Age)
		})
	}
}

func Test_WriteJSON(t *testing.T) {
	t.Parallel()

	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	tests := []test.Expected[test.Request[User], test.Response[string]]{
		{
			Name: "success",
			Args: test.Request[User]{
				Body: User{
					Name: "John Doe",
					Age:  30,
				},
				Code: http.StatusOK,
			},
			Expected: test.Response[string]{
				Code:        http.StatusOK,
				ContentType: httputil.JSON,
				Body:        `{"name":"John Doe","age":30}`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()

			err := httputil.WriteJSON(w, tt.Args.Code, tt.Args.Body)

			require.NoError(t, err)
			require.Equal(t, tt.Expected.Code, w.Code)
			require.Equal(t, tt.Expected.ContentType, w.Header().Get("Content-Type"))
			require.Equal(t, tt.Expected.Body, w.Body.String())
		})
	}
}
