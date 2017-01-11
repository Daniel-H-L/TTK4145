from threading import Thread

i = 0;

def thread1func():
	global i
	for j in range(1000000):
		i += 1
	return 0

def thread2func():
	global i
	for k in range(1000000):
		i -= 1
	return 0

def main():
	thread1 = Thread(target = thread1func, args = (),)
	thread2 = Thread(target = thread2func, args = (),)

	thread1.start()
	thread2.start()

	thread1.join()
	thread2.join()

	print(i)


main()


