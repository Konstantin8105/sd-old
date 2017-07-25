package model

type resultNolinearBuckling int

const (
	diverge resultNolinearBuckling = iota
	converge
)

// TODO: if you have nonlinear elements, then we can use
