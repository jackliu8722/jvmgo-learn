package classpath

import (
    "path/filepath"
    "strings"
)

type ClassPath struct {
    
}

func Parse(cpOption string) *ClassPath {
    cp := &ClassPath{}
    cp.parseBootAndExtClassPath()
    cp.parseUserClassPath(cpOption)
    return cp
}

func (self *ClassPath) parseBootAndExtClassPath(){
    jarLibPath := filepath.Join(options.)
}
