package cmd

import (
	"testing"
)

func TestDeleteCommand(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "with dry-run flag",
			args:    []string{"delete", "--dry-run", "--all"},
			wantErr: false,
		},
		{
			name:    "with filter flag",
			args:    []string{"delete", "--filter", "test", "--dry-run"},
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
