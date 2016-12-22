*** Keywords ***
โอนเงิน
	[arguments]    ${from_acc_no}    ${from_acc_name}    ${to_acc_no}    ${to_acc_name}    ${to_bank_code}    ${to_bank_name}    ${tamt}    ${fee_amt}    ${testcase_id}
	Go to    ${URL}${from_acc_no} 
	Capture Page Screenshot    ${testcase_id}_AccountOverview.jpg
	Click Button    btn_ok
	Transfer Page    ${from_acc_no}    ${to_acc_no}    ${to_bank_code}    ${tamt}    ${testcase_id}
	Confirmation Page       ${from_acc_no}    ${from_acc_name}    ${to_acc_no}    ${to_acc_name}    ${to_bank_name}    ${tamt}    ${fee_amt}    ${testcase_id}
	Successful Page      ${from_acc_no}    ${from_acc_name}    ${to_acc_no}    ${to_acc_name}    ${to_bank_name}    ${tamt}    ${fee_amt}    ${testcase_id}

Transfer Page
	[Arguments]    ${from_acc_no}    ${to_acc_no}    ${to_bank_code}    ${tamt}    ${testcase_id}
	Element Text Should Be    from_acc_no    ${from_acc_no}
	Input Text    to_acc_no    ${to_acc_no}
	Select From List By Value    bank_code    ${to_bank_code}
	Input Text    tamt    ${tamt}
	Capture Page Screenshot    ${testcase_id}_Trf.jpg
	Click Button    btn_ok

Confirmation Page
	[Arguments]    ${from_acc_no}    ${from_acc_name}    ${to_acc_no}    ${to_acc_name}    ${to_bank_name}    ${tamt}    ${fee_amt}    ${testcase_id}
    Element Text Should Be    from_acc_no    ${from_acc_no}
	Element Text Should Be	  from_acc_name    ${from_acc_name}
	Element Text Should Be    to_acc_no    ${to_acc_no}
	Element Text Should Be    to_acc_name    ${to_acc_name}
	Element Text Should Be    bank_name    ${to_bank_name}
	Element Text Should Be    tamt    ${tamt} บาท
	Element Text Should Be    fee_amt    ${fee_amt} บาท
	Capture Page Screenshot    ${testcase_id}_Comfirm.jpg
	Click Button    btn_ok

Successful Page
	[Arguments]   ${from_acc_no}    ${from_acc_name}    ${to_acc_no}    ${to_acc_name}    ${to_bank_name}    ${tamt}    ${fee_amt}    ${testcase_id}
	Element Text Should Be    from_acc_no    ${from_acc_no}
	Element Text Should Be    to_acc_no    ${to_acc_no}
	Element Text Should Be    to_acc_name    ${to_acc_name}
	Element Text Should Be    bank_name    ${to_bank_name}
	Element Text Should Be    tamt    ${tamt} บาท
	Element Text Should Be    fee_amt    ${fee_amt} บาท
	Capture Page Screenshot    ${testcase_id}_Success.jpg
