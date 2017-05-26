// Example: Design an algorithm to print all permutations
// of a string. For simplicity, assume all characters are unique.

package main

import (
    "fmt"
)

func PrintPermutations(str string) {
    var perms []string
    for index, char := range str {
        if index == 0 {
            perms = append(perms, string(char))
        } else {
            perms = addCharToStrings(string(char), perms)
        }
    }
    fmt.Println(perms)
}

func addCharToStrings(char string, perms []string) []string {
    var newPerms []string
    for _, perm := range perms {
        for i := 0; i <= len(perm); i++ {
            newPerms = append(newPerms, perm[:i] + char + perm[i:])
        }
    }
    return newPerms
}

func main() {
    PrintPermutations("abc")
    PrintPermutations("duck")
    PrintPermutations("go")
    PrintPermutations("python")
}
