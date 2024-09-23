// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file proto_bill/v1/bill.proto (package proto_bill.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message proto_bill.v1.BillServiceCreateRequest
 */
export class BillServiceCreateRequest extends Message<BillServiceCreateRequest> {
  /**
   * @generated from field: int32 project_id = 1;
   */
  projectId = 0;

  /**
   * @generated from field: int32 src_user_id = 2;
   */
  srcUserId = 0;

  /**
   * @generated from field: int32 dst_user_id = 3;
   */
  dstUserId = 0;

  /**
   * @generated from field: int32 price = 4;
   */
  price = 0;

  constructor(data?: PartialMessage<BillServiceCreateRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto_bill.v1.BillServiceCreateRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "project_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "src_user_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "dst_user_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "price", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): BillServiceCreateRequest {
    return new BillServiceCreateRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): BillServiceCreateRequest {
    return new BillServiceCreateRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): BillServiceCreateRequest {
    return new BillServiceCreateRequest().fromJsonString(jsonString, options);
  }

  static equals(a: BillServiceCreateRequest | PlainMessage<BillServiceCreateRequest> | undefined, b: BillServiceCreateRequest | PlainMessage<BillServiceCreateRequest> | undefined): boolean {
    return proto3.util.equals(BillServiceCreateRequest, a, b);
  }
}

/**
 * @generated from message proto_bill.v1.BillServiceCreateResponse
 */
export class BillServiceCreateResponse extends Message<BillServiceCreateResponse> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  /**
   * @generated from field: int32 price = 2;
   */
  price = 0;

  /**
   * @generated from field: int32 src_user_id = 3;
   */
  srcUserId = 0;

  /**
   * @generated from field: int32 dst_user_id = 4;
   */
  dstUserId = 0;

  /**
   * @generated from field: int32 project_id = 5;
   */
  projectId = 0;

  constructor(data?: PartialMessage<BillServiceCreateResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto_bill.v1.BillServiceCreateResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "price", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "src_user_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "dst_user_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "project_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): BillServiceCreateResponse {
    return new BillServiceCreateResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): BillServiceCreateResponse {
    return new BillServiceCreateResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): BillServiceCreateResponse {
    return new BillServiceCreateResponse().fromJsonString(jsonString, options);
  }

  static equals(a: BillServiceCreateResponse | PlainMessage<BillServiceCreateResponse> | undefined, b: BillServiceCreateResponse | PlainMessage<BillServiceCreateResponse> | undefined): boolean {
    return proto3.util.equals(BillServiceCreateResponse, a, b);
  }
}

/**
 * @generated from message proto_bill.v1.BillServiceSumrizeByProjectRequest
 */
export class BillServiceSumrizeByProjectRequest extends Message<BillServiceSumrizeByProjectRequest> {
  /**
   * @generated from field: int32 project_id = 1;
   */
  projectId = 0;

  constructor(data?: PartialMessage<BillServiceSumrizeByProjectRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto_bill.v1.BillServiceSumrizeByProjectRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "project_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): BillServiceSumrizeByProjectRequest {
    return new BillServiceSumrizeByProjectRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): BillServiceSumrizeByProjectRequest {
    return new BillServiceSumrizeByProjectRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): BillServiceSumrizeByProjectRequest {
    return new BillServiceSumrizeByProjectRequest().fromJsonString(jsonString, options);
  }

  static equals(a: BillServiceSumrizeByProjectRequest | PlainMessage<BillServiceSumrizeByProjectRequest> | undefined, b: BillServiceSumrizeByProjectRequest | PlainMessage<BillServiceSumrizeByProjectRequest> | undefined): boolean {
    return proto3.util.equals(BillServiceSumrizeByProjectRequest, a, b);
  }
}

/**
 * @generated from message proto_bill.v1.BillServiceSumrizeByProjectResponse
 */
export class BillServiceSumrizeByProjectResponse extends Message<BillServiceSumrizeByProjectResponse> {
  /**
   * @generated from field: repeated proto_bill.v1.SumrizeBill bills = 1;
   */
  bills: SumrizeBill[] = [];

  constructor(data?: PartialMessage<BillServiceSumrizeByProjectResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto_bill.v1.BillServiceSumrizeByProjectResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "bills", kind: "message", T: SumrizeBill, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): BillServiceSumrizeByProjectResponse {
    return new BillServiceSumrizeByProjectResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): BillServiceSumrizeByProjectResponse {
    return new BillServiceSumrizeByProjectResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): BillServiceSumrizeByProjectResponse {
    return new BillServiceSumrizeByProjectResponse().fromJsonString(jsonString, options);
  }

  static equals(a: BillServiceSumrizeByProjectResponse | PlainMessage<BillServiceSumrizeByProjectResponse> | undefined, b: BillServiceSumrizeByProjectResponse | PlainMessage<BillServiceSumrizeByProjectResponse> | undefined): boolean {
    return proto3.util.equals(BillServiceSumrizeByProjectResponse, a, b);
  }
}

/**
 * @generated from message proto_bill.v1.SumrizeBill
 */
export class SumrizeBill extends Message<SumrizeBill> {
  /**
   * @generated from field: string src_user_name = 1;
   */
  srcUserName = "";

  /**
   * @generated from field: string dst_user_name = 2;
   */
  dstUserName = "";

  /**
   * @generated from field: int32 price = 3;
   */
  price = 0;

  constructor(data?: PartialMessage<SumrizeBill>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "proto_bill.v1.SumrizeBill";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "src_user_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "dst_user_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "price", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SumrizeBill {
    return new SumrizeBill().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SumrizeBill {
    return new SumrizeBill().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SumrizeBill {
    return new SumrizeBill().fromJsonString(jsonString, options);
  }

  static equals(a: SumrizeBill | PlainMessage<SumrizeBill> | undefined, b: SumrizeBill | PlainMessage<SumrizeBill> | undefined): boolean {
    return proto3.util.equals(SumrizeBill, a, b);
  }
}

