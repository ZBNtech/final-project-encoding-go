package encoding

import (
	"encoding/json"
	"fmt"
	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"os"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	// ниже реализуйте метод
	// ...
	jsonFile, err := os.ReadFile(j.FileInput)
	if err != nil {
		fmt.Printf("ошибка при чтении %s файла: %s", j.FileInput, err.Error())
		return nil

	}

	err = json.Unmarshal(jsonFile, &j.DockerCompose)

	if err != nil {
		fmt.Printf("ошибка при десериализации %s файла: %s", j.FileInput, err.Error())
		return nil
	}

	yamlFile, err := os.Open(j.FileOutput)
	if err != nil {
		fmt.Printf("ошибка при открытии файла %s: %s", j.FileOutput, err.Error())
		fmt.Printf("создаем файл %s", j.FileOutput)
		yamlFile, err = os.Create(j.FileOutput)
		if err != nil {
			fmt.Printf("ошибка при создании файла: %s: %s", j.FileOutput, err.Error())
			return nil
		}
	}

	yamlData, err := yaml.Marshal(&j.DockerCompose)
	if err != nil {
		fmt.Printf("ошибка при сериализации в yaml: %s", err.Error())
		return nil
	}

	_, err = yamlFile.Write(yamlData)
	if err != nil {
		fmt.Printf("ошибка при записи данных в файл: %s: %s", j.FileOutput, err.Error())
		return nil
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	// Ниже реализуйте метод
	// ...
	yamlFile, err := os.ReadFile(y.FileInput)
	if err != nil {
		fmt.Printf("ошибка при чтении %s файла: %s", y.FileInput, err.Error())
		return nil
	}

	err = yaml.Unmarshal(yamlFile, &y.DockerCompose)
	if err != nil {
		fmt.Printf("ошибка при десериализации %s файла: %s", y.FileInput, err.Error())
		return nil
	}

	jsonFile, err := os.Open(y.FileOutput)
	if err != nil {
		fmt.Printf("ошибка при открытии файла %s: %s", y.FileOutput, err.Error())
		fmt.Printf("создаем файл %s", y.FileOutput)
		jsonFile, err = os.Create(y.FileOutput)
		if err != nil {
			fmt.Printf("ошибка при создании файла: %s: %s", y.FileOutput, err.Error())
			return nil
		}
	}

	defer jsonFile.Close()

	jsonData, err := json.Marshal(&y.DockerCompose)
	if err != nil {
		fmt.Printf("ошибка при сериализации в json: %s", err.Error())
		return nil
	}

	_, err = jsonFile.Write(jsonData)
	if err != nil {
		fmt.Printf("ошибка при записи данных в файл: %s: %s", y.FileOutput, err.Error())
		return nil
	}

	return nil
}
