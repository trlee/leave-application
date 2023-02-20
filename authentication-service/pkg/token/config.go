package token

import "time"

var (
	// Cookie parameters
	CookieName  = "authToken"
	ExpiresDate = time.Now().Add(time.Minute * 60) // 60 min
	MaxAge      = 60 * 60                          // 60 min

	// Token parameter
	Duration = time.Minute * 60                   // duration of paseto token: 60 min
	IV       = "cftheB6xhdMeRgQ7Hrafi9uUWW9R5ExX" // initialization vector
)
