## 0x0BSoD/strNorm
Package for normalizing string

### Example
```go
import (
    "github.com/0x0bsod/strNorm"
    "fmt"
)
func main() {
    str := "/dev/sda3            329229596   184366340   128069656  60% /home"
    norm_string = strNorm.Normalize(str)
    fmt.Println(norm_string)
    // Output: /dev/sda3 329229596 184366340 128069656 60% /home
}
```