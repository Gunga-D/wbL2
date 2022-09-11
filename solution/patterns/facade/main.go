package main

type Account struct {
	login    string
	password string
}

func NewAccount(login string, password string) Account {
	return Account{login: login, password: password}
}

type AccountValidator struct {
}

func NewAccountValidator() AccountValidator {
	return AccountValidator{}
}

func (av AccountValidator) IsValid(original Account, checked Account) bool {
	return original.login == checked.login && original.password == checked.password
}

type AccountsCache struct {
	//Bad idea to store accounts by login
	kernel map[string]Account
}

func NewAccountsCache() AccountsCache {
	return AccountsCache{kernel: make(map[string]Account)}
}

func (ac AccountsCache) Insert(element Account) {
	ac.kernel[element.login] = element
}

func (ac AccountsCache) Get(login string) Account {
	return ac.kernel[login]
}

type Auth struct {
	accounts  AccountsCache
	validator AccountValidator
}

func NewAuth() Auth {
	return Auth{
		accounts:  NewAccountsCache(),
		validator: NewAccountValidator(),
	}
}

func (a Auth) SignUp(newAccount Account) {
	a.accounts.Insert(newAccount)
}

func (a Auth) SignIn(possible Account) bool {
	original := a.accounts.Get(possible.login)
	return a.validator.IsValid(original, possible)
}
