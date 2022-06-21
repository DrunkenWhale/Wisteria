# Wisteria

use `core.StartREPL()` to test it

for example:

```shell

wisteria>  get 114514 
2022/06/21 23:37:56 key [114514] not exist
wisteria>  put 114514 1919810
wisteria>  get 114514
2022/06/21 23:38:02 1.91981e+06
wisteria>  put 114514 "test string" 
wisteria>  get 114514
2022/06/21 23:38:12 test string
wisteria>  put 1 LIST
wisteria>  append 1 1919810
2022/06/21 23:38:21 succeed
wisteria>  append 1 1919810.5178
2022/06/21 23:38:26 succeed
wisteria>  append 1 "type" 
2022/06/21 23:38:30 succeed
wisteria>  get 1
2022/06/21 23:38:33 [1.91981e+06 1.9198105178e+06 type]
wisteria>  put 2 MAP
wisteria>  append 2 key value    
2022/06/21 23:38:52 succeed
wisteria>  append 2 key2 value2
2022/06/21 23:39:00 succeed
wisteria>  append 2 key3 value3 
2022/06/21 23:39:05 succeed
wisteria>  get 2
2022/06/21 23:39:07 map[key:value key2:value2 key3:value3]

```

or you can use some boring `API`

```go

transform.BytesTransformToSomeType()
// get special data struct from byte array 

transform.SomeTypeTransformToBytes()
// transform special data struct to bytes

```