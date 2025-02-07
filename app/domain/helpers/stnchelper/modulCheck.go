package stnchelper

import (
	"fmt"
	"stncCms/app/domain/helpers/stnccollection"
)

func ModulNameUrlCheck(modulName string) string {
	var modulNameReturn string
	words := []string{"sacrifece", "fundraising"}
	_, found := stnccollection.FindSlice(words, modulName)
	if found {
		modulNameReturn = modulName
	} else {
		modulNameReturn = "sacrifece"
	}
	fmt.Println(modulNameReturn)
	return modulNameReturn
}
