#3: Reasons for concurrency and parallelism
* Concurrency is the ability to decide the priority of several tasks which have different start and ending times. Parallelism is the ability to execute several diffent tasks at the same time. The main difference is how many tasks that are running at the same time. Parallelism = several tasks, concurrency: only one task are executed at the time (1).
* Machines have become increasingly multicore in the past decade because of Moorse law. When transistors become smaller and smaller, it makes it possible to add more CPU´s at the same space. This also allow higher performance at lower energy and multiple CPU cores allows the device to work at a much higher clock rate.
* Concurrency help solving problmes that occure when several tasks ask for the same recource. 
* Concurrency do solve some problems, but usually when solving one problem, several new problems occour.
* A process and threads are OS-managed. Green threads and coroutines one the other hand are not managed by the OS. Processes are more or less truly concurrent, as well as threads, but green threads and coroutines are not considered as concurrent. Both threads and processes exist within their own address space, but with treads multi-tasking is pre-emptive. Coroutines are like fibers, but not OS-managed. Almost like the "green" version of fibers (2).
* pthread_create() create a thread. threading.Thread() create a thread. go create coroutine. 
* Is the GIL almost like a RUNTIME for python? 
* Can use multiprocessing module. Is this the same as using several GIL's? 
* GOMAXPROCS sets the maximum number of CPUs that can be excecuting simultaneously and returns the previous setting (3). This funtion change the current setting if n > 1. 

##References:
1. https://www.youtube.com/watch?v=VLq9DfL4g8w
2. http://stackoverflow.com/questions/3324643/processes-threads-green-threads-protothreads-fibers-coroutines-whats-the
3. https://golang.org/src/runtime/debug.go