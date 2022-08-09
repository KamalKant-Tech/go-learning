# go-learning

## Go Module
Go modules commonly consist of one project or library and contain a collection of Go packages that are then released together. Go modules solve many problems with GOPATH , the original system, by allowing users to put their project code in their chosen directory and specify versions of dependencies for each module

### Commands
go mod init <module-name>// to initiate go modules it will go.mod file in project root dir

### Go Packages
To use a custom package we must import it first. The import path is the name of the module appended by the subdirectory of the package and the package name.
