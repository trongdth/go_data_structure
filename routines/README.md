Two styles of concurrency exist: deterministic (well-defined ordering) and non-deterministic (locking/mutual exclusion but order undefined). Go’s goroutines and channels promote deterministic concurrency (e.g. channels with one sender, one receiver), which is easier to reason about.

`1. Concurrency vs Parallelism`
- Parallelism is TRYING to do two things at the same time. Concurrency is arranging it so you can do two things at the same time

`2. Race condition`
- A data race is when two or more goroutines attempt to read and write to the same resource at the same time. Race conditions can create bugs that appear totally random or can never surface as they corrupt data. Atomic functions and mutexes are a way to synchronize the access of shared resources between goroutines.

`3. Semaphore`
- A semaphore is a signaling mechanism, and a thread that is waiting on a semaphore can be signaled by another thread.

`4. Fan in - fan out`
- Fan in: multi channels write the same channel
- Fan out: multi funcs read the same channel until it's closed