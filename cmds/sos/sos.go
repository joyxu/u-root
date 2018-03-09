// Copyright 2018 the u-root Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "sync"

var (
	RWLock   sync.RWMutex
	Registry = make(map[string]int)
)

func read(serviceName string) (port int, exists bool) {
	RWLock.RLock()
	defer RWLock.RUnlock()
	port, exists = Registry[serviceName]
	return
}

func register(serviceName string, portNum int) {
	RWLock.Lock()
	defer RWLock.Unlock()
	Registry[serviceName] = portNum
}

func unregister(serviceName string) {
	RWLock.Lock()
	defer RWLock.Unlock()
	delete(Registry, serviceName)
}

func snapshotRegistry() map[string]int {
	RWLock.RLock()
	defer RWLock.RUnlock()
	snapshot := make(map[string]int)
	for name, port := range Registry {
		snapshot[name] = port
	}
	return snapshot
}

func main() {
	startServer()
}