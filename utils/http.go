package utils

import (
	"fmt"
	"net/http"
)

func CrashIfErrorStatus(res *http.Response) {
	if res.StatusCode >= 400 {
		CrashWithError(
			fmt.Sprintf("status code error: %d %s", res.StatusCode, res.Status),
		)
	}
}
