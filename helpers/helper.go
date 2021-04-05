package helpers

import (
	"code-analyser/utils"
	"github.com/Masterminds/semver"
)
//SeverValidate takes semver and value to be matched
func SeverValidate(semverstring, value string) bool {
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

//SeverValidateFromArray validates semver array and version
func SeverValidateFromArray(semverstrings []string, value string) bool {
	for _, sem := range semverstrings {
		if SeverValidate(sem, value) {
			return true
		}
	}
	return false

}
