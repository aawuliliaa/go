package detail

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

func GetApp(app string) error {
	return errors.Wrap(sql.ErrNoRows, fmt.Sprintf("Get app:%s err,err is", app))
}
func CreateApp(app string) {
	fmt.Printf("Add app:%s\n", app)
	return
}
func AddApp(app string) error {
	err := GetApp(app)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			CreateApp(app)
			return nil
		}
		return fmt.Errorf("Get App:%s err,err is:%v ", app, err)

	}
	return fmt.Errorf("App:%s already exist ", app)
}
