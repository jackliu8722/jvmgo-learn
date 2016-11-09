
package main

import "fmt"
import "jvmgo/cmdline"
import "os"
import "strings"
import "jvmgo/classpath"
import "jvmgo/options"
func main(){
    cmd ,err := cmdline.ParseCommand(os.Args)
    if err != nil {
        cmdline.PrintUsage()
    }else {
        startJVM(cmd)
    }
}

func startJVM(cmd *cmdline.Command){
    //fmt.Printf("classpath: %v\n" ,cmd.Options().Classpath())
    options.InitOptions(cmd.Options())
    cp := classpath.Parse(cmd.Options().Classpath())
    fmt.Printf("classpath:%v class:%v args:%v\n",cmd.Options().Classpath(),cmd.Class(),cmd.Args())

    className := strings.Replace(cmd.Class(),".","/",-1)
    _,classData,err := cp.ReadClass(className)
    if err != nil {
        fmt.Printf("Could not find or load main class %s\n",cmd.Class())
        return
    }
    fmt.Printf("class data:%v\n",classData)
}
