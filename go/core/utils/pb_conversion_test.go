package utils

import (
    "testing";

    pb "github.com/marekgalovic/photon/go/core/protos";

    "github.com/stretchr/testify/suite"
)

type PbConversionTest struct {
    suite.Suite
    interfaceMap map[string]interface{}
    valueInterfacePb map[string]*pb.ValueInterface
}

func TestPbConversion(t *testing.T) {
    suite.Run(t, new(PbConversionTest))
}

func (test *PbConversionTest) SetupTest() {
    test.interfaceMap = map[string]interface{}{
        "bool": true,
        "int": int(1),
        "int32": int32(32),
        "int32_list": []int32{32},
        "int64": int64(64),
        "int64_list": []int64{64},
        "float32": float32(32.32),
        "float32_list": []float32{32.32},
        "float64": float64(64.64),
        "float64_list": []float64{64.64},
        "string": "Foo",
        "bytes": []byte("Bar"),
    }

    test.valueInterfacePb = map[string]*pb.ValueInterface{
        "bool": &pb.ValueInterface{Value: &pb.ValueInterface_ValueBoolean{ValueBoolean: true}},
        "int": &pb.ValueInterface{Value: &pb.ValueInterface_ValueInt64{ValueInt64: int64(1)}},
        "int32": &pb.ValueInterface{Value: &pb.ValueInterface_ValueInt32{ValueInt32: int32(32)}},
        "int32_list": &pb.ValueInterface{Value: &pb.ValueInterface_ListInt32{ListInt32: &pb.ListInt32{Value: []int32{32}}}},
        "int64": &pb.ValueInterface{Value: &pb.ValueInterface_ValueInt64{ValueInt64: int64(64)}},
        "int64_list": &pb.ValueInterface{Value: &pb.ValueInterface_ListInt64{ListInt64: &pb.ListInt64{Value: []int64{64}}}},
        "float32": &pb.ValueInterface{Value: &pb.ValueInterface_ValueFloat32{ValueFloat32: float32(32.32)}},
        "float32_list": &pb.ValueInterface{Value: &pb.ValueInterface_ListFloat32{ListFloat32: &pb.ListFloat32{Value: []float32{32.32}}}},
        "float64": &pb.ValueInterface{Value: &pb.ValueInterface_ValueFloat64{ValueFloat64: float64(64.64)}},
        "float64_list": &pb.ValueInterface{Value: &pb.ValueInterface_ListFloat64{ListFloat64: &pb.ListFloat64{Value: []float64{64.64}}}},
        "string": &pb.ValueInterface{Value: &pb.ValueInterface_ValueString{ValueString: "Foo"}},
        "bytes": &pb.ValueInterface{Value: &pb.ValueInterface_ValueBytes{ValueBytes: []byte("Bar")}},    
    }
}

func (test *PbConversionTest) TestInterfaceMapToValueInterfacePb() {
    valueInterfacePb, err := InterfaceMapToValueInterfacePb(test.interfaceMap)

    test.Nil(err)
    test.Equal(test.valueInterfacePb, valueInterfacePb)
}

func (test *PbConversionTest) TestInterfaceMapToValueInterfacePbReturnsInvalidValueError() {
    _, err := InterfaceMapToValueInterfacePb(map[string]interface{}{"key": []string{"Foo"}})

    test.NotNil(err)
}

func (test *PbConversionTest) TestValueInterfacePbToInterfaceMap() {
    test.interfaceMap["int"] = int64(1)
    interfaceMap, err := ValueInterfacePbToInterfaceMap(test.valueInterfacePb)

    test.Nil(err)
    test.Equal(test.interfaceMap, interfaceMap)
}
