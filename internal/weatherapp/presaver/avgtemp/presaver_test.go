package avgtemp

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"main.go/internal/weatherapp"
)

func TestPreSaver_Save(t *testing.T) {
	type args struct {
		cityResult weatherapp.CityResult
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Positive avg temp",
			args: args{
				cityResult: weatherapp.CityResult{
					Name:        "test_city",
					TempAverage: 5,
				},
			},
		},
		{
			name: "Negative avg temp",
			args: args{
				cityResult: weatherapp.CityResult{
					Name:        "test_city",
					TempAverage: -5,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			avgTempPreSaver := &PreSaver{}
			avgTempPreSaver.Save(tt.args.cityResult)
			assert.InEpsilon(t, tt.args.cityResult.TempAverage, *avgTempPreSaver.highestAvgTemp, 0.001)
		})
	}
}
