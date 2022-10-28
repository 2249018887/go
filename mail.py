#!/usr/bin/python3
import smtplib
from smtplib import SMTP_SSL
from email.mime.text import MIMEText
from email.header import Header
#打开一个文件只读

mail_host = "smtp.exmail.qq.com"                    #邮箱服务
mail_user = "2211472833qq.com"             #发件人
mail_pass = "uiowxtmbbjeiecja"                               #发件人邮箱服务密码
receivers = "milkdog19920214@163.com"               #收件人
message = "test"           #调用msg变量为邮件内容发邮件
subject = u'农险新注册用户'                                       #标题

try:
        smtpObj = SMTP_SSL(mail_host)
        smtpObj.login(mail_user,mail_pass)
        smtpObj.sendmail(mail_user,  receivers, message)
        print ("邮件发送成功")
        smtpObj.quit()
except smtplib.SMTPException as e:
        print ("Error: 无法发送邮件%s" % e)
