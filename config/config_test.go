package config_test

import (
	"app/config"
	"testing"
)

func TestNewConfigFrom(t *testing.T) {
	tests := []struct {
		name    string
		want    *config.Config
		wantErr bool
	}{
		{
			name: "Config",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := config.NewConfigFrom()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConfigFrom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
