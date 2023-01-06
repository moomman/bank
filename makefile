up:
	migrate -path db/migration -database "mysql://root:root@tcp(localhost:3306)/bank" -verbose up 1
down:
	migrate -path db/migration -database "mysql://root:root@tcp(localhost:3306)/bank" -verbose down 1
init:
	migrate create -ext sql -dir db/migration -seq init_schema
sqlc:
	sqlc generate
format: # 格式化并检查代码
	goimports -w . && gofmt -w . && golangci-lint run
#执行所有的单元测试
test:
	go test -v -cover ./...
.phony: up down