import socket
server = socket.socket(socket.AF_INET,socket.SOCK_DGRAM)
print "udp server start"
while True:
    server.sendto("Hello",("10.110.19.215",17800))
    server.close()
