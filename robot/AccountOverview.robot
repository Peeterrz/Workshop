*** Settings ***
Library    Selenium2Library
Suite setup    Open browser    about:blank    chrome
Test template    ค้นหาเลขที่บัญชี 
Suite teardown    Close browser
Resource    ./Config.robot

*** Test cases ***

TC_AccountOverview_1    1234567890    1234567890     นาย ทดสอบ ระบบ     10,000.00    TC_AccountOverview_1
TC_AccountOverview_2    12345678    0012345678    นาย ทดสอบ ระบบ    10,000.00    TC_AccountOverview_2
TC_AccountOverview_3    1234567891    1234567891    นาย ทดสอบ ระบบ    8,000.00    TC_AccountOverview_3
TC_AccountOverview_4    1234567892    1234567892    นาย ทดสอบ ระบบ    8,000.13    TC_AccountOverview_4
TC_AccountOverview_5    1234567893    1234567893    นาย ทดสอบ ระบบ    -1,500.00    TC_AccountOverview_5

*** Keywords ***
ค้นหาเลขที่บัญชี
	[arguments]	   ${acc_no_input}    ${acc_no}    ${acc_name}    ${acc_bal}    ${testcase_id}
	Go to    ${URL}${acc_no_input}
	Element Text Should Be    acc_no    ${acc_no}
	Element Text Should Be    acc_name   ${acc_name}
	Element Text Should Be    id=acc_bal    ${acc_bal} บาท
	Capture Page Screenshot    ${testcase_id}.jpg