# Basics of GOLANG ->
### What are packages in GO ->
  - a package is a way to group functions, and it's made up of all the files in the same directory.
  - When you need your code to do something that might have been implemented by someone else, you can look for a package that has functions you can use in your code.
  - Code functionalities are wrapped inside "packages".

### Intial Step ->
  - Everytime when you create a folder or want to start writing code always initialise the module. 
  - This makes application production ready.
  - `go mod init` --> similar to `npm init`
  - Recommended --> `go mod init <githubRepoUrl>`
  - Generates the `go.mod` file 

### GO is amazing  ->
  - main() : Reserved Keyword. Very similar to c/c++. Acts as an single entry point. 
  - Go is case-sensitive. Packages methods starts with capital letter. 
  - Go automatically cleans up the packages.
  - `go run <package_name>`
  - `go run main.go` -> This is only going to compile and run the main.go  . It is not going to build the executables. 
  - WOW ! You can generate executables of different OS . 
  - `go help` -> lists the series os available commands
  - `GOPATH` -> this gives the GOPATH environment variables. Like in java we need to setup the PATH so that we can use java from anywhere .
        - If the environment variable is unset, GOPATH defaults
          to a subdirectory named "go" in the user's home directory
          ($HOME/go on Unix, %USERPROFILE%\go on Windows),
          unless that directory holds a Go distribution.
          Run "go env GOPATH" to see the current GOPATH. -> Example : /Users/poorvadityabehre/go

  - Golang also have lexar which checks the language syntax. (like c/c++/java). Here Lexar automatically puts the semicolon. But generally good practise to terminate line with semicolon.

  - OVERVIEW : 

  - WHAT ARE TYPES IN GO LANG ->
      - Variables type should be known in advance
      - Everything is a Type. ( like Everything in Java is an object).
      - `TYPES` : String , Bool , Integer (uint8 , uint64 , int8 , int64, uintptr) , Floating (float32 , float64) , Complex.
      - `ADVANCE TYPES` : Array , Slices , Maps, Structs , Pointers , Functions , Channels , Mutex ....
  
  - Creating a variable
      - use keyword "var"
      - `var <variable_Name> dataType`
      - Example: `var username string = "Poorvaditya"`
  
  - `go run .` -> cmnd both compiles and runs the program.
  - 
  