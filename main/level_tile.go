components {
  id: "level_tile"
  component: "/main/level_tile.script"
}
embedded_components {
  id: "model"
  type: "model"
  data: "mesh: \"/assets/KayKit_Medieval_Hexagon_Pack_1.0_EXTRA/Assets/gltf/tiles/base/hex_grass.gltf\"\n"
  "name: \"{{NAME}}\"\n"
  "materials {\n"
  "  name: \"hexagons_medieval\"\n"
  "  material: \"/builtins/materials/model.material\"\n"
  "  textures {\n"
  "    sampler: \"tex0\"\n"
  "    texture: \"/assets/KayKit_Medieval_Hexagon_Pack_1.0_EXTRA/Assets/gltf/buildings/neutral/hexagons_medieval.png\"\n"
  "  }\n"
  "}\n"
  ""
  rotation {
    x: 0.70710677
    w: 0.70710677
  }
}
