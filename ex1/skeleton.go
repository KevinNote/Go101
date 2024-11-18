package main

import (
	"slices"
	"sync/atomic"
	"time"
)

type NS map[string]*PNode

type PNode struct {
	LastUseTime int64 // timestamp
	Loaded      bool
	LoadedTime  int64
	FreqCount   atomic.Int64
}

func (p *PNode) Load() {
	p.Loaded = true
	p.LastUseTime = time.Now().Unix()
	p.LoadedTime = time.Now().Unix()
}

func (p *PNode) Unload() {
	p.Loaded = false
}

func LRU(ns NS, k int) {
  // write here
}

func LFU(ns NS, k int) {
  // write here
}

func FIFO(ns NS, k int) {
  // write here
}
