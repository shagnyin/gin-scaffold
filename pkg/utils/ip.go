package utils

import (
	"net/http"
)

// GetRealIP .
// proxy_set_header X-Real-IP $remote_addr:$server_port;
func GetRealIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-IP")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarder-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}
