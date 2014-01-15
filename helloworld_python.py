# Python 3.3.3 and 2.7.6
# python helloworld_python.py

from threading import Thread, Lock

i = 0
lock = Lock()


def adder():
# In Python you "import" a global variable, instead of "export"ing it when you declare it
	
    	global i
    
    
# In Python 2 this generates a list of integers (which takes time),
# while in Python 3 this is an iterable (which is much faster to generate).
# In python 2, an iterable is created with xrange()

    	for x in range(0, 1000000):
		lock.acquire()
        	i += 1
		lock.release()

        
def subtractor():
	
	global i
	for x in range(0,1000000):
		lock.acquire()
		i -= 1
		lock.release()



def main():
	    adder_thr = Thread(target = adder)
	    subtractor_thr = Thread(target = subtractor)
	    adder_thr.start()    
	    subtractor_thr.start()
	    
	    for x in range(0, 50):
		print(i)
	    adder_thr.join()
	    adder_thr.join()
	    
	    print("Done: " + str(i))


main()
