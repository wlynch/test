package main

import (
	"context"
	"fmt"
	"log"
	"os"

	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/cloudidentity/v1"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	user := os.Args[1]
	group := os.Args[2]

	client, err := cloudidentity.NewService(ctx, option.WithScopes(cloudidentity.CloudIdentityGroupsReadonlyScope))
	if err != nil {
		log.Fatal(err)
	}

	g, err := client.Groups.Lookup().GroupKeyId(group).Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(g.Name)

	membership, err := client.Groups.Memberships.CheckTransitiveMembership(g.Name).Query(fmt.Sprintf("member_key_id == '%s'", user)).Do()
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(membership.HasMembership)
	}

	admin, err := admin.NewService(ctx, option.WithScopes(admin.AdminDirectoryGroupMemberReadonlyScope))
	if err != nil {
		log.Fatal(err)
	}

	member2, err := admin.Members.HasMember(group, user).Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(member2.IsMember)

}
