package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/GoAdminGroup/example/handler"
	_ "github.com/GoAdminGroup/example/theme2"
	_ "github.com/GoAdminGroup/go-admin/adapter/gin" // web framework adapter
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/sqlite" // sql driver
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	_ "github.com/GoAdminGroup/themes/adminlte" // ui theme

	"github.com/GoAdminGroup/components/login"
	"github.com/GoAdminGroup/example/models"
	"github.com/GoAdminGroup/example/pages"
	"github.com/GoAdminGroup/example/tables"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/gin-gonic/gin"
)

var configFile = flag.String("f", "etc/config.json", "the config file")
var port = flag.String("p", "9900", "port")

func main() {
	flag.Parse()
	fmt.Println("port:", *port, "config file:", *configFile)
	startServer()
}

func startServer() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	login.Init(login.Config{
		Theme:         "theme2",
		CaptchaDigits: 4, // 使用图片验证码，这里代表多少个验证码数字
		// 使用腾讯验证码，需提供appID与appSecret
		// TencentWaterProofWallData: login.TencentWaterProofWallData{
		//    AppID:"",
		//    AppSecret: "",
		// }
	})

	r := gin.Default()
	r.GET("/admin/getContract", handler.Contract)

	eng := engine.Default()

	template.AddComp(chartjs.NewChart())
	adminPlugin := admin.NewAdmin(datamodel.Generators)
	if err := eng.AddConfigFromJSON(*configFile).
		AddPlugins(adminPlugin).
		AddGenerators(tables.Generators).
		AddGenerator("external", tables.GetExternalTable).
		Use(r); err != nil {
		panic(err)
	}
	// cfg := config.ReadFromJson("./config.json")
	// fmt.Println(fmt.Sprintf("%+v", cfg.Databases.GroupByDriver()))
	// eng.PluginList
	adminPlugin.SetCaptcha(map[string]string{"driver": login.CaptchaDriverKeyDefault})
	// captcha.Add(login.CaptchaDriverKeyDefault, new(login.DigitsCaptcha))

	// config mysql use
	// models.Init(eng.DefaultConnection(), eng.MysqlConnection())
	models.Init(eng.DefaultConnection(), nil)
	r.Static("/uploads", "./uploads")
	r.Static("/admin/web3", "./html/contractReader")

	eng.HTML("GET", "/admin", pages.DashboardPage)
	eng.HTML("GET", "/admin/form", pages.GetFormContent)
	eng.HTML("GET", "/admin/table", pages.GetTableContent)
	eng.HTMLFile("GET", "/admin/hello", "./html/hello.tmpl", nil)
	eng.HTMLFile("GET", "/admin/dao", "./html/contract.tmpl", nil)

	srv := &http.Server{
		Addr:    ":" + *port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()
	log.Printf("server start listen: http://127.0.0.1:%s/admin\n", *port)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
