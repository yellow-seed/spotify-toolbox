package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "spotify-toolbox",
		Short: "Spotify Podcast bulk delete CLI tool",
		Long: `A CLI tool for managing Spotify podcasts in bulk.
Currently supports deleting podcasts from your library.`,
		Version: "0.1.0",
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.spotify-toolbox.yaml)")
	rootCmd.PersistentFlags().String("client-id", "", "Spotify Client ID")
	rootCmd.PersistentFlags().String("client-secret", "", "Spotify Client Secret")

	if err := viper.BindPFlag("client_id", rootCmd.PersistentFlags().Lookup("client-id")); err != nil {
		fmt.Fprintf(os.Stderr, "Error binding client-id flag: %v\n", err)
		os.Exit(1)
	}
	if err := viper.BindPFlag("client_secret", rootCmd.PersistentFlags().Lookup("client-secret")); err != nil {
		fmt.Fprintf(os.Stderr, "Error binding client-secret flag: %v\n", err)
		os.Exit(1)
	}
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error finding home directory: %v\n", err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".spotify-toolbox")
	}

	viper.SetEnvPrefix("SPOTIFY")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
