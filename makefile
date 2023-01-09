up:
	migrate -path db/migration -database "mysql://root:root@tcp(localhost:3306)/bank" -verbose up 1
down:
	migrate -path db/migration -database "mysql://root:root@tcp(localhost:3306)/bank" -verbose down 1
init:
	migrate create -ext sql -dir db/migration -seq ${d}
sqlc:
	sqlc generate
format: # 格式化并检查代码
	goimports -w . && gofmt -w . && golangci-lint run
#执行所有的单元测试
test:
	go test -v -cover ./...
run:
	go run server/server.go
installMock:
	 go install github.com/golang/mock/mockgen@v1.6.0
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/moonman/mbank/db/sqlc Store
docker-run:
	docker run --name bank -p 8080:8080 --network aro -d -e DbSource="root:root@tcp(mysql-learn:3306)/bank?parseTime=true&loc=Local" bank:latest
del-compose-file:
	docker compose down && docker rmi Container bank-compose-api-1
.phony: up down init sqlc format test run installMock mock