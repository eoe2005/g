package gworker

type Worker interface {
	Init() error
	Execute() error
}
