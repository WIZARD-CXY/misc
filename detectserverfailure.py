#!/usr/bin/env python
# -*- coding: utf-8 -*-

import smtplib
from email.mime.text import MIMEText
import os
import time

hostname = "10.15.199.199"

SMTPserver = 'smtp.126.com'
sender = 'wizardcxy@126.com'
destination = 'wizardcxy@126.com'
password = "2784a48"

failcount=0;

while True:
    response = os.system("ping -c 1 -w2 " + hostname + " > /dev/null 2>&1")

    if response == 0:
        if failcount > 0:
            failcount = 0
            message = 'deeply server is up'
            msg = MIMEText(message)

            msg['Subject'] = '199 Server is up'
            msg['From'] = sender
            msg['To'] = destination
            mailserver = smtplib.SMTP(SMTPserver, 25)
            mailserver.login(sender, password)
            mailserver.sendmail(sender, [sender], msg.as_string())
            mailserver.quit()
    else:
        message = 'deeply server is down'
        print "server is down"
        msg = MIMEText(message)

        msg['Subject'] = '199 Server is down'
        msg['From'] = sender
        msg['To'] = destination

        if failcount < 3 :
            failcount+=1
            mailserver = smtplib.SMTP(SMTPserver, 25)
            mailserver.login(sender, password)
            mailserver.sendmail(sender, [sender], msg.as_string())
            mailserver.quit()

    time.sleep(60)
