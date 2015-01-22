import os
import time
hostname = "10.10.103.223" #example
response = os.system("ping -c 1 -w2 " + hostname + " > /dev/null 2>&1")

#and then check the response...
while True:
	if response == 0:
	          print hostname, 'is up!'
	else:
	          print hostname, 'is down!'
	time.sleep(10)
