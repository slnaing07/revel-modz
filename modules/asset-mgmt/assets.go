package app

import (
	"os"
	"os/exec"
	"strings"

	"github.com/robfig/revel"
)

func init() {

	revel.OnAppStart(runGrunt)
}

// start Grunt with watch
func runGrunt() {

}
