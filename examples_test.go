package jar

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2020 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleReadFile() {
	manifest, err := ReadFile("wagon-file-3.3.2.jar")

	fmt.Printf("Error: %v\n", err)
	fmt.Printf("Specification-Vendor: %s\n", manifest["Specification-Vendor"])
	fmt.Printf("Implementation-Vendor-Id: %s\n", manifest["Implementation-Vendor-Id"])
	fmt.Printf("Specification-Version: %s\n", manifest["Specification-Version"])
}
