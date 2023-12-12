# go-fixedlocationtime
This pckage provides yet another time.Time data type to check timezone in compilation time.

``` go
t := JST{}
```

``` go
jst := New[LocationJST]()
utc := New[LocationUTC]()
if jst == utc {  // Compilation error
}
```
