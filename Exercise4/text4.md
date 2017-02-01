#Exercise 4
* We will probably UDP because...
* We have two different network modules, or at least thats our plan. The fault_handler module changes master if the master falls out and one of the slaves has to step in. The network module will in our implementation not handle this situation. 
* IP-address?
* Non-blocking sockets and select() would be our prefered option because it makes it possible to have several clients (slaves) active in a thread the time. 