package painter

import (
	"image"
	"sync"

	"golang.org/x/exp/shiny/screen"
)

// Receiver отримує текстуру, яка була підготовлена в результаті виконання команд у циелі подій.
type Receiver interface {
	Update(t screen.Texture)
}

// Loop реалізує цикл подій для формування текстури отриманої через виконання операцій отриманих з внутрішньої черги.
type Loop struct {
	Receiver Receiver

	next screen.Texture // текстура, яка зараз формується
	prev screen.Texture // текстура, яка була відправленя останнього разу у Receiver

	Mq messageQueue
	Done chan struct{}
	Stopped bool
}

var size = image.Pt(800, 800)

// Start запускає цикл подій. Цей метод потрібно запустити до того, як викликати на ньому будь-які інші методи.
func (l *Loop) Start(s screen.Screen) {
	l.next, _ = s.NewTexture(size)
	l.prev, _ = s.NewTexture(size)
	l.Mq = messageQueue{}
	go l.eventProcessing()
}

func (l *Loop) eventProcessing() {
	for !l.Stopped || !l.Mq.isEmpty() {
		op := l.Mq.Pull() 
		update := op.Do(l.next)
			if update {
				l.Receiver.Update(l.next)
				l.next, l.prev = l.prev, l.next
			}
		}
		close(l.Done)
	}

// Post додає нову операцію у внутрішню чергу.
func (l *Loop) Post(op Operation) {
	if op != nil {
		l.Mq.Push(op)
	}
}

// StopAndWait сигналізує
func (l *Loop) StopAndWait() {
	l.Post(OperationFunc(func(t screen.Texture) {
		l.Stopped = true
	}))
	l.Stopped = true
	<-l.Done
}

type messageQueue struct {
	Queue []Operation
	mu         sync.Mutex
	blocked    chan struct{}
}

func (Mq *messageQueue) Push(op Operation) {
	Mq.mu.Lock()
	defer Mq.mu.Unlock()

	Mq.Queue = append(Mq.Queue, op)

	if Mq.blocked != nil {
		close(Mq.blocked)
		Mq.blocked = nil
	}
}

func (Mq *messageQueue) Pull() Operation {
	Mq.mu.Lock()
	defer Mq.mu.Unlock()

	for len(Mq.Queue) == 0 {
		Mq.blocked = make(chan struct{})
		Mq.mu.Unlock()
		<-Mq.blocked
		Mq.mu.Lock()
	}

	op := Mq.Queue[0]
	Mq.Queue[0] = nil
	Mq.Queue = Mq.Queue[1:]
	return op
}

func (Mq *messageQueue) isEmpty() bool {
	Mq.mu.Lock()
	defer Mq.mu.Unlock()

	return len(Mq.Queue) == 0
}