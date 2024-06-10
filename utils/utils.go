package utils

func AddAllowedOrigins() map[string]bool {
	return map[string]bool{
		"http://localhost:3000": true,
		"http://localhost:3001": true,
	}
}
