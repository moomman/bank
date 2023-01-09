#!/bin/sh
#必须写在第一行！！！！！

#出现异常立即退出
set -e
echo "run db migrate"

#可以采用等待一段时间的方式，也可以采用loop
sleep 5
/app/migrate -path /app/migration -database "mysql://root:root@tcp(compose-mysql:3306)/bank" -verbose up
echo "start app"

#接受输入的所有参数
exec "$@"