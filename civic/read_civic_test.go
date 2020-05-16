package civic

import (
	"fmt"
	"testing"
)

func TestCivic(t *testing.T) {
	civicVarList := ReadCivic("/cgdata/zhangxi/project/oncogene/variants.with_hgvs.tsv")
	fmt.Println(civicVarList)
	for _, variant := range civicVarList {
		fmt.Println(variant)
	}

}
