package metrics

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMetricNameValue(t *testing.T) {
	type want struct {
		metricname  string
		metricvalue string
		err         error
	}

	tests := []struct {
		name string
		path string
		want want
	}{
		{
			name: "correctValue",
			path: "/update/counter/testCounter/100",
			want: want{
				metricname:  "testCounter",
				metricvalue: "100",
				err:         nil,
			},
		}, {
			name: "correctValue",
			path: "/update/counter/testCounter/",
			want: want{
				metricname:  "testCounter",
				metricvalue: "",
				err:         nil,
			},
		}, {
			name: "correctValue",
			path: "/update/counter/",
			want: want{
				metricname:  "",
				metricvalue: "",
				err:         errors.New("not found"),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			metricname, metricvalue, err := ParseMetricNameValue(test.path)
			println(metricname, metricvalue, err, test.path)
			assert.Equal(t, test.want.metricname, metricname)
			assert.Equal(t, test.want.metricvalue, metricvalue)
			assert.Equal(t, test.want.err, err)
		})
	}
}
