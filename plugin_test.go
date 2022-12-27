package plugin

import (
	"os"
	"testing"

	"github.com/spf13/viper"
)

func Test(t *testing.T) {
	plugin := SkywalkingPlugin{}
	if !plugin.Enabled() {
		t.Fatal("enabled default true")
	}
	if plugin.Order() != defaultConfig["order"] {
		t.Fatal("order should be default")
	}
	if err := plugin.Load(); err == nil {
		t.Fatal("backend address is null")
	}

}

func TestLoad(t *testing.T) {
	plugin := SkywalkingPlugin{}
	viper.Set(configPrefix+"address", "127.0.0.1:11800")
	if err := plugin.Load(); err != nil {
		t.Fatal(err)
	}
	os.Setenv("SW_AGENT_SAMPLE", "bad value")
	if err := plugin.Load(); err == nil {
		t.Fatal("read trace options from env failed")
	}
	os.Setenv("SW_AGENT_SAMPLE", "")

}
