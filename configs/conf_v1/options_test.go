package conf_v1

import (
	"fmt"
	"testing"
)

func TestNewConfig(t *testing.T) {
	conf := NewConfig(ConfigExtendMi(true))
	fmt.Print(conf)
}
