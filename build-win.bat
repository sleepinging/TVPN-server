set GOPATH=%cd%
SET GOOS=windows

go build -o tvpn-server.exe ./src/TVPN-server

pause