package options

import (
    "os"
    "path/filepath"
    "jvmgo/cmdline"
)

var (
    VerboseClass bool
    ThreadStackSize uint
    AbsJavaHome string
    AbsJreLib string
)

func InitOptions(cmdOptions *cmdline.Options){
    VerboseClass = cmdOptions.VerboseClass()
    ThreadStackSize = uint(cmdOptions.Xss)
    initJavaHome(cmdOptions.XuseJavaHome)
}

func initJavaHome(useOsEnv bool){
    jh := "./jre"
    //jh := "/Library/Java/JavaVirtualMachines/jdk1.8.0_66.jdk/Contents/Home/jre"
    if useOsEnv {
        jh = os.Getenv("JAVA_HOME")
        if jh == "" {
            panic("$JAVA_HOME not set!")
        }
        jh += "/jre"
    }
    if absJh, err := filepath.Abs(jh); err == nil {
        AbsJavaHome = absJh
        AbsJreLib = filepath.Join(absJh,"lib")
    }else{
        panic(err)
    }
}
