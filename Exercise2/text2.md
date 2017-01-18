#Mutex and Channels basics
* An atomic operation in concurrent programming are program operations that run completely independently of other processes. 
(An atomic operation is an operation where the processor can both read and write at the same clock cycle. If an operation acts on shared memory and completes in a single step relative to other threads, we call the operation atomic.) 
* A semaphore is variable that is used to control which process that should have access to a (common) shared resource. For example: the variable i in Ex.1. 
* Mutex is short for mutual exclusion, and is a property of concurrent control. The purpose is to prevent race conditions (Ex1). 
* A critical section is that part of the program when shared memory is accessed. 
