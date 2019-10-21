package main

import (
	"encoding/base64"
	"fmt"
)

func main()  {
	// Current version only encodes first word if there are spaces after.
	// Encodes full string without spaces
	var data string
	a := []byte(data)

	fmt.Print("Enter String to Encode: ") 
	fmt.Scanf("%v", &a)


	sEnc := base64.StdEncoding.EncodeToString(a)

	fmt.Print(sEnc)



}