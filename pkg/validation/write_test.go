package validation_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-openapi/testify/v2/require"
	"github.com/go-playground/validator/v10"
	"github.com/yushafro/effective-mobile-tz/pkg/deferfunc"
	"github.com/yushafro/effective-mobile-tz/pkg/logger"
	"github.com/yushafro/effective-mobile-tz/pkg/test"
	"github.com/yushafro/effective-mobile-tz/pkg/validation"
	"go.uber.org/mock/gomock"
)

func TestWriteErrors(t *testing.T) {
	t.Parallel()

	type args struct {
		Name  string `validate:"required"`
		Email string `validate:"required,email"`
		Price int    `validate:"required,gt=0"`
	}

	tests := []test.Expected[args, validation.Resp]{
		{
			Name: "invalid email",
			Args: args{
				Name:  "John Doe",
				Email: "invalid-email",
				Price: 100,
			},
			Expected: validation.Resp{
				"Email": validation.ErrEmail.Error(),
			},
		},
		{
			Name: "invalid price",
			Args: args{
				Name:  "John Doe",
				Email: "jUH6t@example.com",
				Price: -1,
			},
			Expected: validation.Resp{
				"Price": validation.GtError("0").Error(),
			},
		},
		{
			Name: "invalid name, email and price",
			Args: args{
				Name:  "",
				Email: "invalid-email",
				Price: -1,
			},
			Expected: validation.Resp{
				"Name":  validation.ErrRequired.Error(),
				"Email": validation.ErrEmail.Error(),
				"Price": validation.GtError("0").Error(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			mock := logger.NewMockLogger(ctrl)
			ctx := logger.WithLogger(context.Background(), mock)

			validate := validator.New()
			err := validate.Struct(tt.Args)
			require.Error(t, err)

			var ve validator.ValidationErrors
			if errors.As(err, &ve) {
				w := httptest.NewRecorder()

				validation.WriteErrors(ctx, w, ve)

				resp := w.Result()
				deferfunc.Close(ctx, resp.Body.Close, "error closing response body")

				var respBody validation.Resp
				err := json.NewDecoder(resp.Body).Decode(&respBody)
				require.NoError(t, err)

				require.Equal(t, http.StatusBadRequest, resp.StatusCode)
				require.Equal(t, tt.Expected, respBody)
			}
		})
	}
}
