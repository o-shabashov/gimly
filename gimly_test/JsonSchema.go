package gimly_test

const JsonSchema = `
{
  "title": "Generator config",
  "type": "object",
  "properties": {
    "width": {
      "type": "integer"
    },
    "height": {
      "type": "integer"
    },
    "layers": {
      "type": "array",
      "uniqueItems": true,
      "minItems": 1,
      "items": {
        "$ref": "#/definitions/layer"
      }
    },
    "format": {
      "enum": [
        "jpeg",
        "png",
        "gif"
      ]
    },
    "out_path": {
      "type": "string"
    }
  },
  "required": [
    "width",
    "height",
    "layers",
    "format"
  ],
  "definitions": {
    "layer": {
      "type": "object",
      "properties": {
        "background_color": {
          "type": "string",
          "pattern": "^[a-zA-Z0-9]{6}$"
        },
        "background_path": {
          "type": "string"
        },
        "background_layout": {
          "enum": [
            "scale",
            "tile",
            "center"
          ]
        },
        "overlay_path": {
          "type": "string"
        },
        "type": {
          "enum": [
            "background",
            "foreground",
            "image"
          ]
        },
        "path": {
          "type": "string"
        },
        "position": {
          "type": "integer",
          "minimum": 0
        },
        "distortion_type": {
          "enum": [
            "bilinear",
            "affine",
            "polynomial"
          ]
        },
        "distortion_matrix": {
          "type": "array"
        },
        "distortion_order": {
          "enum": [
            1,
            1.5,
            2,
            3,
            4,
            5,
            "1",
            "1.5",
            "2",
            "3",
            "4",
            "5",
            null
          ]
        },
        "numb_points_side": {
          "type": "number"
        },
        "left": {
          "type": "number"
        },
        "top": {
          "type": "number"
        },
        "width": {
          "type": "number"
        },
        "height": {
          "type": "number"
        },
        "design_left": {
          "type": "number"
        },
        "design_top": {
          "type": "number"
        },
        "design_width": {
          "type": "number"
        },
        "design_height": {
          "type": "number"
        }
      },
      "required": [
        "type",
        "position",
        "left",
        "top",
        "width",
        "height",
        "design_left",
        "design_top",
        "design_width",
        "design_height"
      ]
    }
  }
}
`