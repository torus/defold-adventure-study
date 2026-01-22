components {
  id: "player"
  component: "/main/player.script"
}
embedded_components {
  id: "model"
  type: "model"
  data: "mesh: \"/assets/KayKit_Adventurers_2.0_FREE/Characters/gltf/Rogue.glb\"\n"
  "skeleton: \"/assets/KayKit_Adventurers_2.0_FREE/Characters/gltf/Rogue.glb\"\n"
  "animations: \"/assets/KayKit_Adventurers_2.0_FREE/Animations/gltf/Rig_Medium/Rig_Medium_MovementBasic.glb\"\n"
  "default_animation: \"Walking_C\"\n"
  "name: \"{{NAME}}\"\n"
  "materials {\n"
  "  name: \"rogue\"\n"
  "  material: \"/builtins/materials/model_skinned.material\"\n"
  "  textures {\n"
  "    sampler: \"tex0\"\n"
  "    texture: \"/assets/KayKit_Adventurers_2.0_FREE/Assets/gltf/rogue_texture.png\"\n"
  "  }\n"
  "}\n"
  ""
  position {
    z: 0.5
  }
  rotation {
    x: 0.70710677
    w: 0.70710677
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"player\"\n"
  "mask: \"blocker\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "  }\n"
  "  data: 0.5\n"
  "}\n"
  ""
}
