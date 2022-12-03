# Go maps

It's modern package that based on generic maps. The package provides many helpers to operate on thread-safe (only) maps. You don't have to worry about the concurrency moment of this map. Under the hood, realization is based on sync.Map. 

 
# Installation
```
go get github.com/0x9ef/go-maps
```

# Usage
```go
import "github.com/0x9ef/go-maps"
```

## DefaultMap
We can start from the native map realisation. 

### Set
We can set map key with a value. 
```go
m := maps.NewDefaultMap[string, int]()
m.Set("one", 1) // sets for key "one" value "1" 
``` 

### SetIf
We can set map key with a value if predicate function is true. 
```go
m := maps.NewDefaultMap[string, int]()
m.Set("ten", 10)
m.SetIf("one", 1, func(m Map[int, int]) bool {
	return m.Get("ten") == 10 // sets only if "ten" key is equals to 10
})
```  

### Get
We can get a map key and its value.
```go
m := maps.NewDefaultMap[string, int]()
m.Set("one", 1)
val := m.Get("one")
if val == 1 { 
    fmt.Println("found")
}
``` 

### GetOk
We can get a key and identifier if this key was found.
```go
m := maps.NewDefaultMap[string, int]()
val, ok := m.GetOk("one")
if !ok {
    fmt.Println("key was not found")
} 
```   

### Delete
We can delete a key from the map. 
```go
m := maps.NewDefaultMap[string, int]()
m.Set("one", 1)
m.Delete("one")
```

### DeleteIf
We can delete a key from the map if predicate function is true. 
```go
m := maps.NewDefaultMap[string, int]()
m.Set("one", 1)
m.DeleteIf("one", func(m Map[int, int]) bool {
	return !m.Exists("ten") // deletes value only if key "ten" doesn't exists
})
``` 

### Clear
We can clear full map, so this means that we can delete all keys from the map.
```go
m := maps.NewDefaultMap[string, int]()
m.Set("one", 1)
m.Set("two", 2)
m.Set("three", 3)
m.Clear()
``` 

### Keys
We can get all keys from the map.
```go
m := maps.NewDefaultMap[string, int]()
m.Set("one", 1)
m.Set("two", 2)
m.Set("three", 3)
fmt.Println(m.Keys())
// ["one", "two", "three"]
```

### Values
We can get all values from the map.
```go
m := maps.NewDefaultMap[string, int]()
m.Set("one", 1)
m.Set("two", 2)
m.Set("three", 3)
fmt.Println(m.Values())
// [1, 2, 3]
```

### Filter
We can filter map content and return only keys and values that was matched by 
our rules.
```go
m := maps.NewDefaultMap[string, int]()
m.Set("one", 1)
m.Set("two", 2)
m.Set("three", 3)
m.Set("four", 3)
keys, values := m.Filter(func(key string,  value int) bool {
	return key == "three" && value == 3 // we match third element in the map
})
fmt.Println(keys, values)
// ["three"], [1]
```

### Iterate
We can iterate over all keys in the map
```go
m := maps.NewDefaultMap[string, int]()
m.Set("one", 1)
m.Set("two", 2)
m.Set("three", 3)
m.Iterate(func(key string, value int) bool{
    fmt.Println(key, value)
})
// "one" 1
// "two" 2
// "three" 3
```

## License

[MIT](./LICENSE)