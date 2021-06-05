package pkgname

import (
	"errors"
	"strings"
)

type Pkg struct {
	Name, Version, Release, Arch, Type string
}

func Parse(pname string) (Pkg, error) {
	var p Pkg

	if pname == "" {
		return Pkg{}, errors.New("empty input string")
	}
	lastSep := strings.LastIndexByte(pname, '.')
	if lastSep == -1 {
		return Pkg{}, errors.New("no dots in the package name")
	}
	p.Type = pname[lastSep+1:]
	if p.Type == "rpm" {
		return ParseRpm(pname)
	}
	if p.Type == "deb" {
		return ParseDeb(pname)
	}
	return Pkg{}, errors.New("wrong package type; should be rpm or deb")

}

func ParseRpm(pname string) (Pkg, error) {
	var p Pkg

	if pname == "" {
		return Pkg{}, errors.New("empty input string")
	}
	lastSep := strings.LastIndexByte(pname, '.')
	if lastSep == -1 {
		return Pkg{}, errors.New("no dots in the package name")
	}
	p.Type = pname[lastSep+1:]
	if p.Type != "rpm" {
		return Pkg{}, errors.New("not an RPM package name")
	}
	pname = strings.TrimSuffix(pname, ".rpm")
	if len(pname) == 0 { // ".rpm"
		return Pkg{}, errors.New("no arch part in the package name")
	}
	lastSep = strings.LastIndexByte(pname, '.')
	if lastSep == -1 || lastSep == 0 { // "arch.rpm" without dot
		return Pkg{}, errors.New("no release part in the package name")
	}
	p.Arch = pname[lastSep+1:] // "arch.rpm"
	pname = pname[:lastSep]
	lastSep = strings.LastIndexByte(pname, '-')
	if lastSep == -1 || lastSep == 0 { // "rel.arch.rpm" no dash
		return Pkg{}, errors.New("no version part in the package name")
	}
	p.Release = pname[lastSep+1:]
	pname = pname[:lastSep]
	lastSep = strings.LastIndexByte(pname, '-')
	if lastSep == -1 || lastSep == 0 {
		return Pkg{}, errors.New("no name part in the package name")
	}
	p.Version = pname[lastSep+1:]
	p.Name = pname[:lastSep]

	return p, nil
}

func ParseDeb(pname string) (Pkg, error) {
	var p Pkg

	if pname == "" {
		return Pkg{}, errors.New("empty input string")
	}
	lastSep := strings.LastIndexByte(pname, '.')
	if lastSep == -1 {
		return Pkg{}, errors.New("no dots in the package name")
	}
	p.Type = pname[lastSep+1:]
	if p.Type != "deb" {
		return Pkg{}, errors.New("not a Debian package name")
	}
	pname = strings.TrimSuffix(pname, ".deb")
	if len(pname) == 0 { // ".rpm"
		return Pkg{}, errors.New("no arch part in the package name")
	}
	lastSep = strings.LastIndexByte(pname, '_')
	if lastSep == -1 || lastSep == 0 { // "arch.rpm" without dot
		return Pkg{}, errors.New("no release part in the package name")
	}
	p.Arch = pname[lastSep+1:] // "arch.rpm"
	pname = pname[:lastSep]
	lastSep = strings.LastIndexByte(pname, '_')
	if lastSep == -1 || lastSep == 0 {
		return Pkg{}, errors.New("no name part in the package name")
	}
	p.Version = pname[lastSep+1:]
	p.Release = ""
	p.Name = pname[:lastSep]

	return p, nil
}
