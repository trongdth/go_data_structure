
`1. Concurrency vs Parallelism`
- Parallelism is TRYING to do two things at the same time. Concurrency is arranging it so you can do two things at the same time

`2. Data race`
- A data race is when two or more goroutines attempt to read and write to the same resource at the same time. Race conditions can create bugs that appear totally random or can never surface as they corrupt data. Atomic functions and mutexes are a way to synchronize the access of shared resources between goroutines.

