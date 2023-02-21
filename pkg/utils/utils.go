package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) error {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, x); err != nil {
		return err
	}

	return nil
}

func HeaderWriter(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "pkglication/json")
}
