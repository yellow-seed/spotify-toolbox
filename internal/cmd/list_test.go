package cmd

import (
	"testing"
)

func TestListCommand(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "default options",
			args:    []string{"list"},
			wantErr: false,
		},
		{
			name:    "with format flag",
			args:    []string{"list", "--format", "json"},
			wantErr: false,
		},
		{
			name:    "with limit flag",
			args:    []string{"list", "--limit", "100"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootCmd.SetArgs(tt.args)
			err := rootCmd.Execute()
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
