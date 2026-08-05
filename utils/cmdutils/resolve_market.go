package cmdutils

import (
	"fmt"
	"regexp"

	"github.com/jeandeaual/go-locale"
)

func ValidateMarket(market string) error {
	match, _ := regexp.MatchString("^[a-z]{2}-[A-Z]{2}$", market)
	if !match {
		return fmt.Errorf("invalid market format: %s", market)
	}

	return nil
}

func ResolveMarket(market string) (string, error) {
	if market == "" {
		systemLocale, err := locale.GetLocale()
		if err != nil {
			return "", fmt.Errorf("failed to get system locale: %w", err)
		}
		err = ValidateMarket(systemLocale)
		if err != nil {
			return "", fmt.Errorf("invalid system locale: %w", err)
		}

		return systemLocale, nil
	}

	err := ValidateMarket(market)
	if err != nil {
		return "", fmt.Errorf("invalid market: %w", err)
	}

	return market, nil
}
