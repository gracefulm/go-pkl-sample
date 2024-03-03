// Code generated from Pkl module `introduction.of.pkl`. DO NOT EDIT.
package config

type AppConfig struct {
	Host string `pkl:"host"`

	Port uint16 `pkl:"port"`

	LogConfig *LogConfig `pkl:"logConfig"`
}
