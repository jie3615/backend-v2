package main

import (
	"flag"
	"fmt"
	"runtime"

	"metaLand/app/api/internal/config"
	"metaLand/app/api/internal/handler"
	"metaLand/app/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

//var configFile = flag.String("f", "etc/meta_land.yaml", "the config file")

func main() {

	var configPath string
	// get current system
	os := runtime.GOOS
	switch os {
	case "windows":
		configPath = "F:\\xxx\\backend-v2-deving\\app\\api\\etc\\meta_land.yaml" // Please replace this with the actual path under Window
	case "darwin": // macOS
		configPath = "/path/to/your/config/on/macos/meta_land.yaml" // Please replace this with the actual path under macOS
	case "linux":
		configPath = "/path/to/your/config/on/linux/meta_land.yaml" // Please replace this with the actual path under Linux
	default:
		fmt.Printf("Unsupported operating system: %s\n", os)
		return
	}
	configFile := flag.String("f", configPath, "the config file")

	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
