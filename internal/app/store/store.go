package store

type Store interface {
	Users() UsersRepository
	News() NewsRepository
	NewsUsers() NewsUsersRepository
}
