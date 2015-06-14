#!/bin/bash

sed -i "s/zhangsan/lisi/g" `grep zhangsan -rl /modules`
