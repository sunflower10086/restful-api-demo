package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"

	"github.com/spf13/cobra"
	myhttp "github.com/sunflower10086/restful-api-demo/apps/host/http"
	"github.com/sunflower10086/restful-api-demo/apps/host/impl"
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

		// 加载服务实现
		service := impl.NewHostServiceImpl()
		// 注册HTTP服务
		api := myhttp.NewHostHTTPHandler(service)

		g := gin.Default()
		api.Register(g)

		Run(g, fmt.Sprintf("%s:%s", conf.C().App.Host, conf.C().App.Port), conf.C().App.Name)

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
