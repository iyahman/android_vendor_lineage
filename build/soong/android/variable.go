package android
type Product_variables struct {
	Needs_text_relocations struct {
		Cppflags []string
	}
	Uses_generic_camera_parameter_library struct {
		Srcs []string
	}
	Uses_non_treble_camera struct {
		Cflags []string
	}
}

type ProductVariables struct {
	Uses_generic_camera_parameter_library  *bool `json:",omitempty"`
	Specific_camera_parameter_library  *string `json:",omitempty"`
	Needs_text_relocations  *bool `json:",omitempty"`
        Uses_non_treble_camera  *bool `json:",omitempty"`
}
