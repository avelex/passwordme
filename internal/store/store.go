package store

import (
	"errors"

	"github.com/avelex/passwordme/internal/store/passfile"
)

var (
	ErrNotExists = errors.New("password file not exists")
)

type PasswordStore struct {
	rootPath  string
	passfiles []*passfile.PasswordFile
}

func NewPasswordStore(path string) *PasswordStore {
	return &PasswordStore{
		rootPath:  path,
		passfiles: detectPassfiles(path),
	}
}

func (s *PasswordStore) Create(masterPassword string) error {
	file, err := passfile.Create(s.rootPath, masterPassword)
	if err != nil {
		return err
	}

	s.passfiles = append(s.passfiles, file)

	return nil
}

func (s *PasswordStore) Open(masterPassword string) error {
	if len(s.passfiles) == 0 {
		return ErrNotExists
	}

	return s.passfiles[0].Open(masterPassword)
}

func (s *PasswordStore) Opened() bool {
	if len(s.passfiles) == 0 {
		return false
	}

	return s.passfiles[0].Opened()
}

func (s *PasswordStore) Import() {

}

// now it take first password file and append password into it
func (s *PasswordStore) Save(name, password string) error {
	if len(s.passfiles) == 0 {
		return ErrNotExists
	}

	file := s.passfiles[0]

	if err := file.AppendPassword(name, password); err != nil {
		return err
	}

	return nil
}

func (s *PasswordStore) List() ([]string, error) {
	if len(s.passfiles) == 0 {
		return nil, ErrNotExists
	}

	file := s.passfiles[0]

	passwords, err := file.List()
	if err != nil {
		return nil, err
	}

	return passwords, nil
}

func (s *PasswordStore) ShowPassword(name string) (string, error) {
	if len(s.passfiles) == 0 {
		return "", ErrNotExists
	}

	file := s.passfiles[0]

	return file.ShowPassword(name)
}

func (s *PasswordStore) Delete() error {
	return nil
}

func (s *PasswordStore) Exists() bool {
	return len(s.passfiles) > 0
}
