// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package hydraThrift

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type HydraThriftService interface {
	// Parameters:
	//  - S
	AddShip(s *Ship) (err error)
}

type HydraThriftServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewHydraThriftServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *HydraThriftServiceClient {
	return &HydraThriftServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewHydraThriftServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *HydraThriftServiceClient {
	return &HydraThriftServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - S
func (p *HydraThriftServiceClient) AddShip(s *Ship) (err error) {
	if err = p.sendAddShip(s); err != nil {
		return
	}
	return p.recvAddShip()
}

func (p *HydraThriftServiceClient) sendAddShip(s *Ship) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("addShip", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := HydraThriftServiceAddShipArgs{
		S: s,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *HydraThriftServiceClient) recvAddShip() (err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "addShip" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "addShip failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "addShip failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error1 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error2 error
		error2, err = error1.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error2
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "addShip failed: invalid message type")
		return
	}
	result := HydraThriftServiceAddShipResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	return
}

type HydraThriftServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      HydraThriftService
}

func (p *HydraThriftServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *HydraThriftServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *HydraThriftServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewHydraThriftServiceProcessor(handler HydraThriftService) *HydraThriftServiceProcessor {

	self3 := &HydraThriftServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self3.processorMap["addShip"] = &hydraThriftServiceProcessorAddShip{handler: handler}
	return self3
}

func (p *HydraThriftServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x4 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x4.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x4

}

type hydraThriftServiceProcessorAddShip struct {
	handler HydraThriftService
}

func (p *hydraThriftServiceProcessorAddShip) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := HydraThriftServiceAddShipArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("addShip", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := HydraThriftServiceAddShipResult{}
	var err2 error
	if err2 = p.handler.AddShip(args.S); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing addShip: "+err2.Error())
		oprot.WriteMessageBegin("addShip", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	}
	if err2 = oprot.WriteMessageBegin("addShip", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - S
type HydraThriftServiceAddShipArgs struct {
	S *Ship `thrift:"s,1" json:"s"`
}

func NewHydraThriftServiceAddShipArgs() *HydraThriftServiceAddShipArgs {
	return &HydraThriftServiceAddShipArgs{}
}

var HydraThriftServiceAddShipArgs_S_DEFAULT *Ship

func (p *HydraThriftServiceAddShipArgs) GetS() *Ship {
	if !p.IsSetS() {
		return HydraThriftServiceAddShipArgs_S_DEFAULT
	}
	return p.S
}
func (p *HydraThriftServiceAddShipArgs) IsSetS() bool {
	return p.S != nil
}

func (p *HydraThriftServiceAddShipArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *HydraThriftServiceAddShipArgs) readField1(iprot thrift.TProtocol) error {
	p.S = &Ship{}
	if err := p.S.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.S), err)
	}
	return nil
}

func (p *HydraThriftServiceAddShipArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("addShip_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *HydraThriftServiceAddShipArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("s", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:s: ", p), err)
	}
	if err := p.S.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.S), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:s: ", p), err)
	}
	return err
}

func (p *HydraThriftServiceAddShipArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("HydraThriftServiceAddShipArgs(%+v)", *p)
}

type HydraThriftServiceAddShipResult struct {
}

func NewHydraThriftServiceAddShipResult() *HydraThriftServiceAddShipResult {
	return &HydraThriftServiceAddShipResult{}
}

func (p *HydraThriftServiceAddShipResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *HydraThriftServiceAddShipResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("addShip_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *HydraThriftServiceAddShipResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("HydraThriftServiceAddShipResult(%+v)", *p)
}