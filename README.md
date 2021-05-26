### Brain dump

#### How to

- An outline / a path I want to walk

#### Pointers - Value semantics

- scheduler - how does go schedule work?
- how go schedules work
- goroutines and stacks scheduled onto a M - onto a process
- only memory we can write to is the stack, we protect the integrity of our program; we lock it into the goroutine stack
- every function is a data transformation!
- & gives you the address of some variable
- value vs address - both are data
- it was passed by value - no modification was made - value semantics
- inefficiency! memory growth - copying of a value instead of a reference - especially for large structs
- direct memory access through stack frame access

#### Pointers - Pointer semantics

- memory efficiency gained through use of pointer semantics 
- escape to heap analysis 
- side effects - heap escape, data integrity without locking
- `sync.Atomic` / mutex locking/ channels - use an atomic primitive to guarantee atomicity. `Channels` 
  can give you atomicity because one operation is run on the data!
- why the pragma `go:noinline`?
- indirect memory access through pointer access
- modifying the heap, we're out of the stack at this point
- every function works on its copy of data - side effects around mutations! - but, there's a cost. 
  Good data locality with value first semantic. Golang is a value first language, unless it's
  inefficient Unless
    - we are copying lots of data - a large struct
    - stack owned data to heap manipulation - escape to heap so many functions can control it
- unwinding the call stack is self-cleaning - gc does not do this work

#### Pointers - Escape analysis

Where should this value be constructed? Stack, or heap - what is the lifetime of this object? 
Escape analysis identifies variables whose lifetimes will live beyond the lifetime of the function in
which it is declared, and moves the location where the variable is allocated from the stack to the heap. 
Technically we say that b escapes to the heap.

Obviously there is a cost; heap allocated variables have to be garbage collected when they are no longer 
reachable, stack allocated variables are automatically free’d when their function returns.

- pragmas `//go:noescape`
- static code analysis - escape analysis
- not everything can stay on the stack, the compiler inspects the objects lifetime and escapes the ref to the heap
  and is an allocation to the heap - it indicates cost - we're sharing the stack item, which means escaping to 
  the heap. cost incurred is garbage collection latency
- readability means you're not hiding cost!
- value construction - pointer return , make the cost clear - unless it's a return
- `go build -gcflags -m`
- we lose data locality with passing by reference

#### Pointers - Stack growth

- dynamic stack allocation with contiguous stacks; small cost to double the size of the stack
- re-framing the stack after growing the size
- does that mean go does not experience stack overflows?
- interesting stack manipulation demo

#### Pointers - GC

- env GOGC value - 100%
- concurrent mark sweep collector - to work as we reclaim memory
- pacer / pacing - memory pressure - how quickly we fill the heap up before we perform the operation
- an internal GC _latency_ < `100us` to reclaim memory
- scan for references then mark for deletion and create a **write** barrier to ensure integrity

### Reference talks

- [Understanding allocations - the stack and heap](https://www.youtube.com/watch?v=ZMZpH4yT7M0)
- [Escape analysis and memory profiling](https://www.youtube.com/watch?v=2557w0qsDV0)
- [Allocations in Golang](https://medium.com/eureka-engineering/understanding-allocations-in-go-stack-heap-memory-9a2631b5035d)

- increased font size up by 7 
