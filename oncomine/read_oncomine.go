package oncomine

import (
	"fmt"
	"os"
	"strings"

	"github.com/nashera/merge-data/cosmic"
	"github.com/tealeg/xlsx"
)

// VarOncomin variant in oncomin
type VarOncomin struct {
	GeneName string
	Flag     string
	Chr      string
	Start    string
	End      string
	AAchange string
	Ref      string
	Alt      string
	CosmicID string
}

// VarOncominComplete 完善VarOncomin
type VarOncominComplete struct {
	VarOncomin
	cosmic.CosVar
}

// ReadOncomine 读取oncomine表格
func ReadOncomine(sheetPath string) []*VarOncomin {
	xlFile, err := xlsx.OpenFile(sheetPath)
	var varList []*VarOncomin
	if err != nil {
		fmt.Println(err.Error())
	}
	var varSheet = xlFile.Sheet["oncomine_hotspot"]
	var header []string
	for i, row := range varSheet.Rows {
		if i == 0 {
			for _, cell := range row.Cells {
				header = append(header, cell.String())
			}
		} else {
			rowMap := make(map[string]string)
			for j, cell := range row.Cells {
				if j > len(header)-1 {
					break
				}
				rowMap[header[j]] = cell.String()
			}
			varList = append(varList,
				&VarOncomin{
					GeneName: rowMap["Gene"],
					Flag:     rowMap["flag"],
					Chr:      strings.TrimPrefix(rowMap["Chr"], "chr"),
					Start:    rowMap["start"],
					End:      rowMap["end"],
					AAchange: rowMap["AA.change"],
					Ref:      rowMap["Ref"],
					Alt:      rowMap["Alt"],
					CosmicID: rowMap["COSMIC ID"],
				})

		}
	}
	return varList
}

func compareOncomineCosmic(vc1 *VarOncomin, vc2 *cosmic.CosVar) bool {
	if strings.HasPrefix(vc2.MutationGenomePosition, vc1.Chr+":"+vc1.Start) {
		return true
	}
	if vc1.GeneName == vc2.Gene && strings.Contains(vc2.MutationAA, vc1.AAchange) {
		return true
	}
	return false
}

// SearchVarInCosmic search variants of civic in cosmic database
func SearchVarInCosmic(vc *VarOncomin, varCosmicList []*cosmic.CosVar) *VarOncominComplete {
	for _, variant := range varCosmicList {
		if compareOncomineCosmic(vc, variant) {
			return &VarOncominComplete{VarOncomin: *vc, CosVar: *variant}
		}
	}
	return &VarOncominComplete{
		VarOncomin: *vc,
		CosVar:     cosmic.CosVar{}}
}

// PrintVariantList 打印variant list
func PrintVariantList(variants chan *VarOncominComplete, fp *os.File) {
	// fp, err := os.OpenFile(outputFile, os.O_CREATE|os.O_APPEND, 6)
	// if err != nil {
	// 	return err
	// }
	fp.WriteString("Oncomine.GeneName\t")
	fp.WriteString("Oncomine.Flag\t")
	fp.WriteString("Oncomine.Chr\t")
	fp.WriteString("Oncomine.Start\t")
	fp.WriteString("Oncomine.End\t")
	fp.WriteString("Oncomine.AAChange\t")
	fp.WriteString("Oncomine.Ref\t")
	fp.WriteString("Oncomine.Alt\t")
	fp.WriteString("Oncomine.CosmicID\t")

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
		fp.WriteString(variant.Flag + "\t")
		fp.WriteString(variant.Chr + "\t")
		fp.WriteString(variant.Start + "\t")
		fp.WriteString(variant.End + "\t")
		fp.WriteString(variant.AAchange + "\t")
		fp.WriteString(variant.Ref + "\t")
		fp.WriteString(variant.Alt + "\t")
		fp.WriteString(variant.CosmicID + "\t")
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
