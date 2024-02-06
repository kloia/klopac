package helper

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"entrypoint/pkg/command"
	"entrypoint/pkg/flag"
	"entrypoint/pkg/flow"
	"entrypoint/pkg/logger"
	"entrypoint/pkg/option"
	"entrypoint/pkg/shell"
	"fmt"

	"github.com/imdario/mergo"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"

	"io"
	"os"
	"path/filepath"
)

// Reads content of the yaml file and returns it
func ReadFile(filename string) map[string]interface{} {
	log := logger.GetLogger()
	log.Debug("START: READ FILE")
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		log.Debug("Could not open", zap.String("filename", filename))
		return nil
	}

	m := make(map[string]interface{})
	err = yaml.Unmarshal(yamlFile, &m)
	if err != nil {
		log.Debug("Could not decode map", zap.Any("map", m))
	}
	log.Debug("END: READ FILE")
	return m
}

// Writes content to a yaml file

func WriteFile(filename string, data map[string]interface{}, keepHeader bool) error {
	log := logger.GetLogger()
	log.Debug("START: WRITE FILE", zap.String("filename", filename))

	var b bytes.Buffer
	if keepHeader {
		// Başlangıçta '---' işaretlerini ekleyin
		b.WriteString("---\n")
	}

	// YAML içeriğini oluştur
	yamlEncoder := yaml.NewEncoder(&b)
	yamlEncoder.SetIndent(2)
	if err := yamlEncoder.Encode(data); err != nil {
		log.Error("Failed to encode YAML", zap.Error(err))
		return err
	}
	yamlEncoder.Close()

	// Dosyaya yaz
	err := os.WriteFile(filename, b.Bytes(), 0644) // Dosya izinleri 0644 olarak ayarlanır
	if err != nil {
		log.Error("Failed to write file", zap.Error(err), zap.String("filename", filename))
	}

	log.Debug("END: WRITE FILE", zap.String("filename", filename))
	return err
}

// Basically we have two map and we compare them if there are some sort of values that should be changed according to its logic
func Intersection(inputMap, defaultMap map[string]interface{}) (newMap map[string]interface{}) {
	newMap = make(map[string]interface{})
	for inputKey, inputVal := range inputMap {
		if inputInnerMap, ok := inputVal.(map[string]interface{}); ok {
			defaultInnerMap, ok := defaultMap[inputKey].(map[string]interface{})
			if ok {
				newMap[inputKey] = Intersection(inputInnerMap, defaultInnerMap)
			}
		} else {
			fmt.Println("defaultMapinputkey together:", defaultMap[inputKey], inputKey)
			_, ok := defaultMap[inputKey]
			if ok {
				newMap[inputKey] = inputVal
				fmt.Println("newMap[inputKey]:", newMap[inputKey])
			}
		}
	}
	return newMap
}

// Basically takes a interface and varsPath(which is path of the variable files) then it starts to override or leaves unchanged depending to intersection logic
func UpdateValuesFile(valuesModel map[string]interface{}, varsPath string) error {
	log := logger.GetLogger()
	return filepath.Walk(varsPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && filepath.Ext(path) == ".yaml" || filepath.Ext(path) == ".yml" {
				defaultModel := ReadFile(path)
				fmt.Println("defaultModel path:", path)
				intersectionMap := Intersection(valuesModel, defaultModel)
				if isMapNotEmpty(intersectionMap) { // this checks whether intersection map is truly empty or not
					log.Debug("INTERSECTION", zap.Any("map", intersectionMap))
					mergo.Merge(&defaultModel, intersectionMap, mergo.WithOverride)
					WriteFile(path, defaultModel, true)
				}
			}
			return nil
		})
}

// isMapNotEmpty, checks if the map is empty or not
func isMapNotEmpty(m map[string]interface{}) bool {
	for _, value := range m {
		switch v := value.(type) {
		case map[string]interface{}:
			if isMapNotEmpty(v) {
				return true // there is at least one non-empty map
			}
		default:
			return true // there is at least one non-empty value
		}
	}
	return false // map is empty
}

func Untar(tarball, target string) error {
	reader, err := os.Open(tarball)
	log := logger.GetLogger()
	if err != nil {
		log.Debug("Bundle file could not be opened.")
		return err
	}
	defer reader.Close()

	gzf, err := gzip.NewReader(reader)
	if err != nil {
		return err
	}
	defer gzf.Close()

	tarReader := tar.NewReader(gzf)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		path := filepath.Join(target, header.Name)
		info := header.FileInfo()
		if info.IsDir() {
			if err := os.MkdirAll(path, info.Mode()); err != nil {
				return err
			}
			continue
		}

		file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
		if err != nil {
			return err
		}

		if _, err := io.Copy(file, tarReader); err != nil {
			file.Close() // Ensure file is closed in case of error
			return err
		}
		file.Close() // Moved from defer to close immediately after use
	}
	return nil
}

var (
	optionService = option.NewOptionService(flag.NewFlagService())
	flowService   = flow.NewFlowService(shell.NewShellService(command.NewCommandService()))
)

func GetOptionService() *option.OptionService {
	return optionService
}

func GetFlowService() flow.Flow {
	return flowService
}

func GetParam[V comparable](param string) V {
	return *GetOptionService().Params[param].(*V)
}
