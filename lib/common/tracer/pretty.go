package tracer

import (
	"encoding/json"
	"fmt"
)

func PrintJson(data interface{}, name string) error {
	value, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	// Marshall the Colorized JSON

	fmt.Printf("======================= %s =======================\n", name)
	fmt.Println(value)
	fmt.Println("==============================================")
	return nil
}
