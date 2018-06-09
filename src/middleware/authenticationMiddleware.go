package middleware

type authenticationMiddleware struct {
	tokenUsers map[string]string
}

