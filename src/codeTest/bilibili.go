package main

import (
	"encoding/json"
	"fmt"
)

type Request_Impression struct {
	Id string `json:"id"`
	TagId string `json:"tagId"`
	Native *Request_Impression_Native `json:"adNative"`
	BidInfo *Request_Impression_BidInfo `json:"bidInfo"`
	OpenType []int32 `json:"openType"`
	SlotType int32 `json:"slotType"`
	Pmp *Request_Impression_Pmp `json:"pmp"`
}

type Request_Impression_Native struct {
	Titile *Request_Impression_Native_Title `json:"titile"`
	Image *Request_Impression_Native_Image `json:"image"`
	Desc *Request_Impression_Native_Desc `json:"desc"`
}

type Request_Impression_Native_Title struct {
	IsRequired int32 `json:"isRequired"`
	MinTitleLen int32 `json:"minTitleLen"`
	MaxTitleLen int32 `json:"maxTitleLen"`
}
type Request_Impression_Native_Image struct {
	IsRequired int32 `json:"isRequired"`
	ImageWidth int32 `json:"imageWidth"`
	ImageHeight int32 `json:"imageHeight"`
	ImageNum int32 `json:"imageNum"`
}
type Request_Impression_Native_Desc struct {
	IsRequired int32 `json:"isRequired"`
	MinDescLength int32 `json:"minDescLength"`
	MaxDescLength int32 `json:"maxDescLength"`

}

type Request_Impression_BidInfo struct {
	BidType int32 `json:"bidType"`
	BidFloor int32 `json:"bidFloor"`
}

type Request_Impression_Pmp struct {
	DealId string `json:"dealId"`
	BidFloor int32 `json:"bidFloor"`
}

type Request_Impression_Device struct {
	Buvid string `json:"buvid"`
	OsType int32 `json:"osType"`
	Idfa string `json:"idfa"`
	Imei string `json:"imei"`
	AndroidId string `json:"androidId"`
	Ip string `json:"ip"`
	Mac string `json:"mac"`
	Geo *Request_Impression_Device_Geo `json:"geo"`
	ScreenWidth int32 `json:"screenWith"`
	ScreenHeight int32 `json:"screenHeight"`
	Model string `json:"model"`
	Vendor string `json:"vendor"`
	Network string `json:"network"`
	Ua string `json:"ua"`
	Osv string `json:"osv"`
}

type Request_Impression_Device_Geo struct {
	Latitude string `json:"latitude"`
	Longitude string `json:"longitude"`
}
type Request struct {
	Id string	`json:"id"`
	imps []*Request_Impression `json:"imps"`
	device *Request_Impression_Device `json:"device"`
}


func main()  {
	str := `{"device":{"androidId":"3991578f49a42356","buvid":"27391C43-F189-4AB3-B72B-B2C09030C0B212226infoc","geo":{"latitude":"","longitude":""},"idfa":"","imei":"861744038585098","ip":"36.60.255.97","mac":"1C:77:F6:B2:11:F7","model":"OPPO R9m","osType":0,"osv":"5.1","screenHeight":1920,"screenWidth":1080,"ua":"Mozilla/5.0 (Linux; Android 5.1; OPPO R9m Build/LMY47I; wv)","vendor":"OPPO"},"id":"1525663055692q172a22a55a224q69pp","imps":[{"adNative":{"desc":{"isRequired":1,"maxDescLength":6,"minDescLength":2},"image":{"imageHeight":400,"imageNum":1,"imageWidth":640,"isRequired":1},"title":{"isRequired":1,"maxTitleLen":18,"minTitleLen":2}},"bidInfo":{"bidFloor":334,"bidType":0},"id":"3ff7fa69fe0445c299a60b79d5f2899d","openType":[0],"slotType":3,"tagId":"1900"}]}`
	req := &Request{}
	err := json.Unmarshal([]byte(str), req)
	r,_ := json.Marshal(req)

	fmt.Println(req.Id, req.device,err, req,
		r)

	//fmt.Println(string(req))
}
