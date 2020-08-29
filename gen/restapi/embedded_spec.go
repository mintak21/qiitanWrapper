// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "QiitaAPIのラッパーAPIです。",
    "title": "mintak's Qiita Wrapper API",
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/items/sync/{tag}": {
      "get": {
        "description": "指定日付に投稿された記事を一覧で取得。一度に取得できる記事の数は100固定。",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "items"
        ],
        "summary": "同期的にタグの記事を取得。",
        "operationId": "syncTagItems",
        "parameters": [
          {
            "type": "string",
            "description": "取得する記事のタグ",
            "name": "tag",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "format": "date",
            "description": "取得対象の日付",
            "name": "date",
            "in": "query"
          },
          {
            "minimum": 1,
            "type": "integer",
            "default": 1,
            "description": "取得するページ",
            "name": "page",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Items"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/items/trends": {
      "get": {
        "description": "指定月に投稿された記事のうち、ストック数の多い順に100個の記事を取得。",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "items"
        ],
        "summary": "指定月ストック数の多い記事を取得。",
        "operationId": "getMonthlyTrendItems",
        "parameters": [
          {
            "pattern": "[0-9][0-9][0-9][0-9]-[0-9][0-9]",
            "type": "string",
            "description": "対象日付（日付形式で指定し、その月を対象とする）",
            "name": "target_month",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Items"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/items/{tag}": {
      "get": {
        "description": "指定したタグの記事の取得。1度に取得する件数は100固定。",
        "produces": [
          "application/json"
        ],
        "tags": [
          "items"
        ],
        "summary": "指定したタグの記事の取得",
        "operationId": "getTagItems",
        "parameters": [
          {
            "type": "string",
            "description": "取得する記事のタグ",
            "name": "tag",
            "in": "path",
            "required": true
          },
          {
            "minimum": 1,
            "type": "integer",
            "default": 1,
            "description": "取得するページ",
            "name": "page",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Items"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "description": "Error Status model",
      "type": "object",
      "title": "Error",
      "properties": {
        "message": {
          "type": "string",
          "example": "Bad Request"
        }
      },
      "x-examples": {}
    },
    "Item": {
      "type": "object",
      "properties": {
        "created_at": {
          "description": "投稿日時",
          "type": "string",
          "format": "date-time",
          "example": "20200707T124012+09:00"
        },
        "link": {
          "description": "記事のリンクURL",
          "type": "string",
          "example": "https://example.com/test"
        },
        "statistics": {
          "$ref": "#/definitions/Statistics"
        },
        "table_contents": {
          "description": "目次",
          "type": "string",
          "example": "\u003ch1概要\u003c/h1\u003e\u003ch2\u003eその１\u003c/h2\u003e\u003ch2\u003eその２\u003c/h2\u003e"
        },
        "tags": {
          "type": "array",
          "items": {
            "description": "つけられているタグ一覧",
            "type": "string",
            "example": [
              "Terraform",
              "AWS",
              "Lambda"
            ]
          }
        },
        "title": {
          "description": "記事のタイトル",
          "type": "string",
          "example": "Terraformを始める"
        },
        "user": {
          "$ref": "#/definitions/User"
        }
      }
    },
    "Items": {
      "type": "object",
      "properties": {
        "has_next": {
          "description": "次のページがあるかどうか",
          "type": "boolean",
          "x-omitempty": false
        },
        "items": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Item"
          }
        },
        "page": {
          "description": "現在のページ",
          "type": "number",
          "format": "int64",
          "x-omitempty": false,
          "example": 1
        }
      }
    },
    "Statistics": {
      "type": "object",
      "properties": {
        "lgtms": {
          "description": "LGTM数",
          "type": "number",
          "format": "int",
          "x-omitempty": false,
          "example": 2
        },
        "stocks": {
          "description": "ストック数（ストックしているユーザー数）",
          "type": "number",
          "format": "int",
          "example": 1
        }
      }
    },
    "User": {
      "type": "object",
      "properties": {
        "name": {
          "description": "ユーザー名称",
          "type": "string",
          "example": "mintak21"
        },
        "thumbnail_link": {
          "description": "サムネイル画像リンク",
          "type": "string",
          "example": "https://example.com/test.jpg"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Qiitaの記事を操作",
      "name": "items"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "QiitaAPIのラッパーAPIです。",
    "title": "mintak's Qiita Wrapper API",
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/items/sync/{tag}": {
      "get": {
        "description": "指定日付に投稿された記事を一覧で取得。一度に取得できる記事の数は100固定。",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "items"
        ],
        "summary": "同期的にタグの記事を取得。",
        "operationId": "syncTagItems",
        "parameters": [
          {
            "type": "string",
            "description": "取得する記事のタグ",
            "name": "tag",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "format": "date",
            "description": "取得対象の日付",
            "name": "date",
            "in": "query"
          },
          {
            "minimum": 1,
            "type": "integer",
            "default": 1,
            "description": "取得するページ",
            "name": "page",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Items"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/items/trends": {
      "get": {
        "description": "指定月に投稿された記事のうち、ストック数の多い順に100個の記事を取得。",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "items"
        ],
        "summary": "指定月ストック数の多い記事を取得。",
        "operationId": "getMonthlyTrendItems",
        "parameters": [
          {
            "pattern": "[0-9][0-9][0-9][0-9]-[0-9][0-9]",
            "type": "string",
            "description": "対象日付（日付形式で指定し、その月を対象とする）",
            "name": "target_month",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Items"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/items/{tag}": {
      "get": {
        "description": "指定したタグの記事の取得。1度に取得する件数は100固定。",
        "produces": [
          "application/json"
        ],
        "tags": [
          "items"
        ],
        "summary": "指定したタグの記事の取得",
        "operationId": "getTagItems",
        "parameters": [
          {
            "type": "string",
            "description": "取得する記事のタグ",
            "name": "tag",
            "in": "path",
            "required": true
          },
          {
            "minimum": 1,
            "type": "integer",
            "default": 1,
            "description": "取得するページ",
            "name": "page",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Items"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "description": "Error Status model",
      "type": "object",
      "title": "Error",
      "properties": {
        "message": {
          "type": "string",
          "example": "Bad Request"
        }
      },
      "x-examples": {}
    },
    "Item": {
      "type": "object",
      "properties": {
        "created_at": {
          "description": "投稿日時",
          "type": "string",
          "format": "date-time",
          "example": "20200707T124012+09:00"
        },
        "link": {
          "description": "記事のリンクURL",
          "type": "string",
          "example": "https://example.com/test"
        },
        "statistics": {
          "$ref": "#/definitions/Statistics"
        },
        "table_contents": {
          "description": "目次",
          "type": "string",
          "example": "\u003ch1概要\u003c/h1\u003e\u003ch2\u003eその１\u003c/h2\u003e\u003ch2\u003eその２\u003c/h2\u003e"
        },
        "tags": {
          "type": "array",
          "items": {
            "description": "つけられているタグ一覧",
            "type": "string",
            "example": [
              "Terraform",
              "AWS",
              "Lambda"
            ]
          }
        },
        "title": {
          "description": "記事のタイトル",
          "type": "string",
          "example": "Terraformを始める"
        },
        "user": {
          "$ref": "#/definitions/User"
        }
      }
    },
    "Items": {
      "type": "object",
      "properties": {
        "has_next": {
          "description": "次のページがあるかどうか",
          "type": "boolean",
          "x-omitempty": false
        },
        "items": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Item"
          }
        },
        "page": {
          "description": "現在のページ",
          "type": "number",
          "format": "int64",
          "x-omitempty": false,
          "example": 1
        }
      }
    },
    "Statistics": {
      "type": "object",
      "properties": {
        "lgtms": {
          "description": "LGTM数",
          "type": "number",
          "format": "int",
          "x-omitempty": false,
          "example": 2
        },
        "stocks": {
          "description": "ストック数（ストックしているユーザー数）",
          "type": "number",
          "format": "int",
          "example": 1
        }
      }
    },
    "User": {
      "type": "object",
      "properties": {
        "name": {
          "description": "ユーザー名称",
          "type": "string",
          "example": "mintak21"
        },
        "thumbnail_link": {
          "description": "サムネイル画像リンク",
          "type": "string",
          "example": "https://example.com/test.jpg"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Qiitaの記事を操作",
      "name": "items"
    }
  ]
}`))
}
