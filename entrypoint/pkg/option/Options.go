package option

import "entrypoint/pkg/flag"

type optionService struct {
	flag flag.Flag
}

func NewOptionService(f flag.Flag) *optionService {
	return &optionService{flag: f}
}

func (o optionService) Get() map[string]interface{} {
	options := map[string]interface{}{
		"provision":   o.flag.Bool("provision", false, "sadece provision çalıştırır"),
		"validate":    o.flag.Bool("validate", false, "provision ve validate sıra ile çalıştırır"),
		"healthcheck": o.flag.Bool("healthcheck", false, "healthcheck environmentini set ederek sadece finalizer çalıştırır"),
		"websocket":   o.flag.Bool("websocket", false, "websocket"),
		"uri":         o.flag.String("uri", "", "websocket ile çalışacaksa iletişim yapılacak uri"),
		"username":    o.flag.String("username", "", "websocket ile çalışacaksa uri erişimi için kullanılacak username"),
		"password":    o.flag.String("password", "", "websocket ile çalışacaksa uri erişimi için kullanılacak password"),
		"loglevel":    o.flag.String("loglevel", "INFO", "üretilen log'ların seviyesini set eder"),
	}
	o.flag.Parse()
	return options
}
