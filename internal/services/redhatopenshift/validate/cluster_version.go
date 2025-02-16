package validate

import (
	"fmt"
	"regexp"
)

func ClusterVersion(i interface{}, k string) ([]string, []error) {
	v, ok := i.(string)
	if !ok {
		return nil, []error{fmt.Errorf("expected type of %q to be string", k)}
	}

	var errors []error
	if matched := regexp.MustCompile(`^\d+\.\d+\.\d+$`).Match([]byte(v)); !matched {
		errors = append(errors, fmt.Errorf("%q should be in the format `X.Y.Z` (e.g. `4.13.23`)", k))
	}

	return nil, errors
}
