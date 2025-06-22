// Naranza Cogo, Copyright 2025 Andrea Davanzo and contributors, AGPLv3

package cogo

import (
  "bufio"
  "errors"
  "os"
  "reflect"
  "strconv"
  "strings"
)

const Version = "1.2025.2"

func LoadConfig(filename string, out any) error {
  ptrVal := reflect.ValueOf(out)
  if ptrVal.Kind() != reflect.Ptr || ptrVal.IsNil() {
    return errors.New("out must be a non-nil pointer to a struct")
  }
  structVal := ptrVal.Elem()
  if structVal.Kind() != reflect.Struct {
    return errors.New("out must point to a struct")
  }

  file, err := os.Open(filename)
  if err != nil {
    return err
  }
  defer file.Close()

  fieldMap := make(map[string]reflect.Value)
  structType := structVal.Type()
  for i := 0; i < structType.NumField(); i++ {
    field := structType.Field(i)
    fieldMap[field.Name] = structVal.Field(i)
  }

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := strings.TrimSpace(scanner.Text())
    if line != "" && !strings.HasPrefix(line, "#") {

      parts := strings.Fields(line)
      if len(parts) < 3 {
        return errors.New("Invalid config line (less than 3 parts): " + line)
      }

      configKey := parts[0]
      configType := strings.ToLower(parts[1])
      configValue := strings.Join(parts[2:], " ")
      
      field, ok := fieldMap[parts[0]]
      if !ok || !field.CanSet() {
        return errors.New("cannot set field: " + configKey)
      }

      switch configType {
        case "int":
          i, err := strconv.Atoi(configValue)
          if err != nil {
            return errors.New("Invalid int value for key " + configKey)
          }
          field.SetInt(int64(i))
        case "bool":
          b, err := strconv.ParseBool(configValue)
          if err != nil {
            return errors.New("Invalid bool value for key " + configKey)
          }
          field.SetBool(b)
        case "float", "float64":
          f, err := strconv.ParseFloat(configValue, 64)
          if err != nil {
            return errors.New("Invalid float value for key " + configKey)
          }
          field.SetFloat(f)
        case "filemode":
          u, err := strconv.ParseUint(configValue, 8, 32)
          if err != nil {
            return errors.New("Invalid filemode value for key " + configKey)
          }
          field.SetUint(u)
        case "string":
          field.SetString(configValue)
        default:
          return errors.New("unknown config type: " + configType)
      }
    }
  }
  return scanner.Err()
}
