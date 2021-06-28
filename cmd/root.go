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

var v string
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
		booty.ReadYamlInput(cmd, args)
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
		if err := setUpLogs(os.Stdout, v); err != nil {
			return err
		}
		return nil
	}

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bootstrap.yaml)")
	rootCmd.PersistentFlags().StringVarP(&v, "verbosity", "v", log.WarnLevel.String(), "Log level (debug, info, warn, error, fatal, panic")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringP("file", "f", "bootstrap.yaml", "Yaml input")
}

//setUpLogs set the log output ans the log level
func setUpLogs(out io.Writer, level string) error {
	lvl, err := log.ParseLevel(level)
	if err != nil {
		return err
	}

	//log := &log.Logger{
	//	Out:   os.Stdout,
	//	Level: lvl,
	//	Formatter: &myFormatter{log.TextFormatter{
	//		FullTimestamp:          true,
	//		TimestampFormat:        "2006-01-02 15:04:05",
	//		ForceColors:            true,
	//		DisableLevelTruncation: true,
	//	},
	//	},
	//}
	//log.Info("Chuy a test")
	//Formatter := new(log.TextFormatter)
	//Formatter.TimestampFormat = "02-01-2006 15:04:05"

	//log.SetFormatter(Formatter)
	//log.SetOutput(out)
	log.SetLevel(lvl)

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
