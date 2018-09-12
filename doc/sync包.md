# Sync包
## Mutex
mutex为互斥锁，定义如下，其实现了Locker 接口
```
	type Mutex struct {
		state int32
		sema  uint32
	}
	type Locker interface {
		Lock()
		Unlock()
	}
```
一般程序中我们会在定义的结构中加入mutex，来实现互斥访问，例如下面通过mutex来实现安全的map
```
type MyMap struct {
	Data map[int]chan interface{}
	sync.RWMutex
}

func (s *MyMap) Set(key int, value chan interface{}) {
	s.Lock()
	s.Data[key] = value
	s.Unlock()
}

func (s *MyMap) Get(key int) (chan interface{}, bool) {
	s.RLock()
	defer s.RUnlock()
	value, ok := s.Data[key]
	return value, ok
}

func (s *MyMap) Delete(key int) {
	s.Lock()
	delete(s.Data, key)
	s.Unlock()
}

var MapChan MyMap
```
具体Lock和UnLock实现，后面进行分析
## RWMutex
rwmutex为读写互斥锁
## Cond
Cond是现实了一个条件变量，定义如下
```
type Cond struct {
	noCopy noCopy

	// L is held while observing or changing the condition
	L Locker

	notify  notifyList
	checker copyChecker
}

// NewCond returns a new Cond with Locker l.
func NewCond(l Locker) *Cond {
	return &Cond{L: l}
}
func (c *Cond) Broadcast()
func (c *Cond) Signal()
func (c *Cond) Wait()

```
主要方法为Broadcast ,Signal ,Wait
等待通知: wait 
阻塞当前线程，直到收到该条件变量发来的通知
Wait自行解锁c.L并阻塞当前线程，在之后线程恢复执行时，Wait方法会在返回前锁定c.L
单发通知: signal 
让该条件变量向至少一个正在等待它的通知的线程发送通知，表示共享数据的状态已经改变。
Signal唤醒等待c的一个线程（如果存在）。调用者在调用本方法时，建议（但并非必须）保持c.L的锁定。
广播通知: broadcast 
让条件变量给正在等待它的通知的所有线程都发送通知。
Broadcast唤醒所有等待c的线程。调用者在调用本方法时，建议（但并非必须）保持c.L的锁定。
例子可以参考测试用例 

```
func TestCondSignal(t *testing.T) {
	var m Mutex
	c := NewCond(&m)
	n := 2
	running := make(chan bool, n)
	awake := make(chan bool, n)
	for i := 0; i < n; i++ {
		go func() {
			m.Lock()
			running <- true
			c.Wait()
			awake <- true
			m.Unlock()
		}()
	}
	for i := 0; i < n; i++ {
		<-running // Wait for everyone to run.
	}
	for n > 0 {
		select {
		case <-awake:
			t.Fatal("goroutine not asleep")
		default:
		}
		m.Lock()
		c.Signal()
		m.Unlock()
		<-awake // Will deadlock if no goroutine wakes up
		select {
		case <-awake:
			t.Fatal("too many goroutines awake")
		default:
		}
		n--
	}
	c.Signal()
}

```

## WaitGroup
## Pool
## Once