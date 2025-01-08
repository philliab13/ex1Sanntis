## Exercise 1 - Theory questions

### Concepts

What is the difference between _concurrency_ and _parallelism_?

> Concurrency is the process of two or more proceses can happen in overlapping time periods, while paralellism is that they happen in paralell. So paralellism is one kind of concurrency.

What is the difference between a _race condition_ and a _data race_?

> Data race is when multiple threads tries to write and read to the same location at the same time, this is different from race conditions which are that the timing of the program makes something go wrong. Race condition is when two threads update the data at the same time and the correct data doesnt appear. increment and decrement i, but it happens at the same time and we do not get i back.

_Very_ roughly - what does a _scheduler_ do, and how does it do it?

> The schedular selects a thread to run in the list with runnable threads. It can choose a new thread based on some performace measure or choose randomly.

### Engineering

Why would we use multiple threads? What kinds of problems do threads solve?

> We use multiple threads to run multiple processes within the same period of time.

Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads?

> Fibers are threads divided into smaller tasks, this is to run multiple tasks on one operating system thread. We want to use it instead of threads to avoid the problem of getting a lot of callbacks

Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?

> In short, I think making concurrent programs make it so that we can handle multiple things at the same time, it makes it more difficult, but it also makes a much more efficient system. If you were to code three elevators without things happening in the same time period, it would be a slow system.

What do you think is best - _shared variables_ or _message passing_?

> Message passing since we dont override the memory, cooperating vs using the same data, but not neccesarily getting the correct result.
