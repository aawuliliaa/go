package detail

import (
	"database/sql"
	"fmt"
	"strings"
)

func GetAppNoWrap(app string) error {
	return fmt.Errorf("Get app:%s err,err is:%v ", app, sql.ErrNoRows)
}
func CreateAppNoWrap(app string) {
	fmt.Printf("Add app:%s\n", app)
	return
}
func AddAppNoWrap(app string) error {
	err := GetAppNoWrap(app)
	if err != nil {
		if strings.Contains(err.Error(),"no rows") {
			CreateAppNoWrap(app)
			return nil
		}
		return fmt.Errorf("Get App:%s err,err is:%v ", app, err)

	}
	return fmt.Errorf("App:%s already exist ", app)
}
