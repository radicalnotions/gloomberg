package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	apiKeyEtherscan, apiKeyOpensea, cfgFile string
	endpoints, ownWallets                   []string
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "gloomberg",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

//nolint:gochecknoinits
func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// viper.Set("show.all", true)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gloomberg.yaml)")

	// logging
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Show more output")
	_ = viper.BindPFlag("log.verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Show debug output")
	_ = viper.BindPFlag("log.debug", rootCmd.PersistentFlags().Lookup("debug"))

	// rpc node
	rootCmd.PersistentFlags().StringSliceVarP(&endpoints, "endpoints", "e", []string{}, "RPC endpoints")
	_ = viper.BindPFlag("endpoints", rootCmd.Flags().Lookup("endpoints"))

	// wallets
	rootCmd.PersistentFlags().StringSliceVarP(&ownWallets, "wallets", "w", []string{}, "Own wallet addresses")
	_ = viper.BindPFlag("wallets", rootCmd.Flags().Lookup("wallets"))

	// apis
	rootCmd.PersistentFlags().StringVar(&apiKeyEtherscan, "etherscan", "", "Etherscan API Key")
	_ = viper.BindPFlag("api_keys.etherscan", rootCmd.Flags().Lookup("etherscan"))
	rootCmd.PersistentFlags().StringVar(&apiKeyOpensea, "opensea", "", "Opensea API Key")
	_ = viper.BindPFlag("api_keys.opensea", rootCmd.Flags().Lookup("opensea"))

	// // websockets server
	// rootCmd.PersistentFlags().Bool("server", false, "Start websockets server")
	// _ = viper.BindPFlag("server.enabled", rootCmd.Flags().Lookup("server"))
	// rootCmd.PersistentFlags().IP("host", net.IPv4(0, 0, 0, 0), "Websockets server port")
	// _ = viper.BindPFlag("server.host", rootCmd.Flags().Lookup("host"))
	// rootCmd.PersistentFlags().Uint16("port", 42069, "Websockets server port")
	// _ = viper.BindPFlag("server.port", rootCmd.Flags().Lookup("port"))

	// defaults

	// api keys from node providers & other services
	viper.SetDefault("api_keys", map[string]string{"alchemy": "", "infura": "", "moralis": "", "opensea": "", "etherscan": ""})

	// redis settings
	viper.SetDefault("redis.enabled", false)
	viper.SetDefault("redis.host", "127.0.0.1")
	viper.SetDefault("redis.port", 6379)
	viper.SetDefault("redis.database", 0)
	viper.SetDefault("redis.password", "")

	viper.SetDefault("cache.names_ttl", 2*24*time.Hour)
	viper.SetDefault("cache.ens_ttl", 3*24*time.Hour)
	// viper.SetDefault("cache.sales_ttl", 7*24*time.Hour)
	// viper.SetDefault("cache.listings_ttl", 7*24*time.Hour)

	viper.SetDefault("server.websockets.enabled", false)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Find home directory.
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".btv" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".gloomberg.yaml")
	}

	viper.SetConfigType("yaml")

	// backup config
	// viper.SetDefault("config.backup_file", "/tmp/.backup_gloomberg.yaml")

	// logging
	// viper.SetDefault("log.log_file", fmt.Sprint(home, "gloomberg.log"))
	viper.SetDefault("log.log_file", "/tmp/gloomberg.log")

	// environment variables
	viper.SetEnvPrefix("GLOOMBERG")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		//nolint:errorlint
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			// fmt.Printf("config file not found: %s\n", viper.ConfigFileUsed())
		} else {
			// Config file was found but another error was produced
			fmt.Printf("config file error: %s - %s\n", viper.ConfigFileUsed(), err.Error())
		}
	}

	gbl.InitSugaredLogger()
}
