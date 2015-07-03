#add delay to interface

$ tc qdisc add dev $DEV root netem delay 100ms
change delay
$ tc qdisc change dev $DEV root netem delay 10ms

#rate limiting note only egress traffic
$ tc qdisc add dev $DEV root handle 1: tbf rate 256kbit buffer 1600 limit 3000

#change egress rate
tc qdisc change dev $DEV root handle 1: tbf rate 512kbit buffer 1600 limit 3000


#see current interface tc status
tc -s qdisc ls dev $DEV

#tc delete root qdisc
tc qdisc del dev $DEV root

