package fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dd3vahmad/plicense/internals/entity"
)

func LicensePath(key string) (string, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user cache directory %w", err)
	}

	licenseDir := filepath.Join(cacheDir, "plicense", "licenses")
	if err := os.MkdirAll(licenseDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create licenses dir: %w", err)
	}

	path := filepath.Join(licenseDir, "licenses.json")
	if key != "" {
		path = filepath.Join(licenseDir, fmt.Sprintf("%s.json", key))
	}
	return path, nil
}

const baseURL = "https://api.github.com/licenses"

func LicenseList() ([]entity.License, error) {
	path, _ := LicensePath("")

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
	path, _ := LicensePath(key)

	if _, err := os.Stat(path); err == nil {
		file, _ := os.ReadFile(path)

		var license entity.License
		if err := json.Unmarshal(file, &license); err != nil {
			return entity.License{}, fmt.Errorf("failed to decode cached license: %w", err)
		}

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

	return license, nil
}
