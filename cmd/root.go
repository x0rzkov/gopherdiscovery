package cmd

import (
	"fmt"
	"os"
	"reflect"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	log  = logrus.New()
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

// RootCmd is the root command for limo
var RootCmd = &cobra.Command{
	Use:   "gopherdiscovery",
	Short: "Simple Service Discovery for Go and nanomsg.",
	Long:  `Simple Service Discovery for Go and nanomsg.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func die(format string, v ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func date() string {
	return time.Now().Format(time.ANSIC)
}

func interfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}
	ret := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}
	return ret
}
