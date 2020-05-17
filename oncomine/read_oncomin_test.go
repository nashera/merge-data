package oncomine

import (
	"log"
	"os"
	"testing"

	"github.com/nashera/merge-data/cosmic"
)

func TestOncomine(t *testing.T) {
	cosmicTsv := "/cgdata/zhangxi/project/oncogene/cosmic/CosmicMutantExport.tsv.gz"
	cosmicVarList := cosmic.ReadCosmic(cosmicTsv)
	
	
	varList := ReadOncomine("/cgdata/zhangxi/project/oncogene/oncomine.xlsx")
	fp, err := os.OpenFile("/cgdata/zhangxi/project/oncogene/merge-data/oncomineVsComic.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Panic(err)
	}
	ch := make(chan *VarOncominComplete)
	// defer fp.Close()
	// defer close(ch)
	// defer func(ch chan civic.VarCivicComplete, err )

	// var varList []*civic.VarCivicComplete
	go PrintVariantList(ch, fp)
	for _, variant := range varList {
		ch <- SearchVarInCosmic(variant, cosmicVarList)
	}
	fp.Sync()

}
