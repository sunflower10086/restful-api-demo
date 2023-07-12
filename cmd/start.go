package cmd

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sunflower10086/restful-api-demo/apps"
	"github.com/sunflower10086/restful-api-demo/apps/dao/db"
	"golang.org/x/net/context"

	"github.com/spf13/cobra"
	_ "github.com/sunflower10086/restful-api-demo/apps/all"
	myhttp "github.com/sunflower10086/restful-api-demo/apps/host/http"
	"github.com/sunflower10086/restful-api-demo/conf"
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

		apps.Init()

		// 注册HTTP服务
		api := myhttp.NewHostHTTPHandler()
		if err := api.Config(); err != nil {
			return err
		}

		// 启动服务
		g := gin.Default()
		api.RouteRegistry(g)

		Run(g, conf.C().App.HTTPAddr(), conf.C().App.Name)

		return nil
	},
}

func init() {

	StartCmd.PersistentFlags().StringVarP(&configFile, "config", "f", "./etc/demo.toml", "demo config file")
	RootCmd.AddCommand(StartCmd)
}

func Run(r *gin.Engine, Addr, srvName string) {
	srv := &http.Server{
		Addr:    Addr,
		Handler: r,
	}

	// 保证优雅启停
	go func() {
		log.Printf("%s running in %s \n", srvName, srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 相当于监听一下 kill -2 和 kill -9
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT (Ctrl + C)
	// kill -9 is syscall. SIGKILL
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("Shutdown %s ...\n", srvName)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("%s Shutdown err: %v\n", srvName, err)
	}
	// catching ctx.Done(). timeout of 2 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 2 seconds.")
	}
	log.Printf("%s exiting", srvName)
}
