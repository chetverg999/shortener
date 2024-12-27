package handlers

import (
	"fmt"
	"net/url"
)

func parseURL(userURL []byte) error {

	parsedURL, err := url.ParseRequestURI(string(userURL))
	if err != nil {
		fmt.Println("Ошибка при парсинге:", err)
		return err
	}

	// Дополнительная проверка схемы
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		err = fmt.Errorf("Эта схема не поддерживается: %s", parsedURL.Scheme)
		fmt.Println("Ошибка схемы:", err)
		return err
	}

	fmt.Println("URL валиден по схеме:", parsedURL.Scheme)
	return nil
}
