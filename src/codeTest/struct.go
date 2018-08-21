package main

import (
	"fmt"
	"encoding/json"
)

type Response struct {
	Id string `json:"id"`
	SeatBids []*Response_SeatBids `json:"seatBids"`
	BidId string `json:"bidId"`
	ExtraData *Response_ExtraData `json:"extraData"`
}

type Response_SeatBids struct {
	bids []*Response_SeatBids_Bids `json:"bids"`
	seat string `json:"seat"`
}

type Response_SeatBids_Bids struct {
	impId string `json:"impId"`
	bidType int32 `json:"bidType"`
	price int32 `json:"price"`
	bidCreativeId int32 `json:"bidCreativeId"`
	admNative *Response_SeatBids_Bids_AdmNative
}

type Response_SeatBids_Bids_AdmNative struct {
	title string `json:"title"`
	imageUrls []string `json:"imageUrls"`
	imageWidth int `json:"imageWidth"`
	imageHeight int `json:"imageHeight"`
	desc string `json:"desc"`
	linkUrl string `json:"linkUrl"`
	clickUrls string `json:"clickUrls"`
	impUrls string `json:"impUrls"`
	interactionType int32 `json:"interactionType"`
	downloadUrl string `json:"downloadUrl"`
	bundle string `json:"bundle"`
	deeplinkUrl string `json:"deeplinkUrl"`
	advDomain string `json:"advDomain"`
	downloadClickUrls []string `json:"downloadClickUrls"`
}

type Response_ExtraData struct {
	userTags []string `json:"userTags"`
}

func main()  {
	ss := &Response{
		Id:"12",

	}
	bResp, _ := json.Marshal(ss)
	fmt.Println(ss, bResp, string(bResp))
}
