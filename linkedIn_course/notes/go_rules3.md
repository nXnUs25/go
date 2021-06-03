# Go Standard Library

## Overview of standard library in go

Using those packages:

* fmt
* Strings
  * strings
  * strconv
  * unicode
* math
  * math
  * random
* files
  * io/utils
  * os
* networking
  * url
  * http
  * encoding/json
  * encoding/xml

### Formatting with `Fmt` package

#### Simple string output `Print` and `Println`

```go
// TODO: Print a simple string, no newline
fmt.Print("welcome go")
// TODO: Print with a newline
fmt.Println("Welcome to go")

// TODO: Print a string with values
const answer = 42
fmt.Println("The answer is:", answer)

// TODO: Print a string with multiple interspersed values
const a, b, c = 5, 5, 10
fmt.Println("Add", a, "and", b, "equal", c)

// TODO: print a slice of data
items := []int{10, 20, 40, 80}
length, err := fmt.Println(items)
fmt.Println(length, err)
```

```bash
#Output
welcome goWelcome to go
The answer is: 42
Add 5 and 5 equal 10
[10 20 40 80]
14 <nil>
```

#### Advanced formatting with verbs `Printf` and `Sprintf`

```go
 x := 20
 f := 123.45

 // TODO: Basic formatting
 fmt.Printf("%d\n", x)
 fmt.Printf("%x\n", x)

 // TODO: Booleans can be printed as "true" or "false"
 fmt.Printf("%t\n", x > 10)

 // TODO: floating point numbers
 fmt.Printf("%f\n", f)
 fmt.Printf("%e\n", f)

 // TODO: Using explicit argument indexes
 fmt.Printf("%[2]d %[1]d\n", 52, 40)

 // TODO: Argument indexes can be used to print values repeatedly
 fmt.Printf("%d %#[1]o %#[1]x\n", 52)

 // TODO: Print a value in default format
    c := circle{
   radius: 20,
   border: 5,
 }
 fmt.Printf("%v\n", c)
 fmt.Printf("%+v\n", c)
 fmt.Printf("%T\n", c)

 // TODO: Sprintf is the same as Printf, but returns a string
 s := fmt.Sprintf("%[2]d %[1]d\n", 52, 40)
 fmt.Println(s)
```

```bash
# Output
20
14
true
123.450000
1.234500e+02
40 52
52 064 0x34
{20 5}
{radius:20 border:5}
main.circle
40 52
```

#### Numerical formatting and precision `Printf`

```go
f := 123.4567

// TODO: control the decimal precision
fmt.Printf("%.2f\n", f)

// TODO: print with width 10 and default precision
fmt.Printf("%10f\n", f)

// TODO: print with padding and precision
fmt.Printf("%10.3f\n", f)

// TODO: always use a + sign
fmt.Printf("%+10.2f\n", f)

// TODO: pad with 0s instead of spaces
fmt.Printf("%010.4f\n", f)
```

```bash
# Output
123.46
123.456700
   123.457
   +123.46
00123.4567
```

#### Reading standard input from the command line `Stdin` `bufio`

```go
reader := bufio.NewReader(os.Stdin)

fmt.Print("Insert some text: ")
s, _ := reader.ReadString('\n')
fmt.Println(s)
```

```bash
#Output
Insert some text: test
test
```

### Working on Strings

#### Basic String Operations

```go
s := "The quick brown fox jumps over the lazy dog"
 w := "small"
 // Basic string operations

 // TODO: Length of string
 fmt.Println(len(s))

 // TODO: iterate over each character
 for _, ch := range s {
  fmt.Print(string(ch), " ")
 }
 fmt.Println()

 // TODO: iterate over each character without converting them to string
 for index, ch := range w {
  fmt.Printf("index: <%[1]v> %[1]T -- ch: <%[2]v> %[2]T \n", index, ch)
 }
 fmt.Println()

 // TODO: Using operators < > == !=
 fmt.Println("dog" < "cat")
 fmt.Println("dog" < "horse")
 fmt.Println("dog" == "Dog")

 // TODO: Comparing two strings
 result := strings.Compare("dog", "cat")
 fmt.Println(result)
 result = strings.Compare("dog", "dog")
 fmt.Println(result)

 // TODO: EqualFold tests using Unicode case-folding
 fmt.Println(strings.EqualFold("Hello", "hello"))

 // TODO: ToUpper, ToLower, Title
 s1 := strings.ToUpper(s)
 s2 := strings.ToLower(s)
 s3 := strings.Title(s)
 fmt.Println(s1)
 fmt.Println(s2)
 fmt.Println(s3)
```

```bash
43
T h e   q u i c k   b r o w n   f o x   j u m p s   o v e r   t h e   l a z y   d o g
index: <0> int -- ch: <115> int32
index: <1> int -- ch: <109> int32
index: <2> int -- ch: <97> int32
index: <3> int -- ch: <108> int32
index: <4> int -- ch: <108> int32

false
true
false
1
0
true
THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG
the quick brown fox jumps over the lazy dog
The Quick Brown Fox Jumps Over The Lazy Dog
```

#### String searching

```go
 fname := "filename.txt"
 fname2 := "temp_picfile.jpeg"
 vowels := "aeiouAEIOU"
 s := "The quick brown fox jumps over the lazy dog"

 // Common string searching functions

 // Use Contains to see if a substring is in a string
 fmt.Println(strings.Contains(s, "jump"))
 // Use ContainsAny to see if any of the given chars are in the string
 fmt.Println(strings.ContainsAny(s, "abcd"))

 // Use Index to find the offset of the first instance of a substring
 fmt.Println(strings.Index(s, "fox"))
 fmt.Println(strings.Index(s, "cat"))
 // Use IndexAny to find the first instance of any of the given chars
 fmt.Println(strings.IndexAny("grzbl", vowels))
 fmt.Println(strings.IndexAny("Golang!", vowels))

 // HasPrefix and HasSuffix can be used to see if a string starts with
 // or ends with a specific substring
 fmt.Println(strings.HasSuffix(fname, "txt"))
 fmt.Println(strings.HasPrefix(fname2, "temp"))

 // Count returns the number of non-overlapping instances of a substring
 fmt.Println(strings.Count(s, "the"))
 fmt.Println(strings.Count(s, "he"))
```

```bash
#Output
true
true
16
-1
-1
1
true
true
1
2
```

#### String Manipulation

#### Using Map function

```go
// The map function returns a copy of a string with the characters modified
	// according to the mapping function
	shift := 2
	s := "The quick brown fox jumps over the lazy dog"

	// TODO: create the mapping function
	transform := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			value := int('A') + (int(r) - int('A') + shift)
			value = calculateValue(value, 91, 65)
			return rune(value)
		case r >= 'a' && r <= 'z':
			value := int('a') + (int(r) - int('a') + shift)
			value = calculateValue(value, 122, 97)
			return rune(value)
		}
		return r
	}

	// TODO: Encode the message
	encode := strings.Map(transform, s)
	fmt.Println(encode)

	// TODO: Decode the message
	shift = -shift
	decode := strings.Map(transform, encode)
	fmt.Println(decode)
}

func calculateValue(value, asciUpper, asciLower int) int {
	if value > asciUpper { // 91 is the limit of upper cases
		value -= 26
	} else if value < asciLower {
		value += 26
	}
	return value
}
```

```bash
#output
Vjg swkem dtqyp hqz lworu qxgt vjg ncba fqi
The quick brown fox jumps over the lazy dog
```

#### Using string builder

```go
 // TODO: Declare an empty strings.Builder
 var sb strings.Builder

 // TODO: Write some content
 sb.WriteString("This is string 1\n")
 sb.WriteString("This is string 2\n")
 sb.WriteString("This is string 3\n")

 // TODO: Output the concatenated string
 fmt.Println(sb.String())

 // TODO: Examine the builder's capacity
 fmt.Println("Capacity:", sb.Cap())

 // TODO:
 // Grow the capacity - use this if you know in advance how much data
 // you're going to be writing into the buffer to minimize copies
 for i := 0; i <= 10; i++ {
  fmt.Fprintf(&sb, "String %d -- ", i)
 }
 fmt.Println(sb.String())

 // TODO: we can get the length of what the final string will be
 fmt.Println("Builder size is", sb.Len())

 // TODO: The Reset function will reset the builder to original state
 sb.Reset()
 fmt.Println("After Reset:")
 fmt.Println("Capacity:", sb.Cap())
 fmt.Println("Builder size is", sb.Len())
```

```bash
#output
This is string 1
This is string 2
This is string 3

Capacity: 64
This is string 1
This is string 2
This is string 3
String 0 -- String 1 -- String 2 -- String 3 -- String 4 -- String 5 -- String 6 -- String 7 -- String 8 -- String 9 -- String 10 --
Builder size is 184
After Reset:
Capacity: 0
Builder size is 0
```

#### Parsing a string with strconv

```go
ampleint := 100
	samplestr := "250"

	// This does a character conversion, not a numerical one
	fmt.Println("\n// This does a character conversion, not a numerical one")
	newstr := string(sampleint)
	fmt.Println("Int is: :", sampleint)
	fmt.Println("Result of using string():", newstr)

	// The strconv package contains a variety of functions for parsing and formatting
	// numbers, values, and strings
	fmt.Println("\n// The strconv package contains a variety of functions for parsing and formatting")
	fmt.Println("// numbers, values, and strings")

	// Convert an integer to string
	fmt.Println("\n// Convert an integer to string")
	s := strconv.Itoa(sampleint)
	fmt.Printf("%T, %v\n", s, s)

	// Convert a string to integer
	fmt.Println("\n// Convert a string to integer")
	x, err := strconv.Atoi(samplestr)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%T, %v\n", x, x)

	// Other parse functions
	fmt.Println("\n// Other parse functions")
	b, _ := strconv.ParseBool("true")
	fmt.Println("strconv.ParseBool(\"true\")", b)
	f, _ := strconv.ParseFloat("3.14159", 64)
	fmt.Println("strconv.ParseFloat(\"3.14159\", 64)", f)
	i, _ := strconv.ParseInt("-42", 10, 64)
	fmt.Println("strconv.ParseInt(\"-42\", 10, 64)", i)
	u, _ := strconv.ParseUint("42", 10, 64)
	fmt.Println("strconv.ParseUint(\"42\", 10, 64)", u)

	// Other format functions
	fmt.Println("\n// Other format functions")
	s = strconv.FormatBool(true)
	fmt.Println("strconv.FormatBool(true)", s)
	s = strconv.FormatFloat(3.14159, 'E', -1, 64)
	fmt.Println("strconv.FormatFloat(3.14159, 'E', -1, 64)", s)
	s = strconv.FormatInt(-42, 10)
	fmt.Println("strconv.FormatInt(-42, 10)", s)
	s = strconv.FormatUint(42, 10)
	fmt.Println("strconv.FormatUint(42, 10)", s)
```

```bash
# output
// This does a character conversion, not a numerical one
Int is: : 100
Result of using string(): d

// The strconv package contains a variety of functions for parsing and formatting
// numbers, values, and strings

// Convert an integer to string
string, 100

// Convert a string to integer
int, 250

// Other parse functions
strconv.ParseBool("true") true
strconv.ParseFloat("3.14159", 64) 3.14159
strconv.ParseInt("-42", 10, 64) -42
strconv.ParseUint("42", 10, 64) 42

// Other format functions
strconv.FormatBool(true) true
strconv.FormatFloat(3.14159, 'E', -1, 64) 3.14159E+00
strconv.FormatInt(-42, 10) -42
strconv.FormatUint(42, 10) 42
```

#### String tests with Unicode

```go
// Declare a sample string
	fmt.Println("// Declare a sample string")
	const s = "The 'quick' brown fox, jumped over the *LAZY* dog!"
	fmt.Println("Constant s:", s)

	// Init some count variables
	t := `// Init some count variables
	punctCount := 0
	lowerCount, upperCount := 0, 0
	spaceCount := 0
	hexdigitCount := 0`
	fmt.Println("\n", t)
	punctCount := 0
	lowerCount, upperCount := 0, 0
	spaceCount := 0
	hexdigitCount := 0

	// iterate over the characters and call unicode function tests
	fmt.Println("\n", "// iterate over the characters and call unicode function tests")
	for _, ch := range s {
		if unicode.IsPunct(ch) {
			punctCount++
		}
		if unicode.IsLower(ch) {
			lowerCount++
		}
		if unicode.IsUpper(ch) {
			upperCount++
		}
		if unicode.IsSpace(ch) {
			spaceCount++
		}
		if unicode.Is(unicode.Hex_Digit, ch) {
			hexdigitCount++
		}
	}

	t = `// Print the results
	fmt.Println("Punctuation:", punctCount)
	fmt.Println("Lowercase:", lowerCount)
	fmt.Println("Uppercase:", upperCount)
	fmt.Println("Whitespace:", spaceCount)
	fmt.Println("Hex Digits:", hexdigitCount)`
	// Print the results
	fmt.Println("\n", t)
	fmt.Println("Punctuation:", punctCount)
	fmt.Println("Lowercase:", lowerCount)
	fmt.Println("Uppercase:", upperCount)
	fmt.Println("Whitespace:", spaceCount)
	fmt.Println("Hex Digits:", hexdigitCount)
  ```

  ```bash
  # output
  // Declare a sample string
  Constant s: The 'quick' brown fox, jumped over the *LAZY* dog!

  // Init some count variables
	punctCount := 0
	lowerCount, upperCount := 0, 0
	spaceCount := 0
	hexdigitCount := 0

  // iterate over the characters and call unicode function tests

  // Print the results
	fmt.Println("Punctuation:", punctCount)
	fmt.Println("Lowercase:", lowerCount)
	fmt.Println("Uppercase:", upperCount)
	fmt.Println("Whitespace:", spaceCount)
	fmt.Println("Hex Digits:", hexdigitCount)
Punctuation: 6
Lowercase: 31
Uppercase: 5
Whitespace: 8
Hex Digits: 10
```

### Mathematical Operations

### Files and directories

### Networking
