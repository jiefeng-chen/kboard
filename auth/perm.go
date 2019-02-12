package auth

import "errors"

// 权限
type PermCode string
type Mode int8


type PermissionCode interface {
	Set(PermCode, Mode) error
	Check(PermCode) error
}

type permissionCode struct {
	codes map[PermCode]Mode
	Uid  int
}

func (p *permissionCode) Set(code PermCode, mode Mode) error {
	if code == "" {
		return errors.New("permission code is empty")
	}

	return nil
}

func (p *permissionCode) Check(code PermCode) error {
	if _, ok := p.codes[code]; !ok {
		return errors.New("permission code not found")
	}
	return nil
}





