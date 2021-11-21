package reddit

import (
   "github.com/89z/mech"
   "github.com/89z/parse/m3u"
   "net/http"
   "path"
   "sort"
   "strings"
)

type HLS struct {
   ID int
   Resolution string
   URI string
}

func (l Link) HLS() ([]HLS, error) {
   req, err := http.NewRequest("GET", l.Media.Reddit_Video.HLS_URL, nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = ""
   res, err := mech.RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   prefix, _ := path.Split(l.Media.Reddit_Video.HLS_URL)
   var hlss []HLS
   for key, val := range m3u.NewPlaylist(res.Body) {
      hlss = append(hlss, HLS{
         Resolution: val["RESOLUTION"], URI: prefix + key,
      })
   }
   sort.Slice(hlss, func(a, b int) bool {
      switch strings.Compare(hlss[a].Resolution, hlss[b].Resolution) {
      case -1:
         return true
      case 1:
         return false
      }
      return hlss[a].URI < hlss[b].URI
   })
   for i := range hlss {
      hlss[i].ID = i
   }
   return hlss, nil
}
