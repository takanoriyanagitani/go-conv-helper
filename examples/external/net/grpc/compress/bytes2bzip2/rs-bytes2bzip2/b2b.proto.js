/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
(function (global, factory) {
  /* global define, require, module */

  /* AMD */ if (typeof define === "function" && define.amd) {
    define(["protobufjs/minimal"], factory);
  } /* CommonJS */ else if (
    typeof require === "function" && typeof module === "object" && module &&
    module.exports
  ) {
    module.exports = factory(require("protobufjs/minimal"));
  }
})(this, function ($protobuf) {
  "use strict";

  // Common aliases
  var $Reader = $protobuf.Reader,
    $Writer = $protobuf.Writer,
    $util = $protobuf.util;

  // Exported root namespace
  var $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

  $root.bytes2bz = (function () {
    /**
     * Namespace bytes2bz.
     * @exports bytes2bz
     * @namespace
     */
    var bytes2bz = {};

    bytes2bz.v1 = (function () {
      /**
       * Namespace v1.
       * @memberof bytes2bz
       * @namespace
       */
      var v1 = {};

      v1.Compression = (function () {
        /**
         * Properties of a Compression.
         * @memberof bytes2bz.v1
         * @interface ICompression
         */

        /**
         * Constructs a new Compression.
         * @memberof bytes2bz.v1
         * @classdesc Represents a Compression.
         * @implements ICompression
         * @constructor
         * @param {bytes2bz.v1.ICompression=} [properties] Properties to set
         */
        function Compression(properties) {
          if (properties) {
            for (
              var keys = Object.keys(properties), i = 0;
              i < keys.length;
              ++i
            ) {
              if (properties[keys[i]] != null) {
                this[keys[i]] = properties[keys[i]];
              }
            }
          }
        }

        /**
         * Creates a new Compression instance using the specified properties.
         * @function create
         * @memberof bytes2bz.v1.Compression
         * @static
         * @param {bytes2bz.v1.ICompression=} [properties] Properties to set
         * @returns {bytes2bz.v1.Compression} Compression instance
         */
        Compression.create = function create(properties) {
          return new Compression(properties);
        };

        /**
         * Encodes the specified Compression message. Does not implicitly {@link bytes2bz.v1.Compression.verify|verify} messages.
         * @function encode
         * @memberof bytes2bz.v1.Compression
         * @static
         * @param {bytes2bz.v1.ICompression} message Compression message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Compression.encode = function encode(message, writer) {
          if (!writer) {
            writer = $Writer.create();
          }
          return writer;
        };

        /**
         * Encodes the specified Compression message, length delimited. Does not implicitly {@link bytes2bz.v1.Compression.verify|verify} messages.
         * @function encodeDelimited
         * @memberof bytes2bz.v1.Compression
         * @static
         * @param {bytes2bz.v1.ICompression} message Compression message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Compression.encodeDelimited = function encodeDelimited(
          message,
          writer,
        ) {
          return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Compression message from the specified reader or buffer.
         * @function decode
         * @memberof bytes2bz.v1.Compression
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {bytes2bz.v1.Compression} Compression
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Compression.decode = function decode(reader, length) {
          if (!(reader instanceof $Reader)) {
            reader = $Reader.create(reader);
          }
          var end = length === undefined ? reader.len : reader.pos + length,
            message = new $root.bytes2bz.v1.Compression();
          while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
              default:
                reader.skipType(tag & 7);
                break;
            }
          }
          return message;
        };

        /**
         * Decodes a Compression message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof bytes2bz.v1.Compression
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {bytes2bz.v1.Compression} Compression
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Compression.decodeDelimited = function decodeDelimited(reader) {
          if (!(reader instanceof $Reader)) {
            reader = new $Reader(reader);
          }
          return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Compression message.
         * @function verify
         * @memberof bytes2bz.v1.Compression
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Compression.verify = function verify(message) {
          if (typeof message !== "object" || message === null) {
            return "object expected";
          }
          return null;
        };

        /**
         * Creates a Compression message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof bytes2bz.v1.Compression
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {bytes2bz.v1.Compression} Compression
         */
        Compression.fromObject = function fromObject(object) {
          if (object instanceof $root.bytes2bz.v1.Compression) {
            return object;
          }
          return new $root.bytes2bz.v1.Compression();
        };

        /**
         * Creates a plain object from a Compression message. Also converts values to other types if specified.
         * @function toObject
         * @memberof bytes2bz.v1.Compression
         * @static
         * @param {bytes2bz.v1.Compression} message Compression
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Compression.toObject = function toObject() {
          return {};
        };

        /**
         * Converts this Compression to JSON.
         * @function toJSON
         * @memberof bytes2bz.v1.Compression
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Compression.prototype.toJSON = function toJSON() {
          return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for Compression
         * @function getTypeUrl
         * @memberof bytes2bz.v1.Compression
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        Compression.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
          if (typeUrlPrefix === undefined) {
            typeUrlPrefix = "type.googleapis.com";
          }
          return typeUrlPrefix + "/bytes2bz.v1.Compression";
        };

        /**
         * Level enum.
         * @name bytes2bz.v1.Compression.Level
         * @enum {number}
         * @property {number} LEVEL_NUM_UNSPECIFIED=0 LEVEL_NUM_UNSPECIFIED value
         * @property {number} LEVEL_NUM_FAST=1 LEVEL_NUM_FAST value
         * @property {number} LEVEL_NUM_BEST=2 LEVEL_NUM_BEST value
         * @property {number} LEVEL_NUM_0=3 LEVEL_NUM_0 value
         * @property {number} LEVEL_NUM_1=4 LEVEL_NUM_1 value
         * @property {number} LEVEL_NUM_2=5 LEVEL_NUM_2 value
         * @property {number} LEVEL_NUM_3=6 LEVEL_NUM_3 value
         * @property {number} LEVEL_NUM_4=7 LEVEL_NUM_4 value
         * @property {number} LEVEL_NUM_5=8 LEVEL_NUM_5 value
         * @property {number} LEVEL_NUM_6=9 LEVEL_NUM_6 value
         * @property {number} LEVEL_NUM_7=10 LEVEL_NUM_7 value
         * @property {number} LEVEL_NUM_8=11 LEVEL_NUM_8 value
         * @property {number} LEVEL_NUM_9=12 LEVEL_NUM_9 value
         */
        Compression.Level = (function () {
          var valuesById = {}, values = Object.create(valuesById);
          values[valuesById[0] = "LEVEL_NUM_UNSPECIFIED"] = 0;
          values[valuesById[1] = "LEVEL_NUM_FAST"] = 1;
          values[valuesById[2] = "LEVEL_NUM_BEST"] = 2;
          values[valuesById[3] = "LEVEL_NUM_0"] = 3;
          values[valuesById[4] = "LEVEL_NUM_1"] = 4;
          values[valuesById[5] = "LEVEL_NUM_2"] = 5;
          values[valuesById[6] = "LEVEL_NUM_3"] = 6;
          values[valuesById[7] = "LEVEL_NUM_4"] = 7;
          values[valuesById[8] = "LEVEL_NUM_5"] = 8;
          values[valuesById[9] = "LEVEL_NUM_6"] = 9;
          values[valuesById[10] = "LEVEL_NUM_7"] = 10;
          values[valuesById[11] = "LEVEL_NUM_8"] = 11;
          values[valuesById[12] = "LEVEL_NUM_9"] = 12;
          return values;
        })();

        return Compression;
      })();

      v1.CompressRequest = (function () {
        /**
         * Properties of a CompressRequest.
         * @memberof bytes2bz.v1
         * @interface ICompressRequest
         * @property {bytes2bz.v1.Compression.Level|null} [level] CompressRequest level
         * @property {Uint8Array|null} [input] CompressRequest input
         */

        /**
         * Constructs a new CompressRequest.
         * @memberof bytes2bz.v1
         * @classdesc Represents a CompressRequest.
         * @implements ICompressRequest
         * @constructor
         * @param {bytes2bz.v1.ICompressRequest=} [properties] Properties to set
         */
        function CompressRequest(properties) {
          if (properties) {
            for (
              var keys = Object.keys(properties), i = 0;
              i < keys.length;
              ++i
            ) {
              if (properties[keys[i]] != null) {
                this[keys[i]] = properties[keys[i]];
              }
            }
          }
        }

        /**
         * CompressRequest level.
         * @member {bytes2bz.v1.Compression.Level} level
         * @memberof bytes2bz.v1.CompressRequest
         * @instance
         */
        CompressRequest.prototype.level = 0;

        /**
         * CompressRequest input.
         * @member {Uint8Array} input
         * @memberof bytes2bz.v1.CompressRequest
         * @instance
         */
        CompressRequest.prototype.input = $util.newBuffer([]);

        /**
         * Creates a new CompressRequest instance using the specified properties.
         * @function create
         * @memberof bytes2bz.v1.CompressRequest
         * @static
         * @param {bytes2bz.v1.ICompressRequest=} [properties] Properties to set
         * @returns {bytes2bz.v1.CompressRequest} CompressRequest instance
         */
        CompressRequest.create = function create(properties) {
          return new CompressRequest(properties);
        };

        /**
         * Encodes the specified CompressRequest message. Does not implicitly {@link bytes2bz.v1.CompressRequest.verify|verify} messages.
         * @function encode
         * @memberof bytes2bz.v1.CompressRequest
         * @static
         * @param {bytes2bz.v1.ICompressRequest} message CompressRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        CompressRequest.encode = function encode(message, writer) {
          if (!writer) {
            writer = $Writer.create();
          }
          if (
            message.level != null &&
            Object.hasOwnProperty.call(message, "level")
          ) {
            writer.uint32(/* id 1, wireType 0 =*/ 8).int32(message.level);
          }
          if (
            message.input != null &&
            Object.hasOwnProperty.call(message, "input")
          ) {
            writer.uint32(/* id 2, wireType 2 =*/ 18).bytes(message.input);
          }
          return writer;
        };

        /**
         * Encodes the specified CompressRequest message, length delimited. Does not implicitly {@link bytes2bz.v1.CompressRequest.verify|verify} messages.
         * @function encodeDelimited
         * @memberof bytes2bz.v1.CompressRequest
         * @static
         * @param {bytes2bz.v1.ICompressRequest} message CompressRequest message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        CompressRequest.encodeDelimited = function encodeDelimited(
          message,
          writer,
        ) {
          return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a CompressRequest message from the specified reader or buffer.
         * @function decode
         * @memberof bytes2bz.v1.CompressRequest
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {bytes2bz.v1.CompressRequest} CompressRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        CompressRequest.decode = function decode(reader, length) {
          if (!(reader instanceof $Reader)) {
            reader = $Reader.create(reader);
          }
          var end = length === undefined ? reader.len : reader.pos + length,
            message = new $root.bytes2bz.v1.CompressRequest();
          while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
              case 1: {
                message.level = reader.int32();
                break;
              }
              case 2: {
                message.input = reader.bytes();
                break;
              }
              default:
                reader.skipType(tag & 7);
                break;
            }
          }
          return message;
        };

        /**
         * Decodes a CompressRequest message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof bytes2bz.v1.CompressRequest
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {bytes2bz.v1.CompressRequest} CompressRequest
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        CompressRequest.decodeDelimited = function decodeDelimited(reader) {
          if (!(reader instanceof $Reader)) {
            reader = new $Reader(reader);
          }
          return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a CompressRequest message.
         * @function verify
         * @memberof bytes2bz.v1.CompressRequest
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        CompressRequest.verify = function verify(message) {
          if (typeof message !== "object" || message === null) {
            return "object expected";
          }
          if (message.level != null && message.hasOwnProperty("level")) {
            switch (message.level) {
              default:
                return "level: enum value expected";
              case 0:
              case 1:
              case 2:
              case 3:
              case 4:
              case 5:
              case 6:
              case 7:
              case 8:
              case 9:
              case 10:
              case 11:
              case 12:
                break;
            }
          }
          if (message.input != null && message.hasOwnProperty("input")) {
            if (
              !(message.input && typeof message.input.length === "number" ||
                $util.isString(message.input))
            ) {
              return "input: buffer expected";
            }
          }
          return null;
        };

        /**
         * Creates a CompressRequest message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof bytes2bz.v1.CompressRequest
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {bytes2bz.v1.CompressRequest} CompressRequest
         */
        CompressRequest.fromObject = function fromObject(object) {
          if (object instanceof $root.bytes2bz.v1.CompressRequest) {
            return object;
          }
          var message = new $root.bytes2bz.v1.CompressRequest();
          switch (object.level) {
            default:
              if (typeof object.level === "number") {
                message.level = object.level;
                break;
              }
              break;
            case "LEVEL_NUM_UNSPECIFIED":
            case 0:
              message.level = 0;
              break;
            case "LEVEL_NUM_FAST":
            case 1:
              message.level = 1;
              break;
            case "LEVEL_NUM_BEST":
            case 2:
              message.level = 2;
              break;
            case "LEVEL_NUM_0":
            case 3:
              message.level = 3;
              break;
            case "LEVEL_NUM_1":
            case 4:
              message.level = 4;
              break;
            case "LEVEL_NUM_2":
            case 5:
              message.level = 5;
              break;
            case "LEVEL_NUM_3":
            case 6:
              message.level = 6;
              break;
            case "LEVEL_NUM_4":
            case 7:
              message.level = 7;
              break;
            case "LEVEL_NUM_5":
            case 8:
              message.level = 8;
              break;
            case "LEVEL_NUM_6":
            case 9:
              message.level = 9;
              break;
            case "LEVEL_NUM_7":
            case 10:
              message.level = 10;
              break;
            case "LEVEL_NUM_8":
            case 11:
              message.level = 11;
              break;
            case "LEVEL_NUM_9":
            case 12:
              message.level = 12;
              break;
          }
          if (object.input != null) {
            if (typeof object.input === "string") {
              $util.base64.decode(
                object.input,
                message.input = $util.newBuffer(
                  $util.base64.length(object.input),
                ),
                0,
              );
            } else if (object.input.length >= 0) {
              message.input = object.input;
            }
          }
          return message;
        };

        /**
         * Creates a plain object from a CompressRequest message. Also converts values to other types if specified.
         * @function toObject
         * @memberof bytes2bz.v1.CompressRequest
         * @static
         * @param {bytes2bz.v1.CompressRequest} message CompressRequest
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        CompressRequest.toObject = function toObject(message, options) {
          if (!options) {
            options = {};
          }
          var object = {};
          if (options.defaults) {
            object.level = options.enums === String
              ? "LEVEL_NUM_UNSPECIFIED"
              : 0;
            if (options.bytes === String) {
              object.input = "";
            } else {
              object.input = [];
              if (options.bytes !== Array) {
                object.input = $util.newBuffer(object.input);
              }
            }
          }
          if (message.level != null && message.hasOwnProperty("level")) {
            object.level = options.enums === String
              ? $root.bytes2bz.v1.Compression.Level[message.level] === undefined
                ? message.level
                : $root.bytes2bz.v1.Compression.Level[message.level]
              : message.level;
          }
          if (message.input != null && message.hasOwnProperty("input")) {
            object.input = options.bytes === String
              ? $util.base64.encode(message.input, 0, message.input.length)
              : options.bytes === Array
              ? Array.prototype.slice.call(message.input)
              : message.input;
          }
          return object;
        };

        /**
         * Converts this CompressRequest to JSON.
         * @function toJSON
         * @memberof bytes2bz.v1.CompressRequest
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        CompressRequest.prototype.toJSON = function toJSON() {
          return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for CompressRequest
         * @function getTypeUrl
         * @memberof bytes2bz.v1.CompressRequest
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        CompressRequest.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
          if (typeUrlPrefix === undefined) {
            typeUrlPrefix = "type.googleapis.com";
          }
          return typeUrlPrefix + "/bytes2bz.v1.CompressRequest";
        };

        return CompressRequest;
      })();

      v1.CompressResponse = (function () {
        /**
         * Properties of a CompressResponse.
         * @memberof bytes2bz.v1
         * @interface ICompressResponse
         * @property {Uint8Array|null} [compressed] CompressResponse compressed
         */

        /**
         * Constructs a new CompressResponse.
         * @memberof bytes2bz.v1
         * @classdesc Represents a CompressResponse.
         * @implements ICompressResponse
         * @constructor
         * @param {bytes2bz.v1.ICompressResponse=} [properties] Properties to set
         */
        function CompressResponse(properties) {
          if (properties) {
            for (
              var keys = Object.keys(properties), i = 0;
              i < keys.length;
              ++i
            ) {
              if (properties[keys[i]] != null) {
                this[keys[i]] = properties[keys[i]];
              }
            }
          }
        }

        /**
         * CompressResponse compressed.
         * @member {Uint8Array} compressed
         * @memberof bytes2bz.v1.CompressResponse
         * @instance
         */
        CompressResponse.prototype.compressed = $util.newBuffer([]);

        /**
         * Creates a new CompressResponse instance using the specified properties.
         * @function create
         * @memberof bytes2bz.v1.CompressResponse
         * @static
         * @param {bytes2bz.v1.ICompressResponse=} [properties] Properties to set
         * @returns {bytes2bz.v1.CompressResponse} CompressResponse instance
         */
        CompressResponse.create = function create(properties) {
          return new CompressResponse(properties);
        };

        /**
         * Encodes the specified CompressResponse message. Does not implicitly {@link bytes2bz.v1.CompressResponse.verify|verify} messages.
         * @function encode
         * @memberof bytes2bz.v1.CompressResponse
         * @static
         * @param {bytes2bz.v1.ICompressResponse} message CompressResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        CompressResponse.encode = function encode(message, writer) {
          if (!writer) {
            writer = $Writer.create();
          }
          if (
            message.compressed != null &&
            Object.hasOwnProperty.call(message, "compressed")
          ) {
            writer.uint32(/* id 1, wireType 2 =*/ 10).bytes(message.compressed);
          }
          return writer;
        };

        /**
         * Encodes the specified CompressResponse message, length delimited. Does not implicitly {@link bytes2bz.v1.CompressResponse.verify|verify} messages.
         * @function encodeDelimited
         * @memberof bytes2bz.v1.CompressResponse
         * @static
         * @param {bytes2bz.v1.ICompressResponse} message CompressResponse message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        CompressResponse.encodeDelimited = function encodeDelimited(
          message,
          writer,
        ) {
          return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a CompressResponse message from the specified reader or buffer.
         * @function decode
         * @memberof bytes2bz.v1.CompressResponse
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {bytes2bz.v1.CompressResponse} CompressResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        CompressResponse.decode = function decode(reader, length) {
          if (!(reader instanceof $Reader)) {
            reader = $Reader.create(reader);
          }
          var end = length === undefined ? reader.len : reader.pos + length,
            message = new $root.bytes2bz.v1.CompressResponse();
          while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
              case 1: {
                message.compressed = reader.bytes();
                break;
              }
              default:
                reader.skipType(tag & 7);
                break;
            }
          }
          return message;
        };

        /**
         * Decodes a CompressResponse message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof bytes2bz.v1.CompressResponse
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {bytes2bz.v1.CompressResponse} CompressResponse
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        CompressResponse.decodeDelimited = function decodeDelimited(reader) {
          if (!(reader instanceof $Reader)) {
            reader = new $Reader(reader);
          }
          return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a CompressResponse message.
         * @function verify
         * @memberof bytes2bz.v1.CompressResponse
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        CompressResponse.verify = function verify(message) {
          if (typeof message !== "object" || message === null) {
            return "object expected";
          }
          if (
            message.compressed != null && message.hasOwnProperty("compressed")
          ) {
            if (
              !(message.compressed &&
                  typeof message.compressed.length === "number" ||
                $util.isString(message.compressed))
            ) {
              return "compressed: buffer expected";
            }
          }
          return null;
        };

        /**
         * Creates a CompressResponse message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof bytes2bz.v1.CompressResponse
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {bytes2bz.v1.CompressResponse} CompressResponse
         */
        CompressResponse.fromObject = function fromObject(object) {
          if (object instanceof $root.bytes2bz.v1.CompressResponse) {
            return object;
          }
          var message = new $root.bytes2bz.v1.CompressResponse();
          if (object.compressed != null) {
            if (typeof object.compressed === "string") {
              $util.base64.decode(
                object.compressed,
                message.compressed = $util.newBuffer(
                  $util.base64.length(object.compressed),
                ),
                0,
              );
            } else if (object.compressed.length >= 0) {
              message.compressed = object.compressed;
            }
          }
          return message;
        };

        /**
         * Creates a plain object from a CompressResponse message. Also converts values to other types if specified.
         * @function toObject
         * @memberof bytes2bz.v1.CompressResponse
         * @static
         * @param {bytes2bz.v1.CompressResponse} message CompressResponse
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        CompressResponse.toObject = function toObject(message, options) {
          if (!options) {
            options = {};
          }
          var object = {};
          if (options.defaults) {
            if (options.bytes === String) {
              object.compressed = "";
            } else {
              object.compressed = [];
              if (options.bytes !== Array) {
                object.compressed = $util.newBuffer(object.compressed);
              }
            }
          }
          if (
            message.compressed != null && message.hasOwnProperty("compressed")
          ) {
            object.compressed = options.bytes === String
              ? $util.base64.encode(
                message.compressed,
                0,
                message.compressed.length,
              )
              : options.bytes === Array
              ? Array.prototype.slice.call(message.compressed)
              : message.compressed;
          }
          return object;
        };

        /**
         * Converts this CompressResponse to JSON.
         * @function toJSON
         * @memberof bytes2bz.v1.CompressResponse
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        CompressResponse.prototype.toJSON = function toJSON() {
          return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for CompressResponse
         * @function getTypeUrl
         * @memberof bytes2bz.v1.CompressResponse
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        CompressResponse.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
          if (typeUrlPrefix === undefined) {
            typeUrlPrefix = "type.googleapis.com";
          }
          return typeUrlPrefix + "/bytes2bz.v1.CompressResponse";
        };

        return CompressResponse;
      })();

      v1.CompressService = (function () {
        /**
         * Constructs a new CompressService service.
         * @memberof bytes2bz.v1
         * @classdesc Represents a CompressService
         * @extends $protobuf.rpc.Service
         * @constructor
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         */
        function CompressService(rpcImpl, requestDelimited, responseDelimited) {
          $protobuf.rpc.Service.call(
            this,
            rpcImpl,
            requestDelimited,
            responseDelimited,
          );
        }

        (CompressService.prototype = Object.create(
          $protobuf.rpc.Service.prototype,
        )).constructor = CompressService;

        /**
         * Creates new CompressService service using the specified rpc implementation.
         * @function create
         * @memberof bytes2bz.v1.CompressService
         * @static
         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
         * @returns {CompressService} RPC service. Useful where requests and/or responses are streamed.
         */
        CompressService.create = function create(
          rpcImpl,
          requestDelimited,
          responseDelimited,
        ) {
          return new this(rpcImpl, requestDelimited, responseDelimited);
        };

        /**
         * Callback as used by {@link bytes2bz.v1.CompressService#compress}.
         * @memberof bytes2bz.v1.CompressService
         * @typedef CompressCallback
         * @type {function}
         * @param {Error|null} error Error, if any
         * @param {bytes2bz.v1.CompressResponse} [response] CompressResponse
         */

        /**
         * Calls Compress.
         * @function compress
         * @memberof bytes2bz.v1.CompressService
         * @instance
         * @param {bytes2bz.v1.ICompressRequest} request CompressRequest message or plain object
         * @param {bytes2bz.v1.CompressService.CompressCallback} callback Node-style callback called with the error, if any, and CompressResponse
         * @returns {undefined}
         * @variation 1
         */
        Object.defineProperty(
          CompressService.prototype.compress = function compress(
            request,
            callback,
          ) {
            return this.rpcCall(
              compress,
              $root.bytes2bz.v1.CompressRequest,
              $root.bytes2bz.v1.CompressResponse,
              request,
              callback,
            );
          },
          "name",
          { value: "Compress" },
        );

        /**
         * Calls Compress.
         * @function compress
         * @memberof bytes2bz.v1.CompressService
         * @instance
         * @param {bytes2bz.v1.ICompressRequest} request CompressRequest message or plain object
         * @returns {Promise<bytes2bz.v1.CompressResponse>} Promise
         * @variation 2
         */

        return CompressService;
      })();

      return v1;
    })();

    return bytes2bz;
  })();

  return $root;
});
