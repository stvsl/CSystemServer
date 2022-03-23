module stvsljl.com/stvsl/influxdb

go 1.18

require github.com/influxdata/influxdb-client-go/v2 v2.8.0

require stvsljl.com/stvsl/Sql v0.0.0

replace stvsljl.com/stvsl/Sql v0.0.0 => ../db

require (
	github.com/deepmap/oapi-codegen v1.9.1 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/influxdata/line-protocol v0.0.0-20210922203350-b1ad95c89adf // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/mysql v1.3.2 // indirect
	gorm.io/gorm v1.23.3 // indirect
)
