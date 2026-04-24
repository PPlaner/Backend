package handler

import "github.com/gin-gonic/gin"

const refreshCookieName = "refresh_token"

func setRefreshCookie(c *gin.Context, token string) {
	c.SetCookie(
		refreshCookieName,
		token,
		60*60*24*7,
		"/",
		"",
		false,
		true,
	)
}

func clearRefreshCookie(c *gin.Context) {
	c.SetCookie(
		refreshCookieName,
		"",
		-1,
		"/",
		"",
		false,
		true,
	)
}
