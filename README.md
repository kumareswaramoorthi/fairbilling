# Fair Billing

Requirements:
-----------------

1. go 1.12+

How to Run:
-----------------

1. build the application by following command

	go build -o fairbilling main.go

2. run the application by the following command (provide the path of input txt file like the command below) 

	./fairbilling /Users/Name/Desktop/FairBilling/input.txt


Sample Input:
-----------------
14:02:03 ALICE99 Start\
14:02:05 CHARLIE End\
14:02:34 ALICE99 End\
14:02:58 ALICE99 Start\
14:03:02 CHARLIE Start\
14:03:33 ALICE99 Start\
14:03:35 ALICE99 End\
14:03:37 CHARLIE End\
14:04:05 ALICE99 End\
14:04:23 ALICE99 End\
14:04:41 CHARLIE Start\



Output:
-----------------
    ALICE99 4 240 

	CHARLIE 3 37
