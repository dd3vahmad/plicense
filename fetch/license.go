package fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dd3vahmad/plicense/entity"
)

const baseURL = "https://api.github.com/licenses"

func LicenseList(dir string) ([]entity.License, error) {
	path := filepath.Join(dir, "licenses.json")
	os.MkdirAll(filepath.Dir(path), os.ModePerm)

	file, err := os.Open(path)

	if err != nil {
		res, err := http.Get(baseURL)
		if err != nil {
			fmt.Printf("error fetching licenses")
			return []entity.License{}, err
		}

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return []entity.License{}, err
		}

		var data []entity.License
		if err := json.Unmarshal(body, &data); err != nil {
			return []entity.License{}, err
		}

		newFile, _ := os.Create(path)
		defer newFile.Close()

		json.NewEncoder(newFile).Encode(data)

		return data, nil
	} else {
		defer file.Close()
		var licenses []entity.License
		json.NewDecoder(file).Decode(&licenses)

		return licenses, nil
	}
}

func LicenseDetails(key string) (entity.License, error) {
	path := filepath.Join("licenses", fmt.Sprintf("%s.json", key))
	if file, err := os.Open(path); err == nil {
		defer file.Close()

		var license entity.License
		json.NewDecoder(file).Decode(&license)

		return license, nil
	}

	res, err := http.Get(fmt.Sprintf("%s/%s", baseURL, key))
	if err != nil {
		return entity.License{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return entity.License{}, err
	}

	var license entity.License
	if err := json.Unmarshal(body, &license); err != nil {
		return entity.License{}, err
	}

	newLicense, _ := os.Create(path)
	defer newLicense.Close()

	json.NewEncoder(newLicense).Encode(license)

	return license, nil
}
