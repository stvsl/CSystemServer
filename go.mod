module stvsljl.com/CSystemServer

go 1.18

require stvsljl.com/stvsl/Service v0.0.0

require (
	golang.org/x/text v0.3.7 // indirect
	stvsljl.com/stvsl/AES v0.0.0 // indirect
	stvsljl.com/stvsl/Sql v0.0.0 //indirect
)

require (
	github.com/deepmap/oapi-codegen v1.9.1 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.7.7 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.10.1 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.7 // indirect
	github.com/influxdata/influxdb-client-go/v2 v2.8.0 // indirect
	github.com/influxdata/line-protocol v0.0.0-20210922203350-b1ad95c89adf // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	golang.org/x/crypto v0.0.0-20220315160706-3147a52a75dd // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sys v0.0.0-20220319134239-a9b59b0215f8 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/mysql v1.3.2 // indirect
	gorm.io/gorm v1.23.3 // indirect
	stvsljl.com/stvsl/RSA v0.0.0 // indirect
	stvsljl.com/stvsl/influxdb v0.0.0
)

replace stvsljl.com/stvsl/Service v0.0.0 => ./Service

replace stvsljl.com/stvsl/influxdb v0.0.0 => ./tsdb

replace stvsljl.com/stvsl/Sql v0.0.0 => ./db

replace stvsljl.com/stvsl/RSA v0.0.0 => ./rsa

replace stvsljl.com/stvsl/AES v0.0.0 => ./aes
