package handlers

import (
    "fmt"
	"strconv"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"splitwise/splitwisedb"
	"splitwise/schemas"
)

var dbConnection *splitwisedb.SplitwiseDB

func init() {
	dbConnection = splitwisedb.CreateDB()
}

func GetTransactions (c *gin.Context){
	var gid, uid int
	sgid, gok := c.GetQuery("groupid")
	if !gok {
		gid = -1
	} else {
		gid, _ = strconv.Atoi(sgid);
	}
	suid, uok := c.GetQuery("userid")
	if !uok {
		panic("Need user Id")
	} else {
		uid, _ = strconv.Atoi(suid); 
	}

	var txns = dbConnection.SelectTransctions(gid)
	result := make(map[int]schemas.Result)
	for _, txn := range *txns {
		println(txn.Tid, txn.Groupid, txn.Desc, txn.Creator, txn.Totalamount, txn.Owee, txn.Pendingamount)
		groupname := dbConnection.SelectGroups(txn.Groupid)
		var res schemas.Result;
		if txn.Creator == uid {
			t, ok := result[txn.Tid]
			if ok {
				println("at uid old")
				res = schemas.Result{
					Date: calculateTime(txn.Date),
					Group: groupname,
					Desc: txn.Desc,
					Totalamount: txn.Totalamount,
					Pendingamount: txn.Pendingamount + t.Pendingamount,
				}
				result[txn.Tid] = res
			} else {
				println("At uid new")
				res = schemas.Result{
					Date: calculateTime(txn.Date),
					Group: groupname,
					Desc: txn.Desc,
					Totalamount: txn.Totalamount,
					Pendingamount: txn.Pendingamount,
				}
				result[txn.Tid] = res;
			}
		} else if txn.Owee == uid {
			println("at oweee")
			res = schemas.Result{
				Date: calculateTime(txn.Date),
				Group: groupname,
				Desc: txn.Desc,
				Totalamount: txn.Totalamount,
				Pendingamount: -txn.Pendingamount,
			}
			result[txn.Tid] = res;
		}
	} 
	a, _ := json.Marshal(result)
	fmt.Println(string(a))

	c.JSON(200, gin.H{
		"message": string(a),
	})
}

func GetUsers (c *gin.Context){
	users := dbConnection.SelectUsers()
	a, _ := json.Marshal(users)

	c.JSON(200, gin.H{
		"message": string(a),
	})
}

func CreateUser (c *gin.Context){
	var user schemas.User
	if err := c.BindJSON(&user); err != nil {
        panic(err)
    }

	fmt.Printf("Got user - %d %s\n", user.Id, user.Name)
	dbConnection.InsertUser(user);
	c.JSON(200, gin.H{
		"message": "Row Inserted",
	})
}

func CreateGroup (c *gin.Context){
	var groups []schemas.Group
	if err := c.BindJSON(&groups); err != nil {
        panic(err)
    }
	dbConnection.InsertGroup(groups);
	c.JSON(200, gin.H{
		"message": "Row Inserted",
	})
}

func CreateTransaction (c *gin.Context){
	var txns []schemas.Transactions
	if err := c.BindJSON(&txns); err != nil {
        panic(err)
    }

    dbConnection.InsertTransactions(txns);
	c.JSON(200, gin.H{
		"message": "Row Inserted",
	})
}

func calculateTime (date time.Time) string {
	y, m, d := date.Date()
	return strconv.Itoa(d) + "-" + m.String() + "-" + strconv.Itoa(y)   
}