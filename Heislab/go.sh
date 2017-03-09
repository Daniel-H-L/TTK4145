export GOPATH=$(pwd)

go install DriveElevator
go install Master
go install Network
go install Slave

go run main.go