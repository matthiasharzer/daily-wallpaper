package cmdutils

import (
	"fmt"
	"regexp"

	"github.com/jeandeaual/go-locale"
	"golang.org/x/text/language"
)

func bcpToMarketTag(bcp47Tag string) (string, error) {
	tag, err := language.Parse(bcp47Tag)
	if err != nil {
		return "", err
	}
	base, _ := tag.Base()
	region, _ := tag.Region()

	return fmt.Sprintf("%s-%s", base.String(), region.String()), nil
}

func ValidateMarket(market string) error {
	match, _ := regexp.MatchString("^[a-zA-Z]{2}-[a-zA-Z]{2}$", market)
	if !match {
		return fmt.Errorf("invalid market format: %s", market)
	}

	return nil
}

func ResolveMarket(market string) (string, error) {
	if market == "" {
		systemLocaleBcpTag, err := locale.GetLocale()
		if err != nil {
			return "", fmt.Errorf("failed to get system locale: %w", err)
		}

		return bcpToMarketTag(systemLocaleBcpTag)
	}

	err := ValidateMarket(market)
	if err != nil {
		return "", fmt.Errorf("invalid market provided: %w", err)
	}

	return market, nil
}
