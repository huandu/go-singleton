// Copyright 2020 Huan Du. All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

// Package singleton provides a generic func to return a singleton of any type.
package singleton

import (
	"reflect"
	"sync"
)

var cache sync.Map

// Singleton returns a singleton of T.
func Singleton[T any]() (t *T) {
	hash := reflect.TypeOf(t)
	v, ok := cache.Load(hash)

	if ok {
		return v.(*T)
	}

	v = new(T)
	v, _ = cache.LoadOrStore(hash, v)
	return v.(*T)
}
