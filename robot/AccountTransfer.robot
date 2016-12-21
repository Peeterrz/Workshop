*** Settings ***
Library    Selenium2Library
Suite setup    Open browser    about:blank    chrome

*** Test cases ***

TC_AccountTransferForm_1
	โอนเงิน    1234567894    นายทดสอบ ระบบ    2000000001    นายสมชาย เข็มกลัด    0    50.00    0    TC_AccountTransferForm_1

TC_AccountTransferForm_2
	โอนเงิน    1234567895    นายทดสอบ ระบบ    2000000002    นายสมชาย เข็มกลัด    0    49.99    0    TC_AccountTransferForm_2

TC_AccountTransferForm_3
	โอนเงิน    1234567896    นายทดสอบ ระบบ    2000000003    นายสมชาย เข็มกลัด    0    100,000.00    0    TC_AccountTransferForm_3

TC_AccountTransferForm_4
	โอนเงิน    1234567897    นายทดสอบ ระบบ    2000000004    นายสมชาย เข็มกลัด    0    100,000.01    0    TC_AccountTransferForm_4



*** Keywords ***
โอนเงิน
	[arguments]    ${from_acc_no}    ${from_acc_name}    ${to_acc_no}    ${to_acc_name}    ${to_bank_code}    ${tamt}    ${fee_amt}    ${testcase_id}
	#Open browser    about:blank    chrome
	Go to    http://localhost:8080/demobank/${from_acc_no}
	Capture Page Screenshot    ${testcase_id}_AccountOverview.jpg
	Click Button    btn_ok
	Transfer Page    ${from_acc_no}    ${to_acc_no}    ${to_bank_code}    ${tamt}    ${testcase_id}
	Comfirmation Page       ${from_acc_no}    ${from_acc_name}    ${to_acc_no}    ${to_acc_name}    ${to_bank_code}    ${tamt}    ${fee_amt}    ${testcase_id}
	Successful Page      ${from_acc_no}    ${from_acc_name}    ${to_acc_no}    ${to_acc_name}    ${to_bank_code}    ${tamt}    ${fee_amt}    ${testcase_id}

Transfer Page
	[Arguments]    ${from_acc_no}    ${to_acc_no}    ${to_bank_code}    ${tamt}    ${testcase_id}
	Element Text Should Be    from_acc_no    ${from_acc_no}
	Input Text    to_acc_no    ${to_acc_no}
	Select From List By Index    bank_code    ${to_bank_code}
	Input Text    tamt    ${tamt}
	Capture Page Screenshot    ${testcase_id}_Trf.jpg
	Click Button    btn_ok

Confirmation Page
	[Arguments]    ${from_acc_no}    ${from_acc_name}    ${to_acc_no}    ${to_acc_name}    ${to_bank_code}    ${tamt}    ${fee_amt}    ${testcase_id}
  Element Text Should Be    from_acc_no    ${from_acc_no}
	Element Text Should Be	  from_acc_name    ${from_acc_name}
	Element Text Should Be    to_acc_no    ${to_acc_no}
	Element Text Should Be    to_acc_name    ${to_acc_name}
	Element Text Should Be    bank_code    ${to_bank_code}
	Element Text Should Be    tamt    ${tamt}
	Element Text Should Be    fee_amt    ${fee_amt}
	Capture Page Screenshot    ${testcase_id}_Comfirm.jpg
	Click Button    btn_ok

Successful Page
	[Arguments]   ${from_acc_no}    ${from_acc_name}    ${to_acc_no}    ${to_acc_name}    ${to_bank_code}    ${tamt}    ${fee_amt}    ${testcase_id}
	Element Text Should Be    from_acc_no    ${from_acc_no}
	Element Text Should Be    to_acc_no    ${to_acc_no}
	Element Text Should Be    to_acc_name    ${to_acc_name}
	Element Text Should Be    bank_code    ${to_bank_code}
	Element Text Should Be    tamt    ${tamt}
	Element Text Should Be    fee_amt    ${fee_amt}
	Capture Page Screenshot    ${testcase_id}_Success.jpg
