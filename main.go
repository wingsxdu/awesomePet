package main

import (
	"awesomePet/action"
	"awesomePet/api/debug"
	"awesomePet/echarts"
	"awesomePet/gorm_mysql"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"runtime"
)
func main() {
	Init()

	e := echo.New()
	e.Pre(middleware.HTTPSRedirect())
	e.Use(middleware.Logger())
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 4 << 10, // 4 KB
	}))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "echarts",
		Browse: true,
	}))

	user := e.Group("/user")
	visual := e.Group("/visual")
	jwt := e.Group("/jwt")
	jwt.Use(middleware.JWT([]byte("secret")))

	e.POST("/crawler", action.Crawler)

	user.POST("/register", action.Register)
	user.POST("/login", action.Login)
	user.PUT("/reset", action.Reset)

	visual.GET("/test", echarts.BarHandler)

	search := jwt.Group("/search")

	//查看请求信息
	e.GET("/info", func(c echo.Context) error {
		req := c.Request()
		format := `
			<code>
				Protocol: %s<br>
				Host: %s<br>
				Remote Address: %s<br>
				Method: %s<br>
				Path: %s<br>
			</code>
		`
		return c.HTML(http.StatusOK, fmt.Sprintf(format, req.Proto, req.Host, req.RemoteAddr, req.Method, req.URL.Path))
	})

	e.Logger.Fatal(e.StartTLS(":443", "./cert.pem", "./key.pem"))
}

func Init() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 6)
	fmt.Println("run CPUs number:", runtime.NumCPU())
	ReadConfig()
	args := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.Mysql.UserName, c.Mysql.UserPassword, c.Mysql.Address, c.Mysql.Database)
	gorm_mysql.Init(&args)
	grpc.Init()
}

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type Conf struct {
	GRPCAddress string `yaml:"gRPCAddress"`
	Mysql       struct {
		UserName     string `yaml:"UserName"`
		UserPassword string `yaml:"UserPassword"`
		Address      string `yaml:"Address"`
		Database     string `yaml:"Database"`
	}
}
var c = Conf{}

func ReadConfig() {
	data, err := ioutil.ReadFile("config.yaml")
	debug.PanicErr(err)
	fmt.Println(string(data))
	err = yaml.Unmarshal(data, &c)
	debug.PanicErr(err)
}