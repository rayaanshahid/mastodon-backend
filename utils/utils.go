package utils

func AddAllowedOrigins(url string) map[string]bool {
	return map[string]bool{
		url: true,
	}
}
