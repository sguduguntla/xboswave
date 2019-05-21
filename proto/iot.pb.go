// Code generated by protoc-gen-go. DO NOT EDIT.
// source: iot.proto

//This is designed to be included by the main xbos proto file and includes the
//definitions for the XBOS IoT devices
//
//Maintainer: Gabe Fierro
//Version 1.0

package xbospb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type URI struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *URI) Reset()         { *m = URI{} }
func (m *URI) String() string { return proto.CompactTextString(m) }
func (*URI) ProtoMessage()    {}
func (*URI) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{0}
}

func (m *URI) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_URI.Unmarshal(m, b)
}
func (m *URI) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_URI.Marshal(b, m, deterministic)
}
func (m *URI) XXX_Merge(src proto.Message) {
	xxx_messageInfo_URI.Merge(m, src)
}
func (m *URI) XXX_Size() int {
	return xxx_messageInfo_URI.Size(m)
}
func (m *URI) XXX_DiscardUnknown() {
	xxx_messageInfo_URI.DiscardUnknown(m)
}

var xxx_messageInfo_URI proto.InternalMessageInfo

func (m *URI) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *URI) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Triple struct {
	Subject              *URI     `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Predicate            *URI     `protobuf:"bytes,2,opt,name=predicate,proto3" json:"predicate,omitempty"`
	Object               *URI     `protobuf:"bytes,3,opt,name=object,proto3" json:"object,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Triple) Reset()         { *m = Triple{} }
func (m *Triple) String() string { return proto.CompactTextString(m) }
func (*Triple) ProtoMessage()    {}
func (*Triple) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{1}
}

func (m *Triple) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Triple.Unmarshal(m, b)
}
func (m *Triple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Triple.Marshal(b, m, deterministic)
}
func (m *Triple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Triple.Merge(m, src)
}
func (m *Triple) XXX_Size() int {
	return xxx_messageInfo_Triple.Size(m)
}
func (m *Triple) XXX_DiscardUnknown() {
	xxx_messageInfo_Triple.DiscardUnknown(m)
}

var xxx_messageInfo_Triple proto.InternalMessageInfo

func (m *Triple) GetSubject() *URI {
	if m != nil {
		return m.Subject
	}
	return nil
}

func (m *Triple) GetPredicate() *URI {
	if m != nil {
		return m.Predicate
	}
	return nil
}

func (m *Triple) GetObject() *URI {
	if m != nil {
		return m.Object
	}
	return nil
}

type Error struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{2}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type XBOSIoTDeviceState struct {
	// current time at device/service
	//unit:ns
	Time uint64 `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	// unique identifier for this request; used to line up with device state requests
	Requestid int64 `protobuf:"varint,2,opt,name=requestid,proto3" json:"requestid,omitempty"`
	// any error that occured since the last device report. If requestid above is non-zero,
	// then this error corresponds to the request with the given requestid
	Error *Error `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	// XBOS IoT devices
	Thermostat               *Thermostat               `protobuf:"bytes,4,opt,name=thermostat,proto3" json:"thermostat,omitempty"`
	Meter                    *Meter                    `protobuf:"bytes,5,opt,name=meter,proto3" json:"meter,omitempty"`
	Light                    *Light                    `protobuf:"bytes,6,opt,name=light,proto3" json:"light,omitempty"`
	Evse                     *EVSE                     `protobuf:"bytes,7,opt,name=evse,proto3" json:"evse,omitempty"`
	WeatherStation           *WeatherStation           `protobuf:"bytes,8,opt,name=weather_station,json=weatherStation,proto3" json:"weather_station,omitempty"`
	WeatherStationPrediction *WeatherStationPrediction `protobuf:"bytes,9,opt,name=weather_station_prediction,json=weatherStationPrediction,proto3" json:"weather_station_prediction,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                  `json:"-"`
	XXX_unrecognized         []byte                    `json:"-"`
	XXX_sizecache            int32                     `json:"-"`
}

func (m *XBOSIoTDeviceState) Reset()         { *m = XBOSIoTDeviceState{} }
func (m *XBOSIoTDeviceState) String() string { return proto.CompactTextString(m) }
func (*XBOSIoTDeviceState) ProtoMessage()    {}
func (*XBOSIoTDeviceState) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{3}
}

func (m *XBOSIoTDeviceState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XBOSIoTDeviceState.Unmarshal(m, b)
}
func (m *XBOSIoTDeviceState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XBOSIoTDeviceState.Marshal(b, m, deterministic)
}
func (m *XBOSIoTDeviceState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XBOSIoTDeviceState.Merge(m, src)
}
func (m *XBOSIoTDeviceState) XXX_Size() int {
	return xxx_messageInfo_XBOSIoTDeviceState.Size(m)
}
func (m *XBOSIoTDeviceState) XXX_DiscardUnknown() {
	xxx_messageInfo_XBOSIoTDeviceState.DiscardUnknown(m)
}

var xxx_messageInfo_XBOSIoTDeviceState proto.InternalMessageInfo

func (m *XBOSIoTDeviceState) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *XBOSIoTDeviceState) GetRequestid() int64 {
	if m != nil {
		return m.Requestid
	}
	return 0
}

func (m *XBOSIoTDeviceState) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *XBOSIoTDeviceState) GetThermostat() *Thermostat {
	if m != nil {
		return m.Thermostat
	}
	return nil
}

func (m *XBOSIoTDeviceState) GetMeter() *Meter {
	if m != nil {
		return m.Meter
	}
	return nil
}

func (m *XBOSIoTDeviceState) GetLight() *Light {
	if m != nil {
		return m.Light
	}
	return nil
}

func (m *XBOSIoTDeviceState) GetEvse() *EVSE {
	if m != nil {
		return m.Evse
	}
	return nil
}

func (m *XBOSIoTDeviceState) GetWeatherStation() *WeatherStation {
	if m != nil {
		return m.WeatherStation
	}
	return nil
}

func (m *XBOSIoTDeviceState) GetWeatherStationPrediction() *WeatherStationPrediction {
	if m != nil {
		return m.WeatherStationPrediction
	}
	return nil
}

type XBOSIoTDeviceActuation struct {
	// current time at device/service
	//unit:ns
	Time uint64 `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	// unique identifier for this request; used to line up with device state responses
	Requestid int64 `protobuf:"varint,2,opt,name=requestid,proto3" json:"requestid,omitempty"`
	// XBOS IoT devices
	Thermostat           *Thermostat `protobuf:"bytes,3,opt,name=thermostat,proto3" json:"thermostat,omitempty"`
	Meter                *Meter      `protobuf:"bytes,4,opt,name=meter,proto3" json:"meter,omitempty"`
	Light                *Light      `protobuf:"bytes,5,opt,name=light,proto3" json:"light,omitempty"`
	Evse                 *EVSE       `protobuf:"bytes,6,opt,name=evse,proto3" json:"evse,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *XBOSIoTDeviceActuation) Reset()         { *m = XBOSIoTDeviceActuation{} }
func (m *XBOSIoTDeviceActuation) String() string { return proto.CompactTextString(m) }
func (*XBOSIoTDeviceActuation) ProtoMessage()    {}
func (*XBOSIoTDeviceActuation) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{4}
}

func (m *XBOSIoTDeviceActuation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XBOSIoTDeviceActuation.Unmarshal(m, b)
}
func (m *XBOSIoTDeviceActuation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XBOSIoTDeviceActuation.Marshal(b, m, deterministic)
}
func (m *XBOSIoTDeviceActuation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XBOSIoTDeviceActuation.Merge(m, src)
}
func (m *XBOSIoTDeviceActuation) XXX_Size() int {
	return xxx_messageInfo_XBOSIoTDeviceActuation.Size(m)
}
func (m *XBOSIoTDeviceActuation) XXX_DiscardUnknown() {
	xxx_messageInfo_XBOSIoTDeviceActuation.DiscardUnknown(m)
}

var xxx_messageInfo_XBOSIoTDeviceActuation proto.InternalMessageInfo

func (m *XBOSIoTDeviceActuation) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *XBOSIoTDeviceActuation) GetRequestid() int64 {
	if m != nil {
		return m.Requestid
	}
	return 0
}

func (m *XBOSIoTDeviceActuation) GetThermostat() *Thermostat {
	if m != nil {
		return m.Thermostat
	}
	return nil
}

func (m *XBOSIoTDeviceActuation) GetMeter() *Meter {
	if m != nil {
		return m.Meter
	}
	return nil
}

func (m *XBOSIoTDeviceActuation) GetLight() *Light {
	if m != nil {
		return m.Light
	}
	return nil
}

func (m *XBOSIoTDeviceActuation) GetEvse() *EVSE {
	if m != nil {
		return m.Evse
	}
	return nil
}

type XBOSIoTContext struct {
	// current time at device/service
	//unit:ns
	Time uint64 `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	// any triples this device wants to add about itself
	// these triples will be assumed to be generated by the entity
	// who has created/signed this message
	Context              []*Triple `protobuf:"bytes,2,rep,name=context,proto3" json:"context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *XBOSIoTContext) Reset()         { *m = XBOSIoTContext{} }
func (m *XBOSIoTContext) String() string { return proto.CompactTextString(m) }
func (*XBOSIoTContext) ProtoMessage()    {}
func (*XBOSIoTContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{5}
}

func (m *XBOSIoTContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XBOSIoTContext.Unmarshal(m, b)
}
func (m *XBOSIoTContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XBOSIoTContext.Marshal(b, m, deterministic)
}
func (m *XBOSIoTContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XBOSIoTContext.Merge(m, src)
}
func (m *XBOSIoTContext) XXX_Size() int {
	return xxx_messageInfo_XBOSIoTContext.Size(m)
}
func (m *XBOSIoTContext) XXX_DiscardUnknown() {
	xxx_messageInfo_XBOSIoTContext.DiscardUnknown(m)
}

var xxx_messageInfo_XBOSIoTContext proto.InternalMessageInfo

func (m *XBOSIoTContext) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *XBOSIoTContext) GetContext() []*Triple {
	if m != nil {
		return m.Context
	}
	return nil
}

// Thermostat
type Thermostat struct {
	//Current temperature recorded by thermostat
	//unit:celsius
	Temperature *Double `protobuf:"bytes,1,opt,name=temperature,proto3" json:"temperature,omitempty"`
	//unit:% rh
	RelativeHumidity *Double `protobuf:"bytes,2,opt,name=relative_humidity,json=relativeHumidity,proto3" json:"relative_humidity,omitempty"`
	//If true, then the thermostat is in override mode
	//unit: t/f
	Override *Bool `protobuf:"bytes,3,opt,name=override,proto3" json:"override,omitempty"`
	//If true, the fan is on; else it is off
	//unit: t/f
	FanState *Bool `protobuf:"bytes,4,opt,name=fan_state,json=fanState,proto3" json:"fan_state,omitempty"`
	//Current operating mode of the fan
	//unit: xbos/iot/FanMode
	FanMode *FanMode `protobuf:"bytes,5,opt,name=fan_mode,json=fanMode,proto3" json:"fan_mode,omitempty"`
	//Current operating mode of the HVAC
	//unit: xbos/iot/HVACMode
	Mode *HVACMode `protobuf:"bytes,6,opt,name=mode,proto3" json:"mode,omitempty"`
	//Current HVAC state
	//unit: xbos/iot/HVACState
	State *HVACState `protobuf:"bytes,7,opt,name=state,proto3" json:"state,omitempty"`
	//number of heat stages enabled
	EnabledHeatStages *Int32 `protobuf:"bytes,8,opt,name=enabled_heat_stages,json=enabledHeatStages,proto3" json:"enabled_heat_stages,omitempty"`
	//number of cool stages enabled
	EnabledCoolStages *Int32 `protobuf:"bytes,9,opt,name=enabled_cool_stages,json=enabledCoolStages,proto3" json:"enabled_cool_stages,omitempty"`
	//heating setpoint
	//unit: celsius
	HeatingSetpoint *Double `protobuf:"bytes,10,opt,name=heating_setpoint,json=heatingSetpoint,proto3" json:"heating_setpoint,omitempty"`
	//cooling setpoint
	//unit: celsius
	CoolingSetpoint      *Double  `protobuf:"bytes,11,opt,name=cooling_setpoint,json=coolingSetpoint,proto3" json:"cooling_setpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Thermostat) Reset()         { *m = Thermostat{} }
func (m *Thermostat) String() string { return proto.CompactTextString(m) }
func (*Thermostat) ProtoMessage()    {}
func (*Thermostat) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{6}
}

func (m *Thermostat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Thermostat.Unmarshal(m, b)
}
func (m *Thermostat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Thermostat.Marshal(b, m, deterministic)
}
func (m *Thermostat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Thermostat.Merge(m, src)
}
func (m *Thermostat) XXX_Size() int {
	return xxx_messageInfo_Thermostat.Size(m)
}
func (m *Thermostat) XXX_DiscardUnknown() {
	xxx_messageInfo_Thermostat.DiscardUnknown(m)
}

var xxx_messageInfo_Thermostat proto.InternalMessageInfo

func (m *Thermostat) GetTemperature() *Double {
	if m != nil {
		return m.Temperature
	}
	return nil
}

func (m *Thermostat) GetRelativeHumidity() *Double {
	if m != nil {
		return m.RelativeHumidity
	}
	return nil
}

func (m *Thermostat) GetOverride() *Bool {
	if m != nil {
		return m.Override
	}
	return nil
}

func (m *Thermostat) GetFanState() *Bool {
	if m != nil {
		return m.FanState
	}
	return nil
}

func (m *Thermostat) GetFanMode() *FanMode {
	if m != nil {
		return m.FanMode
	}
	return nil
}

func (m *Thermostat) GetMode() *HVACMode {
	if m != nil {
		return m.Mode
	}
	return nil
}

func (m *Thermostat) GetState() *HVACState {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *Thermostat) GetEnabledHeatStages() *Int32 {
	if m != nil {
		return m.EnabledHeatStages
	}
	return nil
}

func (m *Thermostat) GetEnabledCoolStages() *Int32 {
	if m != nil {
		return m.EnabledCoolStages
	}
	return nil
}

func (m *Thermostat) GetHeatingSetpoint() *Double {
	if m != nil {
		return m.HeatingSetpoint
	}
	return nil
}

func (m *Thermostat) GetCoolingSetpoint() *Double {
	if m != nil {
		return m.CoolingSetpoint
	}
	return nil
}

type ThermostatSchedule struct {
	//Map day of week to daily schedule
	ScheduleMap          map[string]*ThermostatScheduleDay `protobuf:"bytes,1,rep,name=scheduleMap,proto3" json:"scheduleMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *ThermostatSchedule) Reset()         { *m = ThermostatSchedule{} }
func (m *ThermostatSchedule) String() string { return proto.CompactTextString(m) }
func (*ThermostatSchedule) ProtoMessage()    {}
func (*ThermostatSchedule) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{7}
}

func (m *ThermostatSchedule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThermostatSchedule.Unmarshal(m, b)
}
func (m *ThermostatSchedule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThermostatSchedule.Marshal(b, m, deterministic)
}
func (m *ThermostatSchedule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThermostatSchedule.Merge(m, src)
}
func (m *ThermostatSchedule) XXX_Size() int {
	return xxx_messageInfo_ThermostatSchedule.Size(m)
}
func (m *ThermostatSchedule) XXX_DiscardUnknown() {
	xxx_messageInfo_ThermostatSchedule.DiscardUnknown(m)
}

var xxx_messageInfo_ThermostatSchedule proto.InternalMessageInfo

func (m *ThermostatSchedule) GetScheduleMap() map[string]*ThermostatScheduleDay {
	if m != nil {
		return m.ScheduleMap
	}
	return nil
}

type ThermostatScheduleDay struct {
	//Daily schedule is Multiple Blocks
	Blocks               []*ThermostatScheduleBlock `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ThermostatScheduleDay) Reset()         { *m = ThermostatScheduleDay{} }
func (m *ThermostatScheduleDay) String() string { return proto.CompactTextString(m) }
func (*ThermostatScheduleDay) ProtoMessage()    {}
func (*ThermostatScheduleDay) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{8}
}

func (m *ThermostatScheduleDay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThermostatScheduleDay.Unmarshal(m, b)
}
func (m *ThermostatScheduleDay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThermostatScheduleDay.Marshal(b, m, deterministic)
}
func (m *ThermostatScheduleDay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThermostatScheduleDay.Merge(m, src)
}
func (m *ThermostatScheduleDay) XXX_Size() int {
	return xxx_messageInfo_ThermostatScheduleDay.Size(m)
}
func (m *ThermostatScheduleDay) XXX_DiscardUnknown() {
	xxx_messageInfo_ThermostatScheduleDay.DiscardUnknown(m)
}

var xxx_messageInfo_ThermostatScheduleDay proto.InternalMessageInfo

func (m *ThermostatScheduleDay) GetBlocks() []*ThermostatScheduleBlock {
	if m != nil {
		return m.Blocks
	}
	return nil
}

type ThermostatScheduleBlock struct {
	//heating setpoint
	//unit: celsius
	HeatingSetpoint *Double `protobuf:"bytes,1,opt,name=heating_setpoint,json=heatingSetpoint,proto3" json:"heating_setpoint,omitempty"`
	//cooling setpoint
	//unit: celsius
	CoolingSetpoint *Double `protobuf:"bytes,2,opt,name=cooling_setpoint,json=coolingSetpoint,proto3" json:"cooling_setpoint,omitempty"`
	//Current system mode of thermostat
	//unit: xbos/iot/HVACMode
	Mode *HVACMode `protobuf:"bytes,3,opt,name=mode,proto3" json:"mode,omitempty"`
	//Time when schedule block takes effect
	//format: RRule
	Time                 string   `protobuf:"bytes,4,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThermostatScheduleBlock) Reset()         { *m = ThermostatScheduleBlock{} }
func (m *ThermostatScheduleBlock) String() string { return proto.CompactTextString(m) }
func (*ThermostatScheduleBlock) ProtoMessage()    {}
func (*ThermostatScheduleBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{9}
}

func (m *ThermostatScheduleBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThermostatScheduleBlock.Unmarshal(m, b)
}
func (m *ThermostatScheduleBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThermostatScheduleBlock.Marshal(b, m, deterministic)
}
func (m *ThermostatScheduleBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThermostatScheduleBlock.Merge(m, src)
}
func (m *ThermostatScheduleBlock) XXX_Size() int {
	return xxx_messageInfo_ThermostatScheduleBlock.Size(m)
}
func (m *ThermostatScheduleBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_ThermostatScheduleBlock.DiscardUnknown(m)
}

var xxx_messageInfo_ThermostatScheduleBlock proto.InternalMessageInfo

func (m *ThermostatScheduleBlock) GetHeatingSetpoint() *Double {
	if m != nil {
		return m.HeatingSetpoint
	}
	return nil
}

func (m *ThermostatScheduleBlock) GetCoolingSetpoint() *Double {
	if m != nil {
		return m.CoolingSetpoint
	}
	return nil
}

func (m *ThermostatScheduleBlock) GetMode() *HVACMode {
	if m != nil {
		return m.Mode
	}
	return nil
}

func (m *ThermostatScheduleBlock) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

type Meter struct {
	//unit: kW
	Power *Double `protobuf:"bytes,1,opt,name=power,proto3" json:"power,omitempty"`
	//unit: V
	Voltage *Double `protobuf:"bytes,2,opt,name=voltage,proto3" json:"voltage,omitempty"`
	//unit: kVA
	ApparentPower *Double `protobuf:"bytes,3,opt,name=apparent_power,json=apparentPower,proto3" json:"apparent_power,omitempty"`
	//unit: KWh
	Energy               *Double  `protobuf:"bytes,4,opt,name=energy,proto3" json:"energy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Meter) Reset()         { *m = Meter{} }
func (m *Meter) String() string { return proto.CompactTextString(m) }
func (*Meter) ProtoMessage()    {}
func (*Meter) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{10}
}

func (m *Meter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Meter.Unmarshal(m, b)
}
func (m *Meter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Meter.Marshal(b, m, deterministic)
}
func (m *Meter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Meter.Merge(m, src)
}
func (m *Meter) XXX_Size() int {
	return xxx_messageInfo_Meter.Size(m)
}
func (m *Meter) XXX_DiscardUnknown() {
	xxx_messageInfo_Meter.DiscardUnknown(m)
}

var xxx_messageInfo_Meter proto.InternalMessageInfo

func (m *Meter) GetPower() *Double {
	if m != nil {
		return m.Power
	}
	return nil
}

func (m *Meter) GetVoltage() *Double {
	if m != nil {
		return m.Voltage
	}
	return nil
}

func (m *Meter) GetApparentPower() *Double {
	if m != nil {
		return m.ApparentPower
	}
	return nil
}

func (m *Meter) GetEnergy() *Double {
	if m != nil {
		return m.Energy
	}
	return nil
}

type Light struct {
	// True if the light is on
	//unit: on/off
	State *Bool `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	// 100 is maximum brightness
	Brightness           *Int64   `protobuf:"bytes,2,opt,name=brightness,proto3" json:"brightness,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Light) Reset()         { *m = Light{} }
func (m *Light) String() string { return proto.CompactTextString(m) }
func (*Light) ProtoMessage()    {}
func (*Light) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{11}
}

func (m *Light) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Light.Unmarshal(m, b)
}
func (m *Light) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Light.Marshal(b, m, deterministic)
}
func (m *Light) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Light.Merge(m, src)
}
func (m *Light) XXX_Size() int {
	return xxx_messageInfo_Light.Size(m)
}
func (m *Light) XXX_DiscardUnknown() {
	xxx_messageInfo_Light.DiscardUnknown(m)
}

var xxx_messageInfo_Light proto.InternalMessageInfo

func (m *Light) GetState() *Bool {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *Light) GetBrightness() *Int64 {
	if m != nil {
		return m.Brightness
	}
	return nil
}

type EVSE struct {
	// maximum allowed current for charging
	//unit: A
	CurrentLimit *Double `protobuf:"bytes,1,opt,name=current_limit,json=currentLimit,proto3" json:"current_limit,omitempty"`
	// active charge current
	//unit: A
	Current *Double `protobuf:"bytes,2,opt,name=current,proto3" json:"current,omitempty"`
	// active charge voltage
	//unit: V
	Voltage *Double `protobuf:"bytes,3,opt,name=voltage,proto3" json:"voltage,omitempty"`
	// seconds left until car is charged
	//unit: seconds
	ChargingTimeLeft *Int32 `protobuf:"bytes,4,opt,name=charging_time_left,json=chargingTimeLeft,proto3" json:"charging_time_left,omitempty"`
	// charge state of the EVSE
	//unit: on/off
	State                *Bool    `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EVSE) Reset()         { *m = EVSE{} }
func (m *EVSE) String() string { return proto.CompactTextString(m) }
func (*EVSE) ProtoMessage()    {}
func (*EVSE) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{12}
}

func (m *EVSE) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EVSE.Unmarshal(m, b)
}
func (m *EVSE) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EVSE.Marshal(b, m, deterministic)
}
func (m *EVSE) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EVSE.Merge(m, src)
}
func (m *EVSE) XXX_Size() int {
	return xxx_messageInfo_EVSE.Size(m)
}
func (m *EVSE) XXX_DiscardUnknown() {
	xxx_messageInfo_EVSE.DiscardUnknown(m)
}

var xxx_messageInfo_EVSE proto.InternalMessageInfo

func (m *EVSE) GetCurrentLimit() *Double {
	if m != nil {
		return m.CurrentLimit
	}
	return nil
}

func (m *EVSE) GetCurrent() *Double {
	if m != nil {
		return m.Current
	}
	return nil
}

func (m *EVSE) GetVoltage() *Double {
	if m != nil {
		return m.Voltage
	}
	return nil
}

func (m *EVSE) GetChargingTimeLeft() *Int32 {
	if m != nil {
		return m.ChargingTimeLeft
	}
	return nil
}

func (m *EVSE) GetState() *Bool {
	if m != nil {
		return m.State
	}
	return nil
}

var E_BrickEquipClass = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*URI)(nil),
	Field:         10000,
	Name:          "xbospb.brick_equip_class",
	Tag:           "bytes,10000,opt,name=brick_equip_class",
	Filename:      "iot.proto",
}

var E_BrickPointClass = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*URI)(nil),
	Field:         50000,
	Name:          "xbospb.brick_point_class",
	Tag:           "bytes,50000,opt,name=brick_point_class",
	Filename:      "iot.proto",
}

func init() {
	proto.RegisterType((*URI)(nil), "xbospb.URI")
	proto.RegisterType((*Triple)(nil), "xbospb.Triple")
	proto.RegisterType((*Error)(nil), "xbospb.Error")
	proto.RegisterType((*XBOSIoTDeviceState)(nil), "xbospb.XBOSIoTDeviceState")
	proto.RegisterType((*XBOSIoTDeviceActuation)(nil), "xbospb.XBOSIoTDeviceActuation")
	proto.RegisterType((*XBOSIoTContext)(nil), "xbospb.XBOSIoTContext")
	proto.RegisterType((*Thermostat)(nil), "xbospb.Thermostat")
	proto.RegisterType((*ThermostatSchedule)(nil), "xbospb.ThermostatSchedule")
	proto.RegisterMapType((map[string]*ThermostatScheduleDay)(nil), "xbospb.ThermostatSchedule.ScheduleMapEntry")
	proto.RegisterType((*ThermostatScheduleDay)(nil), "xbospb.ThermostatScheduleDay")
	proto.RegisterType((*ThermostatScheduleBlock)(nil), "xbospb.ThermostatScheduleBlock")
	proto.RegisterType((*Meter)(nil), "xbospb.Meter")
	proto.RegisterType((*Light)(nil), "xbospb.Light")
	proto.RegisterType((*EVSE)(nil), "xbospb.EVSE")
	proto.RegisterExtension(E_BrickEquipClass)
	proto.RegisterExtension(E_BrickPointClass)
}

func init() { proto.RegisterFile("iot.proto", fileDescriptor_1804728d01c3c43d) }

var fileDescriptor_1804728d01c3c43d = []byte{
	// 1309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x97, 0x4b, 0x73, 0xdb, 0x54,
	0x14, 0xc7, 0xc7, 0xf1, 0xab, 0x3e, 0xce, 0xc3, 0xb9, 0x7d, 0x89, 0xf4, 0x15, 0x54, 0x60, 0xca,
	0x30, 0xa4, 0x33, 0x09, 0x03, 0x34, 0x0c, 0xc3, 0xd4, 0x6e, 0x32, 0xc9, 0x34, 0x69, 0x83, 0x9c,
	0x96, 0xb2, 0x28, 0x1e, 0x59, 0x3e, 0xb1, 0x45, 0x24, 0x5d, 0xf5, 0xea, 0x2a, 0xa9, 0x57, 0xcc,
	0x74, 0xc1, 0x9a, 0xcf, 0xc0, 0x9a, 0x6d, 0x57, 0xec, 0xd8, 0xb1, 0x63, 0xc5, 0x87, 0x60, 0x45,
	0xbf, 0x01, 0x73, 0x5f, 0xb6, 0x65, 0xcb, 0x4c, 0xbb, 0x93, 0xee, 0xf9, 0x9f, 0xdf, 0x7d, 0x9c,
	0xc7, 0x95, 0xa0, 0xe6, 0x53, 0xbe, 0x11, 0x33, 0xca, 0x29, 0xa9, 0xbc, 0xec, 0xd2, 0x24, 0xee,
	0xae, 0x5d, 0x8c, 0xd2, 0x20, 0x70, 0xbb, 0x01, 0xf2, 0x61, 0x8c, 0x89, 0x32, 0xae, 0x5d, 0x3e,
	0x47, 0x97, 0x0f, 0x90, 0x75, 0x12, 0xee, 0x72, 0x9f, 0x46, 0x7a, 0x78, 0xbd, 0x4f, 0x69, 0x3f,
	0xc0, 0xbb, 0xf2, 0xad, 0x9b, 0x9e, 0xdc, 0xed, 0x61, 0xe2, 0x31, 0x3f, 0xe6, 0x94, 0x29, 0x85,
	0x7d, 0x0f, 0x8a, 0x4f, 0x9c, 0x7d, 0x72, 0x1d, 0x6a, 0x91, 0x1b, 0x62, 0x12, 0xbb, 0x1e, 0x5a,
	0x85, 0xf5, 0xc2, 0x9d, 0x9a, 0x33, 0x1e, 0x20, 0x97, 0xa0, 0x7c, 0xe6, 0x06, 0x29, 0x5a, 0x0b,
	0xd2, 0xa2, 0x5e, 0xec, 0x9f, 0xa0, 0x72, 0xcc, 0xfc, 0x38, 0x40, 0xf2, 0x21, 0x54, 0x93, 0xb4,
	0xfb, 0x23, 0x7a, 0x5c, 0xfa, 0xd6, 0x37, 0xeb, 0x1b, 0x6a, 0xb1, 0x1b, 0x4f, 0x9c, 0x7d, 0xc7,
	0xd8, 0xc8, 0xc7, 0x50, 0x8b, 0x19, 0xf6, 0x7c, 0xcf, 0xe5, 0x0a, 0x35, 0x25, 0x1c, 0x5b, 0xc9,
	0x6d, 0xa8, 0x50, 0x05, 0x2c, 0xce, 0xea, 0xb4, 0xc9, 0x7e, 0x0f, 0xca, 0x3b, 0x8c, 0x51, 0x46,
	0x1a, 0x50, 0x0c, 0x93, 0xbe, 0x5e, 0xb7, 0x78, 0xb4, 0x7f, 0x2b, 0x02, 0x79, 0xd6, 0x7c, 0xdc,
	0xde, 0xa7, 0xc7, 0x0f, 0xf0, 0xcc, 0xf7, 0xb0, 0xcd, 0x05, 0x96, 0x40, 0x89, 0xfb, 0xa1, 0xda,
	0x61, 0xc9, 0x91, 0xcf, 0x62, 0xeb, 0x0c, 0x5f, 0xa4, 0x98, 0x70, 0xbf, 0x27, 0x57, 0x55, 0x74,
	0xc6, 0x03, 0xe4, 0x36, 0x94, 0x51, 0xcc, 0xa1, 0xd7, 0xb1, 0x64, 0xd6, 0x21, 0x27, 0x76, 0x94,
	0x8d, 0x6c, 0x02, 0x88, 0xc3, 0x0f, 0xa9, 0x38, 0x7d, 0xab, 0x24, 0x95, 0xc4, 0x28, 0x8f, 0x47,
	0x16, 0x67, 0x42, 0x25, 0xc0, 0x21, 0x72, 0x64, 0x56, 0x39, 0x0b, 0x3e, 0x14, 0x83, 0x8e, 0xb2,
	0x09, 0x51, 0xe0, 0xf7, 0x07, 0xdc, 0xaa, 0x64, 0x45, 0x07, 0x62, 0xd0, 0x51, 0x36, 0xb2, 0x0e,
	0x25, 0x3c, 0x4b, 0xd0, 0xaa, 0x4a, 0xcd, 0xe2, 0x68, 0x85, 0x4f, 0xdb, 0x3b, 0x8e, 0xb4, 0x90,
	0x6f, 0x60, 0x65, 0x2a, 0x3f, 0xac, 0x0b, 0x52, 0x7c, 0xc5, 0x88, 0xbf, 0x53, 0xe6, 0xb6, 0xb2,
	0x3a, 0xcb, 0xe7, 0x99, 0x77, 0xf2, 0x03, 0xac, 0x4d, 0x01, 0x3a, 0x2a, 0x56, 0x92, 0x55, 0x93,
	0xac, 0xf5, 0x7c, 0xd6, 0xd1, 0x48, 0xe7, 0x58, 0xe7, 0x73, 0x2c, 0xf6, 0x3f, 0x05, 0xb8, 0x92,
	0x09, 0xd7, 0x7d, 0x8f, 0xa7, 0x6a, 0xea, 0x77, 0x0f, 0x59, 0x36, 0x1a, 0xc5, 0x77, 0x8b, 0x46,
	0xe9, 0x6d, 0xa2, 0x51, 0x7e, 0x8b, 0x68, 0x54, 0xe6, 0x45, 0xc3, 0x7e, 0x04, 0xcb, 0x7a, 0xaf,
	0x2d, 0x1a, 0x71, 0x7c, 0xc9, 0x73, 0xf7, 0x78, 0x07, 0xaa, 0x9e, 0x32, 0x5b, 0x0b, 0xeb, 0xc5,
	0x3b, 0xf5, 0xcd, 0xe5, 0xd1, 0x16, 0x64, 0xd1, 0x39, 0xc6, 0x6c, 0xff, 0x5a, 0x05, 0x18, 0x6f,
	0x8b, 0x3c, 0x81, 0x3a, 0xc7, 0x30, 0x46, 0xe6, 0xf2, 0x94, 0xa1, 0x2e, 0xc8, 0x91, 0xf3, 0x03,
	0x9a, 0x76, 0x03, 0x6c, 0xde, 0x7e, 0xf5, 0xda, 0xaa, 0x42, 0xb9, 0xcb, 0x7c, 0xef, 0xf4, 0xd5,
	0x6b, 0xeb, 0x12, 0x21, 0xc7, 0x63, 0x9f, 0x4e, 0x1b, 0xa3, 0x84, 0x32, 0x67, 0x92, 0x43, 0x9e,
	0xc3, 0x2a, 0xc3, 0xc0, 0xe5, 0xfe, 0x19, 0x76, 0x06, 0x69, 0xe8, 0xf7, 0x7c, 0x3e, 0xd4, 0x45,
	0x3c, 0x0d, 0x5f, 0xcf, 0xc2, 0x57, 0xc9, 0xca, 0x9e, 0x76, 0x30, 0xe4, 0x86, 0x41, 0x19, 0x03,
	0x79, 0x08, 0x17, 0xe8, 0x19, 0x32, 0xe6, 0xf7, 0x50, 0x87, 0x6c, 0x74, 0x74, 0x4d, 0x4a, 0x83,
	0xe6, 0xfb, 0x59, 0x26, 0x21, 0x8d, 0xc7, 0x5a, 0xde, 0x69, 0xd1, 0x30, 0x74, 0xa3, 0x9e, 0x33,
	0x02, 0x90, 0x5d, 0xa8, 0x9d, 0xb8, 0x91, 0x4c, 0x55, 0xd4, 0x11, 0xcd, 0xd2, 0xae, 0x65, 0x69,
	0x8b, 0x04, 0x76, 0xdd, 0xa8, 0x23, 0xf2, 0x32, 0x4d, 0x9c, 0x0b, 0x27, 0x6e, 0xa4, 0xda, 0xc5,
	0x1e, 0x88, 0xe7, 0x4e, 0x48, 0x7b, 0xa8, 0x63, 0xbe, 0x62, 0x30, 0xbb, 0x6e, 0x74, 0x48, 0x7b,
	0xd8, 0xbc, 0x9e, 0x25, 0x2d, 0x91, 0xba, 0x20, 0x99, 0x25, 0x55, 0x4f, 0x94, 0x8c, 0xb4, 0xa0,
	0x24, 0x29, 0x2a, 0x2b, 0x1a, 0x86, 0xb2, 0xf7, 0xf4, 0x7e, 0x4b, 0x62, 0x6e, 0x64, 0x31, 0xcb,
	0x64, 0x51, 0x0c, 0x8f, 0x38, 0xd2, 0x99, 0x3c, 0x84, 0xb2, 0xda, 0x92, 0xaa, 0xf4, 0xd5, 0x49,
	0x8a, 0x5c, 0x70, 0xd3, 0xce, 0x62, 0x2e, 0x92, 0xd5, 0x71, 0x62, 0x98, 0xed, 0x29, 0x06, 0xf9,
	0x1a, 0x2e, 0x62, 0x24, 0xae, 0x91, 0x5e, 0x67, 0x80, 0x2e, 0x17, 0x87, 0xd5, 0xc7, 0x44, 0xf7,
	0x85, 0x51, 0x6a, 0xef, 0x47, 0x7c, 0x6b, 0xd3, 0x59, 0xd5, 0xca, 0x3d, 0x74, 0x79, 0x5b, 0xea,
	0x26, 0xdd, 0x3d, 0x4a, 0x03, 0xe3, 0x5e, 0xfb, 0x3f, 0xf7, 0x16, 0xa5, 0x81, 0x76, 0x1f, 0x40,
	0x43, 0xcc, 0xea, 0x47, 0xfd, 0x4e, 0x82, 0x3c, 0xa6, 0x7e, 0xc4, 0x2d, 0xc8, 0x4d, 0xa6, 0x4f,
	0xb3, 0x5b, 0xba, 0x49, 0xae, 0xef, 0x69, 0xc7, 0x6c, 0xc6, 0x2a, 0x88, 0xb3, 0xa2, 0xb1, 0x66,
	0x40, 0xcc, 0x24, 0x16, 0x98, 0x99, 0xa9, 0xfe, 0x96, 0x33, 0xb5, 0xb4, 0x63, 0xfe, 0x4c, 0x1a,
	0x6b, 0x06, 0xb6, 0xaf, 0xbd, 0x7a, 0x53, 0x1a, 0xf9, 0xbf, 0x29, 0x2d, 0x92, 0x89, 0xaa, 0xb4,
	0xff, 0x28, 0x00, 0x19, 0xbf, 0xb6, 0xbd, 0x01, 0xf6, 0xd2, 0x00, 0xc9, 0x21, 0xd4, 0x13, 0xfd,
	0x7c, 0xe8, 0xc6, 0x56, 0x41, 0x56, 0xfa, 0x27, 0xb3, 0xcd, 0xca, 0x38, 0x6c, 0xb4, 0xc7, 0xea,
	0x9d, 0x88, 0xb3, 0xa1, 0x33, 0xe9, 0xbf, 0xf6, 0x1c, 0x1a, 0xd3, 0x02, 0x71, 0x39, 0x9e, 0xe2,
	0xd0, 0x5c, 0x8e, 0xa7, 0x38, 0x24, 0x5b, 0x93, 0xd7, 0x79, 0x7d, 0xf3, 0xc6, 0xfc, 0xe9, 0x1e,
	0xb8, 0x43, 0x7d, 0xdb, 0x6f, 0x2f, 0x7c, 0x59, 0xb0, 0x8f, 0xe0, 0x72, 0xae, 0x86, 0x7c, 0x01,
	0x95, 0x6e, 0x40, 0xbd, 0xd3, 0x44, 0xef, 0xe0, 0xd6, 0x7c, 0x64, 0x53, 0xe8, 0x1c, 0x2d, 0xb7,
	0xff, 0x2c, 0xc0, 0xd5, 0x39, 0x1a, 0x72, 0x2f, 0x27, 0x47, 0x72, 0xbb, 0xd9, 0x6c, 0xd0, 0xef,
	0xe5, 0x04, 0x7d, 0x21, 0xdf, 0x75, 0x2a, 0x8a, 0xe4, 0x03, 0x5d, 0xa9, 0xc5, 0xfc, 0x4a, 0xd5,
	0xa5, 0x68, 0x3a, 0x76, 0x49, 0x9e, 0xaa, 0x7c, 0xb6, 0xff, 0x5e, 0x80, 0xb2, 0xbc, 0x2f, 0x48,
	0x0b, 0xca, 0x31, 0x3d, 0x47, 0x36, 0xa7, 0xf9, 0xce, 0x16, 0xfb, 0x91, 0x50, 0x9b, 0xe6, 0xa8,
	0x7c, 0xc9, 0x3e, 0x54, 0xcf, 0x68, 0x20, 0xca, 0x65, 0x4e, 0x9b, 0xbd, 0x95, 0xc5, 0x34, 0xc8,
	0xf2, 0x53, 0xa5, 0x37, 0x20, 0xe3, 0x4f, 0x9e, 0xc3, 0xb2, 0x1b, 0xc7, 0x2e, 0xc3, 0x88, 0x77,
	0xd4, 0xc2, 0x8a, 0xb9, 0xc4, 0x8f, 0xb2, 0xc4, 0xab, 0xe4, 0xf2, 0x7d, 0xe3, 0x96, 0x59, 0xe1,
	0x92, 0xa1, 0xc9, 0x51, 0xb2, 0x0b, 0x15, 0x8c, 0x90, 0xf5, 0x87, 0xba, 0xd7, 0x4e, 0x63, 0x6f,
	0x66, 0xb1, 0x2b, 0x64, 0x69, 0x47, 0xca, 0x0d, 0x4e, 0x7b, 0x6f, 0x5f, 0xc9, 0x16, 0x50, 0x95,
	0xa8, 0xe3, 0xb4, 0x7f, 0x2f, 0x40, 0x59, 0xde, 0xb1, 0xe2, 0x60, 0x55, 0x07, 0x2c, 0xe4, 0x34,
	0xf5, 0xd9, 0xf3, 0x90, 0x4e, 0x22, 0x09, 0x64, 0x77, 0x34, 0x9d, 0xef, 0x5b, 0x80, 0x2e, 0x13,
	0x96, 0x08, 0x93, 0x44, 0x9f, 0xed, 0x64, 0xc7, 0xfa, 0xfc, 0xb3, 0x9c, 0x3e, 0x7a, 0x90, 0x86,
	0x7e, 0xe4, 0x46, 0xde, 0xb8, 0x27, 0x4f, 0x40, 0xb6, 0xd7, 0xb2, 0x2b, 0xaf, 0x93, 0x9a, 0xd2,
	0xfb, 0x0c, 0xed, 0x7f, 0x0b, 0x50, 0x12, 0xb7, 0x3f, 0xd9, 0x82, 0x25, 0x2f, 0x65, 0xf2, 0x34,
	0x03, 0x3f, 0xf4, 0xe7, 0x25, 0xf3, 0xa2, 0x16, 0x1d, 0x08, 0x8d, 0xfc, 0x0c, 0x50, 0xef, 0x73,
	0x12, 0xd8, 0x98, 0x85, 0xd2, 0xe4, 0x4b, 0x31, 0x5f, 0x69, 0xd2, 0xe1, 0x2b, 0x20, 0xde, 0xc0,
	0x65, 0x7d, 0x71, 0x32, 0x22, 0x73, 0x3b, 0x01, 0x9e, 0xf0, 0xe9, 0x2f, 0x1f, 0xd5, 0xba, 0x1b,
	0x46, 0x78, 0xec, 0x87, 0x78, 0x80, 0x27, 0x9c, 0xd8, 0x26, 0x04, 0xe5, 0xd9, 0x10, 0xe8, 0x13,
	0xde, 0xfe, 0x1e, 0x56, 0xe5, 0x51, 0x74, 0xf0, 0x45, 0xea, 0xc7, 0x1d, 0x2f, 0x70, 0x93, 0x84,
	0xdc, 0xda, 0x50, 0x3f, 0x23, 0x1b, 0xe6, 0x67, 0x64, 0xe3, 0x10, 0x93, 0xc4, 0xed, 0xe3, 0xe3,
	0x58, 0x7c, 0xe8, 0x25, 0xd6, 0x2f, 0x8f, 0x66, 0x3f, 0xf5, 0x57, 0x24, 0x67, 0x47, 0x60, 0x5a,
	0x82, 0xb2, 0xfd, 0xcc, 0xa0, 0x65, 0xb5, 0x6a, 0xf4, 0x8d, 0x19, 0xf4, 0xae, 0x8f, 0x41, 0xcf,
	0x80, 0xff, 0xfa, 0xb9, 0x38, 0x8f, 0x7c, 0x24, 0x28, 0x92, 0xdc, 0xad, 0x48, 0xe7, 0xad, 0xff,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x7c, 0x4b, 0xd9, 0x73, 0x0d, 0x00, 0x00,
}
