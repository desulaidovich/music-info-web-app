package render

import (
	"encoding/json"
	"io"
	"net/http"
)

func BindAs[T any](r *http.Request) (*T, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	data := new(T)

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func RenderAs[T any](t *T, statusCode int, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)

	data, err := renderer(t)
	if err != nil {
		return err
	}

	w.Write(data)

	return nil
}

func renderer[T any](t *T) ([]byte, error) {
	data, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	return data, nil
}
