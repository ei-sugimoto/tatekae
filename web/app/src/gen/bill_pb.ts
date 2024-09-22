// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file bill.proto (package protoBill, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message protoBill.Bill
 */
export class Bill extends Message<Bill> {
  /**
   * @generated from field: int32 price = 1;
   */
  price = 0;

  /**
   * @generated from field: int32 projectId = 2;
   */
  projectId = 0;

  /**
   * @generated from field: int32 srcUserId = 3;
   */
  srcUserId = 0;

  /**
   * @generated from field: int32 dstUserId = 4;
   */
  dstUserId = 0;

  /**
   * @generated from field: int32 ID = 5;
   */
  ID = 0;

  constructor(data?: PartialMessage<Bill>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "protoBill.Bill";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "price", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "projectId", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "srcUserId", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "dstUserId", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "ID", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Bill {
    return new Bill().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Bill {
    return new Bill().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Bill {
    return new Bill().fromJsonString(jsonString, options);
  }

  static equals(a: Bill | PlainMessage<Bill> | undefined, b: Bill | PlainMessage<Bill> | undefined): boolean {
    return proto3.util.equals(Bill, a, b);
  }
}

/**
 * @generated from message protoBill.SumrizeBill
 */
export class SumrizeBill extends Message<SumrizeBill> {
  /**
   * @generated from field: string srcUserName = 1;
   */
  srcUserName = "";

  /**
   * @generated from field: string dstUserName = 2;
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
  static readonly typeName = "protoBill.SumrizeBill";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "srcUserName", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "dstUserName", kind: "scalar", T: 9 /* ScalarType.STRING */ },
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

/**
 * @generated from message protoBill.CreateBillRequest
 */
export class CreateBillRequest extends Message<CreateBillRequest> {
  /**
   * @generated from field: int32 price = 1;
   */
  price = 0;

  /**
   * @generated from field: int32 projectId = 2;
   */
  projectId = 0;

  /**
   * @generated from field: int32 srcUserId = 3;
   */
  srcUserId = 0;

  /**
   * @generated from field: int32 dstUserId = 4;
   */
  dstUserId = 0;

  constructor(data?: PartialMessage<CreateBillRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "protoBill.CreateBillRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "price", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "projectId", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "srcUserId", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "dstUserId", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateBillRequest {
    return new CreateBillRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateBillRequest {
    return new CreateBillRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateBillRequest {
    return new CreateBillRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateBillRequest | PlainMessage<CreateBillRequest> | undefined, b: CreateBillRequest | PlainMessage<CreateBillRequest> | undefined): boolean {
    return proto3.util.equals(CreateBillRequest, a, b);
  }
}

/**
 * @generated from message protoBill.CreateBillResponse
 */
export class CreateBillResponse extends Message<CreateBillResponse> {
  /**
   * @generated from field: int32 price = 1;
   */
  price = 0;

  /**
   * @generated from field: int32 projectId = 2;
   */
  projectId = 0;

  /**
   * @generated from field: int32 srcUserId = 3;
   */
  srcUserId = 0;

  /**
   * @generated from field: int32 dstUserId = 4;
   */
  dstUserId = 0;

  /**
   * @generated from field: int32 ID = 5;
   */
  ID = 0;

  constructor(data?: PartialMessage<CreateBillResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "protoBill.CreateBillResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "price", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "projectId", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "srcUserId", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "dstUserId", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "ID", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateBillResponse {
    return new CreateBillResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateBillResponse {
    return new CreateBillResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateBillResponse {
    return new CreateBillResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateBillResponse | PlainMessage<CreateBillResponse> | undefined, b: CreateBillResponse | PlainMessage<CreateBillResponse> | undefined): boolean {
    return proto3.util.equals(CreateBillResponse, a, b);
  }
}

/**
 * @generated from message protoBill.ListByProjectRequest
 */
export class ListByProjectRequest extends Message<ListByProjectRequest> {
  /**
   * @generated from field: int32 projectId = 1;
   */
  projectId = 0;

  constructor(data?: PartialMessage<ListByProjectRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "protoBill.ListByProjectRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "projectId", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListByProjectRequest {
    return new ListByProjectRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListByProjectRequest {
    return new ListByProjectRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListByProjectRequest {
    return new ListByProjectRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListByProjectRequest | PlainMessage<ListByProjectRequest> | undefined, b: ListByProjectRequest | PlainMessage<ListByProjectRequest> | undefined): boolean {
    return proto3.util.equals(ListByProjectRequest, a, b);
  }
}

/**
 * @generated from message protoBill.ListByProjectResponse
 */
export class ListByProjectResponse extends Message<ListByProjectResponse> {
  /**
   * @generated from field: repeated protoBill.Bill bills = 1;
   */
  bills: Bill[] = [];

  constructor(data?: PartialMessage<ListByProjectResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "protoBill.ListByProjectResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "bills", kind: "message", T: Bill, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListByProjectResponse {
    return new ListByProjectResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListByProjectResponse {
    return new ListByProjectResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListByProjectResponse {
    return new ListByProjectResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListByProjectResponse | PlainMessage<ListByProjectResponse> | undefined, b: ListByProjectResponse | PlainMessage<ListByProjectResponse> | undefined): boolean {
    return proto3.util.equals(ListByProjectResponse, a, b);
  }
}

/**
 * @generated from message protoBill.SumrizeRequest
 */
export class SumrizeRequest extends Message<SumrizeRequest> {
  /**
   * @generated from field: int32 projectId = 1;
   */
  projectId = 0;

  constructor(data?: PartialMessage<SumrizeRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "protoBill.SumrizeRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "projectId", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SumrizeRequest {
    return new SumrizeRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SumrizeRequest {
    return new SumrizeRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SumrizeRequest {
    return new SumrizeRequest().fromJsonString(jsonString, options);
  }

  static equals(a: SumrizeRequest | PlainMessage<SumrizeRequest> | undefined, b: SumrizeRequest | PlainMessage<SumrizeRequest> | undefined): boolean {
    return proto3.util.equals(SumrizeRequest, a, b);
  }
}

/**
 * @generated from message protoBill.SumrizeResponse
 */
export class SumrizeResponse extends Message<SumrizeResponse> {
  /**
   * @generated from field: repeated protoBill.SumrizeBill sumrizeBills = 1;
   */
  sumrizeBills: SumrizeBill[] = [];

  constructor(data?: PartialMessage<SumrizeResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "protoBill.SumrizeResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "sumrizeBills", kind: "message", T: SumrizeBill, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SumrizeResponse {
    return new SumrizeResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SumrizeResponse {
    return new SumrizeResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SumrizeResponse {
    return new SumrizeResponse().fromJsonString(jsonString, options);
  }

  static equals(a: SumrizeResponse | PlainMessage<SumrizeResponse> | undefined, b: SumrizeResponse | PlainMessage<SumrizeResponse> | undefined): boolean {
    return proto3.util.equals(SumrizeResponse, a, b);
  }
}

