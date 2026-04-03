package application

import (
	"encoding/json"
	"strings"
)

type ValidationError struct {
	Errs map[string][]string
}

func (ve ValidationError) Error() string {
	msgs := []string{}
	for field, errs := range ve.Errs {
		for _, err := range errs {
			msgs = append(msgs, field+": "+err)
		}
	}
	return strings.Join(msgs, ", ")
}

func (ve ValidationError) ToJSON() []byte {
	b, _ := json.Marshal(ve.Errs)
	return b
}
