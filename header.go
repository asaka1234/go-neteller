package go_neteller

import "fmt"

func getHeaders() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}

// notice:  input params is fixed
func getAuthHeaders(auth string) map[string]string {
	return map[string]string{
		"Content-Type":  "application/json",
		"charset":       "utf-8",
		"Authorization": fmt.Sprintf("Basic %s", auth),
		"Simulator":     "EXTERNAL", //TODO 这个做什么的?
	}
}
