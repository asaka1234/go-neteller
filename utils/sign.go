package utils

import (
	"encoding/base64"
)

// http://paysafegroup.github.io/neteller_rest_api_v1/#/introduction/technical-introduction/oauth-authentication
// Sample Request #1 (using HTTP Basic Authorization)
// https://developer.paysafe.com/en/neteller-api-1/
// base64-encode(string username:password)
func Sign(clientId, clientSecret string) string {
	authStr := base64.StdEncoding.EncodeToString([]byte(clientId + ":" + clientSecret))
	//TODO 要确认下是否需要escape
	//authStr = url.QueryEscape(authStr)
	return authStr
}
