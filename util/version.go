package util

import (
	"fmt"

	"github.com/hashicorp/go-version"
)

/* compareVersions function to compare two versions. It returns
   if vesrion1 < version2 return 1
   if vesrion1 > version2 return -1
   else return 0
*/
func CompareVersions(version1, version2 string) int {

	v1, err := version.NewVersion(version1)

	if err != nil {
		fmt.Println("error to create the newversion :", err)
	}
	v2, err := version.NewVersion(version2)
	if err != nil {
		fmt.Println("error to create the newversion :", err)
	}

	if v1.LessThan(v2) {
		return 1
	} else if v2.LessThan(v1) {
		return -1
	}
	return 0
}
