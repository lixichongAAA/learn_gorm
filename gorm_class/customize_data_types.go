package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type CInfo struct {
	Name string
	Age  int
}

func (c CInfo) Value() (driver.Value, error) {
	str, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return string(str), nil
}

func (c *CInfo) Scan(value interface{}) error {
	str, ok := value.([]byte)
	if !ok {
		return errors.New("类型不匹配!")
	}
	json.Unmarshal(str, c)
	return nil
}

type Args []string

func (a Args) Value() (driver.Value, error) {
	if len(a) > 0 {
		var str string = a[0]
		for _, v := range a[1:] {
			str += "," + v
		}
		return str, nil
	} else {
		return "", nil
	}
}

func (a *Args) Scan(value interface{}) error {
	str, ok := value.([]byte)
	if !ok {
		return errors.New("类型无法解析")
	}
	*a = strings.Split(string(str), ",")
	return nil
}

type Cuser struct {
	ID   int
	Info CInfo
	Args Args
}

func Custome() {
	GLOBLE_DB.AutoMigrate(&Cuser{})
	GLOBLE_DB.Create(&Cuser{Info: CInfo{Name: "Lxc", Age: 23}, Args: []string{"1", "2", "3", "4"}})

	var u Cuser
	GLOBLE_DB.Last(&u)
	fmt.Println(u)
}
