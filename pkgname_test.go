package pkgname_test

import (
	"testing"

	"github.com/pavelanni/pkgname"
)

func TestEmptyString(t *testing.T) {
	got, goterr := pkgname.Parse("")
	want := pkgname.Pkg{}
	if got != want {
		t.Errorf("got: %s\nwant: %s\n", got, want)
	}
	if goterr == nil {
		t.Error("error should be not nil")
	}
}

func TestNoDots(t *testing.T) {
	got, goterr := pkgname.Parse("gjkflgj")
	want := pkgname.Pkg{}
	if got != want {
		t.Errorf("got: %s\nwant: %s\n", got, want)
	}
	if goterr == nil {
		t.Error("error should be not nil")
	}
}

func TestWrongType(t *testing.T) {
	got, goterr := pkgname.Parse("gjkflgj.exe")
	want := pkgname.Pkg{}
	if got != want {
		t.Errorf("got: %s\nwant: %s\n", got, want)
	}
	if goterr == nil {
		t.Error("error should be not nil")
	}
}

func TestEmptyStringRpm(t *testing.T) {
	got, goterr := pkgname.ParseRpm("")
	want := pkgname.Pkg{}
	if got != want {
		t.Errorf("got: %s\nwant: %s\n", got, want)
	}
	if goterr == nil {
		t.Error("error should be not nil")
	}
}

func TestNoDotsRpm(t *testing.T) {
	got, goterr := pkgname.ParseRpm("gjkflgj")
	want := pkgname.Pkg{}
	if got != want {
		t.Errorf("got: %s\nwant: %s\n", got, want)
	}
	if goterr == nil {
		t.Error("error should be not nil")
	}
}

func TestEmptyStringDeb(t *testing.T) {
	got, goterr := pkgname.ParseDeb("")
	want := pkgname.Pkg{}
	if got != want {
		t.Errorf("got: %s\nwant: %s\n", got, want)
	}
	if goterr == nil {
		t.Error("error should be not nil")
	}
}

func TestNoDotsDeb(t *testing.T) {
	got, goterr := pkgname.ParseDeb("gjkflgj")
	want := pkgname.Pkg{}
	if got != want {
		t.Errorf("got: %s\nwant: %s\n", got, want)
	}
	if goterr == nil {
		t.Error("error should be not nil")
	}
}

func TestNotRpm(t *testing.T) {
	got, goterr := pkgname.ParseRpm("bash-2.5_amd64.deb")
	want := pkgname.Pkg{}
	if got != want {
		t.Errorf("got: %s\nwant: %s\n", got, want)
	}
	if goterr == nil {
		t.Error("error should be not nil")
	}
}

func TestOnlyRpmExt(t *testing.T) {
	got, goterr := pkgname.ParseRpm(".rpm")
	want := pkgname.Pkg{}
	if got != want {
		t.Errorf("got: %s\nwant: %s\n", got, want)
	}
	if goterr == nil {
		t.Error("error should be not nil")
	}
}

func TestEmptyName(t *testing.T) {
	got, goterr := pkgname.ParseRpm("4.4-2.20200513gitc570c61.fc33.x86_64.rpm")
	want := pkgname.Pkg{}
	if goterr == nil {
		t.Error("error should be not nil")
	}
	if got != want {
		t.Errorf("got: %s\nwant: %s\n", got, want)
	}
}

func TestEmptyNameWithDash(t *testing.T) {
	got, goterr := pkgname.ParseRpm("-4.4-2.20200513gitc570c61.fc33.x86_64.rpm")
	want := pkgname.Pkg{}
	if goterr == nil {
		t.Error("error should be not nil")
	}
	if got != want {
		t.Errorf("got: %s\nwant: %s\n", got, want)
	}
}

func TestEmptyVersion(t *testing.T) {
	got, goterr := pkgname.ParseRpm("2.20200513gitc570c61.fc33.x86_64.rpm")
	want := pkgname.Pkg{}
	if goterr == nil {
		t.Error("error should be not nil")
	}
	if got != want {
		t.Errorf("got: %s\nwant: %s\n", got, want)
	}
}

func TestEmptyVersionWithDash(t *testing.T) {
	got, goterr := pkgname.ParseRpm("-2.20200513gitc570c61.fc33.x86_64.rpm")
	want := pkgname.Pkg{}
	if goterr == nil {
		t.Error("error should be not nil")
	}
	if got != want {
		t.Errorf("got: %s\nwant: %s\n", got, want)
	}
}

func TestEmptyReleaseWithDot(t *testing.T) {
	got, goterr := pkgname.ParseRpm(".x86_64.rpm")
	want := pkgname.Pkg{}
	if goterr == nil {
		t.Error("error should be not nil")
	}
	if got != want {
		t.Errorf("got: %s\nwant: %s\n", got, want)
	}
}

func TestEmptyRelease(t *testing.T) {
	got, goterr := pkgname.ParseRpm("x86_64.rpm")
	want := pkgname.Pkg{}
	if goterr == nil {
		t.Error("error should be not nil")
	}
	if got != want {
		t.Errorf("got: %s\nwant: %s\n", got, want)
	}
}

func TestValidName(t *testing.T) {
	got, goterr := pkgname.ParseRpm("squashfs-tools-4.4-2.20200513gitc570c61.fc33.x86_64.rpm")
	want := pkgname.Pkg{
		Name:    "squashfs-tools",
		Version: "4.4",
		Release: "2.20200513gitc570c61.fc33",
		Arch:    "x86_64",
		Type:    "rpm",
	}
	if goterr != nil {
		t.Error("error should be nil")
	}
	if got != want {
		t.Errorf("got: %s\nwant: %s\n", got, want)
	}
}

func TestNotDeb(t *testing.T) {
	got, goterr := pkgname.ParseDeb("bash-2.5_amd64.rpm")
	want := pkgname.Pkg{}
	if got != want {
		t.Errorf("got: %s\nwant: %s\n", got, want)
	}
	if goterr == nil {
		t.Error("error should be not nil")
	}
}

func TestOnlyDebExt(t *testing.T) {
	got, goterr := pkgname.ParseDeb(".deb")
	want := pkgname.Pkg{}
	if got != want {
		t.Errorf("got: %s\nwant: %s\n", got, want)
	}
	if goterr == nil {
		t.Error("error should be not nil")
	}
}

func TestEmptyNameDeb(t *testing.T) {
	got, goterr := pkgname.ParseRpm("4.4-2.20200513gitc570c61.fc33.x86_64.rpm")
	want := pkgname.Pkg{}
	if goterr == nil {
		t.Error("error should be not nil")
	}
	if got != want {
		t.Errorf("got: %s\nwant: %s\n", got, want)
	}
}

func TestEmptyNameWithUnderscoreDeb(t *testing.T) {
	got, goterr := pkgname.ParseDeb("_10.2.0-5ubuntu1~20.04_amd64.deb")
	want := pkgname.Pkg{}
	if goterr == nil {
		t.Error("error should be not nil")
	}
	if got != want {
		t.Errorf("got: %s\nwant: %s\n", got, want)
	}
}

func TestEmptyVersionDeb(t *testing.T) {
	got, goterr := pkgname.ParseDeb("amd64.deb")
	want := pkgname.Pkg{}
	if goterr == nil {
		t.Error("error should be not nil")
	}
	if got != want {
		t.Errorf("got: %s\nwant: %s\n", got, want)
	}
}

func TestEmptyVersionWithUnderscoreDeb(t *testing.T) {
	got, goterr := pkgname.ParseDeb("_amd64.deb")
	want := pkgname.Pkg{}
	if goterr == nil {
		t.Error("error should be not nil")
	}
	if got != want {
		t.Errorf("got: %s\nwant: %s\n", got, want)
	}
}

func TestValidNameDeb(t *testing.T) {
	got, goterr := pkgname.ParseDeb("gcc-10-base_10.2.0-5ubuntu1~20.04_amd64.deb")
	want := pkgname.Pkg{
		Name:    "gcc-10-base",
		Version: "10.2.0-5ubuntu1~20.04",
		Release: "",
		Arch:    "amd64",
		Type:    "deb",
	}
	if goterr != nil {
		t.Error("error should be nil")
	}
	if got != want {
		t.Errorf("got: %s\nwant: %s\n", got, want)
	}
}
