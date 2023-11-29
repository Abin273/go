- The go mod init command initializes a new Go module in the current directory

```
go mod init module_name 

```


### Run the Go Code:

```
go run file_name.go

```

### to take build

- If we want to build an executable binary instead, we can use the go build command,

```
go build file_name.go

```

-This will create an executable file named file_name (eg:- if file_name was hello.go then  "hello") (or hello.exe on Windows), and you can run it directly using,

 #### On Unix-like systems

```
./hello 

```

 #### On On Windows systems

```
hello 

```