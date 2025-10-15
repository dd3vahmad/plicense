package fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/dd3vahmad/plicense/ui"
)

const baseURL = "https://api.github.com/licenses"

func LicenseList() ([]ui.License, error) {
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

	return data, nil
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
