package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Name")
	name, _ := reader.ReadString('\n')

	fmt.Println("address")
	address, _ := reader.ReadString('\n')

	name = strings.TrimSpace(name)
	address = strings.TrimSpace(address)

	userAddressMap := make(map[string]string)

	userAddressMap["name"] = name
	userAddressMap["address"] = address

	jsonFile, err := json.Marshal(userAddressMap)
	if err != nil {
		fmt.Println("bad input")
	}

	fmt.Println("json file", string(jsonFile))
}
