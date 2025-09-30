
package wcat

import (
  "fmt"
)

const (
  value = "Toast: "
)

func Toast(in string) string {
  return fmt.Sprintf("%#v", value + in)
}

