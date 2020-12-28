// Package middleware handles pre/post business logic.
// This package is tight coupling to gin framework, ignore
// me and create your own if not using gin.
package middleware

import "github.com/gin-gonic/gin"

// WithCacheHeaderControl sets Cache Control header to manage cache
// duration on client side
func WithCacheHeaderControl(durationSec string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "max-age="+durationSec)
		c.Next()
	}
}

// WithSecurityHeaderControl sets security header value
// for more details about the security related issues: https://infosec.mozilla.org/guidelines/web_security
func WithSecurityHeaderControl() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Strict-Transport-Security", "max-age=63072000; includeSubDomains;")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("Referrer-Policy", "no-referrer-when-downgrade")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Next()
	}
}
