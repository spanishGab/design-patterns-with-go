package iterator

type UserCollection struct {
	users []*User
}

func (userCollection *UserCollection) createIterator() Iterator {
	return &UserIterator{
		users: userCollection.users,
	}
}
