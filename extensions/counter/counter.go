package counter

import (
	"fmt"
	"sync/atomic"

	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/counter", new(RootModule))
}

// RootModule is created once per test run.
// Shared state must live here, NOT in ModuleInstance.
type RootModule struct {
	val int64 // shared across all VUs, must use atomic operations only
}

// ModuleInstance is created per VU — should NOT contain mutable shared state.
type ModuleInstance struct {
	root *RootModule
	vu   modules.VU
}

func (m *RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &ModuleInstance{root: m, vu: vu}
}

func (mi *ModuleInstance) Exports() modules.Exports {
	return modules.Exports{
		Named: map[string]interface{}{
			"nextId": mi.NextId,
			"reset":  mi.Reset,
			"value":  mi.Value,
		},
	}
}

// NextId returns the next ID in sequence — thread-safe using atomic.
func (mi *ModuleInstance) NextId(pad int) string {
	id := atomic.AddInt64(&mi.root.val, 1)
	return fmt.Sprintf("%0*d", pad, id)
}

// Value returns the current value without incrementing.
func (mi *ModuleInstance) Value() int64 {
	return atomic.LoadInt64(&mi.root.val)
}

// Reset sets the counter back to 0 — should only be used in setup/teardown.
func (mi *ModuleInstance) Reset() {
	atomic.StoreInt64(&mi.root.val, 0)
}
