package mimetype_ext

import "testing"

func TestGetExtension(t *testing.T) {
	ext, _ := GetExtension("image/png")

	expect := "png"

	if ext[0] != expect {
		t.Errorf("got: %v, want: %v", ext[0], expect)
	}
}

func TestGetMimetype(t *testing.T) {
	ext, _ := GetMimetype("png")

	expect := "image/png"

	if ext[0] != expect {
		t.Errorf("got: %v, want: %v", ext[0], expect)
	}

	ext, _ = GetMimetype(".png")

	if ext[0] != expect {
		t.Errorf("got: %v, want: %v", ext[0], expect)
	}
}
