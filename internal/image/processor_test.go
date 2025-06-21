package imageproc

import "testing"

func TestAllowedName_Matching(t *testing.T) {
	valid := []string{
		"v1_001_img.jpg",
		"v5_0.5_some.png",
		"v9_0.5_Art_abc.jpeg",
		"v10_2.3_xyz.webp",
	}

	invalid := []string{
		"cover.jpg",
		"v_001_image.jpg",
		"v1_1_img.jpg",   // не 3 цифры и не дробное
		"v1_abc_img.jpg", // не цифры
		"v10_03_xyz.png", // 2 цифры вместо 3
		"chapter1.png",
	}

	for _, name := range valid {
		if !allowedName.MatchString(name) {
			t.Errorf("Expected valid: %s", name)
		}
	}

	for _, name := range invalid {
		if allowedName.MatchString(name) {
			t.Errorf("Expected invalid: %s", name)
		}
	}
}
