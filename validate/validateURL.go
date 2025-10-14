package validate

import (
	"fmt"
	"net/url"
)

// Проверка корректности URL
func ValidateURL(s string) error {
	u, err := url.Parse(s)
	if err != nil {
		return err
	}

	if u.Scheme == "" || u.Host == "" {
		return fmt.Errorf("URL должен содержать схему и хост (пример: https://example.com)")
	}

	return nil
}
