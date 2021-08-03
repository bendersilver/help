package sync

import "sync"

// Mutex - sync.Mutex
type Mutex struct{ sync.Mutex }

// RWMutex - sync.RWMutex
type RWMutex struct{ sync.RWMutex }

// Map - sync.Map
type Map struct{ sync.Map }
