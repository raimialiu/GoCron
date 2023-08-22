package lib

import "sync"

type Promise struct {
	wg sync.WaitGroup
}

func NewPromise() Promise {
	wg := sync.WaitGroup{}
	return Promise{
		wg: wg,
	}
}

func (p *Promise) WhenAll() {

}
