package frm

import (
	"fmt"
	"testing"

	"github.com/lsytj0413/myinctl/pkg/helper"
)

func Test(t *testing.T) {
	r := NewReader()
	fmt.Println(r.ReadFile(helper.JoinWithProjectAbsPath("./test/user_accounts.frm")))
}
