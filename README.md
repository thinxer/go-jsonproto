#JSON Proto

Generate a skeleton Go struct code given a JSON example.

NOTE: The resulting code will not always be valid and may need further editing.

##Installation

    go install github.com/thinxer/go-jsonproto

##Usage

    go-jsonproto < example.json

##Example

Input:

```JSON
{
   "source" : "<a href=\"http://app.weibo.com/t/feed/9ksdit\" rel=\"nofollow\">iPhone客户端</a>",
   "favorited" : false,
   "_id" : 3518474036703576,
   "original_pic" : "http://ww4.sinaimg.cn/large/55178631jw1dzec9tkejkj.jpg",
   "mid" : "3518474036703576",
   "truncated" : false,
   "created_at" : "Sat Dec 01 17:10:44 +0800 2012",
   "text" : "百无聊赖等人中，貌似遇见熟人了，对脸盲来说最大的困惑在于既叫不出名字同时又想不起来到底是幼儿园同学还是小学初中或者高中同学，就差直接叫住人家问：你还认识我啊。可是我胖成这样了人家估计也认不出了。 我在#麦当劳 莫愁路店# http://t.cn/zj5kMAC",
   "in_reply_to_user_id" : "",
   "id" : 3518474036703576,
   "in_reply_to_status_id" : "",
   "comments_count" : 28,
   "idstr" : "3518474036703576",
   "attitudes_count" : 0,
   "bmiddle_pic" : "http://ww4.sinaimg.cn/bmiddle/55178631jw1dzec9tkejkj.jpg",
   "reposts_count" : 0,
   "visible" : {
      "list_id" : 0,
      "type" : 0
   },
   "geo" : {
      "coordinates" : [
         32.042191,
         118.776932
      ],
      "type" : "Point"
   },
   "thumbnail_pic" : "http://ww4.sinaimg.cn/thumbnail/55178631jw1dzec9tkejkj.jpg",
   "annotations" : [
      {
         "place" : {
            "poiid" : "B2094757D46DA7FF459D",
            "lat" : 32.042191,
            "title" : "麦当劳 莫愁路店",
            "type" : "checkin",
            "lon" : 118.776932
         }
      }
   ],
   "mlevel" : 0,
   "user_id" : 1427605041,
   "in_reply_to_screen_name" : ""
}
```

Output:

```Go
struct {
    RepostsCount int64 `json:"reposts_count"`
    Truncated bool
    ThumbnailPic string `json:"thumbnail_pic"`
    AttitudesCount int64 `json:"attitudes_count"`
    Idstr string
    Geo struct {
        Type string
        Coordinates []float64
    }
    Id int64 `json:"_id"`
    InReplyToStatusId string `json:"in_reply_to_status_id"`
    BmiddlePic string `json:"bmiddle_pic"`
    InReplyToScreenName string `json:"in_reply_to_screen_name"`
    Annotations []struct {
        Place struct {
            Lat float64
            Lon float64
            Type string
            Poiid string
            Title string
        }

    }
    OriginalPic string `json:"original_pic"`
    Visible struct {
        Type int64
        ListId int64 `json:"list_id"`
    }
    Id int64
    UserId int64 `json:"user_id"`
    Mid string
    Source string
    Favorited bool
    CreatedAt string `json:"created_at"`
    Mlevel int64
    Text string
    InReplyToUserId string `json:"in_reply_to_user_id"`
    CommentsCount int64 `json:"comments_count"`
}
```
