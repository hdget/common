package biz

import "sync"

type transactor interface {
	Ref(tx any)  // 引用, 因为可能嵌套NewTransactor, 这里加入计数
	Unref() bool // 解引用, 如果计数为0，表示需要Finalize，并Destroy tx
	Get() any    // 获取Tx
	Destroy()    // 销毁tx
}

type safeTxImpl struct {
	mu    sync.RWMutex
	tx    any
	count int
}

func newTransactor() transactor {
	return &safeTxImpl{
		mu: sync.RWMutex{},
	}
}

func (t *safeTxImpl) Ref(tx any) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.tx == nil {
		t.tx = tx
	}
	t.count++
}

func (t *safeTxImpl) Unref() bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.count > 0 {
		t.count--
	}

	// 为0表示需要Commit Or Rollback
	return t.count <= 0
}

func (t *safeTxImpl) Get() any {
	t.mu.RLock() // 使用读锁，允许多个 Get 调用并发执行
	defer t.mu.RUnlock()
	return t.tx
}

func (t *safeTxImpl) Destroy() {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.count == 0 { // 检查计数是否为 0
		t.tx = nil
	}
}
