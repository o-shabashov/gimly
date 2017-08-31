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

const RequestPartialDistort = `
{
  "width": 500,
  "height": 500,
  "format": "jpeg",
  "layers": [
    {
      "path": "https:\/\/catalog-fm.vsemayki.ru\/106579268557ac6ec461cfc8.83182403",
      "type": "image",
      "position": 0,
      "left": 0,
      "top": 0,
      "width": 500,
      "height": 500,
      "design_width": 500,
      "design_height": 500,
      "design_left": 0,
      "design_top": 0,
      "overlay_width": 500,
      "overlay_height": 500,
      "overlay_left": 0,
      "overlay_top": 0
    }, {
      "top": 25.625,
      "left": 55.625,
      "type": "image",
      "width": 443.2549,
      "height": 458.125,
      "position": 2,
      "design_top": 179.9375088625,
      "design_left": 0,
      "design_width": 443.2549,
      "design_height": 243.888891,
      "distortion_type": "partial",
      "distortion_order": null,
      "numb_points_side": 2,
      "distortion_matrix": [0, 152.500008625, 66.87501922574, 114.3749835625, 70.62500000523, 152.500008625, 74.37498078472, 138.7499823125, 85.62501177417, 152.500008625, 85.62501177417, 152.500008625, 100.62497921762, 152.500008625, 99.99998980862, 155.6250166875, 115.62499098656, 152.500008625, 115.62499098656, 152.500008625, 177.50000628932, 152.500008625, 177.50000628932, 128.749984, 240.00001100108, 152.500008625, 241.87497922808, 97.499995, 295.00000982837, 152.500008625, 293.74998668488, 64.374992, 324.37499963176, 152.500008625, 320.62501885227, 39.99999325, 443.2549, 152.500008625, 332.49999492523, 0, 0, 457.361122375, 66.87501922574, 419.2375175, 70.62500000523, 457.361122375, 81.87498666919, 441.2500114375, 85.62501177417, 457.361122375, 93.12501765864, 454.9999919375, 100.62497921762, 457.361122375, 107.49999569309, 458.125, 115.62499098656, 457.361122375, 123.12499687103, 454.9999919375, 177.50000628932, 457.361122375, 177.50000628932, 433.6125179375, 240.00001100108, 457.361122375, 231.87501570761, 401.8749923125, 295.00000982837, 457.361122375, 261.250005511, 379.3750075625, 324.37499963176, 457.361122375, 292.50000786688, 356.7374936875, 443.2549, 457.361122375, 332.49999492523, 304.862488125],
      "path": "http:\/\/generator.fm.vsemayki.ru\/699546773594538f0dd7144.33205872",
      "overlay_width": 500,
      "overlay_height": 500,
      "overlay_left": -55.625,
      "overlay_top": -25.625
    }, {
      "path": "https:\/\/catalog-fm.vsemayki.ru\/96406336557ac6ec5aa4b14.88194089",
      "type": "image",
      "position": 3,
      "left": 0,
      "top": 0,
      "width": 500,
      "height": 500,
      "design_width": 500,
      "design_height": 500,
      "design_left": 0,
      "design_top": 0,
      "overlay_width": 500,
      "overlay_height": 500,
      "overlay_left": 0,
      "overlay_top": 0
    }
  ]
}
`
