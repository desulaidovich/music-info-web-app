package v1

import (
	"net/http"
)

type (
	Signature interface {
		Name() string
		Method() string
		Pattern() string
		Handler() http.HandlerFunc
	}

	SignatureList []Signature
)
