package oncokb

import (
	"log"
	"os"
	"strings"

	"github.com/nashera/merge-data/cosmic"
	"github.com/nashera/merge-data/tsvfile"
)

// VarOncokb Variant in oncokb
type VarOncokb struct {
	GeneName string
	AAChange string
}

// VarOncokbComplete 合并cosmic数据
type VarOncokbComplete struct {
	VarOncokb
	cosmic.CosVar
}

// ReadOncokb 读取oncokb数据库数据
func ReadOncokb(oncokbPath string) []*VarOncokb {
	lines, errors, err := tsvfile.TsvLines(oncokbPath)
	if err != nil {
		log.Fatal(err)
	}
	go func(errs chan error) {
		err := <-errs
		if err != nil {
			log.Fatal(err)
		}
	}(errors)
	var varList []*VarOncokb
	for line := range lines {
		w := strings.Split(line, "\t")
		varList = append(varList,
			&VarOncokb{
				GeneName: w[0],
				AAChange: w[1],
			})

		// fmt.Printf("%+v\n", string(line))
	}
	return varList

}

func compareOncokbCosmic(vc1 *VarOncokb, vc2 *cosmic.CosVar) bool {
	if vc1.GeneName == vc2.Gene && strings.Contains(vc2.MutationAA, vc1.AAChange) {
		return true
	}
	return false
}

// SearchVarInCosmic search variants of civic in cosmic database
func SearchVarInCosmic(vc *VarOncokb, varCosmicList []*cosmic.CosVar) *VarOncokbComplete {
	for _, variant := range varCosmicList {
		if compareOncokbCosmic(vc, variant) {
			return &VarOncokbComplete{VarOncokb: *vc, CosVar: *variant}
		}
	}
	return &VarOncokbComplete{
		VarOncokb: *vc,
		CosVar:    cosmic.CosVar{}}
}

// PrintVariantList 打印variant list
func PrintVariantList(variants chan *VarOncokbComplete, fp *os.File) {
	// fp, err := os.OpenFile(outputFile, os.O_CREATE|os.O_APPEND, 6)
	// if err != nil {
	// 	return err
	// }
	fp.WriteString("Oncokb.GeneName\t")
	fp.WriteString("Oncokb.AAChange\t")

	fp.WriteString("cosmic.AccessionNumber")
	fp.WriteString("cosmic.Gene")
	fp.WriteString("cosmic.GeneCDSLength\t")
	fp.WriteString("cosmic.SampleName\t")
	fp.WriteString("cosmic.MutationCDS\t")
	fp.WriteString("cosmic.MutationAA\t")
	fp.WriteString("cosmic.MutationGenomePosition\t")
	fp.WriteString("cosmic.HGVSC\t")
	fp.WriteString("cosmic.HGVSG\t")
	fp.WriteString("cosmic.GenomicMutationID\n")
	for variant := range variants {
		fp.WriteString(variant.GeneName + "\t")
		fp.WriteString(variant.AAChange + "\t")

		fp.WriteString(variant.AccessionNumber + "\t")
		fp.WriteString(variant.Gene + "\t")
		fp.WriteString(variant.GeneCDSLength + "\t")
		fp.WriteString(variant.SampleName + "\t")
		fp.WriteString(variant.MutationCDS + "\t")
		fp.WriteString(variant.MutationAA + "\t")
		fp.WriteString(variant.MutationGenomePosition + "\t")
		fp.WriteString(variant.HGVSC + "\t")
		fp.WriteString(variant.HGVSG + "\t")
		fp.WriteString(variant.GenomicMutationID + "\n")
	}

}
