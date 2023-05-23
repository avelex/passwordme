package store

type PasswordStore struct {
}

func (s *PasswordStore) Create() {

}

func (s *PasswordStore) Import() {

}

func (s *PasswordStore) Save() error {
	return nil
}

func (s *PasswordStore) List() []string {
	return []string{}
}

func (s *PasswordStore) Delete() error {
	return nil
}
