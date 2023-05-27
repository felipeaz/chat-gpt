package build

import (
	"flag"
)

const (
	_openaiTokenFlag            = "openaitoken"
	_openaiTokenFlagDescription = "open ai api token used to make client calls"
)

type Config struct {
	Flags *Flags
}

type Flags struct {
	OpenAIToken string
}

func MakeConfig() *Config {
	return &Config{
		Flags: makeFlags(),
	}
}

func makeFlags() *Flags {
	cfg := new(Flags)
	flag.StringVar(&cfg.OpenAIToken, _openaiTokenFlag, "", _openaiTokenFlagDescription)
	flag.Parse()
	return cfg
}
