package passfile

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/avelex/passwordme/internal/crypto/aes"
)

const (
	EXTENSION     = ".me"
	_DEFAULT_PERM = 0600
)

var (
	ErrPermissionDenied = errors.New("permission denied")
)

type PasswordFile struct {
	opened     bool
	path       string
	masterHash []byte

	mutex     sync.RWMutex
	passwords map[string]string
}

func Create(dst, masterPassword string) (*PasswordFile, error) {
	path := filepath.Join(dst, time.Now().Format(time.RFC3339)+EXTENSION)
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	if err := f.Close(); err != nil {
		return nil, err
	}

	h := sha256.Sum256([]byte(masterPassword))

	return &PasswordFile{
		opened:     true,
		path:       path,
		masterHash: h[:],
		mutex:      sync.RWMutex{},
		passwords:  make(map[string]string),
	}, nil
}

func New(path string) *PasswordFile {
	return &PasswordFile{
		opened:     false,
		path:       path,
		masterHash: make([]byte, 0),
		mutex:      sync.RWMutex{},
		passwords:  make(map[string]string),
	}
}

func (f *PasswordFile) Open(masterPassword string) error {
	if f.opened {
		return nil
	}

	h := sha256.Sum256([]byte(masterPassword))
	f.masterHash = h[:]

	decrypted, err := f.ReadFull()
	if err != nil {
		f.masterHash = make([]byte, 0)
		return err
	}

	split := strings.Split(string(decrypted), "\n")
	passwords := make(map[string]string, len(split))

	for _, v := range split {
		passwordSplit := strings.Split(v, " ")
		if len(passwordSplit) != 2 {
			continue
		}
		// TODO: add split len check
		name := passwordSplit[0]
		password := passwordSplit[1]
		passwords[name] = password
	}

	f.passwords = passwords
	f.opened = true

	return nil
}

func (f *PasswordFile) AppendPassword(name, password string) error {
	if !f.opened {
		return ErrPermissionDenied
	}

	data, err := f.ReadFull()
	if err != nil {
		return err
	}

	buff := bytes.NewBuffer(data)
	if _, err := buff.WriteString(name + " " + password + "\n"); err != nil {
		return err
	}

	if _, err := f.Write(buff.Bytes()); err != nil {
		return err
	}

	f.mutex.Lock()
	f.passwords[name] = password
	f.mutex.Unlock()

	return nil
}

func (f *PasswordFile) ShowPassword(name string) (string, error) {
	if !f.opened {
		return "", ErrPermissionDenied
	}

	f.mutex.RLock()
	password, ok := f.passwords[name]
	f.mutex.RUnlock()

	if !ok {
		return "", errors.New("unknown password")
	}

	return password, nil
}

func (f *PasswordFile) List() ([]string, error) {
	if !f.opened {
		return nil, ErrPermissionDenied
	}

	list := make([]string, 0, len(f.passwords))

	f.mutex.RLock()
	for k := range f.passwords {
		list = append(list, k)
	}
	f.mutex.RUnlock()

	return list, nil
}

func (f *PasswordFile) ReadFull() ([]byte, error) {
	data, err := os.ReadFile(f.path)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return data, nil
	}

	decrypted, err := aes.Decrypt(data, f.masterHash)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}

func (f *PasswordFile) Write(data []byte) (n int, err error) {
	encrypted, err := aes.Encrypt(data, f.masterHash)
	if err != nil {
		return 0, err
	}

	if err := os.WriteFile(f.path, encrypted, _DEFAULT_PERM); err != nil {
		return 0, err
	}

	return len(encrypted), nil
}

func (f *PasswordFile) Opened() bool {
	return f.opened
}
