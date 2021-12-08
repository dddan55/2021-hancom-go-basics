package main

import (
	"fmt"

	"github.com/jcs1910/sky_membership"
)

func main() {
	// Register a member
	new_member := sky_membership.Register("Kim", "Male", 30)
	fmt.Println("main ==> ", new_member)

	// Sponsor
	new_member.Sponsor(1000)

	currentAsset := new_member.CheckAsset()
	fmt.Println(currentAsset)

	// Withdraw asset
	err := new_member.WithdrawAsset(2000)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(new_member)

}
