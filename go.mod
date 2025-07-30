module github.com/Qjoyboy/tasks-service

go 1.24.4

require (
	google.golang.org/grpc v1.73.0
	gorm.io/gorm v1.30.0
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.7.5 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	golang.org/x/crypto v0.39.0 // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sync v0.15.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250324211829-b45e905df463 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)

require (
	github.com/Qjoyboy/project-proto v0.4.4-0.20250718111222-7132168753e3
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/text v0.26.0 // indirect
	gorm.io/driver/postgres v1.6.0
)

replace github.com/Qjoyboy/tasks-service => ../tasks-service
