package gocommerce

import (
	"net"
	"strconv"
)

func parseHostPort(hp string) (string, int, error) {
	host, port, err := net.SplitHostPort(hp)
	if err != nil {
		return "", 0, err
	}
	p, err := strconv.Atoi(port)
	if err != nil {
		return "", 0, err
	}
	return host, p, nil
}
