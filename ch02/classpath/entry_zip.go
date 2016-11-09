package classpath

import (
    "archive/zip"
    "errors"
    "io/ioutil"
    "path/filepath"
)

type ZipEntry struct {
    obsZip string
    zipRC *zip.ReadCloser
}

func newZipEntry(path string) *ZipEntry{
    obsZip,err := filepath.Abs(path)
    if err != nil {
        panic(err)
    }

    return &ZipEntry{obsZip,nil}
}

func (self *ZipEntry) readClass(className string)(Entry,[]byte,error){
    if self.zipRC == nil {
        err := self.openJar()
        if err != nil {
            return self,nil,err
        }
    }

    classFile := self.findClass(className)

    if classFile == nil {
        return self,nil,errors.New("class not found: " + className)
    }

    data,err := readClass(classFile)
    return self,data,err
}

func (self *ZipEntry) openJar() error {
    r ,err := zip.OpenReader(self.obsZip)
    if err == nil {
        self.zipRC = r
    }
    return err
}

func (self *ZipEntry) findClass(className string) *Zip.File {
    for _,f := range self.zipRC.File {
        if f.Name == className {
            return f
        }
    }
    return nil
}

func readClass(classFile *zip.File) ([]byte,error) {
    rc ,err := classFile.Open()
    if err != nil {
        return nil,err
    }

    data,err := ioutil.ReadAll(rc)
    rc.Close()

    if err != nil {
        return nil,err
    }

    return data,nil
}

func (self *ZipEntry) String() string {
    return self.obsZip
}
