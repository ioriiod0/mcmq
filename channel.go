package mcmq

import (
    "time"
    "log"
    "errors"
)


type PullReq struct {
    ErrCh chan error
    MsgCh chan *Msg 
}

type PushReq struct {
    ErrCh chan error
    Body []byte
}

type CommitReq struct {
    ID uint64
}

type Channel struct {
    queue Queue
    id uint64

    Name string
    PullCh chan PullReq
    PushCh chan PushReq
    CommitCh chan CommitReq
    Quit chan bool
}


func NewChannel(name string,pushQueueSize int,q Queue) *Channel {
    return &Channel{
        Name:name,
        queue:Queue,
        PullCh:make(chan PullReq),
        PushCh:make(chan PushReq,pushQueueSize), //currently,we only allow one consumer one time.
        CommitCh: make(chan CommitReq),
        Quit:make(chan bool),
    }
}

func (c * Channel) push(msg *Msg) (err error) {
    defer func() {
         if r := recover(); r != nil {
            switch x := r.(type) {
            case string:
                err = errors.New(x)
            case error:
                err = x
            default:
                log.Println("unkown panic:",r)
                err = errors.New("Unknown panic")
            }
        }
    }()

    if err != c.queue.Enque(msg);err != nil {
        return err
    } else {
        return nil
    }
}


func (c * Channel) pull() (msg *Msg,err error) {
    defer func() {
         if r := recover(); r != nil {
            switch x := r.(type) {
            case string:
                err = errors.New(x)
            case error:
                err = x
            default:
                log.Println("unkown panic:",r)
                err = errors.New("Unknown panic")
            }
            msg = nil
        }
    }()

    if v,err != c.queue.Front();err != nil {
        return nil,err
    } else {
        return v.(*Msg),nil
    }
}


func (c * Channel) commit() (err error) {
    defer func() {
         if r := recover(); r != nil {
            switch x := r.(type) {
            case string:
                err = errors.New(x)
            case error:
                err = x
            default:
                log.Println("unkown panic:",r)
                err = errors.New("Unknown panic")
            }
        }
    }()

    if _,err != c.queue.Deque();err != nil {
        return err
    } else {
        return nil
    }
    
}


func (c *Channel) Run() {

    defer func () {
        if err := c.queue.Save();err != nil {
            log.Println("save queue err:",err)
        } 
    }()

LOOP:
    for {
        select {
        case r := <-c.PushCh:
            msg := &Msg{
                Channel:c.Name,
                Timestramp:time.Now().Unix(),
                ID:c.id,
                Body:r.Body,
            }
            c.id ++
            r.ErrCh <- c.push(msg)
          
        case r := <-c.PullCh:
            if msg,err := c.pull();err != nil {
                r.ErrCh <- err
            } else if msg == nil {

                select {
                case r := <- c.PullCh:

                case r := <-Quit:

                case <- time.After()
                }

            } else {
                r.MsgCh <- msg
            }

        case r := <-c.CommitCh:
            r.ErrCh <- c.commit()

        case <-Quit:
            break LOOP
        }

    }


}