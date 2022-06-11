package tracer

import (
	"encoding/json"
	"fmt"
)

func PrintJson(data []byte, name string) error {
	value, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	fmt.Printf("======================= %s =======================\n", name)
	fmt.Println(value)
	fmt.Println("==============================================")
	return nil
}
