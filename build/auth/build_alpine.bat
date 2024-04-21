set GOOS=linux
go build -o auth_alpine ../../cmd/auth/main.go 
set GOOS=windows
pause