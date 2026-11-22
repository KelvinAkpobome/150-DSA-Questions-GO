## Heap Algorithms

I benchmarked my generic Go heap/priority-queue at [priority-queue.go](./priority-queue.go) and found that repeated `Push()` and `Pop()` operations perform **zero heap allocations** when sufficient slice capacity is available. The benchmark recorded **22.77 ns/op, 0 B/op, and 0 allocs/op**. This shows that the implementation can handle repeated heap operations without additional memory allocations after initialization.


| Problem | Solution | Explanation |
|---|---|---|
| Connect Ropes With Minimum Cost | [Code](./min-ropes.go) | N/A |