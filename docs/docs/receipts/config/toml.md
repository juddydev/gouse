# Toml

## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong1009/gouse/config"
)
```
## Functions


### SampleConfigToml

```go
func SampleConfigToml() {
	type Configuration struct {
		Mysql struct {
			Host     string
			Username string
			Password string
			Database string
		}
	}

	var myConf Configuration
	err := config.Toml("conf.toml", &myConf)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Mysql Host:", myConf.Mysql.Host)
	fmt.Println("Mysql Username:", myConf.Mysql.Username)
	fmt.Println("Mysql Password:", myConf.Mysql.Password)
	fmt.Println("Mysql Database:", myConf.Mysql.Database)
}
```
