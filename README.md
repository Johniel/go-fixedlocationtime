# go-fixedlocationtime
This pckage provide yen another Time data type to check timezone in compilation time.

``` go
t := JST{}
```

``` go
jst := New[LocationJST]()
utc := New[LocationUTC]()
if jst == utc {  // Compilation error
  // do something
}
```
