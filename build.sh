cd server
cd src
set GOARM=6
export GOARM
set GOARCH=arm
export GOARCH
set GOOS=linux
export GOOS 
go build main.go