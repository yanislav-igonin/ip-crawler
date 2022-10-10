package http

import (
	"net"
	"net/http"
)

func Request(ip string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://"+ip, nil)
	if err != nil {
		panic(err)
	}
	_, err = client.Do(req)
	if err != nil {
		// Check for timeout
		if err, ok := err.(net.Error); ok && err.Timeout() {
			return "timeout"
		}
		// Something else, need to check manualy later
		return "error"
	}
	return "ok"
}
