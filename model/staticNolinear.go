package model

type resultNolinearBuckling int

const (
	diverge resultNolinearBuckling = iota
	converge
)
