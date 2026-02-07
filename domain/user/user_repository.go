package user

type UserRepository interface {
	FindByID(id string) (*User, error)
  FindByGoogleSub(sub string) (*User, error)
	Create(user *User) error
}

type IdentityRepository interface {
	FindByProviderSub(provider, sub string) (*Identity, error)
	Create(indentity *Identity) error
}
