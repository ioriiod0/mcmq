package mcmq

type Queue interface {
	Front() (interface{}, error)
	Enque(interface{}) error
	Deque() (interface{}, error)
	Save() error
	Load() error
}
