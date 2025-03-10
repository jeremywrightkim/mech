package main

import (
   "flag"
   "fmt"
   "github.com/89z/format"
   "github.com/89z/mech/bandcamp"
   "net/http"
   "os"
   "strings"
   "time"
)

func main() {
   var (
      info, verbose bool
      sleep time.Duration
   )
   flag.BoolVar(&info, "i", false, "info only")
   flag.DurationVar(&sleep, "s", time.Second, "sleep")
   flag.BoolVar(&verbose, "v", false, "verbose")
   flag.Parse()
   if flag.NArg() == 0 {
      fmt.Println("bandcamp [flags] [track or album]")
      flag.PrintDefaults()
      return
   }
   if verbose {
      bandcamp.LogLevel = 2
   }
   addr := flag.Arg(0)
   data, err := bandcamp.NewDataTralbum(addr)
   if err != nil {
      panic(err)
   }
   for _, track := range data.TrackInfo {
      if info {
         fmt.Printf("%+v\n", track)
      } else {
         addr, ok := track.File.MP3_128()
         if ok {
            fmt.Println("GET", addr)
            res, err := http.Get(addr)
            if err != nil {
               panic(err)
            }
            defer res.Body.Close()
            name := data.Artist + "-" + track.Title + ".mp3"
            file, err := os.Create(strings.Map(format.Clean, name))
            if err != nil {
               panic(err)
            }
            defer file.Close()
            if _, err := file.ReadFrom(res.Body); err != nil {
               panic(err)
            }
            time.Sleep(sleep)
         }
      }
   }
}
