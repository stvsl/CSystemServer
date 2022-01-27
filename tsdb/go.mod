module stvsljl.com/stvsl/influxdb

go 1.17

require github.com/influxdata/influxdb-client-go/v2 v2.6.0

require stvsljl.com/stvsl/Sql v0.0.0

replace stvsljl.com/stvsl/Sql v0.0.0 => ../db

require (
	github.com/deepmap/oapi-codegen v1.8.2 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/influxdata/line-protocol v0.0.0-20200327222509-2487e7298839 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	gorm.io/driver/mysql v1.2.3 // indirect
	gorm.io/gorm v1.22.5 // indirect
)
