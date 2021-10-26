# juuri
A GraphQL API vulnerability scanner made with Go. Contributions are welcome!

Currently work in progress, but the aim is to check what kind of data the API exposes with e.g. introspection.

## Build
```
go get
go build
```

## Run
```
./juuri <options> https://endpoint/api
```
### Options
Show debug messages
```
-debug=true
```
Output to file
```
-file=output.txt
```
