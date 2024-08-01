package httpgrpc

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
)

type Address struct {
	IsSsl bool   `json:"ssl"`
	Host  string `json:"host"`
	Port  uint16 `json:"port"`
	addr  net.Addr
}

func (a *Address) Request(method string, path string, parameters map[string]string, body io.Reader) (*http.Request, error) {
	var urlstr string
	if a.IsSsl {
		urlstr = "https://"
	} else {
		urlstr = "http://"
	}
	var paramstr string
	if parameters == nil {
		paramstr = ""
	} else if 0 < len(parameters) {
		x := make([]string, len(parameters))
		i := 0
		for k, v := range parameters {
			x[i] = fmt.Sprintf("%s=%s", url.QueryEscape(k), url.QueryEscape(v))
			i++
		}
		paramstr = "?" + strings.Join(x, "&")
	} else {
		paramstr = ""
	}
	return http.NewRequest(method, fmt.Sprintf("%s%s:%d%s%s", urlstr, a.Host, a.Port, path, paramstr), body)
}
