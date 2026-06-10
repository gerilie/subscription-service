package httputil

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/yushafro/effective-mobile-tz/pkg/deferfunc"
)

// ErrMismatch indicates that the request body contains a JSON type mismatch
// or malformed JSON that cannot be properly decoded.
var (
	ErrMismatch = errors.New("invalid request body: type mismatch or malformed JSON")

	ErrReadBody   = errors.New("read request body")
	ErrDecodeBody = errors.New("decode request body")

	ErrEncodeBody = errors.New("encode response body")
	ErrWriteBody  = errors.New("write response body")
)

// DecodeJSON reads and decodes JSON request body into dst.
// It writes appropriate HTTP errors for malformed input.
func DecodeJSON[T any](ctx context.Context, w http.ResponseWriter, r *http.Request, dst *T) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrReadBody, err)
	}
	defer deferfunc.Close(ctx, r.Body.Close, "close request body")

	if err := json.Unmarshal(body, dst); err != nil {
		var unmarshalErr *json.UnmarshalTypeError
		var syntaxErr *json.SyntaxError
		if errors.As(err, &unmarshalErr) || errors.As(err, &syntaxErr) {
			http.Error(
				w,
				ErrMismatch.Error(),
				http.StatusBadRequest,
			)

			return fmt.Errorf("%w: %w", ErrMismatch, err)
		}

		http.Error(w, ErrInternalServer.Error(), http.StatusInternalServerError)

		return fmt.Errorf("%w: %w", ErrDecodeBody, err)
	}

	return nil
}

// WriteJSON encodes v to JSON and writes it to the response with given status code.
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		http.Error(w, ErrInternalServer.Error(), http.StatusInternalServerError)

		return fmt.Errorf("%w: %w", ErrEncodeBody, err)
	}

	w.Header().Set(ContentType, JSON)
	w.WriteHeader(status)

	_, err = w.Write(data)
	if err != nil {
		http.Error(w, ErrInternalServer.Error(), http.StatusInternalServerError)

		return fmt.Errorf("%w: %w", ErrWriteBody, err)
	}

	return nil
}
