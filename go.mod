module cmd/websvr/main.go

go 1.15

require (
	127.0.0.1/version v0.0.0-00010101000000-000000000000
	go.uber.org/zap v1.16.0
)

replace (
	127.0.0.1/version => ./version
)
