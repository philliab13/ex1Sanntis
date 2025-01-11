3. When I ran two threads in C, the shared the CPU and some of them might run two times and we get race conditions. When we increment and decrement at the same time we wont necceseraliy get the correct value.

What does GOMAXPROCS do? What happens if you set it to 1? It chooses how many OS-threads can run simoultanelesly. If you set it to one, only one process may run at the same time, so you dont get race conditions. The same thing happens in go when I allow them to run simoultanelesly, but not when I run one thread at a time.

4.  In c, as far as I understand, Mutex are used when you have concrete variables or such that more than one thread need to read and write to. While Semaphores are used when you need more flexibility and some threads are working on another thread and you cant simply allow one thread at a time. Therefore I use mutex since we have a single variable, i, which one thread should modify at a time.

Made interleaving possible by locking and unlocking every time they run the for-loop.
