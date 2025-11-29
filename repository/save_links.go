package repository

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/VysMax/links-list-server/entities"
)

func (r *Repository) SaveLinks(linksSet entities.LinksToSave) error {
	resp, err := json.Marshal(linksSet)
	if err != nil {
		return err
	}

	err = os.WriteFile("./saved_links/set_"+strconv.Itoa(linksSet.LinksNum)+".json", resp, 0644)
	if err != nil {
		return err
	}

	entities.LastNum = linksSet.LinksNum

	return nil
}
