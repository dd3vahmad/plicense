package fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dd3vahmad/plicense/ui"
)

const baseURL = "https://api.github.com/licenses"

func LicenseList(dir string) ([]ui.License, error) {
	path := filepath.Join(dir, "licenses.json")
	file, err := os.Open(path)

	if err != nil {
		res, err := http.Get(baseURL)
		if err != nil {
			fmt.Printf("error fetching licenses")
			return []ui.License{}, err
		}

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return []ui.License{}, err
		}

		var data []ui.License
		if err := json.Unmarshal(body, &data); err != nil {
			return []ui.License{}, err
		}

		newFile, _ := os.Create(path)
		defer newFile.Close()

		json.NewEncoder(newFile).Encode(data)

		return data, nil
	} else {
		defer file.Close()
		var licenses []ui.License
		json.NewDecoder(file).Decode(&licenses)

		return licenses, nil
	}
}

func LicenseDetails(key string) (ui.License, error) {
	res, err := http.Get(fmt.Sprintf("%s/%s", baseURL, key))
	if err != nil {
		return ui.License{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ui.License{}, err
	}

	var license ui.License
	if err := json.Unmarshal(body, &license); err != nil {
		return ui.License{}, err
	}

	return license, nil
}
