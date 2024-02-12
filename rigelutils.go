package rigel

import (
	"fmt"

	"github.com/remiges-tech/rigel/types"
)

// GetSchemaPath constructs the base key for a schema in etcd based on the provided appName, moduleName and version.
func GetSchemaPath(appName string, moduleName string, version int) string {
	return fmt.Sprintf("%s/%s/%s/%d/", rigelPrefix, appName, moduleName, version)
}

// GetSchemaFieldsPath constructs the path for a schema based on the provided appName, moduleName and version.
func GetSchemaFieldsPath(appName string, moduleName string, version int) string {
	return fmt.Sprintf("%s/%s/%s/%d/fields", rigelPrefix, appName, moduleName, version)
}

// getSchemaDescriptionPath constructs the path for a schema based on the provided appName, moduleName and version.
func GetSchemaDescriptionPath(appName string, moduleName string, version int) string {
	return fmt.Sprintf("%s/%s/%s/%d/description", rigelPrefix, appName, moduleName, version)
}

// GetConfPath constructs the path for a configuration based on the provided appName, moduleName and version.
func GetConfPath(appName string, moduleName string, version int, namedConfig string) string {
	return fmt.Sprintf("%s/%s/%s/%d/config/%s", rigelPrefix, appName, moduleName, version, namedConfig)
}

// GetConfKeyPath constructs the path for a configuration based on the provided appName, moduleName, version, namedConfig, and confKey.
func GetConfKeyPath(appName string, moduleName string, version int, namedConfig string, confKey string) string {
	// return fmt.Sprintf("%s/%s/%s/%d/config/%s/%s", rigelPrefix, appName, moduleName, version, namedConfig, confKey)
	return fmt.Sprintf("/remiges/rigel/%s/%s/%d/config/%s/keys/%s", appName, moduleName, version, namedConfig, confKey)

}

func ValidateValueAgainstConstraints(value string, field *types.Field) bool {
	// Convert the value to the correct type
	val, err := convertToType(value, field.Type)
	if err != nil {
		return false
	}

	// Check the constraints
	if field.Constraints != nil {
		if field.Constraints.Min != nil {
			switch field.Type {
			case "int":
				if val.(int) < *field.Constraints.Min {
					return false
				}
			case "string":
				if len(val.(string)) < *field.Constraints.Min {
					return false
				}
			case "float":
				if val.(float64) < float64(*field.Constraints.Min) {
					return false
				}
			}
		}
		if field.Constraints.Max != nil {
			switch field.Type {
			case "int":
				if val.(int) > *field.Constraints.Max {
					return false
				}
			case "string":
				if len(val.(string)) > *field.Constraints.Max {
					return false
				}
			case "float":
				if val.(float64) > float64(*field.Constraints.Max) {
					return false
				}
			}
		}
		if field.Constraints.Enum != nil {
			if field.Type == "string" {
				found := false
				for _, v := range field.Constraints.Enum {
					if v == val.(string) {
						found = true
						break
					}
				}
				if !found {
					return false
				}
			}
		}
	}

	return true
}
