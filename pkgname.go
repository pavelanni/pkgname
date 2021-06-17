// Package pkgname provides functions to parse Linux package names
// in RPM and Debian formats.
package pkgname

import (
	"errors"
	"strings"
)

type Pkg struct {
	Name, Version, Arch, Type string
}

// Parse parses the package name passed as a string
// and returns a Pkg structure that includes package's
// Name, Version, Release (for RPM), Arch, and Type (can be rpm or deb).
// It calls ParseRpm or ParseDeb depending on the package type.
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

// ParseRpm parses the RPM package name passed as a string
// and returns a Pkg structure that includes package's
// Name, Version, Release, Arch, and Type (which is rpm in this case).
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
	release := pname[lastSep+1:]
	pname = pname[:lastSep]
	lastSep = strings.LastIndexByte(pname, '-')
	if lastSep == -1 || lastSep == 0 {
		return Pkg{}, errors.New("no name part in the package name")
	}
	p.Version = pname[lastSep+1:] + "-" + release
	p.Name = pname[:lastSep]

	return p, nil
}

// ParseDeb parses the RPM package name passed as a string
// and returns a Pkg structure that includes package's
// Name, Version, Release (which is empty for deb) Arch, and Type (which is deb in this case).
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
	p.Name = pname[:lastSep]

	return p, nil
}
