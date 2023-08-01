package schemas

import (
    "time"
)

type CustomTime struct {
    time.Time
}

func (t *CustomTime) UnmarshalJSON(b []byte) (err error) {
    date, err := time.Parse(`"2006-01-02T15:04:05.000-0700"`, string(b))
    if err != nil {
        return err
    }
    t.Time = date
    return
}

type User struct {
    Id     int  `json:"id"`
    Name  string  `json:"name"`
}

type Group struct {
    Groupid     int  `json:"id"`
    Name  string  `json:"name"`
    Member int  `json:"member"`
}

type Transactions struct {
    Tid     int  `json:"id"`
    Groupid  int  `json:"groupid"`
    Desc string  `json:"description"`
	Date time.Time `json:"date"`
	Creator int `json:"creator"`
    Totalamount  int `json:"totalamount"`
    Owee int  `json:"owee"`
    Pendingamount int  `json:"pendingamount"`
}

type Result struct {
    Date string `json:"date"`
    Group string `json:"groupname"`
    Desc string `json:"desc"`
    Totalamount int `json:"totalamount"`
    Pendingamount int `json:"pendingamount"`
}

