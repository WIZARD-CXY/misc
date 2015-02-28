#!/usr/bin/env python 
# coding=utf-8 
 
import os
import re
 
ssh_config_file = '~/.ssh/config'
 
def parse_config():
    entry_list = []
    entry_name = host_name = user_name = ''
 
    conf = os.path.expanduser(ssh_config_file)
 
    if not os.path.exists(conf):
        print 'No such file exists: "%s"!' % conf
        return entry_list, 1
 
    fp = open(conf, 'r')
 
    for line in fp:
        line = line.strip()
 
        if not line or line.startswith('#'):
            continue
 
        if line.startswith('Host '):
            if entry_name:
                entry_address = '%s@%s' % (user_name, host_name)
                entry_list.append((entry_name, entry_address))
 
            entry_name = line.split()[1]
 
        if line.startswith('HostName '):
            host_name = line.split()[1]
 
        if line.startswith('User '):
            user_name = line.split()[1]
 
    if entry_name:
        entry_address = '%s@%s' % (user_name, host_name)
        entry_list.append((entry_name, entry_address))
 
    fp.close()
 
    return entry_list, 0
 
def ssh_helper():
    entry_id = 0
    entry_list, status_code = parse_config()
 
    if not entry_list and status_code == 1:
        return
 
 
    print '+-----+------------------------------+------------------------------------------+'
    print '| id  | name                         | address                                  |'
    print '+-----+------------------------------+------------------------------------------+'
 
    for entry_id, entry in enumerate(entry_list):
        entry_name, entry_addr = entry
 
        print '| %-3d | %-28s | %-40s |' % (entry_id+1, entry_name, entry_addr)
 
    print '+-----+------------------------------+------------------------------------------+'
    print '''
Tips: Press a number betwwen 1 and %d to select the host to connect, or "q" \
to quit.''' % (entry_id+1)
 
    select = ''
 
    while select != 'q':
        select = raw_input('\n# ')
 
        if select == 'q':
            break
 
        try:
            os.system('ssh %s' % entry_list[int(select)-1][0])
        except (ValueError, IndexError):
            print 'You must press a number between 0 and %d' % entry_id
 
if __name__ == '__main__':
    ssh_helper()
