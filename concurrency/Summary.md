# Summary

> Make it work, make it right, make it fast. from Kent Beck

- `goroutines`, the basic unit of concurrency in Go, which let us check more
than one website at the same time.
- `anonymous functions`, which we used to start each of the concurrent processes
that check websites.
- `channels`, to help organize and control the communication between the
different processes, allowing us to avoid a race condition bug.
- `the race detector` which helped us debug problems with concurrent code
