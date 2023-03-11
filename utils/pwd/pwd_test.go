package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	fmt.Println(HashPwd("123456"))
}
func TestCheckPwd(t *testing.T) {
	fmt.Println(CheckPwd("$2a$04$twxYjFwO7YBrM.7jvMWzHOgc3JHUNGrtrFOtGTYD8V7dihHcZHVM6", "123456"))
}
