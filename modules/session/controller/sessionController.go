package controller

import (
	"time"
	"github.com/labstack/echo"
	"github.com/nu7hatch/gouuid"
	"net/http"
	"physiobot/modules/session/rest"
)

func CreateSession(c echo.Context) error {
	session := new(rest.NewSession)

	identification, err := uuid.NewV4()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	session.SessionID = identification.String()
	session.IPAddress = GetIP(c.Request())
	session.Date = time.Now()

	return c.JSON(http.StatusOK, session)
}

// GetIP gets a requests IP address by reading off the forwarded-for
// header (for proxies) and falls back to use the remote address.
func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}
