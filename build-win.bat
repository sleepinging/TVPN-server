set GOPATH=%cd%
SET GOOS=windows

go build -o myvpn-server.exe ./src/mytap

pause