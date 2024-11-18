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
	var times []int64
	for _, node := range ns {
		if !node.Loaded {
			continue
		}
		times = append(times, node.LastUseTime)
	}
	if len(times) <= k {
		return
	}
	slices.Sort(times)

	seg := times[len(times)-k]

	for _, node := range ns {
		if node.LastUseTime >= seg || !node.Loaded {
			continue
		}
		node.Unload()
	}
}

func LFU(ns NS, k int) {
	var freqs []int64
	for _, node := range ns {
		if !node.Loaded {
			continue
		}
		freqs = append(freqs, node.FreqCount.Load())
	}
	if len(freqs) <= k {
		return
	}
	slices.Sort(freqs)

	seg := freqs[len(freqs)-k]

	for _, node := range ns {
		if node.FreqCount.Load() >= seg || !node.Loaded {
			continue
		}
		node.Unload()
	}
}

func FIFO(ns NS, k int) {
	var times []int64
	for _, node := range ns {
		if !node.Loaded {
			continue
		}
		times = append(times, node.LoadedTime)
	}
	if len(times) <= k {
		return
	}
	slices.Sort(times)

	seg := times[len(times)-k]

	for _, node := range ns {
		if node.LoadedTime >= seg || !node.Loaded {
			continue
		}
		node.Unload()
	}
}
