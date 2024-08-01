package iterator

import "fmt"

func IteratorClient() {
	user1 := &User{
		name: "user1",
		age:  30,
	}
	user2 := &User{
		name: "user2",
		age:  20,
	}

	userCollection := &UserCollection{
		users: []*User{user1, user2},
	}
	iterator := userCollection.createIterator()

	for iterator.hasNext() {
		fmt.Printf("User is %+v\n", iterator.getNext())
	}
}
