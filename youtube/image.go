package youtube

import (
   "strings"
)

var Images = []Image{
   {Width:120, Height:90, Base:"default.jpg"},
   {Width:120, Height:90, Base:"1.jpg"},
   {Width:120, Height:90, Base:"2.jpg"},
   {Width:120, Height:90, Base:"3.jpg"},
   {Width:120, Height:90, Base:"default.webp"},
   {Width:120, Height:90, Base:"1.webp"},
   {Width:120, Height:90, Base:"2.webp"},
   {Width:120, Height:90, Base:"3.webp"},
   {Width:320, Height:180, Base:"mq1.jpg", Crop:true},
   {Width:320, Height:180, Base:"mq2.jpg", Crop:true},
   {Width:320, Height:180, Base:"mq3.jpg", Crop:true},
   {Width:320, Height:180, Base:"mqdefault.jpg"},
   {Width:320, Height:180, Base:"mq1.webp", Crop:true},
   {Width:320, Height:180, Base:"mq2.webp", Crop:true},
   {Width:320, Height:180, Base:"mq3.webp", Crop:true},
   {Width:320, Height:180, Base:"mqdefault.webp"},
   {Width:480, Height:360, Base:"0.jpg"},
   {Width:480, Height:360, Base:"hqdefault.jpg"},
   {Width:480, Height:360, Base:"hq1.jpg"},
   {Width:480, Height:360, Base:"hq2.jpg"},
   {Width:480, Height:360, Base:"hq3.jpg"},
   {Width:480, Height:360, Base:"0.webp"},
   {Width:480, Height:360, Base:"hqdefault.webp"},
   {Width:480, Height:360, Base:"hq1.webp"},
   {Width:480, Height:360, Base:"hq2.webp"},
   {Width:480, Height:360, Base:"hq3.webp"},
   {Width:640, Height:480, Base:"sddefault.jpg"},
   {Width:640, Height:480, Base:"sd1.jpg"},
   {Width:640, Height:480, Base:"sd2.jpg"},
   {Width:640, Height:480, Base:"sd3.jpg"},
   {Width:640, Height:480, Base:"sddefault.webp"},
   {Width:640, Height:480, Base:"sd1.webp"},
   {Width:640, Height:480, Base:"sd2.webp"},
   {Width:640, Height:480, Base:"sd3.webp"},
   {Width:1280, Height:720, Base:"hq720.jpg"},
   {Width:1280, Height:720, Base:"maxresdefault.jpg"},
   {Width:1280, Height:720, Base:"maxres1.jpg"},
   {Width:1280, Height:720, Base:"maxres2.jpg"},
   {Width:1280, Height:720, Base:"maxres3.jpg"},
   {Width:1280, Height:720, Base:"hq720.webp"},
   {Width:1280, Height:720, Base:"maxresdefault.webp"},
   {Width:1280, Height:720, Base:"maxres1.webp"},
   {Width:1280, Height:720, Base:"maxres2.webp"},
   {Width:1280, Height:720, Base:"maxres3.webp"},
}

type Image struct {
   Width int
   Height int
   Base string
   Crop bool
}

func (i Image) Format(id string) string {
   var addr strings.Builder
   addr.WriteString("http://i.ytimg.com/vi")
   if strings.HasSuffix(i.Base, ".webp") {
      addr.WriteString("_webp")
   }
   addr.WriteByte('/')
   addr.WriteString(id)
   addr.WriteByte('/')
   addr.WriteString(i.Base)
   return addr.String()
}
