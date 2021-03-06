package main

import (
	"awesomePet/action"
	"awesomePet/api"
	"awesomePet/echarts"
	"awesomePet/gorm_mysql"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net/http"
	"runtime"
)

// init()函数会在每个包完成初始化后自动执行，并且执行优先级比 main 函数高。
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU() * 6)
	fmt.Println("run CPUs number:", runtime.NumCPU())
	var c = Conf{}
	c.ReadConfig()
	fmt.Println("\n初始化完成")
}

func main() {
	//Init()
	e := echo.New()
	e.Pre(middleware.HTTPSRedirect())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 2 << 10, // 2 KB
	}))
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "echarts",
		Browse: true,
	}))

	u := e.Group("/user")
	u.POST("/register", action.Register)
	u.POST("/login", action.Login)
	u.PUT("/reset", action.Reset)

	v := e.Group("/visual")
	v.GET("/test", echarts.TotalHandler)

	jwt := e.Group("/jwt")
	jwt.Use(middleware.JWT([]byte("yourSecret")))

	user := jwt.Group("/user")
	user.POST("/profile", action.ProfilePhoto)
	user.GET("/profile", action.ThumbnailProfilePhoto)
	user.GET("/info", action.GetUserInfo)
	user.PUT("/info", action.UpdateUserInfo)
	user.DELETE("/info", action.DeleteUser)
	user.GET("/blog", action.GetUserBlog)

	pet := jwt.Group("/pet")
	pet.GET("/blog", action.GetBlogById)
	pet.POST("/blog", action.UploadBlog)
	pet.PUT("/blog", action.UpdateBlogContext)
	pet.DELETE("/blog/:id", action.DeleteBlog)

	//search := jwt.Group("/search")

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

func (c *Conf) ReadConfig() {
	data, err := ioutil.ReadFile("config.yaml")
	api.PanicErr(err)
	fmt.Println(string(data))
	err = yaml.Unmarshal(data, &c)
	api.PanicErr(err)
	args := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Mysql.UserName, c.Mysql.UserPassword, c.Mysql.Address, c.Mysql.Database)
	fmt.Println(args)
	gorm_mysql.Init(&args)
}
