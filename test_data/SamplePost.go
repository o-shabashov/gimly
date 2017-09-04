package test_data

const RequestPolynomialDistort = `
{
  "width": 700,
  "height": 700,
  "format": "jpeg",
  "layers": [
    {
      "path": "http:\/\/catalog.fm.vsemayki.ru\/414436646597807f42f48b3.27402821",
      "type": "image",
      "position": 0,
      "left": 0,
      "top": 0,
      "width": 100,
      "height": 100,
      "design_width": 100,
      "design_height": 100,
      "design_left": 0,
      "design_top": 0
    },
    {
      "top": 2.125,
      "left": -14,
      "type": "image",
      "width": 128.43669,
      "height": 91,
      "position": 1,
      "design_top": 1.78571,
      "design_left": 0.095837030974765,
      "design_width": 99.905645938047,
      "design_height": 96.97802,
      "distortion_type": "polynomial",
      "numb_points_side": 7,
      "distortion_matrix": [0.09732,1.78571,0,9.06593,16.74777,1.78571],
      "overlay_path": "http:\/\/catalog.fm.vsemayki.ru\/20283848485940a9c5b6b982.28126856",
      "path": "http:\/\/generator.fm.vsemayki.ru\/147673297059510edb058342.10982043"
    },
    {
      "path": "http:\/\/catalog.fm.vsemayki.ru\/633179955592e5c32f28729.10500083",
      "type": "image",
      "position": 2,
      "left": 0,
      "top": 0,
      "width": 100,
      "height": 100,
      "design_width": 100,
      "design_height": 100,
      "design_left": 0,
      "design_top": 0
    }
  ]
}
`

const BackgroundLayer = `
{
  "background_path": "http:\/\/catalog.fm.vsemayki.ru\/633179955592e5c32f28729.10500083",
  "type": "background",
  "background_layout": "tile",
  "position": 0,
  "left": 0,
  "top": 0,
  "width": 100,
  "height": 100,
  "design_width": 100,
  "design_height": 100,
  "design_left": 0,
  "design_top": 0
}
`

const MainLayer = `
{
  "top": 2.125,
  "left": -14,
  "type": "image",
  "width": 128.43669,
  "height": 91,
  "position": 1,
  "design_top": 1.78571,
  "design_left": 0.095837030974765,
  "design_width": 99.905645938047,
  "design_height": 96.97802,
  "distortion_type": "polynomial",
  "distortion_order": 0,
  "numb_points_side": 7,
  "distortion_matrix": [1,0.09732,1.78571,0,9.06593,16.74777,1.78571,0.09732,1.78571,0,9.06593,16.74777,1.78571],
  "overlay_path": "http:\/\/catalog.fm.vsemayki.ru\/20283848485940a9c5b6b982.28126856",
  "path": "http:\/\/generator.fm.vsemayki.ru\/147673297059510edb058342.10982043"
}
`

const OverlayLayer = `
{
  "top": 2.125,
  "left": -14,
  "type": "image",
  "width": 128.43669,
  "height": 91,
  "position": 1,
  "design_top": 1.78571,
  "design_left": 0.095837030974765,
  "design_width": 99.905645938047,
  "design_height": 96.97802,
  "distortion_type": "polynomial",
  "distortion_order": 0,
  "numb_points_side": 7,
  "distortion_matrix": [0.09732,1.78571,0,9.06593,16.74777,1.78571,16.74777,1.78571],
  "overlay_path": "http:\/\/catalog.fm.vsemayki.ru\/20283848485940a9c5b6b982.28126856",
  "overlay_width": 100,
  "overlay_height": 100,
  "overlay_left": 10,
  "overlay_top": 10
}
`

const BuildLayer = `
{
  "background_path": "http:\/\/catalog.fm.vsemayki.ru\/633179955592e5c32f28729.10500083",
  "background_layout": "tile",
  "top": 2.125,
  "left": -14,
  "type": "image",
  "width": 128.43669,
  "height": 91,
  "position": 0,
  "design_top": 1.78571,
  "design_left": 0.095837030974765,
  "design_width": 99.905645938047,
  "design_height": 96.97802,
  "distortion_type": "polynomial",
  "distortion_order": 0,
  "numb_points_side": 7,
  "distortion_matrix": [1,0.09732,1.78571,0,9.06593,16.74777,1.78571,0.09732,1.78571,0,9.06593,16.74777,1.78571],
  "overlay_path": "http:\/\/catalog.fm.vsemayki.ru\/20283848485940a9c5b6b982.28126856",
  "path": "http:\/\/generator.fm.vsemayki.ru\/147673297059510edb058342.10982043",
  "overlay_width": 100,
  "overlay_height": 100,
  "overlay_left": 10,
  "overlay_top": 10
}
`

const PartialDistortLayer = `
{
    "top": 5.125,
    "left": 11.125,
    "type": "image",
    "width": 88.65098,
    "height": 91.625,
    "position": 2,
    "design_top": 39.276946,
    "design_left": 0,
    "design_width": 100,
    "design_height": 53.23632,
    "distortion_type": "partial",
    "distortion_order": null,
    "numb_points_side": 2,
    "distortion_matrix": [0,33.28786,15.08726,24.96589,15.93327,33.28786,16.77928,30.28649,19.31733,33.28786,19.31733,
    33.28786,22.70138,33.28786,22.56038,33.96999,26.08544,33.28786,26.08544,33.28786,40.04468,33.28786,40.04468,
    28.10368,54.14492,33.28786,54.56792,21.2824,66.55313,33.28786,66.27112,14.05184,73.18024,33.28786,72.33423,8.73124,
    100,33.28786,75.01327,0,0,99.83326,15.08726,91.5116,15.93327,99.83326,18.47131,96.31651,19.31733,99.83326,21.00936,
    99.31787,22.70138,99.83326,24.25241,100,26.08544,99.83326,27.77747,99.31787,40.04468,99.83326,40.04468,94.64939,
    54.14492,99.83326,52.31189,87.72169,66.55313,99.83326,58.939,82.81037,73.18024,99.83326,65.98912,77.86903,100,
    99.83326,75.01327,66.5457],
    "path": "http://generator.fm.vsemayki.ru/699546773594538f0dd7144.33205872"
    }
`
