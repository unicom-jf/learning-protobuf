// @generated by protoc-gen-es v2.2.3 with parameter "target=ts"
// @generated from file addressbook.proto (package learning, syntax proto3)
/* eslint-disable */

import type { GenEnum, GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { enumDesc, fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import type { Timestamp } from "@bufbuild/protobuf/wkt";
import { file_google_protobuf_timestamp } from "@bufbuild/protobuf/wkt";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file addressbook.proto.
 */
export const file_addressbook: GenFile = /*@__PURE__*/
  fileDesc("ChFhZGRyZXNzYm9vay5wcm90bxIIbGVhcm5pbmci0wEKBlBlcnNvbhIMCgRuYW1lGAEgASgJEgoKAmlkGAIgASgFEg0KBWVtYWlsGAMgASgJEiwKBnBob25lcxgEIAMoCzIcLmxlYXJuaW5nLlBlcnNvbi5QaG9uZU51bWJlchIwCgxsYXN0X3VwZGF0ZWQYBSABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wGkAKC1Bob25lTnVtYmVyEg4KBm51bWJlchgBIAEoCRIhCgR0eXBlGAIgASgOMhMubGVhcm5pbmcuUGhvbmVUeXBlIi8KC0FkZHJlc3NCb29rEiAKBnBlb3BsZRgBIAMoCzIQLmxlYXJuaW5nLlBlcnNvbiIHCgVFbXB0eSIvCg1TZXJ2aWNlUmVzdWx0EgwKBGluZm8YASABKAkSEAoIZXJyb3JfaWQYAiABKAUqaAoJUGhvbmVUeXBlEhoKFlBIT05FX1RZUEVfVU5TUEVDSUZJRUQQABIVChFQSE9ORV9UWVBFX01PQklMRRABEhMKD1BIT05FX1RZUEVfSE9NRRACEhMKD1BIT05FX1RZUEVfV09SSxADMoMBChJBZGRyZXNzQm9va1NlcnZpY2USOgoIc2F2ZUJvb2sSFS5sZWFybmluZy5BZGRyZXNzQm9vaxoXLmxlYXJuaW5nLlNlcnZpY2VSZXN1bHQSMQoHR2V0Qm9vaxIPLmxlYXJuaW5nLkVtcHR5GhUubGVhcm5pbmcuQWRkcmVzc0Jvb2tCA1oBLmIGcHJvdG8z", [file_google_protobuf_timestamp]);

/**
 * @generated from message learning.Person
 */
export type Person = Message<"learning.Person"> & {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * Unique ID number for this person.
   *
   * @generated from field: int32 id = 2;
   */
  id: number;

  /**
   * @generated from field: string email = 3;
   */
  email: string;

  /**
   * @generated from field: repeated learning.Person.PhoneNumber phones = 4;
   */
  phones: Person_PhoneNumber[];

  /**
   * @generated from field: google.protobuf.Timestamp last_updated = 5;
   */
  lastUpdated?: Timestamp;
};

/**
 * Describes the message learning.Person.
 * Use `create(PersonSchema)` to create a new message.
 */
export const PersonSchema: GenMessage<Person> = /*@__PURE__*/
  messageDesc(file_addressbook, 0);

/**
 * @generated from message learning.Person.PhoneNumber
 */
export type Person_PhoneNumber = Message<"learning.Person.PhoneNumber"> & {
  /**
   * @generated from field: string number = 1;
   */
  number: string;

  /**
   * @generated from field: learning.PhoneType type = 2;
   */
  type: PhoneType;
};

/**
 * Describes the message learning.Person.PhoneNumber.
 * Use `create(Person_PhoneNumberSchema)` to create a new message.
 */
export const Person_PhoneNumberSchema: GenMessage<Person_PhoneNumber> = /*@__PURE__*/
  messageDesc(file_addressbook, 0, 0);

/**
 * Our address book file is just one of these.
 *
 * @generated from message learning.AddressBook
 */
export type AddressBook = Message<"learning.AddressBook"> & {
  /**
   * @generated from field: repeated learning.Person people = 1;
   */
  people: Person[];
};

/**
 * Describes the message learning.AddressBook.
 * Use `create(AddressBookSchema)` to create a new message.
 */
export const AddressBookSchema: GenMessage<AddressBook> = /*@__PURE__*/
  messageDesc(file_addressbook, 1);

/**
 * @generated from message learning.Empty
 */
export type Empty = Message<"learning.Empty"> & {
};

/**
 * Describes the message learning.Empty.
 * Use `create(EmptySchema)` to create a new message.
 */
export const EmptySchema: GenMessage<Empty> = /*@__PURE__*/
  messageDesc(file_addressbook, 2);

/**
 * @generated from message learning.ServiceResult
 */
export type ServiceResult = Message<"learning.ServiceResult"> & {
  /**
   * @generated from field: string info = 1;
   */
  info: string;

  /**
   * 0 for no error
   *
   * @generated from field: int32 error_id = 2;
   */
  errorId: number;
};

/**
 * Describes the message learning.ServiceResult.
 * Use `create(ServiceResultSchema)` to create a new message.
 */
export const ServiceResultSchema: GenMessage<ServiceResult> = /*@__PURE__*/
  messageDesc(file_addressbook, 3);

/**
 * @generated from enum learning.PhoneType
 */
export enum PhoneType {
  /**
   * @generated from enum value: PHONE_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: PHONE_TYPE_MOBILE = 1;
   */
  MOBILE = 1,

  /**
   * @generated from enum value: PHONE_TYPE_HOME = 2;
   */
  HOME = 2,

  /**
   * @generated from enum value: PHONE_TYPE_WORK = 3;
   */
  WORK = 3,
}

/**
 * Describes the enum learning.PhoneType.
 */
export const PhoneTypeSchema: GenEnum<PhoneType> = /*@__PURE__*/
  enumDesc(file_addressbook, 0);

/**
 * @generated from service learning.AddressBookService
 */
export const AddressBookService: GenService<{
  /**
   * @generated from rpc learning.AddressBookService.saveBook
   */
  saveBook: {
    methodKind: "unary";
    input: typeof AddressBookSchema;
    output: typeof ServiceResultSchema;
  },
  /**
   * @generated from rpc learning.AddressBookService.GetBook
   */
  getBook: {
    methodKind: "unary";
    input: typeof EmptySchema;
    output: typeof AddressBookSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_addressbook, 0);

