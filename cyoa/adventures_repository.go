package cyoa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// AdventureRepository doc
type AdventureRepository struct {
	Adventures []Adventure
}

// NewAdventureRepository doc
func NewAdventureRepository(fileName string) (AdventureRepository, error) {
	data, err := initData(fileName) // TODO: Melhorar
	if err != nil {
		return AdventureRepository{}, err
	}

	return AdventureRepository{data}, err
}

// GetBySlug doc
func (ar AdventureRepository) GetBySlug(slug string) (Adventure, error) {
	for _, v := range ar.Adventures {
		if v.Slug == slug {
			return v, nil
		}
	}
	return Adventure{}, fmt.Errorf("Adventure not found")
}

// GetAllSlugs doc
func (ar AdventureRepository) GetAllSlugs() []string {
	slugs := make([]string, len(ar.Adventures))
	for i, v := range ar.Adventures {
		slugs[i] = v.Slug
	}
	return slugs
}

func initData(fileName string) ([]Adventure, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	result := make(map[string]Adventure)
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	i := 0
	adventures := make([]Adventure, len(result))
	for key, value := range result {
		value.Slug = key
		adventures[i] = value
		i++
	}

	return adventures, nil
}
