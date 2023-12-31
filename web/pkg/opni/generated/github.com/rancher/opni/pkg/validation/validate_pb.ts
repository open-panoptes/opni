// @generated by protoc-gen-es v1.3.0 with parameter "target=ts,import_extension=none,ts_nocheck=false"
// @generated from file github.com/rancher/opni/pkg/validation/validate.proto (package validate, syntax proto3)
/* eslint-disable */

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from message validate.PEMBlock
 */
export class PEMBlock extends Message<PEMBlock> {
  /**
   * @generated from field: string type = 1;
   */
  type = "";

  /**
   * @generated from field: map<string, string> headers = 2;
   */
  headers: { [key: string]: string } = {};

  /**
   * @generated from field: bytes bytes = 3;
   */
  bytes = new Uint8Array(0);

  constructor(data?: PartialMessage<PEMBlock>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "validate.PEMBlock";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "headers", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
    { no: 3, name: "bytes", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PEMBlock {
    return new PEMBlock().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PEMBlock {
    return new PEMBlock().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PEMBlock {
    return new PEMBlock().fromJsonString(jsonString, options);
  }

  static equals(a: PEMBlock | PlainMessage<PEMBlock> | undefined, b: PEMBlock | PlainMessage<PEMBlock> | undefined): boolean {
    return proto3.util.equals(PEMBlock, a, b);
  }
}

/**
 * @generated from message validate.X509Certificate
 */
export class X509Certificate extends Message<X509Certificate> {
  /**
   * @generated from field: bytes raw = 1;
   */
  raw = new Uint8Array(0);

  /**
   * @generated from field: bool isCA = 2;
   */
  isCA = false;

  /**
   * @generated from field: string issuer = 3;
   */
  issuer = "";

  /**
   * @generated from field: string subject = 4;
   */
  subject = "";

  /**
   * @generated from field: google.protobuf.Timestamp notBefore = 5;
   */
  notBefore?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp notAfter = 6;
   */
  notAfter?: Timestamp;

  /**
   * @generated from field: string alg = 7;
   */
  alg = "";

  constructor(data?: PartialMessage<X509Certificate>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "validate.X509Certificate";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "raw", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 2, name: "isCA", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "issuer", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "subject", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "notBefore", kind: "message", T: Timestamp },
    { no: 6, name: "notAfter", kind: "message", T: Timestamp },
    { no: 7, name: "alg", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): X509Certificate {
    return new X509Certificate().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): X509Certificate {
    return new X509Certificate().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): X509Certificate {
    return new X509Certificate().fromJsonString(jsonString, options);
  }

  static equals(a: X509Certificate | PlainMessage<X509Certificate> | undefined, b: X509Certificate | PlainMessage<X509Certificate> | undefined): boolean {
    return proto3.util.equals(X509Certificate, a, b);
  }
}

