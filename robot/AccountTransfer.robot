*** Settings ***
Library    Selenium2Library
Suite setup    Open browser    about:blank    chrome
Test template    โอนเงิน
Suite teardown    Close browser
Resource    ./Config.robot
Resource    ./Transfer.robot

*** Test cases ***

TC_AccountTransfer_1    1234567894    นาย ทดสอบ ระบบ    2000000001    นาย สมชาย เข็มกลัด    0    ธนาคารสมเกียรติ จำกัด    50.00    0.00    TC_AccountTransfer_1
#TC_AccountTransfer_2    1234567895    นาย ทดสอบ ระบบ    2000000002    นาย สมชาย เข็มกลัด   0    ธนาคารสมเกียรติ จำกัด    49.99    0.00    TC_AccountTransfer_2
TC_AccountTransfer_3    1234567896    นาย ทดสอบ ระบบ    2000000003    นาย สมชาย เข็มกลัด    0    ธนาคารสมเกียรติ จำกัด    100,000.00    0.00    TC_AccountTransfer_3
#TC_AccountTransfer_4    1234567897    นาย ทดสอบ ระบบ    2000000004    นาย สมชาย เข็มกลัด    0    ธนาคารสมเกียรติ จำกัด    100,000.01    0.00   TC_AccountTransfer_4