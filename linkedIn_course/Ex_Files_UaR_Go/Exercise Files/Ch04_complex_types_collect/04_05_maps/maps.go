package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)
	fmt.Println(states)

	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"
	fmt.Println(states)
	fmt.Printf("Type: <%T> Values: %v\n", states, states)

	california := states["CA"]
	fmt.Printf("Type: <%T> Values: %v\n", california, california)

	delete(states, "OR")
	fmt.Printf("Type: <%T> Values: %v\n", states, states)

	states["OR"] = "Oregon"
	states["NY"] = "New York"

	fmt.Printf("Type: <%T> Values: %v\n", states, states)

	for k, v := range states {
	fmt.Printf("Type: <%T> Key %v -- Values: %v\n", states, k, v)
	}

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	for x := range keys {
		fmt.Printf("%v: %v\n", keys[x], states[keys[x]])
	}

	fmt.Printf("+v :: %+v\n", states)
	fmt.Printf("v  :: %v\n", states)
}