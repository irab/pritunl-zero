package proxy

import (
	"github.com/Sirupsen/logrus"
	"github.com/pritunl/pritunl-zero/node"
	"github.com/pritunl/pritunl-zero/utils"
	"net/http"
	"strings"
)

func WriteError(w http.ResponseWriter, r *http.Request, code int, err error) {
	http.Error(w, utils.GetStatusMessage(code), code)

	logrus.WithFields(logrus.Fields{
		"client": node.Self.GetRemoteAddr(r),
		"error":  err,
	}).Error("proxy: Serve error")
}

func stripCookie(r *http.Request) {
	cookie := r.Header.Get("Cookie")
	start := strings.Index(cookie, "pritunl-zero=")
	if start != -1 {
		str := cookie[start:]
		end := strings.Index(str, ";")
		if end != -1 {
			if len(str) > end+1 && string(str[end+1]) == " " {
				end += 1
			}
			cookie = cookie[:start] + cookie[start+end+1:]
		} else {
			cookie = cookie[:start]
		}
	}

	cookie = strings.TrimSpace(cookie)

	if len(cookie) > 0 {
		r.Header.Set("Cookie", cookie)
	} else {
		r.Header.Del("Cookie")
	}
}
