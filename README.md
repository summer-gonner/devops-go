# devops-go


goctl rpc protoc sys.proto --go_out=./ --go-grpc_out=./ --zrpc_out=.

goctl api go --api user.api --dir ../ --style=goZerocls
goctl model mysql ddl -src sys_user.sql -dir ./model/sysmodel