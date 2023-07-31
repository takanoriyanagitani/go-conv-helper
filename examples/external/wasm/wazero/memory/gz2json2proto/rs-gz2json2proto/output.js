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

  $root.output = (function () {
    /**
     * Namespace output.
     * @exports output
     * @namespace
     */
    var output = {};

    output.v1 = (function () {
      /**
       * Namespace v1.
       * @memberof output
       * @namespace
       */
      var v1 = {};

      v1.Id = (function () {
        /**
         * Properties of an Id.
         * @memberof output.v1
         * @interface IId
         * @property {string|null} [id] Id id
         */

        /**
         * Constructs a new Id.
         * @memberof output.v1
         * @classdesc Represents an Id.
         * @implements IId
         * @constructor
         * @param {output.v1.IId=} [properties] Properties to set
         */
        function Id(properties) {
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
         * Id id.
         * @member {string} id
         * @memberof output.v1.Id
         * @instance
         */
        Id.prototype.id = "";

        /**
         * Creates a new Id instance using the specified properties.
         * @function create
         * @memberof output.v1.Id
         * @static
         * @param {output.v1.IId=} [properties] Properties to set
         * @returns {output.v1.Id} Id instance
         */
        Id.create = function create(properties) {
          return new Id(properties);
        };

        /**
         * Encodes the specified Id message. Does not implicitly {@link output.v1.Id.verify|verify} messages.
         * @function encode
         * @memberof output.v1.Id
         * @static
         * @param {output.v1.IId} message Id message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Id.encode = function encode(message, writer) {
          if (!writer) {
            writer = $Writer.create();
          }
          if (message.id != null && Object.hasOwnProperty.call(message, "id")) {
            writer.uint32(/* id 1, wireType 2 =*/ 10).string(message.id);
          }
          return writer;
        };

        /**
         * Encodes the specified Id message, length delimited. Does not implicitly {@link output.v1.Id.verify|verify} messages.
         * @function encodeDelimited
         * @memberof output.v1.Id
         * @static
         * @param {output.v1.IId} message Id message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Id.encodeDelimited = function encodeDelimited(message, writer) {
          return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an Id message from the specified reader or buffer.
         * @function decode
         * @memberof output.v1.Id
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {output.v1.Id} Id
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Id.decode = function decode(reader, length) {
          if (!(reader instanceof $Reader)) {
            reader = $Reader.create(reader);
          }
          var end = length === undefined ? reader.len : reader.pos + length,
            message = new $root.output.v1.Id();
          while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
              case 1: {
                message.id = reader.string();
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
         * Decodes an Id message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof output.v1.Id
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {output.v1.Id} Id
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Id.decodeDelimited = function decodeDelimited(reader) {
          if (!(reader instanceof $Reader)) {
            reader = new $Reader(reader);
          }
          return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an Id message.
         * @function verify
         * @memberof output.v1.Id
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Id.verify = function verify(message) {
          if (typeof message !== "object" || message === null) {
            return "object expected";
          }
          if (message.id != null && message.hasOwnProperty("id")) {
            if (!$util.isString(message.id)) {
              return "id: string expected";
            }
          }
          return null;
        };

        /**
         * Creates an Id message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof output.v1.Id
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {output.v1.Id} Id
         */
        Id.fromObject = function fromObject(object) {
          if (object instanceof $root.output.v1.Id) {
            return object;
          }
          var message = new $root.output.v1.Id();
          if (object.id != null) {
            message.id = String(object.id);
          }
          return message;
        };

        /**
         * Creates a plain object from an Id message. Also converts values to other types if specified.
         * @function toObject
         * @memberof output.v1.Id
         * @static
         * @param {output.v1.Id} message Id
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Id.toObject = function toObject(message, options) {
          if (!options) {
            options = {};
          }
          var object = {};
          if (options.defaults) {
            object.id = "";
          }
          if (message.id != null && message.hasOwnProperty("id")) {
            object.id = message.id;
          }
          return object;
        };

        /**
         * Converts this Id to JSON.
         * @function toJSON
         * @memberof output.v1.Id
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Id.prototype.toJSON = function toJSON() {
          return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for Id
         * @function getTypeUrl
         * @memberof output.v1.Id
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        Id.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
          if (typeUrlPrefix === undefined) {
            typeUrlPrefix = "type.googleapis.com";
          }
          return typeUrlPrefix + "/output.v1.Id";
        };

        return Id;
      })();

      v1.Name = (function () {
        /**
         * Properties of a Name.
         * @memberof output.v1
         * @interface IName
         * @property {string|null} [name] Name name
         */

        /**
         * Constructs a new Name.
         * @memberof output.v1
         * @classdesc Represents a Name.
         * @implements IName
         * @constructor
         * @param {output.v1.IName=} [properties] Properties to set
         */
        function Name(properties) {
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
         * Name name.
         * @member {string} name
         * @memberof output.v1.Name
         * @instance
         */
        Name.prototype.name = "";

        /**
         * Creates a new Name instance using the specified properties.
         * @function create
         * @memberof output.v1.Name
         * @static
         * @param {output.v1.IName=} [properties] Properties to set
         * @returns {output.v1.Name} Name instance
         */
        Name.create = function create(properties) {
          return new Name(properties);
        };

        /**
         * Encodes the specified Name message. Does not implicitly {@link output.v1.Name.verify|verify} messages.
         * @function encode
         * @memberof output.v1.Name
         * @static
         * @param {output.v1.IName} message Name message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Name.encode = function encode(message, writer) {
          if (!writer) {
            writer = $Writer.create();
          }
          if (
            message.name != null && Object.hasOwnProperty.call(message, "name")
          ) {
            writer.uint32(/* id 1, wireType 2 =*/ 10).string(message.name);
          }
          return writer;
        };

        /**
         * Encodes the specified Name message, length delimited. Does not implicitly {@link output.v1.Name.verify|verify} messages.
         * @function encodeDelimited
         * @memberof output.v1.Name
         * @static
         * @param {output.v1.IName} message Name message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Name.encodeDelimited = function encodeDelimited(message, writer) {
          return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Name message from the specified reader or buffer.
         * @function decode
         * @memberof output.v1.Name
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {output.v1.Name} Name
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Name.decode = function decode(reader, length) {
          if (!(reader instanceof $Reader)) {
            reader = $Reader.create(reader);
          }
          var end = length === undefined ? reader.len : reader.pos + length,
            message = new $root.output.v1.Name();
          while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
              case 1: {
                message.name = reader.string();
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
         * Decodes a Name message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof output.v1.Name
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {output.v1.Name} Name
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Name.decodeDelimited = function decodeDelimited(reader) {
          if (!(reader instanceof $Reader)) {
            reader = new $Reader(reader);
          }
          return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Name message.
         * @function verify
         * @memberof output.v1.Name
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Name.verify = function verify(message) {
          if (typeof message !== "object" || message === null) {
            return "object expected";
          }
          if (message.name != null && message.hasOwnProperty("name")) {
            if (!$util.isString(message.name)) {
              return "name: string expected";
            }
          }
          return null;
        };

        /**
         * Creates a Name message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof output.v1.Name
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {output.v1.Name} Name
         */
        Name.fromObject = function fromObject(object) {
          if (object instanceof $root.output.v1.Name) {
            return object;
          }
          var message = new $root.output.v1.Name();
          if (object.name != null) {
            message.name = String(object.name);
          }
          return message;
        };

        /**
         * Creates a plain object from a Name message. Also converts values to other types if specified.
         * @function toObject
         * @memberof output.v1.Name
         * @static
         * @param {output.v1.Name} message Name
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Name.toObject = function toObject(message, options) {
          if (!options) {
            options = {};
          }
          var object = {};
          if (options.defaults) {
            object.name = "";
          }
          if (message.name != null && message.hasOwnProperty("name")) {
            object.name = message.name;
          }
          return object;
        };

        /**
         * Converts this Name to JSON.
         * @function toJSON
         * @memberof output.v1.Name
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Name.prototype.toJSON = function toJSON() {
          return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for Name
         * @function getTypeUrl
         * @memberof output.v1.Name
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        Name.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
          if (typeUrlPrefix === undefined) {
            typeUrlPrefix = "type.googleapis.com";
          }
          return typeUrlPrefix + "/output.v1.Name";
        };

        return Name;
      })();

      v1.Device = (function () {
        /**
         * Properties of a Device.
         * @memberof output.v1
         * @interface IDevice
         * @property {output.v1.IId|null} [id] Device id
         * @property {output.v1.IName|null} [name] Device name
         */

        /**
         * Constructs a new Device.
         * @memberof output.v1
         * @classdesc Represents a Device.
         * @implements IDevice
         * @constructor
         * @param {output.v1.IDevice=} [properties] Properties to set
         */
        function Device(properties) {
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
         * Device id.
         * @member {output.v1.IId|null|undefined} id
         * @memberof output.v1.Device
         * @instance
         */
        Device.prototype.id = null;

        /**
         * Device name.
         * @member {output.v1.IName|null|undefined} name
         * @memberof output.v1.Device
         * @instance
         */
        Device.prototype.name = null;

        /**
         * Creates a new Device instance using the specified properties.
         * @function create
         * @memberof output.v1.Device
         * @static
         * @param {output.v1.IDevice=} [properties] Properties to set
         * @returns {output.v1.Device} Device instance
         */
        Device.create = function create(properties) {
          return new Device(properties);
        };

        /**
         * Encodes the specified Device message. Does not implicitly {@link output.v1.Device.verify|verify} messages.
         * @function encode
         * @memberof output.v1.Device
         * @static
         * @param {output.v1.IDevice} message Device message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Device.encode = function encode(message, writer) {
          if (!writer) {
            writer = $Writer.create();
          }
          if (message.id != null && Object.hasOwnProperty.call(message, "id")) {
            $root.output.v1.Id.encode(
              message.id,
              writer.uint32(/* id 1, wireType 2 =*/ 10).fork(),
            ).ldelim();
          }
          if (
            message.name != null && Object.hasOwnProperty.call(message, "name")
          ) {
            $root.output.v1.Name.encode(
              message.name,
              writer.uint32(/* id 2, wireType 2 =*/ 18).fork(),
            ).ldelim();
          }
          return writer;
        };

        /**
         * Encodes the specified Device message, length delimited. Does not implicitly {@link output.v1.Device.verify|verify} messages.
         * @function encodeDelimited
         * @memberof output.v1.Device
         * @static
         * @param {output.v1.IDevice} message Device message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Device.encodeDelimited = function encodeDelimited(message, writer) {
          return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Device message from the specified reader or buffer.
         * @function decode
         * @memberof output.v1.Device
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {output.v1.Device} Device
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Device.decode = function decode(reader, length) {
          if (!(reader instanceof $Reader)) {
            reader = $Reader.create(reader);
          }
          var end = length === undefined ? reader.len : reader.pos + length,
            message = new $root.output.v1.Device();
          while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
              case 1: {
                message.id = $root.output.v1.Id.decode(reader, reader.uint32());
                break;
              }
              case 2: {
                message.name = $root.output.v1.Name.decode(
                  reader,
                  reader.uint32(),
                );
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
         * Decodes a Device message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof output.v1.Device
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {output.v1.Device} Device
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Device.decodeDelimited = function decodeDelimited(reader) {
          if (!(reader instanceof $Reader)) {
            reader = new $Reader(reader);
          }
          return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Device message.
         * @function verify
         * @memberof output.v1.Device
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Device.verify = function verify(message) {
          if (typeof message !== "object" || message === null) {
            return "object expected";
          }
          if (message.id != null && message.hasOwnProperty("id")) {
            var error = $root.output.v1.Id.verify(message.id);
            if (error) {
              return "id." + error;
            }
          }
          if (message.name != null && message.hasOwnProperty("name")) {
            var error = $root.output.v1.Name.verify(message.name);
            if (error) {
              return "name." + error;
            }
          }
          return null;
        };

        /**
         * Creates a Device message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof output.v1.Device
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {output.v1.Device} Device
         */
        Device.fromObject = function fromObject(object) {
          if (object instanceof $root.output.v1.Device) {
            return object;
          }
          var message = new $root.output.v1.Device();
          if (object.id != null) {
            if (typeof object.id !== "object") {
              throw TypeError(".output.v1.Device.id: object expected");
            }
            message.id = $root.output.v1.Id.fromObject(object.id);
          }
          if (object.name != null) {
            if (typeof object.name !== "object") {
              throw TypeError(".output.v1.Device.name: object expected");
            }
            message.name = $root.output.v1.Name.fromObject(object.name);
          }
          return message;
        };

        /**
         * Creates a plain object from a Device message. Also converts values to other types if specified.
         * @function toObject
         * @memberof output.v1.Device
         * @static
         * @param {output.v1.Device} message Device
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Device.toObject = function toObject(message, options) {
          if (!options) {
            options = {};
          }
          var object = {};
          if (options.defaults) {
            object.id = null;
            object.name = null;
          }
          if (message.id != null && message.hasOwnProperty("id")) {
            object.id = $root.output.v1.Id.toObject(message.id, options);
          }
          if (message.name != null && message.hasOwnProperty("name")) {
            object.name = $root.output.v1.Name.toObject(message.name, options);
          }
          return object;
        };

        /**
         * Converts this Device to JSON.
         * @function toJSON
         * @memberof output.v1.Device
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Device.prototype.toJSON = function toJSON() {
          return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for Device
         * @function getTypeUrl
         * @memberof output.v1.Device
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        Device.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
          if (typeUrlPrefix === undefined) {
            typeUrlPrefix = "type.googleapis.com";
          }
          return typeUrlPrefix + "/output.v1.Device";
        };

        return Device;
      })();

      v1.Current = (function () {
        /**
         * Properties of a Current.
         * @memberof output.v1
         * @interface ICurrent
         * @property {number|null} [current] Current current
         */

        /**
         * Constructs a new Current.
         * @memberof output.v1
         * @classdesc Represents a Current.
         * @implements ICurrent
         * @constructor
         * @param {output.v1.ICurrent=} [properties] Properties to set
         */
        function Current(properties) {
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
         * Current current.
         * @member {number} current
         * @memberof output.v1.Current
         * @instance
         */
        Current.prototype.current = 0;

        /**
         * Creates a new Current instance using the specified properties.
         * @function create
         * @memberof output.v1.Current
         * @static
         * @param {output.v1.ICurrent=} [properties] Properties to set
         * @returns {output.v1.Current} Current instance
         */
        Current.create = function create(properties) {
          return new Current(properties);
        };

        /**
         * Encodes the specified Current message. Does not implicitly {@link output.v1.Current.verify|verify} messages.
         * @function encode
         * @memberof output.v1.Current
         * @static
         * @param {output.v1.ICurrent} message Current message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Current.encode = function encode(message, writer) {
          if (!writer) {
            writer = $Writer.create();
          }
          if (
            message.current != null &&
            Object.hasOwnProperty.call(message, "current")
          ) {
            writer.uint32(/* id 1, wireType 1 =*/ 9).double(message.current);
          }
          return writer;
        };

        /**
         * Encodes the specified Current message, length delimited. Does not implicitly {@link output.v1.Current.verify|verify} messages.
         * @function encodeDelimited
         * @memberof output.v1.Current
         * @static
         * @param {output.v1.ICurrent} message Current message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Current.encodeDelimited = function encodeDelimited(message, writer) {
          return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Current message from the specified reader or buffer.
         * @function decode
         * @memberof output.v1.Current
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {output.v1.Current} Current
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Current.decode = function decode(reader, length) {
          if (!(reader instanceof $Reader)) {
            reader = $Reader.create(reader);
          }
          var end = length === undefined ? reader.len : reader.pos + length,
            message = new $root.output.v1.Current();
          while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
              case 1: {
                message.current = reader.double();
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
         * Decodes a Current message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof output.v1.Current
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {output.v1.Current} Current
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Current.decodeDelimited = function decodeDelimited(reader) {
          if (!(reader instanceof $Reader)) {
            reader = new $Reader(reader);
          }
          return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Current message.
         * @function verify
         * @memberof output.v1.Current
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Current.verify = function verify(message) {
          if (typeof message !== "object" || message === null) {
            return "object expected";
          }
          if (message.current != null && message.hasOwnProperty("current")) {
            if (typeof message.current !== "number") {
              return "current: number expected";
            }
          }
          return null;
        };

        /**
         * Creates a Current message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof output.v1.Current
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {output.v1.Current} Current
         */
        Current.fromObject = function fromObject(object) {
          if (object instanceof $root.output.v1.Current) {
            return object;
          }
          var message = new $root.output.v1.Current();
          if (object.current != null) {
            message.current = Number(object.current);
          }
          return message;
        };

        /**
         * Creates a plain object from a Current message. Also converts values to other types if specified.
         * @function toObject
         * @memberof output.v1.Current
         * @static
         * @param {output.v1.Current} message Current
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Current.toObject = function toObject(message, options) {
          if (!options) {
            options = {};
          }
          var object = {};
          if (options.defaults) {
            object.current = 0;
          }
          if (message.current != null && message.hasOwnProperty("current")) {
            object.current = options.json && !isFinite(message.current)
              ? String(message.current)
              : message.current;
          }
          return object;
        };

        /**
         * Converts this Current to JSON.
         * @function toJSON
         * @memberof output.v1.Current
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Current.prototype.toJSON = function toJSON() {
          return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for Current
         * @function getTypeUrl
         * @memberof output.v1.Current
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        Current.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
          if (typeUrlPrefix === undefined) {
            typeUrlPrefix = "type.googleapis.com";
          }
          return typeUrlPrefix + "/output.v1.Current";
        };

        return Current;
      })();

      v1.Timestamp = (function () {
        /**
         * Properties of a Timestamp.
         * @memberof output.v1
         * @interface ITimestamp
         * @property {number|Long|null} [unixtimeSeconds] Timestamp unixtimeSeconds
         */

        /**
         * Constructs a new Timestamp.
         * @memberof output.v1
         * @classdesc Represents a Timestamp.
         * @implements ITimestamp
         * @constructor
         * @param {output.v1.ITimestamp=} [properties] Properties to set
         */
        function Timestamp(properties) {
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
         * Timestamp unixtimeSeconds.
         * @member {number|Long} unixtimeSeconds
         * @memberof output.v1.Timestamp
         * @instance
         */
        Timestamp.prototype.unixtimeSeconds = $util.Long
          ? $util.Long.fromBits(0, 0, false)
          : 0;

        /**
         * Creates a new Timestamp instance using the specified properties.
         * @function create
         * @memberof output.v1.Timestamp
         * @static
         * @param {output.v1.ITimestamp=} [properties] Properties to set
         * @returns {output.v1.Timestamp} Timestamp instance
         */
        Timestamp.create = function create(properties) {
          return new Timestamp(properties);
        };

        /**
         * Encodes the specified Timestamp message. Does not implicitly {@link output.v1.Timestamp.verify|verify} messages.
         * @function encode
         * @memberof output.v1.Timestamp
         * @static
         * @param {output.v1.ITimestamp} message Timestamp message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Timestamp.encode = function encode(message, writer) {
          if (!writer) {
            writer = $Writer.create();
          }
          if (
            message.unixtimeSeconds != null &&
            Object.hasOwnProperty.call(message, "unixtimeSeconds")
          ) {
            writer.uint32(/* id 1, wireType 1 =*/ 9).sfixed64(
              message.unixtimeSeconds,
            );
          }
          return writer;
        };

        /**
         * Encodes the specified Timestamp message, length delimited. Does not implicitly {@link output.v1.Timestamp.verify|verify} messages.
         * @function encodeDelimited
         * @memberof output.v1.Timestamp
         * @static
         * @param {output.v1.ITimestamp} message Timestamp message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Timestamp.encodeDelimited = function encodeDelimited(message, writer) {
          return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Timestamp message from the specified reader or buffer.
         * @function decode
         * @memberof output.v1.Timestamp
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {output.v1.Timestamp} Timestamp
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Timestamp.decode = function decode(reader, length) {
          if (!(reader instanceof $Reader)) {
            reader = $Reader.create(reader);
          }
          var end = length === undefined ? reader.len : reader.pos + length,
            message = new $root.output.v1.Timestamp();
          while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
              case 1: {
                message.unixtimeSeconds = reader.sfixed64();
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
         * Decodes a Timestamp message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof output.v1.Timestamp
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {output.v1.Timestamp} Timestamp
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Timestamp.decodeDelimited = function decodeDelimited(reader) {
          if (!(reader instanceof $Reader)) {
            reader = new $Reader(reader);
          }
          return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Timestamp message.
         * @function verify
         * @memberof output.v1.Timestamp
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Timestamp.verify = function verify(message) {
          if (typeof message !== "object" || message === null) {
            return "object expected";
          }
          if (
            message.unixtimeSeconds != null &&
            message.hasOwnProperty("unixtimeSeconds")
          ) {
            if (
              !$util.isInteger(message.unixtimeSeconds) &&
              !(message.unixtimeSeconds &&
                $util.isInteger(message.unixtimeSeconds.low) &&
                $util.isInteger(message.unixtimeSeconds.high))
            ) {
              return "unixtimeSeconds: integer|Long expected";
            }
          }
          return null;
        };

        /**
         * Creates a Timestamp message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof output.v1.Timestamp
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {output.v1.Timestamp} Timestamp
         */
        Timestamp.fromObject = function fromObject(object) {
          if (object instanceof $root.output.v1.Timestamp) {
            return object;
          }
          var message = new $root.output.v1.Timestamp();
          if (object.unixtimeSeconds != null) {
            if ($util.Long) {
              (message.unixtimeSeconds = $util.Long.fromValue(
                object.unixtimeSeconds,
              )).unsigned = false;
            } else if (typeof object.unixtimeSeconds === "string") {
              message.unixtimeSeconds = parseInt(object.unixtimeSeconds, 10);
            } else if (typeof object.unixtimeSeconds === "number") {
              message.unixtimeSeconds = object.unixtimeSeconds;
            } else if (typeof object.unixtimeSeconds === "object") {
              message.unixtimeSeconds = new $util.LongBits(
                object.unixtimeSeconds.low >>> 0,
                object.unixtimeSeconds.high >>> 0,
              ).toNumber();
            }
          }
          return message;
        };

        /**
         * Creates a plain object from a Timestamp message. Also converts values to other types if specified.
         * @function toObject
         * @memberof output.v1.Timestamp
         * @static
         * @param {output.v1.Timestamp} message Timestamp
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Timestamp.toObject = function toObject(message, options) {
          if (!options) {
            options = {};
          }
          var object = {};
          if (options.defaults) {
            if ($util.Long) {
              var long = new $util.Long(0, 0, false);
              object.unixtimeSeconds = options.longs === String
                ? long.toString()
                : options.longs === Number
                ? long.toNumber()
                : long;
            } else {
              object.unixtimeSeconds = options.longs === String ? "0" : 0;
            }
          }
          if (
            message.unixtimeSeconds != null &&
            message.hasOwnProperty("unixtimeSeconds")
          ) {
            if (typeof message.unixtimeSeconds === "number") {
              object.unixtimeSeconds = options.longs === String
                ? String(message.unixtimeSeconds)
                : message.unixtimeSeconds;
            } else {
              object.unixtimeSeconds = options.longs === String
                ? $util.Long.prototype.toString.call(message.unixtimeSeconds)
                : options.longs === Number
                ? new $util.LongBits(
                  message.unixtimeSeconds.low >>> 0,
                  message.unixtimeSeconds.high >>> 0,
                ).toNumber()
                : message.unixtimeSeconds;
            }
          }
          return object;
        };

        /**
         * Converts this Timestamp to JSON.
         * @function toJSON
         * @memberof output.v1.Timestamp
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Timestamp.prototype.toJSON = function toJSON() {
          return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for Timestamp
         * @function getTypeUrl
         * @memberof output.v1.Timestamp
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        Timestamp.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
          if (typeUrlPrefix === undefined) {
            typeUrlPrefix = "type.googleapis.com";
          }
          return typeUrlPrefix + "/output.v1.Timestamp";
        };

        return Timestamp;
      })();

      v1.Detail = (function () {
        /**
         * Properties of a Detail.
         * @memberof output.v1
         * @interface IDetail
         * @property {number|null} [timestampNanos] Detail timestampNanos
         * @property {number|null} [voltage] Detail voltage
         */

        /**
         * Constructs a new Detail.
         * @memberof output.v1
         * @classdesc Represents a Detail.
         * @implements IDetail
         * @constructor
         * @param {output.v1.IDetail=} [properties] Properties to set
         */
        function Detail(properties) {
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
         * Detail timestampNanos.
         * @member {number} timestampNanos
         * @memberof output.v1.Detail
         * @instance
         */
        Detail.prototype.timestampNanos = 0;

        /**
         * Detail voltage.
         * @member {number} voltage
         * @memberof output.v1.Detail
         * @instance
         */
        Detail.prototype.voltage = 0;

        /**
         * Creates a new Detail instance using the specified properties.
         * @function create
         * @memberof output.v1.Detail
         * @static
         * @param {output.v1.IDetail=} [properties] Properties to set
         * @returns {output.v1.Detail} Detail instance
         */
        Detail.create = function create(properties) {
          return new Detail(properties);
        };

        /**
         * Encodes the specified Detail message. Does not implicitly {@link output.v1.Detail.verify|verify} messages.
         * @function encode
         * @memberof output.v1.Detail
         * @static
         * @param {output.v1.IDetail} message Detail message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Detail.encode = function encode(message, writer) {
          if (!writer) {
            writer = $Writer.create();
          }
          if (
            message.timestampNanos != null &&
            Object.hasOwnProperty.call(message, "timestampNanos")
          ) {
            writer.uint32(/* id 1, wireType 5 =*/ 13).fixed32(
              message.timestampNanos,
            );
          }
          if (
            message.voltage != null &&
            Object.hasOwnProperty.call(message, "voltage")
          ) {
            writer.uint32(/* id 2, wireType 5 =*/ 21).float(message.voltage);
          }
          return writer;
        };

        /**
         * Encodes the specified Detail message, length delimited. Does not implicitly {@link output.v1.Detail.verify|verify} messages.
         * @function encodeDelimited
         * @memberof output.v1.Detail
         * @static
         * @param {output.v1.IDetail} message Detail message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Detail.encodeDelimited = function encodeDelimited(message, writer) {
          return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Detail message from the specified reader or buffer.
         * @function decode
         * @memberof output.v1.Detail
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {output.v1.Detail} Detail
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Detail.decode = function decode(reader, length) {
          if (!(reader instanceof $Reader)) {
            reader = $Reader.create(reader);
          }
          var end = length === undefined ? reader.len : reader.pos + length,
            message = new $root.output.v1.Detail();
          while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
              case 1: {
                message.timestampNanos = reader.fixed32();
                break;
              }
              case 2: {
                message.voltage = reader.float();
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
         * Decodes a Detail message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof output.v1.Detail
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {output.v1.Detail} Detail
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Detail.decodeDelimited = function decodeDelimited(reader) {
          if (!(reader instanceof $Reader)) {
            reader = new $Reader(reader);
          }
          return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Detail message.
         * @function verify
         * @memberof output.v1.Detail
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Detail.verify = function verify(message) {
          if (typeof message !== "object" || message === null) {
            return "object expected";
          }
          if (
            message.timestampNanos != null &&
            message.hasOwnProperty("timestampNanos")
          ) {
            if (!$util.isInteger(message.timestampNanos)) {
              return "timestampNanos: integer expected";
            }
          }
          if (message.voltage != null && message.hasOwnProperty("voltage")) {
            if (typeof message.voltage !== "number") {
              return "voltage: number expected";
            }
          }
          return null;
        };

        /**
         * Creates a Detail message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof output.v1.Detail
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {output.v1.Detail} Detail
         */
        Detail.fromObject = function fromObject(object) {
          if (object instanceof $root.output.v1.Detail) {
            return object;
          }
          var message = new $root.output.v1.Detail();
          if (object.timestampNanos != null) {
            message.timestampNanos = object.timestampNanos >>> 0;
          }
          if (object.voltage != null) {
            message.voltage = Number(object.voltage);
          }
          return message;
        };

        /**
         * Creates a plain object from a Detail message. Also converts values to other types if specified.
         * @function toObject
         * @memberof output.v1.Detail
         * @static
         * @param {output.v1.Detail} message Detail
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Detail.toObject = function toObject(message, options) {
          if (!options) {
            options = {};
          }
          var object = {};
          if (options.defaults) {
            object.timestampNanos = 0;
            object.voltage = 0;
          }
          if (
            message.timestampNanos != null &&
            message.hasOwnProperty("timestampNanos")
          ) {
            object.timestampNanos = message.timestampNanos;
          }
          if (message.voltage != null && message.hasOwnProperty("voltage")) {
            object.voltage = options.json && !isFinite(message.voltage)
              ? String(message.voltage)
              : message.voltage;
          }
          return object;
        };

        /**
         * Converts this Detail to JSON.
         * @function toJSON
         * @memberof output.v1.Detail
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Detail.prototype.toJSON = function toJSON() {
          return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for Detail
         * @function getTypeUrl
         * @memberof output.v1.Detail
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        Detail.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
          if (typeUrlPrefix === undefined) {
            typeUrlPrefix = "type.googleapis.com";
          }
          return typeUrlPrefix + "/output.v1.Detail";
        };

        return Detail;
      })();

      v1.Output = (function () {
        /**
         * Properties of an Output.
         * @memberof output.v1
         * @interface IOutput
         * @property {output.v1.IDevice|null} [device] Output device
         * @property {output.v1.ITimestamp|null} [timestamp] Output timestamp
         * @property {output.v1.ICurrent|null} [current] Output current
         * @property {Array.<output.v1.IDetail>|null} [data] Output data
         */

        /**
         * Constructs a new Output.
         * @memberof output.v1
         * @classdesc Represents an Output.
         * @implements IOutput
         * @constructor
         * @param {output.v1.IOutput=} [properties] Properties to set
         */
        function Output(properties) {
          this.data = [];
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
         * Output device.
         * @member {output.v1.IDevice|null|undefined} device
         * @memberof output.v1.Output
         * @instance
         */
        Output.prototype.device = null;

        /**
         * Output timestamp.
         * @member {output.v1.ITimestamp|null|undefined} timestamp
         * @memberof output.v1.Output
         * @instance
         */
        Output.prototype.timestamp = null;

        /**
         * Output current.
         * @member {output.v1.ICurrent|null|undefined} current
         * @memberof output.v1.Output
         * @instance
         */
        Output.prototype.current = null;

        /**
         * Output data.
         * @member {Array.<output.v1.IDetail>} data
         * @memberof output.v1.Output
         * @instance
         */
        Output.prototype.data = $util.emptyArray;

        /**
         * Creates a new Output instance using the specified properties.
         * @function create
         * @memberof output.v1.Output
         * @static
         * @param {output.v1.IOutput=} [properties] Properties to set
         * @returns {output.v1.Output} Output instance
         */
        Output.create = function create(properties) {
          return new Output(properties);
        };

        /**
         * Encodes the specified Output message. Does not implicitly {@link output.v1.Output.verify|verify} messages.
         * @function encode
         * @memberof output.v1.Output
         * @static
         * @param {output.v1.IOutput} message Output message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Output.encode = function encode(message, writer) {
          if (!writer) {
            writer = $Writer.create();
          }
          if (
            message.device != null &&
            Object.hasOwnProperty.call(message, "device")
          ) {
            $root.output.v1.Device.encode(
              message.device,
              writer.uint32(/* id 1, wireType 2 =*/ 10).fork(),
            ).ldelim();
          }
          if (
            message.timestamp != null &&
            Object.hasOwnProperty.call(message, "timestamp")
          ) {
            $root.output.v1.Timestamp.encode(
              message.timestamp,
              writer.uint32(/* id 2, wireType 2 =*/ 18).fork(),
            ).ldelim();
          }
          if (
            message.current != null &&
            Object.hasOwnProperty.call(message, "current")
          ) {
            $root.output.v1.Current.encode(
              message.current,
              writer.uint32(/* id 3, wireType 2 =*/ 26).fork(),
            ).ldelim();
          }
          if (message.data != null && message.data.length) {
            for (var i = 0; i < message.data.length; ++i) {
              $root.output.v1.Detail.encode(
                message.data[i],
                writer.uint32(/* id 4, wireType 2 =*/ 34).fork(),
              ).ldelim();
            }
          }
          return writer;
        };

        /**
         * Encodes the specified Output message, length delimited. Does not implicitly {@link output.v1.Output.verify|verify} messages.
         * @function encodeDelimited
         * @memberof output.v1.Output
         * @static
         * @param {output.v1.IOutput} message Output message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Output.encodeDelimited = function encodeDelimited(message, writer) {
          return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an Output message from the specified reader or buffer.
         * @function decode
         * @memberof output.v1.Output
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {output.v1.Output} Output
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Output.decode = function decode(reader, length) {
          if (!(reader instanceof $Reader)) {
            reader = $Reader.create(reader);
          }
          var end = length === undefined ? reader.len : reader.pos + length,
            message = new $root.output.v1.Output();
          while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
              case 1: {
                message.device = $root.output.v1.Device.decode(
                  reader,
                  reader.uint32(),
                );
                break;
              }
              case 2: {
                message.timestamp = $root.output.v1.Timestamp.decode(
                  reader,
                  reader.uint32(),
                );
                break;
              }
              case 3: {
                message.current = $root.output.v1.Current.decode(
                  reader,
                  reader.uint32(),
                );
                break;
              }
              case 4: {
                if (!(message.data && message.data.length)) {
                  message.data = [];
                }
                message.data.push(
                  $root.output.v1.Detail.decode(reader, reader.uint32()),
                );
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
         * Decodes an Output message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof output.v1.Output
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {output.v1.Output} Output
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Output.decodeDelimited = function decodeDelimited(reader) {
          if (!(reader instanceof $Reader)) {
            reader = new $Reader(reader);
          }
          return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an Output message.
         * @function verify
         * @memberof output.v1.Output
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Output.verify = function verify(message) {
          if (typeof message !== "object" || message === null) {
            return "object expected";
          }
          if (message.device != null && message.hasOwnProperty("device")) {
            var error = $root.output.v1.Device.verify(message.device);
            if (error) {
              return "device." + error;
            }
          }
          if (
            message.timestamp != null && message.hasOwnProperty("timestamp")
          ) {
            var error = $root.output.v1.Timestamp.verify(message.timestamp);
            if (error) {
              return "timestamp." + error;
            }
          }
          if (message.current != null && message.hasOwnProperty("current")) {
            var error = $root.output.v1.Current.verify(message.current);
            if (error) {
              return "current." + error;
            }
          }
          if (message.data != null && message.hasOwnProperty("data")) {
            if (!Array.isArray(message.data)) {
              return "data: array expected";
            }
            for (var i = 0; i < message.data.length; ++i) {
              var error = $root.output.v1.Detail.verify(message.data[i]);
              if (error) {
                return "data." + error;
              }
            }
          }
          return null;
        };

        /**
         * Creates an Output message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof output.v1.Output
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {output.v1.Output} Output
         */
        Output.fromObject = function fromObject(object) {
          if (object instanceof $root.output.v1.Output) {
            return object;
          }
          var message = new $root.output.v1.Output();
          if (object.device != null) {
            if (typeof object.device !== "object") {
              throw TypeError(".output.v1.Output.device: object expected");
            }
            message.device = $root.output.v1.Device.fromObject(object.device);
          }
          if (object.timestamp != null) {
            if (typeof object.timestamp !== "object") {
              throw TypeError(".output.v1.Output.timestamp: object expected");
            }
            message.timestamp = $root.output.v1.Timestamp.fromObject(
              object.timestamp,
            );
          }
          if (object.current != null) {
            if (typeof object.current !== "object") {
              throw TypeError(".output.v1.Output.current: object expected");
            }
            message.current = $root.output.v1.Current.fromObject(
              object.current,
            );
          }
          if (object.data) {
            if (!Array.isArray(object.data)) {
              throw TypeError(".output.v1.Output.data: array expected");
            }
            message.data = [];
            for (var i = 0; i < object.data.length; ++i) {
              if (typeof object.data[i] !== "object") {
                throw TypeError(".output.v1.Output.data: object expected");
              }
              message.data[i] = $root.output.v1.Detail.fromObject(
                object.data[i],
              );
            }
          }
          return message;
        };

        /**
         * Creates a plain object from an Output message. Also converts values to other types if specified.
         * @function toObject
         * @memberof output.v1.Output
         * @static
         * @param {output.v1.Output} message Output
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Output.toObject = function toObject(message, options) {
          if (!options) {
            options = {};
          }
          var object = {};
          if (options.arrays || options.defaults) {
            object.data = [];
          }
          if (options.defaults) {
            object.device = null;
            object.timestamp = null;
            object.current = null;
          }
          if (message.device != null && message.hasOwnProperty("device")) {
            object.device = $root.output.v1.Device.toObject(
              message.device,
              options,
            );
          }
          if (
            message.timestamp != null && message.hasOwnProperty("timestamp")
          ) {
            object.timestamp = $root.output.v1.Timestamp.toObject(
              message.timestamp,
              options,
            );
          }
          if (message.current != null && message.hasOwnProperty("current")) {
            object.current = $root.output.v1.Current.toObject(
              message.current,
              options,
            );
          }
          if (message.data && message.data.length) {
            object.data = [];
            for (var j = 0; j < message.data.length; ++j) {
              object.data[j] = $root.output.v1.Detail.toObject(
                message.data[j],
                options,
              );
            }
          }
          return object;
        };

        /**
         * Converts this Output to JSON.
         * @function toJSON
         * @memberof output.v1.Output
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Output.prototype.toJSON = function toJSON() {
          return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for Output
         * @function getTypeUrl
         * @memberof output.v1.Output
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        Output.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
          if (typeUrlPrefix === undefined) {
            typeUrlPrefix = "type.googleapis.com";
          }
          return typeUrlPrefix + "/output.v1.Output";
        };

        return Output;
      })();

      return v1;
    })();

    return output;
  })();

  return $root;
});
