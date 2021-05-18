package helpers

import (
	"code-analyser/utils"
	"github.com/Masterminds/semver"
	"strings"
)

//SemverValidate takes semver and value to be matched
func SemverValidate(semverstring, value string) bool {
	if strings.Contains(value, "^") {
		value = strings.Replace(value, "^", "", 1)
	}
	if strings.Contains(value, "~") {
		value = strings.Replace(value, "~", "", 1)
	}
	c, err := semver.NewConstraint(semverstring)
	if err != nil {
		utils.Logger(err, semverstring)
		return false
	}

	v, err := semver.NewVersion(value)
	if err != nil {
		utils.Logger(err, value)
		return false
	}

	// Validate a version against a constraint.
	a, _ := c.Validate(v)
	return a
}

//SemverValidateFromArray validates semver array and version
func SemverValidateFromArray(semverstrings []string, value string) bool {
	for _, sem := range semverstrings {
		if SemverValidate(sem, value) {
			return true
		}
	}
	return false

}


