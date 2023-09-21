# fileio
read and write to files in a simpler way

Instead of reading and writing file contents as []byte, for cirtain use cases, reading and writing as a string is a simpler approach

here is an example of this in action.

test.txt
```
hello from text file!

```

list.text
```
-eat
-sleep
-school
-study
-code

```

main.go
```
package main

import (
	"fmt"

	"github.com/KingCosma/fileio"
)

func main() {
	x, err := fileio.ReadString("test.txt") //store the string from test.txt in a variable
	if err != nil {                         //error handling
		panic(err)
	}
	fmt.Println(x) // print x

	fileio.WriteString("test.txt", "This text file is changed") //overwrite the contents

	x, err = fileio.ReadString("test.txt") //change the value of x to the updated file contents
	if err != nil {                        //error handling again :)
		panic(err)
	}
	fmt.Println(x)                                     // print x
	arr, err1 := fileio.ReadAndSplit("list.txt", "\n") //read the contents list.txt as an array seperated by "\n" or newline
	if err1 != nil {                                   //error handeling :)
		panic(err1)
	}
	for _, item := range arr { //for loop to iterate over the array
		fmt.Println(item) //print each item in a newline
	}
}

```
