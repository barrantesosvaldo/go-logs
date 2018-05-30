# go-logs
Create logs for Information, Warnings, Errors and Debugs.

It allows the creation of log files with the keywords: **INFO**, **WARN**, **ERRO** and **DEBU** to record information, warnings, errors and debugs.
A new file is generated every exact hour.

## Install
go get https://github.com/barrantesosvaldo/go-logs

## Usage
First you have to initialize the logs with:
```go
gologs.Init()
```
or you can do it with:
```go
gologs.Init("FileName")
```
By default the names of the generated files will have their names with the date and time (minutes and seconds in zero) in which they were generated

## Example

```go
package main

import (
    "github.com/barrantesosvaldo/go-logs"
)

func main() {
    // initialize logs
    gologs.Init()
  
    gologs.Info.Println("This is an information log")
    gologs.Warning.Println("This is a warning log")
    gologs.Error.Println("This is an error log")
    gologs.Debug.Println("This is a debug log")
}
```
