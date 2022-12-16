package plugin

import (
	"fmt"

	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
	plugin "github.com/jsmzr/boot-plugin"
	"github.com/spf13/viper"
)

type SkywalkingPlugin struct{}

const configPrefix = "boot.skywalking"

var defaultConfig map[string]interface{} = map[string]interface{}{"enabled": true, "order": 40, "name": "boot-skywalking"}

var t *go2sky.Tracer

func (s *SkywalkingPlugin) Load() error {
	address := viper.GetString(configPrefix + ".address")
	if address == "" {
		return fmt.Errorf("skywalking backend address is null")
	}
	r, err := reporter.NewGRPCReporter(address)
	if err != nil {
		return err
	}
	name := viper.GetString(configPrefix + ".name")
	if tracer, err := go2sky.NewTracer(name, go2sky.WithReporter(r)); err != nil {
		return err
	} else {
		t = tracer

		return nil
	}
}

func (s *SkywalkingPlugin) Enabled() bool {
	return viper.GetBool(configPrefix + ".enabled")
}

func (s *SkywalkingPlugin) Order() int {
	return viper.GetInt(configPrefix + ".order")
}

func GetTracer() *go2sky.Tracer {
	return t
}

func init() {
	for key := range defaultConfig {
		viper.SetDefault(configPrefix+"."+key, defaultConfig[key])
	}
	plugin.Register("skywalking", &SkywalkingPlugin{})
}
