package main

import (
	"fmt"

	"github.com/pavelanni/pkgname"
)

var pkgList = []string{
	"dnf-plugin-spacewalk-2.8.5-11.module+el8.1.0+3455+3ddf2832.noarch.rpm",
	"pt-sans-fonts-20141121-18.fc33.noarch.rpm",
	"python3-canberra-0-0.25.git88c53cd.fc33.noarch.rpm",
	"python3-pycurl-7.43.0.5-6.fc33.x86_64.rpm",
	"shim-x64-15-8.x86_64.rpm",
	"skkdic-20200128-2.T1339.fc33.noarch.rpm",
	"squashfs-tools-4.4-2.20200513gitc570c61.fc33.x86_64.rpm",
	"4.4-2.20200513gitc570c61.fc33.x86_64.rpm",
	"-4.4-2.20200513gitc570c61.fc33.x86_64.rpm",
	"2.20200513gitc570c61.fc33.x86_64.rpm",
	".x86_64.rpm",
	"x86_64.rpm",
	".rpm",
	"rpm",
	"bash_5.0-6ubuntu1.1_amd64.deb",
	"dpkg_1.19.7ubuntu3_amd64.deb",
	"gcc-10-base_10.2.0-5ubuntu1~20.04_amd64.deb",
	"libapt-pkg6.0_2.0.5_amd64.deb",
}

func main() {

	for _, pname := range pkgList {
		p, err := pkgname.Parse(pname)
		if err != nil {
			fmt.Println(pname, err)
			continue
		}
		fmt.Println(pname, p)
	}
}
