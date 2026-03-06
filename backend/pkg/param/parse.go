package param

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func ParseToDate(param []string, r *http.Request) ([]time.Time, error) {
	var params []time.Time

	url := r.URL.Query()
	for _, p := range param {
		paramParse := url.Get(p)
		if paramParse == "" {
			return nil, fmt.Errorf("empty param %s", p)
		}

		date, err := time.Parse("2006-01-02", paramParse)
		if err != nil {
			return nil, fmt.Errorf("error parse to date : %s", p)
		}

		params = append(params, date)
	}

	return params, nil
}

func ParseToUuid(param []string, r *http.Request) ([]uuid.UUID, error) {
	var params []uuid.UUID

	url := r.URL.Query()
	for _, p := range param {
		paramParse := url.Get(p)
		if paramParse == "" {
			return nil, fmt.Errorf("empty param %s", p)
		}

		uid, err := uuid.Parse(paramParse)
		if err != nil {
			return nil, fmt.Errorf("error parse to uuid : %s", p)
		}

		params = append(params, uid)
	}

	return params, nil
}

func Parse(param []string, r *http.Request) ([]string, error) {
	var params []string

	url := r.URL.Query()
	for _, p := range param {
		paramParse := url.Get(p)
		if paramParse == "" {
			return nil, fmt.Errorf("empty param %s", p)
		}

		params = append(params, p)
	}

	return params, nil
}
