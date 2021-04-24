package doc

//go:generate go-bindata -nometadata -ignore=.DS_Store -pkg doc -o bindata.go swagger-ui/... thor.yaml

import (
	yaml "gopkg.in/yaml.v2"
)

//Version open api version
func Version() string {
	return version
}

var version string

type openAPIInfo struct {
	Info struct {
		Version string
	}
}

func init() {
	var oai openAPIInfo
	if err := yaml.Unmarshal(MustAsset("thor.yaml"), &oai); err != nil {
		panic(err)
	}
	version = oai.Info.Version
}
