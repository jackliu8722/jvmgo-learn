
package main

import "fmt"
import "jvmgo/cmdline"
import "os"
import "strings"
import "jvmgo/classpath"
import (
    "jvmgo/options"
    "jvmgo/classfile"
)
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
    //fmt.Println(cp.CompositeEntry.String())
    fmt.Printf("classpath:%v class:%v args:%v\n",cmd.Options().Classpath(),cmd.Class(),cmd.Args())

    className := strings.Replace(cmd.Class(),".","/",-1)
    _,classData,err := cp.ReadClass(className)
    if err != nil {
        fmt.Printf("Could not find or load main class %s\n",cmd.Class())
        return
    }
    //fmt.Printf("class data:%v\n",classData)

    cf, err := classfile.Parse(classData)
    if err != nil {
        panic(err)
    }
    printClassInfo(cf)
}

func printClassInfo(cf *classfile.ClassFile) {
    fmt.Printf("version: %v.%v\n",cf.MajorVersion(),cf.MinorVersion())
    fmt.Printf("classname: %v\n",cf.ClassName())
    fmt.Printf("super classname: %v\n",cf.SupperClassName())

    fmt.Printf("field count: %v\n",len(cf.Fields()))
    for _,f := range cf.Fields() {
        fmt.Printf("    %v \n",f.Name())
    }

    fmt.Printf("method count: %v\n",len(cf.Methods()))
    for _, m := range cf.Methods() {
        fmt.Printf("    %v\n",m.Name())
    }
}
