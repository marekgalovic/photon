package utils

import (
    "fmt";
    "reflect";

    pb "github.com/marekgalovic/serving/server/protos"
)

func InterfaceMapToValueInterfacePb(features map[string]interface{}) (map[string]*pb.ValueInterface, error) {
    result := make(map[string]*pb.ValueInterface, len(features))

    for key, value := range features {
        if value == nil {
            result[key] = nil
            continue
        }

        switch value.(type) {
        case bool:
            result[key] = &pb.ValueInterface{Value: &pb.ValueInterface_ValueBoolean{ValueBoolean: value.(bool)}}
        case int:
            result[key] = &pb.ValueInterface{Value: &pb.ValueInterface_ValueInt64{ValueInt64: int64(value.(int))}}
        case int32:
            result[key] = &pb.ValueInterface{Value: &pb.ValueInterface_ValueInt32{ValueInt32: value.(int32)}}
        case int64:
            result[key] = &pb.ValueInterface{Value: &pb.ValueInterface_ValueInt64{ValueInt64: value.(int64)}}
        case float32:
            result[key] = &pb.ValueInterface{Value: &pb.ValueInterface_ValueFloat32{ValueFloat32: value.(float32)}}
        case float64:
            result[key] = &pb.ValueInterface{Value: &pb.ValueInterface_ValueFloat64{ValueFloat64: value.(float64)}}
        case string:
            result[key] = &pb.ValueInterface{Value: &pb.ValueInterface_ValueString{ValueString: value.(string)}}
        case []byte:
            result[key] = &pb.ValueInterface{Value: &pb.ValueInterface_ValueBytes{ValueBytes: value.([]byte)}}
        default:
            return nil, fmt.Errorf("Unsupported type %s for key %s", reflect.TypeOf(value).String(), key)
        } 
    }

    return result, nil
}

func ValueInterfacePbToInterfaceMap(features map[string]*pb.ValueInterface) (map[string]interface{}, error) {
    result := make(map[string]interface{}, len(features))

    for key, value := range features {
        if value == nil {
            result[key] = nil
            continue
        }

        switch value.Value.(type) {
        case *pb.ValueInterface_ValueBoolean:
            result[key] = value.GetValueBoolean()
        case *pb.ValueInterface_ValueInt32:
            result[key] = value.GetValueInt32()
        case *pb.ValueInterface_ValueInt64:
            result[key] = value.GetValueInt64()
        case *pb.ValueInterface_ValueFloat32:
            result[key] = value.GetValueFloat32()
        case *pb.ValueInterface_ValueFloat64:
            result[key] = value.GetValueFloat64()
        case *pb.ValueInterface_ValueString:
            result[key] = value.GetValueString()
        case *pb.ValueInterface_ValueBytes:
            result[key] = value.GetValueBytes()
        default:
            return nil, fmt.Errorf("Unsupoprted type %s for key %s", reflect.TypeOf(value.GetValue()).String(), key)
        }
        result[key] = value.GetValue()
    }

    return result, nil
}
