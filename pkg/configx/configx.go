package configx

import (
	"encoding/json"
	"fmt"

	"github.com/blackhorseya/mundo/pkg/logging"
	"github.com/blackhorseya/mundo/pkg/netx"
)

// Config defines the config struct.
type Config struct {
	ID   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`

	Log  logging.Config `json:"log" yaml:"log"`
	HTTP HTTP           `json:"http" yaml:"http"`

	LineBot struct {
		Secret string `json:"secret" yaml:"secret"`
		Token  string `json:"token" yaml:"token"`
	} `json:"line_bot" yaml:"lineBot"`

	OpenAI struct {
		Key string `json:"key" yaml:"key"`
	} `json:"openai" yaml:"openai"`

	Storage struct {
		Mongodb struct {
			DSN string `json:"dsn" yaml:"dsn"`
		} `json:"mongodb" yaml:"mongodb"`
	} `json:"storage" yaml:"storage"`
}

func (x *Config) String() string {
	bytes, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		return err.Error()
	}

	return string(bytes)
}

// HTTP defines the http struct.
type HTTP struct {
	URL  string `json:"url" yaml:"url"`
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
	Mode string `json:"mode" yaml:"mode"`
}

// GetAddr is used to get the http address.
func (http *HTTP) GetAddr() string {
	if http.Host == "" {
		http.Host = "0.0.0.0"
	}

	if http.Port == 0 {
		http.Port = netx.GetAvailablePort()
	}

	return fmt.Sprintf("%s:%d", http.Host, http.Port)
}
