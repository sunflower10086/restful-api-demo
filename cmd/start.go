package cmd

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/sunflower10086/restful-api-demo/apps"
	_ "github.com/sunflower10086/restful-api-demo/apps/all"
	"github.com/sunflower10086/restful-api-demo/apps/dao/db"
	"github.com/sunflower10086/restful-api-demo/conf"
	"github.com/sunflower10086/restful-api-demo/protocol"
)

var (
	configFile string
)

var StartCmd = &cobra.Command{
	Use:     "start",
	Long:    "demo API后端",
	Short:   "demo API后端",
	Example: "demo API后端 commands",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 加载配置文件
		if err := conf.LoadConfigFromToml(configFile); err != nil {
			return err
		}

		// 启动与mysql的连接
		if err := db.Init(); err != nil {
			return err
		}

		// 注册服务到IOC
		// _ "github.com/sunflower10086/restful-api-demo/apps/host/impl"
		// 通过这个方法去注册一个服务
		//apps.HostService = impl.NewHostServiceImpl()

		// 注册服务到IOC
		apps.InitImpl()

		// 以后通过Ioc注册中心自动注册HTTP handler
		// 注册HTTP服务
		//api := myhttp.NewHostHTTPHandler()
		//if err := api.Config(); err != nil {
		//	return err
		//}

		//// 注册一个gin的实例
		//g := gin.Default()
		//
		//// 注册所有的 HTTP handler 方法
		//if err := apps.InitGinHandler(g); err != nil {
		//	return err
		//}
		//
		//Run(g, conf.C().App.HTTPAddr(), conf.C().App.Name)

		master := newMaster()

		// 相当于监听一下 kill -2 和 kill -9
		quit := make(chan os.Signal)
		// kill (no param) default send syscanll.SIGTERM
		// kill -2 is syscall.SIGINT (Ctrl + C)
		// kill -9 is syscall.SIGKILL
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		go master.WaitStop(quit)

		return master.Start()
	},
}

func newMaster() *master {
	return &master{
		http: protocol.NewHTTPService(),
	}
}

type master struct {
	http *protocol.HTTPService
}

func (m *master) Start() error {
	if err := m.http.Start(); err != nil {
		return err
	}
	return nil
}

func (m *master) Stop() {
	log.Printf("Shutdown %s ...\n", m.http.Conf.App.Name)
}

func (m *master) WaitStop(quit <-chan os.Signal) {
	for v := range quit {
		switch v {
		default:
			m.http.L.Printf("received signal: %s", v)
			m.http.Stop()
		}
	}
}

func init() {

	StartCmd.PersistentFlags().StringVarP(&configFile, "config", "f", "./etc/demo.toml", "demo config file")
	RootCmd.AddCommand(StartCmd)
}
