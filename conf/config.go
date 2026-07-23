package conf

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-yaml"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Default() *Config {
	return &Config{
		App: &application{
			Host:      "127.0.0.1",
			Port:      8080,
			Domain:    "127.0.0.1",
			UploadDir: "./uploads",
		},
		Mysql: &mySql{
			Host:     "127.0.0.1",
			Port:     3306,
			Username: "root",
			Password: "Root123.",
			Database: "test",
			Debug:    false,
		},
		Redis: &redisConf{
			Host:     "127.0.0.1",
			Port:     6379,
			Password: "",
			DB:       0,
		},
	}
}

type Config struct {
	App   *application `json:"app" yaml:"app"`
	Mysql *mySql       `json:"mysql" yaml:"mysql"`
	Redis *redisConf   `json:"redis" yaml:"redis"`
}

func (c *Config) Yaml() string {
	cj, _ := yaml.Marshal(c)
	return string(cj)
}

type application struct {
	Host      string `json:"host" yaml:"host"`
	Port      int    `json:"port" yaml:"port"`
	Domain    string `json:"domain" yaml:"domain"`
	UploadDir string `json:"uploaddir" yaml:"uploaddir"`

	server *gin.Engine
	root   gin.IRouter
	lock   sync.Mutex
}

func (a *application) GinServer() *gin.Engine {
	if a.server == nil {
		a.server = gin.Default()
	}
	return a.server
}

func (a *application) GinRootServer() gin.IRouter {
	r := a.GinServer()
	if a.root == nil {
		a.root = r.Group("vblog").Group("apps").Group("v1")
	}
	return a.root
}

func (a *application) Address() string {
	return fmt.Sprintf("%s:%d", a.Host, a.Port)
}

func (a *application) Start() error {
	r := a.GinServer()
	return r.Run(a.Address())
}

type mySql struct {
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Database string `json:"database" yaml:"database"`
	Debug    bool   `json:"debug" yaml:"debug"`

	db   *gorm.DB
	lock sync.Mutex
}

func (m *mySql) DSN() string {

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", m.Username, m.Password, m.Host, m.Port, m.Database)
}

func (m *mySql) GetDB() *gorm.DB {
	m.lock.Lock()
	defer m.lock.Unlock()
	if m.db == nil {
		db, err := gorm.Open(mysql.Open(m.DSN()), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		m.db = db
		if m.Debug {
			m.db = db.Debug()
		}
	}
	return m.db
}

type redisConf struct {
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	Password string `json:"password" yaml:"password"`
	DB       int    `json:"db" yaml:"db"`

	client *redis.Client
	lock   sync.Mutex
}

func (r *redisConf) GetClient() *redis.Client {
	r.lock.Lock()
	defer r.lock.Unlock()

	if r.client == nil {
		r.client = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
			Password: r.Password,
			DB:       r.DB,
		})
	}
	return r.client
}
