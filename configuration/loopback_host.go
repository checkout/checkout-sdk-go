package configuration

import (
	"net"
	"strings"
)

// isLoopbackHost reports whether host is localhost (ASCII, any case) or a loopback IP address
// (127.0.0.0/8, ::1, or 127.0.0.0/8 in IPv4-mapped IPv6 form). It only parses host and opens no
// connection. The length check keeps the
// comparison ASCII: strings.EqualFold alone also matches Unicode case folds such as U+017F for "s".
func isLoopbackHost(host string) bool {
	if len(host) == len("localhost") && strings.EqualFold(host, "localhost") {
		return true
	}
	ip := net.ParseIP(host)
	return ip != nil && ip.IsLoopback()
}
