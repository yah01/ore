# Ore

Thanks to Go 1.18, Now we can implement Option and Result like Rust!

## Get Started

### Option
You can use Option to express "nullable", just like all of the gophers did with pointer:

```golang
type Map[K comparable, V any] struct {
	inner map[K]V
}

func (mp *Map[K, V]) Get(key K) Option[V] {
	v, ok := mp.inner[key]
	if !ok {
		return None[V]()
	}

	return Some(v)
}

func main() {
    // init a Map

    if v := mp.Get("hello"); v.Some() {
        fmt.Println("we got value:", v.Value())
    } else {
        fmt.Println("failed to get value")
    }
}
```

### Result
Result is "value or error":
```golang
func (mp *Map[K, V]) Get(key K) Result[V] {
	v, ok := mp.inner[key]
	if !ok {
		return mp.err(errors.New("key not found"))
	}

	return Ok(v)
}

func main() {
    // init a Map

    if res := mp.Get("hello"); res.IsErr() {
		fmt.Println("failed to get value of `hello`, err =", res.Err())
	} else {
		fmt.Println("value of `hello`:", res.Value())
	}
}
```

Refer to [examples](examples/) for more.

