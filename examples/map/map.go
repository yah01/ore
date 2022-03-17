package main

import (
	"errors"
	"fmt"

	. "github.com/yah01/ore"
)

type Map[K comparable, V any] struct {
	inner map[K]V
}

func (mp *Map[K, V]) Get(key K) Result[V] {
	v, ok := mp.inner[key]
	if !ok {
		return mp.err(errors.New("key not found"))
	}

	return Ok(v)
}

func (mp *Map[K, V]) Insert(key K, value V) EResult {
	mp.inner[key] = value
	return OkResult()
}

func (mp *Map[K, V]) Remove(key K, value V) EResult {
	delete(mp.inner, key)
	return OkResult()
}

func (mp *Map[K, V]) err(err error) Result[V] {
	return Err[V](err)
}

func main() {
	mp := Map[string, string]{inner: map[string]string{}}

	if err := mp.Get("hello").Err(); err != nil {
		fmt.Println("failed to get `hello`, err:", err)
	}

	if res := mp.Insert("hello", "world"); res.IsErr() {
		fmt.Println("failed to insert data, err:", res.Err())
	}

	if res := mp.Get("hello"); res.IsErr() {
		fmt.Println("failed to get value of `hello`, err =", res.Err())
	} else {
		fmt.Println("value of `hello`:", res.Value())
	}
}
