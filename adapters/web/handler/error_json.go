package handler

import "encoding/json"

func jsonError(msg string) []byte {
	error := struct {
		Message string `json:"message"`
	} {
		msg,
	}
	result, err := json.Marshal(error)
	if err != nil {
		return []byte(err.Error())
	}
	return result
}
