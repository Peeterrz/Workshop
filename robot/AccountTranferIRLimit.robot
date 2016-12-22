*** Settings ***
Library    Selenium2Library
#Library    String
Suite setup    Open browser    about:blank    chrome
Suite teardown    Close browser
Resource    ./Config.robot
Resource    ./Transfer.robot


*** Test cases ***
คิดค่าธรรมเนียมโอนเงินข้ามเขตเมื่อโอนเงินข้ามเขตตั้งแต่ 11ครั้งเป็นต้นไป
	โอนเงินข้ามเขต 9 ครั้ง
	โอนเงิน      1234567899    นาย ทดสอบ ระบบ    1000000002    นาง สมศรี มีชัย    0     ธนาคารสมเกียรติ จำกัด    50.00    10.00     TC_AccountTransferIR_3

*** Keywords ***
โอนเงินข้ามเขต 9 ครั้ง
	:FOR    ${INDEX}    IN RANGE    1    10
    \    โอนเงิน     1234567899    นาย ทดสอบ ระบบ    1000000002    นาง สมศรี มีชัย    0    ธนาคารสมเกียรติ จำกัด    50.00    0.00    TC_AccountTransferIR_2_${INDEX}
 