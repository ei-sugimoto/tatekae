// @generated by protoc-gen-connect-query v0.4.1 with parameter "target=ts"
// @generated from file bill.proto (package protoBill, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { createQueryService } from "@bufbuild/connect-query";
import { MethodKind } from "@bufbuild/protobuf";
import { CreateBillRequest, CreateBillResponse, ListByProjectRequest, ListByProjectResponse, SumrizeRequest, SumrizeResponse } from "./bill_pb.js";

export const typeName = "protoBill.BillService";

/**
 * @generated from rpc protoBill.BillService.Create
 */
export const create = createQueryService({
  service: {
    methods: {
      create: {
        name: "Create",
        kind: MethodKind.Unary,
        I: CreateBillRequest,
        O: CreateBillResponse,
      },
    },
    typeName: "protoBill.BillService",
  },
}).create;

/**
 * @generated from rpc protoBill.BillService.ListByProject
 */
export const listByProject = createQueryService({
  service: {
    methods: {
      listByProject: {
        name: "ListByProject",
        kind: MethodKind.Unary,
        I: ListByProjectRequest,
        O: ListByProjectResponse,
      },
    },
    typeName: "protoBill.BillService",
  },
}).listByProject;

/**
 * @generated from rpc protoBill.BillService.SumrizeByProject
 */
export const sumrizeByProject = createQueryService({
  service: {
    methods: {
      sumrizeByProject: {
        name: "SumrizeByProject",
        kind: MethodKind.Unary,
        I: SumrizeRequest,
        O: SumrizeResponse,
      },
    },
    typeName: "protoBill.BillService",
  },
}).sumrizeByProject;
