// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "email": "adecandra1500@gmail.com"
        },
        "license": {
            "url": "http://github.com/AdeCandra12"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/ins": {
            "post": {
                "description": "Input data presensi.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Presensi"
                ],
                "summary": "Insert data presensi.",
                "parameters": [
                    {
                        "description": "Payload Body [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.Presensi"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Presensi"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/presensi": {
            "get": {
                "description": "Mengambil semua data presensi.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Presensi"
                ],
                "summary": "Get All Data Presensi.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Presensi"
                        }
                    }
                }
            }
        },
        "/presensi/{id}": {
            "get": {
                "description": "Ambil per ID data presensi.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Presensi"
                ],
                "summary": "Get By ID Data Presensi.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Presensi"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.JamKerja": {
            "type": "object",
            "properties": {
                "durasi": {
                    "type": "integer",
                    "example": 8
                },
                "gmt": {
                    "type": "integer",
                    "example": 7
                },
                "hari": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "Senin",
                        "Selasa",
                        "Rabu",
                        "Kamis",
                        "Jumat",
                        "Sabtu",
                        "Minggu"
                    ]
                },
                "jam_keluar": {
                    "type": "string",
                    "example": "16:00"
                },
                "jam_masuk": {
                    "type": "string",
                    "example": "08:00"
                },
                "piket_tim": {
                    "type": "string",
                    "example": "Piket Z"
                },
                "shift": {
                    "type": "integer",
                    "example": 2
                }
            }
        },
        "controller.Karyawan": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "hari_kerja": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "Senin",
                        "Selasa",
                        "Rabu",
                        "Kamis",
                        "Jumat",
                        "Sabtu",
                        "Minggu"
                    ]
                },
                "jabatan": {
                    "type": "string",
                    "example": "Anonymous"
                },
                "jam_kerja": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controller.JamKerja"
                    }
                },
                "nama": {
                    "type": "string",
                    "example": "Tes Swagger"
                },
                "phone_number": {
                    "type": "string",
                    "example": "08123456789"
                }
            }
        },
        "controller.Presensi": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "biodata": {
                    "$ref": "#/definitions/controller.Karyawan"
                },
                "checkin": {
                    "description": "Datetime     primitive.DateTime ` + "`" + `bson:\"datetime,omitempty\" json:\"datetime,omitempty\"` + "`" + `",
                    "type": "string",
                    "example": "MASUK"
                },
                "latitude": {
                    "type": "number",
                    "example": 123.11
                },
                "location": {
                    "type": "string",
                    "example": "Bandung"
                },
                "longitude": {
                    "type": "number",
                    "example": 123.11
                },
                "phone_number": {
                    "type": "string",
                    "example": "08123456789"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "pemrog3-ulbi.herokuapp.com",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "TES SWAG",
	Description:      "This is a sample swagger for Fiber",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
