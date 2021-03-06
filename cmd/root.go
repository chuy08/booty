package cmd

import (
	"booty/pkg/booty"
	"io"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Verbose string
var cfgFile string

type myFormatter struct {
	log.TextFormatter
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "booty",
	Short: "Bootstrap",
	Long:  `Bootstrap Docker configuration files from yaml templates.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		var p = booty.Things{
			Args:      args,
			Cmd:       cmd,
			Verbosity: Verbose,
		}

		booty.ReadYamlInput(p)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if err := setUpLogs(os.Stdout, Verbose); err != nil {
			return err
		}
		return nil
	}

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bootstrap.yaml)")
	rootCmd.PersistentFlags().StringVarP(&Verbose, "verbosity", "v", log.InfoLevel.String(), "Log level (debug, info, warn, error, fatal, panic")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringP("file", "f", "bootstrap.yaml", "Yaml input")
}

//setUpLogs set the log output ans the log level
func setUpLogs(out io.Writer, level string) error {
	log.Debug("Inside setUpLogs")
	lvl, err := log.ParseLevel(level)
	if err != nil {
		return err
	}

	log.SetLevel(lvl)

	log.SetFormatter(&log.TextFormatter{
		ForceColors:     false,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return nil
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Error(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".booty" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".booty")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Info("Using config file:", viper.ConfigFileUsed())
	}
}
