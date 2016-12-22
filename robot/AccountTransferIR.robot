*** Settings ***
Library    Selenium2Library
Suite setup    Open browser    about:blank    chrome
Test template    โอนเงิน
Suite teardown    Close browser
Resource    ./Config.robot
Resource    ./Transfer.robot

*** Test cases ***

TC_AccountTransferIR_1    1234567898    นาย ทดสอบ ระบบ    1000000001    นาง สมศรี มีชัย    0    ธนาคารสมเกียรติ จำกัด    50.00    0.00    TC_AccountTransferIR_1