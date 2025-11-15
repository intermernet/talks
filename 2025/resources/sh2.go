// sh2Protocol implements the Sensor Hub 2 (SH-2) application protocol.
type sh2Protocol struct {
	device               *Device
	transport            *shtp
	cmdSeq               uint8
	waiting              bool
	lastCmd              uint8
	pendingConfigRequest bool
	pendingConfigSensor  SensorID
	receivedConfig       SensorConfig
	configReady          bool
	configBuf            [17]byte                    // Reusable buffer for setSensorConfig
	commandBuf           [3 + commandParamCount]byte // Reusable buffer for sendCommand
}
