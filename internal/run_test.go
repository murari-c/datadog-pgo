package internal

import (
	"log/slog"
	"testing"
	"time"
)

func TestExtractMergedProfile(t *testing.T) {
	type args struct {
		from                 time.Duration
		timeout              time.Duration
		profiles             int
		fail                 bool
		shouldUsePGOEndpoint bool
		log                  *slog.Logger
		inputQueries         []string
		envkeys              map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    *MergedProfile
		wantErr bool
	}{
		{
			name: "invalid extraction - missing tokens",
			args: args{
				from:                 24 * time.Hour,
				timeout:              10 * time.Second,
				profiles:             5,
				fail:                 false,
				shouldUsePGOEndpoint: true,
				log:                  slog.Default(),
				inputQueries:         []string{"query1", "query2"},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid extraction - invalid tokens",
			args: args{
				from:                 24 * time.Hour,
				timeout:              10 * time.Second,
				profiles:             5,
				fail:                 false,
				shouldUsePGOEndpoint: true,
				log:                  slog.Default(),
				inputQueries:         []string{"query1", "query2"},
				envkeys:              map[string]string{"DD_API_KEY": "test_api_key", "DD_APP_KEY": "test_app_key"},
			},
			wantErr: true,
		},
		{
			name: "valid extraction",
			args: args{
				from:                 24 * time.Hour,
				timeout:              10 * time.Second,
				profiles:             5,
				fail:                 false,
				shouldUsePGOEndpoint: true,
				log:                  slog.Default(),
				inputQueries:         []string{"service:prod-user-service-v2 env:production"},
				envkeys:              map[string]string{"DD_API_KEY": "<valid>", "DD_APP_KEY": "<valid>"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.args.envkeys {
				t.Setenv(k, v)
			}

			_, err := ExtractMergedProfile(tt.args.from, tt.args.timeout, tt.args.profiles, tt.args.fail, tt.args.shouldUsePGOEndpoint, tt.args.log, tt.args.inputQueries...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractMergedProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
