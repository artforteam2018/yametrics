package metrics

import (
	"errors"
	"strings"
)

func ParseMetricNameValue(path string) (string, string, error) {
	splittedValues := strings.Split(path, "/")
	if len(splittedValues) != 5 {
		return "", "", errors.New("not found")
	}

	metricName, metricValue := splittedValues[3], splittedValues[4]

	return metricName, metricValue, nil
}
