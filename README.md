<h2> Backend Shopee Test

<h3> How to Run </h3>
Make sure you have install docker and also docker compose then run the following command the the project root

> docker-compose up

<h3> API EndPoints </h3>
api testing is done by the help of RESTY (https://github.com/micha/resty)

 - Insert Exchange Currency [http://localhost:8080/api/v1/exchange](http://localhost:8080/api/v1/exchange) [post]
	 
	 

	> **Request** 
$ POST /exchange '{"from" : "IDR" , "to" : "USD"}'  
	>**Response**
	>{"id":"5a72f494-159c-440b-a352-d2de61ae1d18", "result":"success"}

- Insert Exchange Daily Rate [http://localhost:8080/api/v1/daily](http://localhost:8080/api/v1/daily) [post]

	> **Request**
$	POST /daily '{"date":"2019-07-05", "from":"IDR", "To":"USD", "rate":"0.000071"}'
	> **Response**
	>{"id":"3e9f89fc-32de-4521-b76c-2f5686776e91", "result":"success"}
- View Exchange Rate from Last 7 Days
	>**Request**
	>$ POST /last7 '{"from":"IDR", "to":"USD", "date":"2019-07-05"}'
	
	>**Response**
	>{"data":{"from":"IDR","to":"USD","average":0.000071,"variance":0,"data":[{"Date":"2019-07-05T00:00:00Z","Rate":0.000071}]},"result":"success"}
	
- View Tracked Exchange Rates [http://localhost:8080/api/v1/tracked](http://localhost:8080/api/v1/tracked) [post]
	> **Request**
	> $ POST /tracked '{"Date":"2019-07-05", "Exchanges":[{"From":"IDR","To":"USD"}]}'
	
	> **Response**
	> {"data":[{"from":"IDR","to":"USD","rate":"insufficient data","7-day-avg":""}],"result":"success"}

- Remove Exchange Rates [http://localhost:8080/api/v1/remove](http://localhost:8080/api/v1/remove) [post]
	> **Request**
	> $ POST /remove '{"Exchanges":[{"From":"IDR","To":"USD"}]}'
	
	> **Response**
	>  {"data":true,"result":"success"}
	
