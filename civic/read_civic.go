package civic

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nashera/merge-data/cosmic"
	"github.com/nashera/merge-data/tsvfile"
)

// VarCivic Variant in civic
type VarCivic struct {
	VariantID  string
	Chromosome string
	ChromStart string
	ChromEnd   string
	Name       string
	EntrezName string
	Score      string
	Hgvs       string
}

// VarCivicComplete complete version of civic variant
type VarCivicComplete struct {
	VarCivic
	cosmic.CosVar
}

// ReadCivic 读取civic数据库数据
func ReadCivic(civicPath string) []*VarCivic {
	fmt.Println(civicPath)
	lines, errors, err := tsvfile.TsvLines(civicPath)
	if err != nil {
		log.Fatal(err)
	}
	go func(errs chan error) {
		err := <-errs
		if err != nil {
			log.Fatal(err)
		}
	}(errors)
	var varList []*VarCivic
	for line := range lines {
		w := strings.Split(line, "\t")
		varList = append(varList,
			&VarCivic{
				VariantID:  strings.Trim(w[0], "\""),
				Chromosome: strings.Trim(w[1], `"`),
				ChromStart: strings.Trim(w[2], `"`),
				ChromEnd:   strings.Trim(w[3], `"`),
				Name:       strings.Trim(w[4], `"`),
				EntrezName: strings.Trim(w[5], `"`), // GeneName
				Score:      strings.Trim(w[6], `"`),
				Hgvs:       strings.Trim(w[7], `"`),
			})

		// fmt.Printf("%+v\n", string(line))
	}
	return varList

}

func compareCivicCosmic(vc1 *VarCivic, vc2 *cosmic.CosVar) bool {
	if strings.HasPrefix(vc2.MutationGenomePosition, vc1.Chromosome+":"+vc1.ChromStart) {
		return true
	}
	if vc1.EntrezName == vc2.Gene && strings.Contains(vc2.MutationAA, vc1.Name) {
		return true
	}
	return false
}

// SearchCivicVar search variants of civic in cosmic database
func SearchCivicVar(vc *VarCivic, varCosmicList []*cosmic.CosVar) *VarCivicComplete {
	for _, variant := range varCosmicList {
		if compareCivicCosmic(vc, variant) {
			return &VarCivicComplete{VarCivic: *vc, CosVar: *variant}
		}
		// return &VarCivicComplete{VarCivic: *vc, CosVar: *variant}
	}
	return &VarCivicComplete{
		VarCivic: *vc,
		CosVar:   cosmic.CosVar{}}
}

// PrintVariantList 打印variant list
func PrintVariantList(variants chan *VarCivicComplete, fp *os.File) {
	// fp, err := os.OpenFile(outputFile, os.O_CREATE|os.O_APPEND, 6)
	// if err != nil {
	// 	return err
	// }
	fp.WriteString("civic.VariantId\t")
	fp.WriteString("civic.Chromosome\t")
	fp.WriteString("civic.ChromStart\t")
	fp.WriteString("civic.ChromEnd\t")
	fp.WriteString("civic.ChromName\t")
	fp.WriteString("civic.EntrezName\t")
	fp.WriteString("civic.Score\t")
	fp.WriteString("civic.Hgvs\t")
	fp.WriteString("cosmic.AccessionNumber\t")
	fp.WriteString("cosmic.Gene\t")
	fp.WriteString("cosmic.GeneCDSLength\t")
	fp.WriteString("cosmic.SampleName\t")
	fp.WriteString("cosmic.MutationCDS\t")
	fp.WriteString("cosmic.MutationAA\t")
	fp.WriteString("cosmic.MutationGenomePosition\t")
	fp.WriteString("cosmic.HGVSC\t")
	fp.WriteString("cosmic.HGVSG\t")
	fp.WriteString("cosmic.GenomicMutationID\n")
	for variant := range variants {
		fp.WriteString(variant.VariantID + "\t")
		fp.WriteString(variant.Chromosome + "\t")
		fp.WriteString(variant.ChromStart + "\t")
		fp.WriteString(variant.ChromEnd + "\t")
		fp.WriteString(variant.Name + "\t")
		fp.WriteString(variant.EntrezName + "\t")
		fp.WriteString(variant.Score + "\t")
		fp.WriteString(variant.Hgvs + "\t")
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
