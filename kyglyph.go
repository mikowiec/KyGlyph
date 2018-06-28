// Copyright (c) 2018 Nikolay Kravets. All Rights Reserved.
// See License.txt for license information.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/mikowiec/kyglyph/log"
	"github.com/mikowiec/kyglyph/utils"
	"github.com/mikowiec/kyglyph/api"
	"github.com/mikowiec/kyglyph/tests"
	"github.com/mikowiec/kyglyph/model"
)

var cfgConfigFile string
var cfgEmail string
var cfgPassword string
var cfgTeamName string
var cfgRole string
var cfgCmdRun bool


func main() {

    cmdInParse()

    utils.ConfigLoad(cfgConfigFile)

}



