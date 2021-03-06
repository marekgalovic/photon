# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: core.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_message "photon.ValueInterface" do
    oneof :value do
      optional :value_boolean, :bool, 1
      optional :value_int32, :int32, 2
      optional :list_int32, :message, 3, "photon.ListInt32"
      optional :value_int64, :int64, 4
      optional :list_int64, :message, 5, "photon.ListInt64"
      optional :value_float32, :float, 6
      optional :list_float32, :message, 7, "photon.ListFloat32"
      optional :value_float64, :double, 8
      optional :list_float64, :message, 9, "photon.ListFloat64"
      optional :value_string, :string, 10
      optional :value_bytes, :bytes, 11
    end
  end
  add_message "photon.ListInt32" do
    repeated :value, :int32, 1
  end
  add_message "photon.ListInt64" do
    repeated :value, :int64, 1
  end
  add_message "photon.ListFloat32" do
    repeated :value, :float, 1
  end
  add_message "photon.ListFloat64" do
    repeated :value, :double, 1
  end
  add_message "photon.EmptyRequest" do
  end
  add_message "photon.EmptyResponse" do
  end
  add_message "photon.DataBlob" do
    oneof :part do
      optional :name, :string, 1
      optional :data, :bytes, 2
    end
  end
end

module Photon
  ValueInterface = Google::Protobuf::DescriptorPool.generated_pool.lookup("photon.ValueInterface").msgclass
  ListInt32 = Google::Protobuf::DescriptorPool.generated_pool.lookup("photon.ListInt32").msgclass
  ListInt64 = Google::Protobuf::DescriptorPool.generated_pool.lookup("photon.ListInt64").msgclass
  ListFloat32 = Google::Protobuf::DescriptorPool.generated_pool.lookup("photon.ListFloat32").msgclass
  ListFloat64 = Google::Protobuf::DescriptorPool.generated_pool.lookup("photon.ListFloat64").msgclass
  EmptyRequest = Google::Protobuf::DescriptorPool.generated_pool.lookup("photon.EmptyRequest").msgclass
  EmptyResponse = Google::Protobuf::DescriptorPool.generated_pool.lookup("photon.EmptyResponse").msgclass
  DataBlob = Google::Protobuf::DescriptorPool.generated_pool.lookup("photon.DataBlob").msgclass
end
