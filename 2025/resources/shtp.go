// shtp implements the Sensor Hub Transport Protocol layer.
type shtp struct {
	hal      *halI2C
	handlers map[uint8]shtpHandler
	seq      [8]uint8
	rx       [maxTransferIn]byte
	tx       [maxTransferOut]byte // Reusable transmit buffer
}
