package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Info struct {
	gorm.Model
	Id          string  `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	NodeId      string  `sql:"type:uuid`
	Load1       float64 `json:"load_1"`
	Load5       float64 `json:"load_5"`
	Load15      float64 `json:"load_15"`
	IpStatus    bool    `json:"ip_status"`
	Thread      uint64  `json:"thread"`
	Process     uint64  `json:"process"`
	NetworkTx   uint64  `json:"network_tx"`
	NetworkRx   uint64  `json:"network_rx"`
	NetworkIn   uint64  `json:"network_in"`
	NetworkOut  uint64  `json:"network_out"`
	Ping10010   float64 `json:"ping_10010"`
	Ping10086   float64 `json:"ping_10086"`
	Ping189     float64 `json:"ping_189"`
	Time10010   uint64  `json:"time_10010"`
	Time10086   uint64  `json:"time_10086"`
	Time189     uint64  `json:"time_189"`
	TCP         uint64  `json:"tcp"`
	UDP         uint64  `json:"udp"`
	CPU         float64 `json:"cpu"`
	MemoryTotal uint64  `json:"memory_total"`
	MemoryUsed  uint64  `json:"memory_used"`
	SwapTotal   uint64  `json:"swap_total"`
	SwapUsed    uint64  `json:"swap_used"`
	Uptime      uint64  `json:"uptime"`
	HddTotal    uint64  `json:"hdd_total"`
	HddUsed     uint64  `json:"hdd_used"`
}

type Status struct {
	Load1       float64 `json:"load_1"`
	Load5       float64 `json:"load_5"`
	Load15      float64 `json:"load_15"`
	IpStatus    bool    `json:"ip_status"`
	Thread      uint64  `json:"thread"`
	Process     uint64  `json:"process"`
	NetworkTx   uint64  `json:"network_tx"`
	NetworkRx   uint64  `json:"network_rx"`
	NetworkIn   uint64  `json:"network_in"`
	NetworkOut  uint64  `json:"network_out"`
	Ping10010   float64 `json:"ping_10010"`
	Ping10086   float64 `json:"ping_10086"`
	Ping189     float64 `json:"ping_189"`
	Time10010   uint64  `json:"time_10010"`
	Time10086   uint64  `json:"time_10086"`
	Time189     uint64  `json:"time_189"`
	TCP         uint64  `json:"tcp"`
	UDP         uint64  `json:"udp"`
	CPU         float64 `json:"cpu"`
	MemoryTotal uint64  `json:"memory_total"`
	MemoryUsed  uint64  `json:"memory_used"`
	SwapTotal   uint64  `json:"swap_total"`
	SwapUsed    uint64  `json:"swap_used"`
	Uptime      uint64  `json:"uptime"`
	HddTotal    uint64  `json:"hdd_total"`
	HddUsed     uint64  `json:"hdd_used"`
}

func (info *Info) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	info.Id = uuid.New().String()
	return
}

func NewDefaultInfo() Info {
	return Info{
		Load1:       0.0,
		Load5:       0.0,
		Load15:      0.0,
		IpStatus:    false,
		Thread:      0,
		Process:     0,
		NetworkTx:   0,
		NetworkRx:   0,
		NetworkIn:   0,
		NetworkOut:  0,
		Ping10010:   0.0,
		Ping10086:   0.0,
		Ping189:     0.0,
		Time10010:   0,
		Time10086:   0,
		Time189:     0,
		TCP:         0,
		UDP:         0,
		CPU:         0.0,
		MemoryTotal: 0,
		MemoryUsed:  0,
		SwapTotal:   0,
		SwapUsed:    0,
		Uptime:      0,
		HddTotal:    0,
		HddUsed:     0,
	}
}

func NewDefaultStatus() Status {
	return Status{
		Load1:       0.0,
		Load5:       0.0,
		Load15:      0.0,
		IpStatus:    false,
		Thread:      0,
		Process:     0,
		NetworkTx:   0,
		NetworkRx:   0,
		NetworkIn:   0,
		NetworkOut:  0,
		Ping10010:   0.0,
		Ping10086:   0.0,
		Ping189:     0.0,
		Time10010:   0,
		Time10086:   0,
		Time189:     0,
		TCP:         0,
		UDP:         0,
		CPU:         0.0,
		MemoryTotal: 0,
		MemoryUsed:  0,
		SwapTotal:   0,
		SwapUsed:    0,
		Uptime:      0,
		HddTotal:    0,
		HddUsed:     0,
	}
}
