// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = "{\n    \"swagger\": \"2.0\",\n    \"info\": {\n        \"title\": \"IMDB API SERVICE\",\n        \"version\": \"1.0\",\n        \"contact\": {\n            \"name\": \"Afdhalul IHsan\",\n            \"email\": \"afdhalulhsn74@gmail.com\"\n        }\n    },\n    \"schemes\": [\n        \"http\"\n    ],\n    \"consumes\": [\n        \"application/json\"\n    ],\n    \"produces\": [\n        \"application/json\"\n    ],\n    \"paths\": {\n        \"/v1/detail/{ImdbId}\": {\n            \"get\": {\n                \"operationId\": \"OmdbMovieService_GetDatilMovie\",\n                \"responses\": {\n                    \"200\": {\n                        \"description\": \"A successful response.\",\n                        \"schema\": {\n                            \"$ref\": \"#/definitions/protoResponseGetDetailMovie\"\n                        }\n                    },\n                    \"404\": {\n                        \"description\": \"Returned when the resource does not exist.\",\n                        \"schema\": {\n                            \"type\": \"string\",\n                            \"format\": \"string\"\n                        }\n                    },\n                    \"default\": {\n                        \"description\": \"An unexpected error response.\",\n                        \"schema\": {\n                            \"$ref\": \"#/definitions/runtimeError\"\n                        }\n                    }\n                },\n                \"parameters\": [\n                    {\n                        \"name\": \"ImdbId\",\n                        \"in\": \"path\",\n                        \"required\": true,\n                        \"type\": \"string\"\n                    }\n                ],\n                \"tags\": [\n                    \"OmdbMovieService\"\n                ]\n            }\n        },\n        \"/v1/getlist\": {\n            \"post\": {\n                \"operationId\": \"OmdbMovieService_GetListMovieOmdb\",\n                \"responses\": {\n                    \"200\": {\n                        \"description\": \"A successful response.\",\n                        \"schema\": {\n                            \"$ref\": \"#/definitions/protoResponseListFilm\"\n                        }\n                    },\n                    \"404\": {\n                        \"description\": \"Returned when the resource does not exist.\",\n                        \"schema\": {\n                            \"type\": \"string\",\n                            \"format\": \"string\"\n                        }\n                    },\n                    \"default\": {\n                        \"description\": \"An unexpected error response.\",\n                        \"schema\": {\n                            \"$ref\": \"#/definitions/runtimeError\"\n                        }\n                    }\n                },\n                \"parameters\": [\n                    {\n                        \"name\": \"body\",\n                        \"in\": \"body\",\n                        \"required\": true,\n                        \"schema\": {\n                            \"$ref\": \"#/definitions/protoGetListMovieRequest\"\n                        }\n                    }\n                ],\n                \"tags\": [\n                    \"OmdbMovieService\"\n                ]\n            }\n        }\n    },\n    \"definitions\": {\n        \"protoDataFilmList\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"Title\": {\n                    \"type\": \"string\"\n                },\n                \"Year\": {\n                    \"type\": \"string\"\n                },\n                \"ImdbId\": {\n                    \"type\": \"string\"\n                },\n                \"Type\": {\n                    \"type\": \"string\"\n                },\n                \"Poster\": {\n                    \"type\": \"string\"\n                },\n                \"Detail\": {\n                    \"$ref\": \"#/definitions/protoResponseGetDetailMovie\"\n                }\n            }\n        },\n        \"protoGetListMovieRequest\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"Keyword\": {\n                    \"type\": \"string\"\n                },\n                \"Page\": {\n                    \"type\": \"integer\",\n                    \"format\": \"int32\"\n                }\n            }\n        },\n        \"protoRatings\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"Source\": {\n                    \"type\": \"string\"\n                },\n                \"Value\": {\n                    \"type\": \"string\"\n                }\n            }\n        },\n        \"protoResponseGetDetailMovie\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"Title\": {\n                    \"type\": \"string\"\n                },\n                \"Year\": {\n                    \"type\": \"string\"\n                },\n                \"Rated\": {\n                    \"type\": \"string\"\n                },\n                \"Released\": {\n                    \"type\": \"string\"\n                },\n                \"Runtime\": {\n                    \"type\": \"string\"\n                },\n                \"Genre\": {\n                    \"type\": \"string\"\n                },\n                \"Director\": {\n                    \"type\": \"string\"\n                },\n                \"Writer\": {\n                    \"type\": \"string\"\n                },\n                \"Actors\": {\n                    \"type\": \"string\"\n                },\n                \"Plot\": {\n                    \"type\": \"string\"\n                },\n                \"Language\": {\n                    \"type\": \"string\"\n                },\n                \"Country\": {\n                    \"type\": \"string\"\n                },\n                \"Awards\": {\n                    \"type\": \"string\"\n                },\n                \"Poster\": {\n                    \"type\": \"string\"\n                },\n                \"Metascore\": {\n                    \"type\": \"string\"\n                },\n                \"imdbRating\": {\n                    \"type\": \"string\"\n                },\n                \"imdbVotes\": {\n                    \"type\": \"string\"\n                },\n                \"imdbID\": {\n                    \"type\": \"string\"\n                },\n                \"Type\": {\n                    \"type\": \"string\"\n                },\n                \"DVD\": {\n                    \"type\": \"string\"\n                },\n                \"BoxOffice\": {\n                    \"type\": \"string\"\n                },\n                \"Production\": {\n                    \"type\": \"string\"\n                },\n                \"Website\": {\n                    \"type\": \"string\"\n                },\n                \"Response\": {\n                    \"type\": \"string\"\n                },\n                \"Rating\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"$ref\": \"#/definitions/protoRatings\"\n                    }\n                }\n            }\n        },\n        \"protoResponseListFilm\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"ListFilm\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"$ref\": \"#/definitions/protoDataFilmList\"\n                    }\n                },\n                \"TotalResult\": {\n                    \"type\": \"string\"\n                }\n            }\n        },\n        \"protobufAny\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"type_url\": {\n                    \"type\": \"string\",\n                    \"description\": \"A URL/resource name that uniquely identifies the type of the serialized\\nprotocol buffer message. This string must contain at least\\none \\\"/\\\" character. The last segment of the URL's path must represent\\nthe fully qualified name of the type (as in\\n`path/google.protobuf.Duration`). The name should be in a canonical form\\n(e.g., leading \\\".\\\" is not accepted).\\n\\nIn practice, teams usually precompile into the binary all types that they\\nexpect it to use in the context of Any. However, for URLs which use the\\nscheme `http`, `https`, or no scheme, one can optionally set up a type\\nserver that maps type URLs to message definitions as follows:\\n\\n* If no scheme is provided, `https` is assumed.\\n* An HTTP GET on the URL must yield a [google.protobuf.Type][]\\n  value in binary format, or produce an error.\\n* Applications are allowed to cache lookup results based on the\\n  URL, or have them precompiled into a binary to avoid any\\n  lookup. Therefore, binary compatibility needs to be preserved\\n  on changes to types. (Use versioned type names to manage\\n  breaking changes.)\\n\\nNote: this functionality is not currently available in the official\\nprotobuf release, and it is not used for type URLs beginning with\\ntype.googleapis.com.\\n\\nSchemes other than `http`, `https` (or the empty scheme) might be\\nused with implementation specific semantics.\"\n                },\n                \"value\": {\n                    \"type\": \"string\",\n                    \"format\": \"byte\",\n                    \"description\": \"Must be a valid serialized protocol buffer of the above specified type.\"\n                }\n            },\n            \"description\": \"`Any` contains an arbitrary serialized protocol buffer message along with a\\nURL that describes the type of the serialized message.\\n\\nProtobuf library provides support to pack/unpack Any values in the form\\nof utility functions or additional generated methods of the Any type.\\n\\nExample 1: Pack and unpack a message in C++.\\n\\n    Foo foo = ...;\\n    Any any;\\n    any.PackFrom(foo);\\n    ...\\n    if (any.UnpackTo(\\u0026foo)) {\\n      ...\\n    }\\n\\nExample 2: Pack and unpack a message in Java.\\n\\n    Foo foo = ...;\\n    Any any = Any.pack(foo);\\n    ...\\n    if (any.is(Foo.class)) {\\n      foo = any.unpack(Foo.class);\\n    }\\n\\n Example 3: Pack and unpack a message in Python.\\n\\n    foo = Foo(...)\\n    any = Any()\\n    any.Pack(foo)\\n    ...\\n    if any.Is(Foo.DESCRIPTOR):\\n      any.Unpack(foo)\\n      ...\\n\\n Example 4: Pack and unpack a message in Go\\n\\n     foo := \\u0026pb.Foo{...}\\n     any, err := anypb.New(foo)\\n     if err != nil {\\n       ...\\n     }\\n     ...\\n     foo := \\u0026pb.Foo{}\\n     if err := any.UnmarshalTo(foo); err != nil {\\n       ...\\n     }\\n\\nThe pack methods provided by protobuf library will by default use\\n'type.googleapis.com/full.type.name' as the type URL and the unpack\\nmethods only use the fully qualified type name after the last '/'\\nin the type URL, for example \\\"foo.bar.com/x/y.z\\\" will yield type\\nname \\\"y.z\\\".\\n\\n\\nJSON\\n====\\nThe JSON representation of an `Any` value uses the regular\\nrepresentation of the deserialized, embedded message, with an\\nadditional field `@type` which contains the type URL. Example:\\n\\n    package google.profile;\\n    message Person {\\n      string first_name = 1;\\n      string last_name = 2;\\n    }\\n\\n    {\\n      \\\"@type\\\": \\\"type.googleapis.com/google.profile.Person\\\",\\n      \\\"firstName\\\": \\u003cstring\\u003e,\\n      \\\"lastName\\\": \\u003cstring\\u003e\\n    }\\n\\nIf the embedded message type is well-known and has a custom JSON\\nrepresentation, that representation will be embedded adding a field\\n`value` which holds the custom JSON in addition to the `@type`\\nfield. Example (for message [google.protobuf.Duration][]):\\n\\n    {\\n      \\\"@type\\\": \\\"type.googleapis.com/google.protobuf.Duration\\\",\\n      \\\"value\\\": \\\"1.212s\\\"\\n    }\"\n        },\n        \"runtimeError\": {\n            \"type\": \"object\",\n            \"properties\": {\n                \"error\": {\n                    \"type\": \"string\"\n                },\n                \"code\": {\n                    \"type\": \"integer\",\n                    \"format\": \"int32\"\n                },\n                \"message\": {\n                    \"type\": \"string\"\n                },\n                \"details\": {\n                    \"type\": \"array\",\n                    \"items\": {\n                        \"$ref\": \"#/definitions/protobufAny\"\n                    }\n                }\n            }\n        }\n    }\n}\n"

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}