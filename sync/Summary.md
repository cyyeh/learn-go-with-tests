# Summary

- We've covered a few things from the sync package
    - `Mutex` allows us to add locks to our data
    - `Waitgroup` is a means of waiting for goroutines to finish jobs
- `go vet`
    - Remember to use go vet in your build scripts as it can alert you to some subtle bugs in your code before they hit your poor users.
- Dont use embedding because it's convenient
    - Think about the effect embedding has on your public API.
    - Do you really want to expose these methods and have people coupling their own code to them?
    - With respect to mutexes, this could be potentially disastrous in very unpredictable and weird ways, imagine some nefarious code unlocking a mutex when it shouldn't be; this would cause some very strange bugs that will be hard to track down.

A common Go newbie mistake is to over-use channels and goroutines just because it's possible, and/or because it's fun. Don't be afraid to use a sync.Mutex if that fits your problem best. Go is pragmatic in letting you use the tools that solve your problem best and not forcing you into one style of code.
Paraphrasing:
- Use channels when passing ownership of data 
- Use mutexes for managing state