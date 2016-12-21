*** Settings ***
Library    Selenium2Library
Suite setup    Open browser    about:blank    chrome

*** Test cases ***

TC_AccountOverview_1
	ค้นหาเลขที่บัญชี    1234567890    1234567890    นายทดสอบ ระบบ    10,000.00    TC_AccountOverview_1

TC_AccountOverview_2
	ค้นหาเลขที่บัญชี    123456789    1234567890    นายทดสอบ ระบบ    10,000.00    TC_AccountOverview_2

*** Keywords ***
ค้นหาเลขที่บัญชี
	[arguments]	   ${acc_no_input}    ${acc_no}    ${acc_name}    ${acc_bal}    ${testcase_id}
	Go to    http://localhost:8080/demobank/${acc_no_input}
	#Capture Page Screenshot    ${testcase_id}.jpg
	Element Text Should Be	acc_no    ${acc_no}
	Element Text Should Be	acc_name   ${acc_name}
	Element Text Should Be	acc_bal    ${acc_bal}
	Capture Page Screenshot    ${testcase_id}.jpg