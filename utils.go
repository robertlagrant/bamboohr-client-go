package bamboohr_client

import (
	"regexp"
	"strings"
)

func parsePayRate(payRate string) (string, string, error) {
	r := regexp.MustCompile(`(?P<precurrency>[\D]*)(?P<salary>\d+(?:[\.,]\d\d)?)(?P<postcurrency>.*)`)

	salary := r.ReplaceAllString(payRate, "${salary}")
	preCurrency := r.ReplaceAllString(payRate, "${precurrency}")
	postCurrency := r.ReplaceAllString(payRate, "${postcurrency}")

	preAndPost := strings.TrimSpace(preCurrency) + strings.TrimSpace(postCurrency)

	return salary, preAndPost, nil
}
