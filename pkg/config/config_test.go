package config

import (
	"encoding/json"
	"path"
	"testing"
)

func TestPath(t *testing.T) {
	srcFile := "device/sdk/app.toml"

	filename := path.Base(srcFile)            // app.toml
	ext := path.Ext(srcFile)                  // .toml
	dir := path.Dir(srcFile)                  // device/sdk
	name := filename[:len(filename)-len(ext)] // app

	if filename != "app.toml" || ext != ".toml" || dir != "device/sdk" || name != "app" {
		t.Fatal("not match expected")
	}
}

func TestMustLoadConfig(t *testing.T) {
	want := `{"etcd":{"Name":"etcd","Retries":3,"SmartDNS":"group.etcd.www.cn","Timeout":500},"mysql":{"Service":{"Name":"mysql","Retries":3,"SmartDNS":"group.mysql.www.cn","Timeout":500},"Read":{"Host":"127.0.0.1:3306","UserName":"root","Password":"root","DBName":"go_gin_api","MaxIdleConn":60,"MaxOpenConn":10},"Write":{"Host":"127.0.0.1:3306","UserName":"root","Password":"root","DBName":"go_gin_api","MaxIdleConn":60,"MaxOpenConn":10}},"redis":{"Service":{"Name":"redis","Retries":3,"SmartDNS":"group.redis.www.cn","Timeout":500},"Config":{"Host":"127.0.0.1:6379","Password":"2222","DB":"0","MaxIdle":1000,"MaxActive":10,"IdleTimeout":6000,"MaxRetries":3,"PoolSize":10}}}`

	type args struct {
		confLocation string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{"./testdata"}, want},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MustLoadConfig(tt.args.confLocation)
			b, err := json.Marshal(ConfigMap)
			if err != nil {
				t.Fatal(err)
			}

			if string(b) != tt.want {
				t.Fatal("not match expected")
			}
		})
	}
}
