package counter

import (
	"fmt"
	"sync/atomic"

	"go.k6.io/k6/js/modules"
)

// ลงทะเบียน module ชื่อ "k6/x/counter"
func init() {
	modules.Register("k6/x/counter", new(RootModule))
}

type RootModule struct{}
type Counter struct {
	val int64
}

func (*RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &Counter{}
}

func (c *Counter) Exports() modules.Exports {
	return modules.Exports{
		Named: map[string]interface{}{
			"nextId": c.NextId,
		},
	}
}

// atomic → thread-safe ข้าม VU ทุกตัว
func (c *Counter) NextId(pad int) string {
	id := atomic.AddInt64(&c.val, 1)
	return fmt.Sprintf("%0*d", pad, id)
}
