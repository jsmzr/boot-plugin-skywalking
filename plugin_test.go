package plugin

import "testing"

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
