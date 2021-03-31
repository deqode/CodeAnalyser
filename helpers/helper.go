package helpers

import (
	"code-analyser/utils"
	"github.com/Masterminds/semver"
)

func SeverValidate(semverstring, value string) bool {
	c, err := semver.NewConstraint(semverstring)
	if err != nil {
		utils.Logger(err)
		return false
	}

	v, err := semver.NewVersion(value)
	if err != nil {
		utils.Logger(err)
		return false
	}

	// Validate a version against a constraint.
	a, _ := c.Validate(v)
	return a
}
