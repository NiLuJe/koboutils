package kobo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
)

// VersionCompare compares two firmware versions.
// a < b = -1
// a = b = 0
// a > b = 1
func VersionCompare(a, b string) int {
	aspl, bspl := strSplitInt(a), strSplitInt(b)
	if len(aspl) != len(bspl) {
		return 0
	}
	for i := range aspl {
		switch {
		case aspl[i] > bspl[i]:
			return 1
		case bspl[i] > aspl[i]:
			return -1
		}
	}
	return 0
}

// ParseKoboVersion gets the info from the .kobo/version file.
func ParseKoboVersion(kpath string) (serial, version, id string, err error) {
	vbuf, err := ioutil.ReadFile(filepath.Join(kpath, ".kobo", "version"))
	if err != nil {
		return "", "", "", err
	}
	spl := strings.Split(strings.TrimSpace(string(vbuf)), ",")
	if len(spl) != 6 {
		return "", "", "", fmt.Errorf("length of split version file should be 6, got %d", len(spl))
	}
	return spl[0], spl[2], spl[5], nil
}

// ParseKoboAffiliate parses the affiliate from the .kobo/affiliate.conf file.
func ParseKoboAffiliate(kpath string) (affiliate string, err error) {
	abuf, err := ioutil.ReadFile(filepath.Join(kpath, ".kobo", "affiliate.conf"))
	if err != nil {
		return "", err
	}
	m := regexp.MustCompile(`affiliate *= *([a-zA-Z0-9]+)`).FindStringSubmatch(string(abuf))
	if len(m) != 2 {
		return "", errors.New("could not parse affiliate.conf")
	}
	return m[1], nil
}
