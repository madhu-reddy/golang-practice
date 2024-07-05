1) Access struct fields with a dot.
2) You can also use dots with struct pointers - the pointers are automatically dereferenced.



Modules to work on

package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	userUtils "os/user"

	"github.com/joho/godotenv"
	flag "github.com/spf13/pflag"
	"golang.org/x/sys/execabs"
)
