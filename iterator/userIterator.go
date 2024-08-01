package iterator

type UserIterator struct {
	index int
	users []*User
}

func (userIterator *UserIterator) hasNext() bool {
	return userIterator.index < len(userIterator.users)
}

func (userIterator *UserIterator) getNext() *User {
	if userIterator.hasNext() {
		user := userIterator.users[userIterator.index]
		userIterator.index++
		return user
	}
	return nil
}
