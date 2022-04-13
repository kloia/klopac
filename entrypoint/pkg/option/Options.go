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
		"provision":   o.flag.Bool("provision", false, "It executes provisioner"),
		"validate":    o.flag.Bool("validate", false, "It executes both provisioner and validator"),
		"healthcheck": o.flag.Bool("healthcheck", false, "It executes finalizer"),
		"websocket":   o.flag.Bool("websocket", false, "It helps to make use of websocket connection - required uri, username, password"),
		"uri":         o.flag.String("uri", "", "websocket uri"),
		"username":    o.flag.String("username", "", "username for websocket connection"),
		"password":    o.flag.String("password", "", "password for websocket connection"),
		"loglevel":    o.flag.String("loglevel", "INFO", "It sets the level of the producing logs"),
		"valuesFile":  o.flag.String("valuesFile", "/data/values.yaml", "Value File"),
		"dataPath":    o.flag.String("files", "/data/", "Variable File"),
		"bundleFile":  o.flag.String("bundleFile", "/data/bundle.tar.gz", "Bundle File"),
	}
	o.flag.Parse()
}

func NewOptionService(f flag.Flag) *OptionService {
	service := OptionService{flag: f}
	service.setFlags()
	return &service
}
