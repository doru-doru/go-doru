package main

import (
	logging "github.com/ipfs/go-log/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const daemonName = "dorud"

var (
	log = logging.Logger("dorud")

	config = &cmd.Config
)

var rootCmd = &cobra.Command{
	Use: daemonName,
	Short: "Doru daemon",
	Long:  "Doru daemon grows consistent data tries.",
	PersistentPreRun: func(c *cobra.Command, args []string) {
		config.Viper.SetConfigType("yaml")
		cmd.ExpandConfigVars(config.Viper, config.Flags)

	},
	Run: func(c *cobra.Command, args []string) {

	},
}

func init() {

}

func main() {
	cmd.ErrCheck(rootCmd.Execute())
}
