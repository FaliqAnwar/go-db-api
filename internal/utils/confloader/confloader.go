package confloader

import (
	"errors"
	"fmt"
	"go-db-api/internal/model"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

var ErrConfigFileNotFound = errors.New("config file not found")

type (
	viperLoader struct {
		viper                *viper.Viper
		envPrefix            string
		configFileName       string
		configFileSearchPath []string
	}

	options func(*viperLoader)
)

func WithConfigFileName(fileName string) options {
	return func(vl *viperLoader) {
		vl.configFileName = fileName
	}
}

func WithConfigFileSearchPath(paths ...string) options {
	return func(vl *viperLoader) {
		vl.configFileSearchPath = append(vl.configFileSearchPath, paths...)
	}
}

func New(envPrefix string, opts ...options) *viperLoader {
	v := &viperLoader{
		viper:                viper.New(),
		envPrefix:            envPrefix,
		configFileName:       "config",
		configFileSearchPath: []string{"."},
	}

	for _, opt := range opts {
		opt(v)
	}

	return v
}

func (v viperLoader) Load(cfg any) (err error) {
	decOption := func(dc *mapstructure.DecoderConfig) {}
	err = v.loadFromFileAndEnv()
	if err != nil {
		if errors.Is(err, ErrConfigFileNotFound) {
			err = fmt.Errorf("%w: no '%s' file found on search paths", ErrConfigFileNotFound, v.configFileName)
			return
		}

		return err
	}

	err = v.viper.Unmarshal(&cfg, decOption)

	return
}

func (v viperLoader) loadFromFileAndEnv() (err error) {
	v.viper.SetConfigName(v.configFileName)
	for _, path := range v.configFileSearchPath {
		v.viper.AddConfigPath(path)
	}
	v.viper.SetEnvPrefix(v.envPrefix)
	v.viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.viper.AutomaticEnv()
	return v.viper.ReadInConfig()
}

func MustLoad(appName string) (config model.Config, err error) {
	cfg := New(
		appName,
		WithConfigFileName("config/config.local"),
		WithConfigFileSearchPath(fmt.Sprintf("/Users/abdulfaliqanwar/Documents/belajar/golang/repository/%s", appName)),
	)

	err = cfg.Load(&config)

	return
}
