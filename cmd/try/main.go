package main

import (
    "log"
    "path/filepath"
    "rr"
)

func main() {
    log.Println(filepath.Ext("a.js/xxx"), filepath.Ext("b.js?v=111"), filepath.Ext("c.json#rgerg=df.js"))
    log.Println(
        rr.F("efwe/ewg.ss?geg").GetName(),
    )
}
