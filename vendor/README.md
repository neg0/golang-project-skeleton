# Vendor
You should be using vgo `go mod init` for any Golang project since version 1.2. This folder contains all dependencies for the application.

Since Go 1.2 we won't be using golang dep or care about `GOPATH`. Project should be instantiated using go module:

    ~$: go mod init <name-of-repository>

