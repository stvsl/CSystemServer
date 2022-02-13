module stvsljl.com/stvsl/Service

go 1.17

require github.com/gin-gonic/gin v1.7.7

require stvsljl.com/stvsl/influxdb v0.0.0

require stvsljl.com/stvsl/Sql v0.0.0

require stvsljl.com/stvsl/RSA v0.0.0

require (
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

replace stvsljl.com/stvsl/RSA v0.0.0 => ../rsa

replace stvsljl.com/stvsl/Sql v0.0.0 => ../db

replace stvsljl.com/stvsl/influxdb v0.0.0 => ../tsdb

require (
	github.com/deepmap/oapi-codegen v1.9.1 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.10.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/influxdata/influxdb-client-go/v2 v2.7.0 // indirect
	github.com/influxdata/line-protocol v0.0.0-20210922203350-b1ad95c89adf // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/ugorji/go/codec v1.2.6 // indirect
	golang.org/x/crypto v0.0.0-20220210151621-f4118a5b28e2 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20220209214540-3681064d5158 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/mysql v1.2.3 // indirect
	gorm.io/gorm v1.22.5 // indirect
)
