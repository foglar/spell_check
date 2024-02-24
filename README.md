# Spellchecker

Simple spellcheck package in go

```shell
git clone https://github.com/foglar/spell_check.git
```

## Example usage

There are already bundled some wordlist, `words.txt` **(EN dictionary)**, `words_alpha.txt` **(EN dictinary alphabet)** and `czech.txt` **(CZ dictionary)**.

```go
package main

import (
 "fmt"
 "github.com/foglar/spell_check/"
 "log"
)

func main() {
 sc, err := spellchecker.NewSpellChecker("./words.txt")
 if err != nil {
  log.Fatalf("Error creating SpellChecker: %v", err)
 }

 word := "exprezzion"
 closestWords := sc.Check(word, 10)

 fmt.Printf("Closest words to '%s': %v\n", word, closestWords)
}
```

## Performace

- wordlists are ordered alphabeticly
- CPU: AMD Ryzen 7 5825U with Radeon Graphics (16) @ 4.546GHz

### words.txt *pneumonoultramicroscopicsilicovolcanoconiosis*

| Test Case       | Real Time | User Time | Sys Time |
| --------------- | --------- | --------- | -------- |
| **performance** | *1,297s*  | 1,660s    | 0,232s   |
| power saver     | *3,823s*  | 4,739s    | 0,533s   |
| balanced        | *1,942s*  | 2,485s    | 0,282s   |

### words_alpha.txt *pneumonoultramicroscopicsilicovolcanoconiosis*

| Test Case       | Real Time | User Time | Sys Time |
| --------------- | --------- | --------- | -------- |
| **performance** | *1,059s*  | 1,338s    | 0,227s   |
| balanced        | *1,554s*  | 1,955s    | 0,272s   |
| power saver     | *3,085s*  | 3,863s    | 0,443s   |

### words.txt *ad*

| Test Case       | Real Time | User Time | Sys Time |
| --------------- | --------- | --------- | -------- |
| **performance** | *0,308s*  | 0,374s    | 0,151s   |
| power saver     | *0,322s*  | 0,368s    | 0,178s   |
| balanced        | *0,318s*  | 0,384s    | 0,145s   |

### words_alpha.txt *ad*

| Test Case       | Real Time | User Time | Sys Time |
| --------------- | --------- | --------- | -------- |
| **performance** | *0,331s*  | 0,371s    | 0,198s   |
| balanced        | *0,422s*  | 0,508s    | 0,258s   |
| power saver     | *0,779s*  | 0,922s    | 0,342s   |

## Reference

- [WagnerFisher algoritm explanation video](https://www.youtube.com/watch?v=d-Eq6x1yssU)
- [WagnerFIsher Python](https://github.com/b001io/wagner-fischer/tree/main)
