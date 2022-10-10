package http

import (
	"net"
	"net/http"
)

func Request(ip string) (string, int) {
	res, err := http.Get("http://" + ip)
	if err != nil {
		// Check for timeout
		if err, ok := err.(net.Error); ok && err.Timeout() {
			return "timeout", 0
		}
		// Something else, need to check manualy later
		return "error", 0
	}
	return "ok", res.StatusCode
}
