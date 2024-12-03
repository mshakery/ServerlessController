package main

import (
	"context"
	"fmt"
	"github.com/mshakery/ServerlessController/etcdMiddleware"
	"github.com/mshakery/ServerlessController/protos"
)

func main() {
	users := []protos.User{
		{Token: "xaxaxaxa", Uid: "0", UserName: "masoud"},
		{Token: "ihihihih", Uid: "1", UserName: "roy"},
		{Token: "lolololo", Uid: "2", UserName: "arshia"},
	}

	cli, err := etcdMiddleware.ConnectToEtcd()
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		ku := fmt.Sprintf("/cluster/resources/user/%s", user.Token)
		err = etcdMiddleware.WriteToEtcdFromPb(cli, context.Background(), ku, &user)
		if err != nil {
			fmt.Printf("ik kan schrijf het %s kunde niet! wat verveland! want: %s", user.Uid, err.Error())
			return
		}
		fmt.Printf("ik schrijf het %s kunde op de database. Wat leuk! ", user.UserName)
	}

}
