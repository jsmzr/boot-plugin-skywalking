package plugin

import (
	"errors"

	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
	config "github.com/jsmzr/boot-config"
	plugin "github.com/jsmzr/boot-plugin"
)

type SkywalkingPlugin struct{}

type SkywalkingProperties struct {
	Address *string
	Enabled *bool
	Name    *string
}

var t *go2sky.Tracer

func (s *SkywalkingPlugin) Load() error {
	var properties SkywalkingProperties
	if err := config.Resolve("boot.skywalking", &properties); err != nil {
		return err
	}
	if properties.Name == nil {
		return errors.New("skywalking name is nil")
	}
	if properties.Address == nil {
		return errors.New("skywalking address should not be nil")
	}
	r, err := reporter.NewGRPCReporter(*properties.Address)
	if err != nil {
		return err
	}
	defer r.Close()
	tracer, err := go2sky.NewTracer(*properties.Name, go2sky.WithReporter(r))
	if err != nil {
		return err
	}
	t = tracer
	return nil
}

func (s *SkywalkingPlugin) Enabled() bool {
	enabled, ok := config.Get("boot.skywalking.enabeld")
	if ok {
		return enabled.Bool()
	}
	return true
}

func (s *SkywalkingPlugin) Order() int {
	return 200
}

func GetTracer() *go2sky.Tracer {
	return t
}

func init() {
	plugin.Register("skywalking", &SkywalkingPlugin{})
}
