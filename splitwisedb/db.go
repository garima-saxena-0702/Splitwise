package splitwisedb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	
	"splitwise/schemas"
)

type SplitwiseDB struct {
	connection *sql.DB
}

func CreateDB() *SplitwiseDB {
	println("Starting db...")
	
	connStr := "user=setu dbname=splitwisedb password=password host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	//Create the required tables
	createUser := "CREATE TABLE USERS ("+
		"id INTEGER PRIMARY KEY,"+
		"name TEXT NOT NULL);"
	createGroups := "CREATE TABLE GROUPS ("+
		"groupid INTEGER NOT NULL,"+
		"groupname TEXT NOT NULL,"+
		"member INTEGER references USERS(id) NOT NULL);"
	createTransactions := "CREATE TABLE TRANSACTIONS ("+
		"tid INTEGER NOT NULL,"+
		"groupid INTEGER NOT NULL,"+
		"description TEXT NOT NULL,"+
		"date date NOT NULL,"+
		"creator INTEGER references USERS(id) NOT NULL,"+
		"totalamount INTEGER NOT NULL,"+
		"owee INTEGER references USERS(id) NOT NULL,"+
		"pendingamount INTEGER NOT NULL);"
	_, err = db.Query(createUser)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Query(createGroups)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Query(createTransactions)
	if err != nil {
		log.Fatal(err)
	}

	return &SplitwiseDB{connection: db}
}

func (db SplitwiseDB) Ping() {
	err := db.connection.Ping()
	if err == nil {
		println("pinged!")
	} else {
		log.Fatal(err)
	}
}

func (db SplitwiseDB) SelectUsers() *[]schemas.User {
	rows, err := db.connection.Query("SELECT * FROM users")
    if err != nil {
		println("Error!")
        log.Fatal(err)
    }
    defer rows.Close()
    
	var users []schemas.User 
	for rows.Next() {
		var user schemas.User
        err := rows.Scan(&user.Id, &user.Name)
        if err != nil {
            panic(err)
        }
		users = append(users, user)
	}

	return &users
}

func (db SplitwiseDB) InsertUser(user schemas.User) {
	sqlStr := fmt.Sprintf("INSERT INTO USERS VALUES (%d, '%s')", user.Id, user.Name)
	_, err := db.connection.Query(sqlStr)

	if err != nil {
		panic(err)
	}
}

func (db SplitwiseDB) InsertGroup(groups []schemas.Group) {
	sqlStr := "INSERT INTO GROUPS VALUES "
	for i, group := range groups {
		sqlStr += fmt.Sprintf("(%d, '%s', %d)",group.Groupid, group.Name, group.Member)
		if i < (len(groups) - 1) {
			sqlStr += ","
		}
	}
	println(sqlStr)
	_, err := db.connection.Query(sqlStr)

	if err != nil {
		panic(err)
	}
}

func (db SplitwiseDB) InsertTransactions(txns []schemas.Transactions) {
	sqlStr := "INSERT INTO TRANSACTIONS VALUES "
	for i, txn := range txns {
		sqlStr += fmt.Sprintf("(%d, %d, '%s', '%d%02d%02d %02d:%02d:%02d', %d, %d, %d, %d)", 
							txn.Tid, txn.Groupid, txn.Desc, 
							txn.Date.Year(), txn.Date.Month(), txn.Date.Day(), txn.Date.Hour(), 
						    txn.Date.Minute(), txn.Date.Second(), txn.Creator,
						    txn.Totalamount, txn.Owee, txn.Pendingamount)
		if i < (len(txns) - 1) {
			sqlStr += ","
		}
	}

	println(sqlStr)
	_, err := db.connection.Query(sqlStr)

	if err != nil {
		panic(err)
	}
}

func (db SplitwiseDB) SelectGroups(groupid int) string {
	sqlStr := fmt.Sprintf("SELECT * FROM Groups where groupid=%d", groupid)
	rows, err := db.connection.Query(sqlStr)
		if err != nil {
			println("Error!")
			log.Fatal(err)
		}
		defer rows.Close()
		
		for rows.Next() {
			var group schemas.Group
			err := rows.Scan(&group.Groupid, &group.Name, &group.Member)
			if err != nil {
				panic(err)
			}
			return group.Name
		}
		return ""
}


func (db SplitwiseDB) SelectTransctions(gid int) *[]schemas.Transactions {
	var sqlStr string
	if gid != -1 {
		sqlStr = fmt.Sprintf("SELECT * FROM TRANSACTIONS WHERE groupid=%d", gid)
	} else {
		sqlStr = "SELECT * FROM TRANSACTIONS"
	}

	rows, err := db.connection.Query(sqlStr)
    if err != nil {
		println("Error!")
        log.Fatal(err)
    }
    defer rows.Close()

	var txns []schemas.Transactions
	for rows.Next() {
		var txn schemas.Transactions
        err := rows.Scan(&txn.Tid, &txn.Groupid, &txn.Desc, &txn.Date, &txn.Creator, &txn.Totalamount, &txn.Owee, &txn.Pendingamount)
        if err != nil {
            panic(err)
        }
		txns = append(txns, txn)
	}

	return &txns
}
