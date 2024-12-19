package handlers

import (
	"fmt"
	"net/url"
)

func parseURL(userURL []byte) error {

	parsedURL, err := url.ParseRequestURI(string(userURL))
	if err != nil {
		fmt.Println("Parse Error:", err)
		return err
	}

	// Дополнительная проверка схемы
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		err = fmt.Errorf("unsupported URL scheme: %s", parsedURL.Scheme)
		fmt.Println("Scheme Error:", err)
		return err
	}

	fmt.Println("Parsed URL is valid with scheme:", parsedURL.Scheme)
	return nil
}
