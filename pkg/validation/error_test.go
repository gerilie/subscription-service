package validation_test

import (
	"testing"

	"github.com/go-openapi/testify/v2/require"
	"github.com/yushafro/effective-mobile-tz/pkg/test"
	"github.com/yushafro/effective-mobile-tz/pkg/validation"
)

type errExpected struct {
	message string
	err     error
}

func TestMinError(t *testing.T) {
	t.Parallel()

	tests := []test.Expected[string, errExpected]{
		{
			Name: "valid min",
			Args: "10",
			Expected: errExpected{
				message: validation.MinError("10").Error(),
				err:     validation.ErrMin,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			err := validation.MinError(tt.Args)
			require.Equal(t, tt.Expected.message, err.Error())
			require.ErrorIs(t, err, tt.Expected.err)
		})
	}
}

func TestMaxError(t *testing.T) {
	t.Parallel()

	tests := []test.Expected[string, errExpected]{
		{
			Name: "valid max",
			Args: "10",
			Expected: errExpected{
				message: validation.MaxError("10").Error(),
				err:     validation.ErrMax,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			err := validation.MaxError(tt.Args)
			require.Equal(t, tt.Expected.message, err.Error())
			require.ErrorIs(t, err, tt.Expected.err)
		})
	}
}

func TestGtError(t *testing.T) {
	t.Parallel()

	tests := []test.Expected[string, errExpected]{
		{
			Name: "valid gt",
			Args: "10",
			Expected: errExpected{
				message: validation.GtError("10").Error(),
				err:     validation.ErrGt,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			err := validation.GtError(tt.Args)
			require.Equal(t, tt.Expected.message, err.Error())
			require.ErrorIs(t, err, tt.Expected.err)
		})
	}
}

func TestGteError(t *testing.T) {
	t.Parallel()

	tests := []test.Expected[string, errExpected]{
		{
			Name: "valid gte",
			Args: "10",
			Expected: errExpected{
				message: validation.GteError("10").Error(),
				err:     validation.ErrGte,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			err := validation.GteError(tt.Args)
			require.Equal(t, tt.Expected.message, err.Error())
			require.ErrorIs(t, err, tt.Expected.err)
		})
	}
}

func TestRequiredWithError(t *testing.T) {
	t.Parallel()

	tests := []test.Expected[string, errExpected]{
		{
			Name: "valid required with",
			Args: "confirm_password",
			Expected: errExpected{
				message: validation.RequiredWithError("confirm_password").Error(),
				err:     validation.ErrRequiredWith,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			err := validation.RequiredWithError(tt.Args)
			require.Equal(t, tt.Expected.message, err.Error())
			require.ErrorIs(t, err, tt.Expected.err)
		})
	}
}
