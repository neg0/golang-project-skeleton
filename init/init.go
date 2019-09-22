package init

import "github.com/neg0/golang-project-skeleton/init/config"

func init() {
	_, err := config.Instance()
	if err != nil {
		panic(err)
	}
}
