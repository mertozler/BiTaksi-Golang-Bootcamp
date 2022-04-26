package utils

import "testing"

func TestHaversine(t *testing.T) {
	type args struct {
		latitude1  float64
		longitude1 float64
		latitude2  float64
		longitude2 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Test case", args{40.58889619,
			29.4638355, 40.581087, 29.46146}, 0.89},
		{"Test case", args{40.94289771,
			29.0390297, 41.2144912, 29.09410704}, 30.55},
		{"Test case", args{41.12084618,
			29.07659061, 40.93636002, 29.07228447}, 20.52},
		{"Test case", args{41.07010113,
			28.84269383, 41.08403379, 29.21033874}, 30.86},
		{"Test case", args{41.02230329,
			29.05332029, 40.94898954, 28.86497529}, 17.79},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Haversine(tt.args.latitude1, tt.args.longitude1, tt.args.latitude2, tt.args.longitude2); got != tt.want {
				t.Errorf("Haversine() = %v, want %v", got, tt.want)
			}
		})
	}
}
