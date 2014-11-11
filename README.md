# Skrif

### An experimenting scrolls in Go, without the Hash.

## Usage

```golang
import "github.com/ys/skrif"

lol := skrif.New()
lol.SetGlobalContext("app=lol")

lol.Println("yolo=true") 
//=> "app=lol yolo=true"

lol.Context("somuch=true", func(lol *Skrif) {
        lol.Println("yolo=true") 
})
//=> "app=lol somuch=true yolo=true"

lol.Context("somuch=true", func(lol *Skrif) {
        lol.Context("RTFM=false", func(lol *Skrif) {
                lol.Println("yolo=true") 
        })
})
//=> "app=lol somuch=true RTFM=false yolo=true"
```