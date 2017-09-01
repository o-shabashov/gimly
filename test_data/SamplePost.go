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

