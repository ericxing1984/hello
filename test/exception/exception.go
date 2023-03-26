package exception

import (
	"fmt"
	"runtime/debug"
)

type Address struct {
	Name        string `json:"name" gorm:"type:varchar(50);not null;"`
	Phone       string `json:"phone" gorm:"type:varchar(50);not null;default:'';"`
	Address     string `json:"address" gorm:"type:varchar(100);not null;default:'';"`
	SupplierID  int64  `json:"supplier_id,string" `
	Description string `json:"description" gorm:"type:varchar(100);not null;default:'';"`
}

func test() {
	j := 0
	i := 1 / j
	fmt.Println(i)

}
func TestException() {
	defer func() {
		if e := recover(); e != nil {
			s := string(debug.Stack())
			fmt.Printf("err=%v, stack=%s\n", e, s)
		}
	}()

	test()

}
