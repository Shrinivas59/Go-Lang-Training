# DAY 3 – DATA STRUCTURES & POINTERS IN GO


# Introduction

Today, we will learn how Go stores collections of data and how pointers help us access and modify values in memory.

Topics covered:

Arrays and Slices
Maps
Range
Pointers
Practical Project – Word Frequency Counter

# 1. ARRAYS IN GO

An array is a collection of elements of the same data type. Its size is fixed when it is created.

Example:

```
package main

import "fmt"

func main() {
    var numbers [5]int

    numbers[0] = 10
    numbers[1] = 20
    numbers[2] = 30
    numbers[3] = 40
    numbers[4] = 50

    fmt.Println("Numbers:", numbers)
    fmt.Println("Length:", len(numbers))
}
```

Output:

```
Numbers: [10 20 30 40 50]
Length: 5
```

Important Points
[5]int means an integer array containing 5 elements.
Array size cannot be changed after declaration.
Array indexes start from 0.
len() returns the number of elements.


# 2. SLICES IN GO

A slice is similar to an array, but its size can grow when we add more elements.

Slices are commonly used when working with collections in Go.

Example:

```
package main

import "fmt"

func main() {
    fruits := []string{"Apple", "Banana", "Mango"}

    fmt.Println("Original:", fruits)

    fruits = append(fruits, "Orange")

    fmt.Println("After append:", fruits)
    fmt.Println("Length:", len(fruits))
    fmt.Println("Capacity:", cap(fruits))
}
```

Output:

```
Original: [Apple Banana Mango]
After append: [Apple Banana Mango Orange]
Length: 4
Capacity: 6
```


Note: Capacity may vary depending on Go's allocation strategy.

# Important Slice Concepts

# A. make() – Create a slice

```
numbers := make([]int, 3, 5)

fmt.Println(numbers)
fmt.Println("Length:", len(numbers))
fmt.Println("Capacity:", cap(numbers))

```


Output:

```
[0 0 0]
Length: 3
Capacity: 5

```


# B. append() – Add elements

```
numbers := []int{10, 20, 30}
numbers = append(numbers, 40)

fmt.Println(numbers)

```


Output:

```
[10 20 30 40]

```

# C. Slicing syntax – Get part of a slice

```
numbers := []int{10, 20, 30, 40, 50}

fmt.Println(numbers[1:4])

```
Output:

```
[20 30 40]

````
Here:

1 is the starting index (included).
4 is the ending index (excluded).
len() vs cap()
Function	Meaning
len()	Number of elements currently in the slice
cap()	Available capacity before a larger underlying array is needed

Remember: Array = Fixed size. Slice = Flexible size.

# 3. MAPS IN GO

A map stores data in key-value pairs.

For example:

Name → John
Age → 25
City → Chennai

Maps help us find values easily using their keys.

Example
```
package main

import "fmt"

func main() {
    student := make(map[string]int)

    student["John"] = 90
    student["David"] = 85
    student["Sam"] = 95

    fmt.Println("Student Marks:", student)
    fmt.Println("John's Marks:", student["John"])
}
```
Output

```
Student Marks: map[David:85 John:90 Sam:95]
John's Marks: 90
```

# Adding, Updating and Deleting Values
```
student := map[string]int{
    "John": 90,
    "David": 85,
}

student["Sam"] = 95
student["John"] = 99

delete(student, "David")

fmt.Println(student)

```
Output:

```
map[John:99 Sam:95]

```

# Comma OK Idiom

Used to check whether a key exists in a map.

```
package main

import "fmt"

func main() {
    student := map[string]int{
        "John": 90,
        "Sam": 95,
    }

    marks, ok := student["John"]

    if ok {
        fmt.Println("Student found:", marks)
    } else {
        fmt.Println("Student not found")
    }
}
```


Output:

```
Student found: 90

```
Remember:

marks stores the value.
ok is true if the key exists.
ok is false if the key does not exist.


# 4. RANGE IN GO

range is used to iterate (loop) through collections like arrays, slices, maps and strings.

# A. Range with slices

```
fruits := []string{"Apple", "Banana", "Mango"}

for index, value := range fruits {
    fmt.Println(index, value)
}

```


Output:

```
0 Apple
1 Banana
2 Mango

```

# B. Range with maps
```
student := map[string]int{
    "John": 90,
    "Sam": 95,
}

for key, value := range student {
    fmt.Println(key, value)
}

```


Output (order may vary):

```
John 90
Sam 95
````


Note: Map iteration order is not guaranteed in Go.

# C. Range with strings
```
name := "GoLang"

for index, value := range name {
    fmt.Println(index, string(value))
}
```

Output:

```
0 G
1 o
2 L
3 a
4 n
5 g

```


Remember: With strings, range returns the byte index and the Unicode character (rune).

# 5. POINTERS IN GO

A pointer is a variable that stores the memory address of another variable.

Pointers allow us to access and modify the original value without passing a copy of that value.

# Important Symbols

& – Gets the memory address of a variable.
* – Accesses or modifies the value stored at that address (dereferencing).

Example:
```
package main

import "fmt"

func main() {
    num := 10

    ptr := &num

    fmt.Println("Original Value:", num)
    fmt.Println("Memory Address:", ptr)
    fmt.Println("Pointer Value:", *ptr)

    *ptr = 50

    fmt.Println("Updated Value:", num)
}
```

Output:

```
Original Value: 10
Memory Address: 0x...
Pointer Value: 10
Updated Value: 50   

```
The actual memory address will be different each time.

How Pointers Work

num := 10 creates a variable.
ptr := &num stores its memory address.
*ptr accesses the original value.
*ptr = 50 changes the original value to 50.
Pointers with Functions

Pointers are useful when we want a function to modify the original variable.

```
package main

import "fmt"

func update(num *int) {
    *num = 100
}

func main() {
    value := 10

    update(&value)

    fmt.Println("Updated value:", value)
}
```

Output:

```
Updated value: 100
```

Explanation:

&value passes the memory address.
num *int receives an integer pointer.
*num = 100 modifies the original value.


# 6. Day 3 Practical Project: Word Frequency Counter


Now let's combine everything we learned into one practical project.

Goal: Read a paragraph and count how many times each word appears.

For example:

Input:

```
Go is simple and Go is powerful
```


Expected output:

```
Go: 2
is: 2
simple: 1
and: 1
powerful: 1

```

# Complete Go program

```


package main

import (
    "fmt"
    "strings"
)

func main() {

    // Input paragraph
    paragraph := "Go is simple and Go is powerful"

    // Convert paragraph into words
    words := strings.Fields(paragraph)

    // Create a map to store word counts
    wordCount := make(map[string]int)

    // Iterate through words
    for _, word := range words {

        // Increase word count
        wordCount[word]++
    }

    // Display word frequencies
    fmt.Println("Word Frequency:")

    for word, count := range wordCount {
        fmt.Println(word, ":", count)
    }
}


```

Output:

```
Word Frequency:
Go : 2
is : 2
simple : 1
and : 1
powerful : 1

```
