import socket
client =socket.socket(socket.AF_INET,socket.SOCK_DGRAM)
data,(addr, port) = server.recvfrom( 1024 )
print data,add,port
client.close()
