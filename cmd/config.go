package cmd

import "github.com/spf13/viper"

// Flag describes a command flag - credit textile
type Flag struct {
	Key          string
	DefaultValue interface{}
}

// Config describes a command configuration params and file info - credit textile
type Config struct {
	Viper     *viper.Viper
	File      string
	Directory string
	Name      string
	Flags     map[string]Flag
	EnvPrefix string
}
