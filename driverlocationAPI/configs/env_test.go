package configs

import "testing"

func TestEnvMongoURI(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		//want is only mongodb because of security
		{"Get ENV file", "mongodb"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EnvMongoURI(); got[:7] != tt.want {
				t.Errorf("EnvMongoURI() = %v, want %v", got, tt.want)
			}
		})
	}
}
