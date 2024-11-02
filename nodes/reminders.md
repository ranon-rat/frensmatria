for running the sql script i will have to execute this

```sh
cat init.sql|sqlite3 database.db
```

when compiling i will have to set something like this

``` sh
go env -w CGO_ENABLED=1
```
