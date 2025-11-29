package repository

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/VysMax/links-list-server/entities"
)

func (r *Repository) GetLinks(linkNums entities.LinksNumRequest) error {
	var linkSet entities.LinksToSave
	uniqueLinks := make(map[string]string)

	for _, num := range linkNums.LinksList {

		file, err := os.ReadFile("./saved_links/set_" + strconv.Itoa(num) + ".json")
		if err != nil {
			fmt.Printf("error reading file:%v\n", err)
			continue
		}

		err = json.Unmarshal(file, &linkSet)
		if err != nil {
			return err
		}

		maps.Copy(uniqueLinks, linkSet.Links)
	}

	rowsForPdf := make([]string, 0, len(uniqueLinks))

	for uniqueLink := range uniqueLinks {
		rowsForPdf = append(rowsForPdf, uniqueLink)
	}
	sort.Strings(rowsForPdf)
	var sb strings.Builder

	for _, v := range rowsForPdf {
		sb.WriteString(v + " - " + uniqueLinks[v] + "\n\n")
	}

	fmt.Println(sb.String())
	return nil
}
