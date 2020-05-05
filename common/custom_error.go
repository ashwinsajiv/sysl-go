package common

import (
	"context"
	"fmt"
	"strconv"

	"github.com/anz-bank/pkg/log"
)

type CustomError map[string]string

func (e CustomError) Error() string {
	return fmt.Sprintf("%s(%#v)", e["name"], e)
}

func (e CustomError) HTTPError(ctx context.Context) *HTTPError {
	httpStatusString := getOrDefault(e, "http_status", "500")
	httpStatus, err := strconv.Atoi(httpStatusString)
	if err != nil {
		log.Error(ctx, err, fmt.Sprintf("invalid http_status: %s for: %s", httpStatusString, e["name"]))
		httpStatus = 500 //nolint: // TODO: use constant for internal server error
	}
	httpCode := getOrDefault(e, "http_code", "")
	httpMessage := getOrDefault(e, "http_message", "")
	return &HTTPError{httpStatus, httpCode, httpMessage}
}

func getOrDefault(m map[string]string, key string, dflt string) string {
	value, ok := m[key]
	if ok {
		return value
	}
	return dflt
}
