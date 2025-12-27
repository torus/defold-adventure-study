components {
  id: "player"
  component: "/main/player.script"
}
embedded_components {
  id: "model"
  type: "model"
  data: "mesh: \"/builtins/assets/meshes/sphere.dae\"\n"
  "name: \"{{NAME}}\"\n"
  "materials {\n"
  "  name: \"default\"\n"
  "  material: \"/builtins/materials/model.material\"\n"
  "  textures {\n"
  "    sampler: \"tex0\"\n"
  "    texture: \"/builtins/assets/images/logo/logo_256.png\"\n"
  "  }\n"
  "}\n"
  ""
}
