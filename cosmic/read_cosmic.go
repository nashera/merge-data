package cosmic

import (
	"log"
	"strings"

	"github.com/nashera/merge-data/mygzip"
)

// CosVar a object of variant on cosmic
type CosVar struct {
	AccessionNumber        string // Accession Number
	Gene                   string // Gene name
	GeneCDSLength          string // Gene CDS length
	SampleName             string // Sample name
	MutationCDS            string // Mutation CDS
	MutationAA             string // Mutation AA
	MutationGenomePosition string // Mutation genome position $26
	HGVSC                  string // HGVSC $39
	HGVSG                  string // HGVSG $40
	GenomicMutationID      string // GENOMIC_MUTATION_ID $17
}

// ReadCosmic 读取cosmic tsv文件到 CosmicVar splice
func ReadCosmic(cosmicPath string) []*CosVar {
	lines, errors, err := mygzip.GZLines(cosmicPath)
	if err != nil {
		log.Fatal(err)
	}
	go func(errs chan error) {
		err := <-errs
		if err != nil {
			log.Fatal(err)
		}
	}(errors)
	num := 0
	var header []string
	var cosmicVarList []*CosVar
	for line := range lines {
		num++
		if num == 1 {
			header = strings.Split(string(line), "\t")
		} else {
			rowMap := make(map[string]string)
			for i, v := range strings.Split(string(line), "\t") {
				if i > len(header)-1 {
					break
				}
				rowMap[header[i]] = v
			}
			cosmicVarList = append(cosmicVarList,
				&CosVar{
					AccessionNumber:        rowMap["Accession Number"],
					Gene:                   rowMap["Gene name"],
					GeneCDSLength:          rowMap["Gene CDS length"],
					SampleName:             rowMap["Sample name"],
					MutationCDS:            rowMap["Mutation CDS"],
					MutationAA:             rowMap["Mutation AA"],
					MutationGenomePosition: rowMap["Mutation genome position"],
					HGVSC:                  rowMap["HGVSC"],
					HGVSG:                  rowMap["HGVSG"],
					GenomicMutationID:      rowMap["GENOMIC_MUTATION_ID"],
				})
		}

		// fmt.Printf("%+v\n", string(line))
	}
	return cosmicVarList

}

// ReadMergedCosmic 读取cosmic tsv文件到 CosmicVar splice
func ReadMergedCosmic(cosmicPath string) []*CosVar {
	lines, errors, err := mygzip.GZLines(cosmicPath)
	if err != nil {
		log.Fatal(err)
	}
	go func(errs chan error) {
		err := <-errs
		if err != nil {
			log.Fatal(err)
		}
	}(errors)
	num := 0
	var header []string
	var cosmicVarList []*CosVar
	for line := range lines {
		num++
		if num == 1 {
			header = strings.Split(string(line), "\t")
		} else {
			rowMap := make(map[string]string)
			for i, v := range strings.Split(string(line), "\t") {
				if i > len(header)-1 {
					break
				}
				rowMap[header[i]] = v
			}
			cosmicVarList = append(cosmicVarList,
				&CosVar{
					AccessionNumber:        rowMap["Accession Number"],
					Gene:                   rowMap["Gene name"],
					GeneCDSLength:          rowMap["Gene CDS length"],
					SampleName:             rowMap["Sample name"],
					MutationCDS:            rowMap["Mutation CDS"],
					MutationAA:             rowMap["Mutation AA"],
					MutationGenomePosition: rowMap["Mutation genome position"],
					HGVSC:                  rowMap["HGVSC"],
					HGVSG:                  rowMap["HGVSG"],
					GenomicMutationID:      rowMap["GENOMIC_MUTATION_ID"],
				})
		}

		// fmt.Printf("%+v\n", string(line))
	}
	return cosmicVarList

}
