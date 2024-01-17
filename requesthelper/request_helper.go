package requesthelper

import (
	"errors"
	"io"
	"net/http"
)

func MakeRequestGet(url string) ([]byte, error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, errors.New("error obtaining locations from api")
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		return nil, errors.New("response failed ")
	}

	if err != nil {
		return nil, errors.New("error reading locations from api")
	}

	return body, nil
}
