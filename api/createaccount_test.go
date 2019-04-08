package api

import (
	"testing"
)

func TestAccountValid(t *testing.T) {
	var accountFormat = map[string]bool{
		"valid":     true,
		" valid":    false,
		"_valid":    false,
		"xxx/valid": false,
		"xxx&valid": false,
		"xxx*alid":  false,
		"xxx alid":  false,
	}

	for k, v := range accountFormat {
		if accountValid(k) != v {
			t.Errorf("`%s` format should %+v", k, v)
		}
	}
}
