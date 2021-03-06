package option

import (
	"entrypoint/pkg/flag"
)

type OptionService struct {
	flag   flag.Flag
	Params map[string]interface{}
}

// It defines args' values and their specified names
func (o *OptionService) setFlags() {
	o.Params = map[string]interface{}{
		"provision":     o.flag.Bool("provision", false, "It executes provisioner"),
		"validate":      o.flag.Bool("validate", false, "It executes both provisioner and validator"),
		"healthcheck":   o.flag.Bool("healthcheck", false, "It executes finalizer"),
		"websocket":     o.flag.Bool("websocket", false, "It helps to make use of websocket connection - required uri, username, password"),
		"uri":           o.flag.String("uri", "", "websocket uri"),
		"username":      o.flag.String("username", "", "username for websocket connection"),
		"password":      o.flag.String("password", "", "password for websocket connection"),
		"logLevel":      o.flag.String("logLevel", "INFO", "It sets the level of the producing logs"),
		"logFile":       o.flag.String("logFile", "/data/entrypoint.log", "Log file of entrypoint"),
		"valuesFile":    o.flag.String("valuesFile", "/data/values.yaml", "Value File"),
		"dataPath":      o.flag.String("dataPath", "/data/", "Data File Path"),
		"varsPath":      o.flag.String("varsPath", "/data/vars", "Variable Directory"),
		"manifestsPath": o.flag.String("manifestsPath", "/data/manifests", "Manifests Directory"),
		"bundleFile":    o.flag.String("bundleFile", "/data/bundle/bundle.tar.gz", "Bundle File"),
	}
	o.flag.Parse()
}

func NewOptionService(f flag.Flag) *OptionService {
	service := OptionService{flag: f}
	service.setFlags()
	return &service
}
