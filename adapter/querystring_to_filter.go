package adapter

import "net/url"

func QueryStringToMap(query url.Values) map[string][]string {
	return map[string][]string(query)
}
