package mutex

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

type RecursiveMutex struct {
	sync.Mutex
	owner     int64 // 持有锁的goroutineId
	recursion int32 // 重入次数
}

func (rm *RecursiveMutex) Lock() {
	gid := GoId()
	if atomic.LoadInt64(&rm.owner) == gid {
		rm.recursion++
		return
	}

	rm.Mutex.Lock()
	atomic.StoreInt64(&rm.owner, gid)
	rm.recursion = 1
}

func (rm *RecursiveMutex) UnLock() {
	gid := GoId()
	if atomic.LoadInt64(&rm.owner) != gid {
		panic(fmt.Sprintf("wrong the owner(%d): %d!", rm.owner, gid))
	}
	rm.recursion--
	if rm.recursion != 0 {
		return
	}
	atomic.StoreInt64(&rm.owner, -1)
	rm.Mutex.Unlock()
}

func GoId() int64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return int64(id)
}
