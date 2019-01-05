package types

import "time"

type MysqlAggPacket struct {
	Id   int       `db:"id"`
	Time time.Time `db:"time"`

	GtwId string `db:"gtw_id"`

	Modulation string `db:"modulation"`
	DataRate   string `db:"datarate"`
	Bitrate    uint32 `db:"bitrate"`
	CodingRate string `db:"coding_rate"`

	SNR       float32 `db:"snr"`
	RSSI      float32 `db:"rssi"`
	Frequency float32 `db:"frequency"`

	Latitude  float64 `db:"latitude"`
	Longitude float64 `db:"longitude"`
	Altitude  float64 `db:"altitude"`
}
