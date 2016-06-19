package dexcom

// A Command specifies an operation to be performed by the Dexcom CGM receiver.
type Command byte

//go:generate stringer -type Command

const (
	PING Command = 10 + iota
	READ_FIRMWARE_HEADER
	_
	_
	_
	READ_DATABASE_PARTITION_INFO
	READ_DATABASE_PAGE_RANGE
	READ_DATABASE_PAGES
	READ_DATABASE_PAGE_HEADER
	_
	_
	_
	_
	_
	_
	READ_TRANSMITTER_ID
	WRITE_TRANSMITTER_ID
	READ_LANGUAGE
	WRITE_LANGUAGE
	READ_DISPLAY_TIME_OFFSET
	WRITE_DISPLAY_TIME_OFFSET
	READ_RTC
	RESET_RECEIVER
	READ_BATTERY_LEVEL
	READ_SYSTEM_TIME
	READ_SYSTEM_TIME_OFFSET
	WRITE_SYSTEM_TIME
	READ_GLUCOSE_UNIT
	WRITE_GLUCOSE_UNIT
	READ_BLINDED_MODE
	WRITE_BLINDED_MODE
	READ_CLOCK_MODE
	WRITE_CLOCK_MODE
	READ_DEVICE_MODE
	_
	ERASE_DATABASE
	SHUTDOWN_RECEIVER
	WRITE_PC_PARAMETERS
	READ_BATTERY_STATE
	READ_HARDWARE_BOARD_ID
	_
	_
	_
	_
	READ_FIRMWARE_SETTINGS
	READ_ENABLE_SETUP_WIZARD_FLAG
	_
	READ_SETUP_WIZARD_STATE
)
