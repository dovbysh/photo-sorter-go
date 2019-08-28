package exif

// sudo apt-get install -y libexif-dev
import (
	"fmt"
	"github.com/xiam/exif"
	"testing"
)

func TestTags(t *testing.T) {
	data, err := exif.Read("../testdata/IMG_20190822_212629.jpg")
	if err != nil {
		t.Fatal(err)
	}
	for key, val := range data.Tags {
		fmt.Printf("%s = %s\n", key, val)
	}

}
