# dbmock

`import github.com/memememomo/dbmock`

dbmock is a mock generator for Go.

## Example

```go
package main

import (
	"fmt"
	
	"github.com/memememomo/dbmock"
)

type User struct {
	ID uint64
	Name string
}

// Implement DBMapper 
func (u *User) ToDB() error {
	// Save to DB...
	return nil
}

// Mock setting
func user(i uint64) dbmock.DBMapper {
	return dbmock.Mock(&User{
		ID: i + 1,
		Name: fmt.Sprintf("Name_%d", i + 1),
	})
}

func UserMock() *Generator {
	return NewGenerator(user)
}

func main() {
	generator := UserMock()
	
	// Single - Only Mock
	user := generator.Single(0, nil).(*User)
	
	// Multi - 5 mocks
	mocks := generator.Multi(5, nil)
	
	// Single with overwriter
	user := generator.Single(1, func(i uint64, mapper DBMapper) DBMapper {
		u := mapper.(*User)
		u.Name = "Overwrite_Name"
		return u
	})
}
```