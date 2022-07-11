package synchronizer

type ForumThread struct {
	Tid          uint   `gorm:"column:tid" json:"tid"`
	Fid          uint   `gorm:"column:fid" json:"fid"`
	Posttableid  uint16 `gorm:"column:posttableid" json:"posttableid"`
	Typeid       uint16 `gorm:"column:typeid" json:"typeid"`
	Sortid       uint16 `gorm:"column:sortid" json:"sortid"`
	Readperm     uint8  `gorm:"column:readperm" json:"readperm"`
	Price        int16  `gorm:"column:price" json:"price"`
	Author       string `gorm:"column:author" json:"author"`
	Authorid     uint   `gorm:"column:authorid" json:"authorid"`
	Subject      string `gorm:"column:subject" json:"subject"`
	Dateline     uint   `gorm:"column:dateline" json:"dateline"`
	Lastpost     uint   `gorm:"column:lastpost" json:"lastpost"`
	Lastposter   string `gorm:"column:lastposter" json:"lastposter"`
	Views        uint   `gorm:"column:views" json:"views"`
	Replies      uint   `gorm:"column:replies" json:"replies"`
	Displayorder int8   `gorm:"column:displayorder" json:"displayorder"`
	Highlight    int8   `gorm:"column:highlight" json:"highlight"`
	Digest       int8   `gorm:"column:digest" json:"digest"`
	Rate         int8   `gorm:"column:rate" json:"rate"`
	Special      int8   `gorm:"column:special" json:"special"`
	Attachment   int8   `gorm:"column:attachment" json:"attachment"`
	Moderated    int8   `gorm:"column:moderated" json:"moderated"`
	Closed       uint   `gorm:"column:closed" json:"closed"`
	Stickreply   uint8  `gorm:"column:stickreply" json:"stickreply"`
	Recommends   int16  `gorm:"column:recommends" json:"recommends"`
	RecommendAdd int16  `gorm:"column:recommend_add" json:"recommend_add"`
	RecommendSub int16  `gorm:"column:recommend_sub" json:"recommend_sub"`
	Heats        uint   `gorm:"column:heats" json:"heats"`
	Status       uint16 `gorm:"column:status" json:"status"`
	Isgroup      int8   `gorm:"column:isgroup" json:"isgroup"`
	Favtimes     int    `gorm:"column:favtimes" json:"favtimes"`
	Sharetimes   int    `gorm:"column:sharetimes" json:"sharetimes"`
	Stamp        int8   `gorm:"column:stamp" json:"stamp"`
	Icon         int8   `gorm:"column:icon" json:"icon"`
	Pushedaid    int    `gorm:"column:pushedaid" json:"pushedaid"`
	Cover        int16  `gorm:"column:cover" json:"cover"`
	Replycredit  int16  `gorm:"column:replycredit" json:"replycredit"`
	Relatebytag  string `gorm:"column:relatebytag" json:"relatebytag"`
	Maxposition  uint   `gorm:"column:maxposition" json:"maxposition"`
	Bgcolor      string `gorm:"column:bgcolor" json:"bgcolor"`
	Comments     uint   `gorm:"column:comments" json:"comments"`
	Hidden       uint16 `gorm:"column:hidden" json:"hidden"`
}

func (f *ForumThread) CollectName() string {
	return "dz.forum_thread"
}
