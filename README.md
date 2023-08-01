## Build Instructions
* sudo docker build . -t setu-splitwise:latest
* sudo docker run -p 8080:8080 setu-splitwise:latest

## Tables
- User Table - contain user info
***schema***
>&ensp; Id int
>&ensp; Name string

- Group Table - contain group info
***schema***
>&ensp; Id int
>&ensp; Name String
>&ensp; Member int

- Transaction Table - contain transaction info
***schema***
>&ensp; Id int
>&ensp; Groupid string
>&ensp; Desc string
>&ensp; Date Time
>&ensp; Creator int
>&ensp; Totalamount int
>&ensp; Owee int
>&ensp; Pendingamount int

## APIs
***Create User POST***
*curl -X POST http://localhost:8080/user --data '{"id": 1, "name": "Clark"}'
*curl -X POST http://localhost:8080/user --data '{"id": 2, "name": "Dave"}'
*curl -X POST http://localhost:8080/user --data '{"id": 3, "name": "Ava"}'

***Get Users GET***
*curl http://localhost:8080/users

***Create Group POST***
*curl -X POST http://localhost:8080/group --data '[{"id": 1, "name": "Friends", "member": 1}]'

***Create Transaction POST***
*curl -X POST http://localhost:8080/transaction --data '[{"id": 1, "groupid": 1, "description": "Breakfast", "date": "2022-02-23T00:00:00Z", "creator": 1, "totalamount": 150, "owee": 2, "pendingamount": 50}, {"id": 1, "groupid": 1, "description": "Breakfast", "date": "2022-02-23T00:00:00Z", "creator": 1, "totalamount": 150, "owee": 3, "pendingamount": 50}]'
*curl -X POST http://localhost:8080/transaction --data '[{"id": 2, "groupid": 1, "description": "Breakfast", "date": "2022-02-23T00:00:00Z", "creator": 2, "totalamount": 100, "owee": 1, "pendingamount": 20}, {"id": 2, "groupid": 1, "description": "Breakfast", "date": "2022-02-23T00:00:00Z", "creator": 2, "totalamount": 100, "owee": 3, "pendingamount": 50}]'

***Get Transactions GET***
*curl http://localhost:8080/transactions?userid=1
*curl http://localhost:8080/transactions?userid=2
*curl http://localhost:8080/transactions?userid=3
