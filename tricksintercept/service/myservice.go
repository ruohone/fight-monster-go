package service

import (
	"fight.monster.go/tricksintercept/entity"
	"fight.monster.go/tricksintercept/intercept"
	"fmt"
)

func MyTe() {
	c := intercept.Intercept{}
	r, err := c.Cache("testkey", hh, 1, 3)
	if err != nil {
		fmt.Println(err)
		return
	}
	u := r.(entity.User)
	fmt.Println(u.UserName)

	r, err = c.Cache("testkey", hh, 1, 3, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)

	r, err = c.Cache("testkey", hh, 1, 3, 4)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)

	r, err = c.Cache("testkey", hh, 1, 3, 4, 32)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)

}

func hh(age, userId int) (entity.User, error) {
	return entity.User{UserName: "ruo", UserId: userId, Age: age}, nil
}
