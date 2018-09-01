// Copyright (C) 2018 BitWhite Team <info@bitwhite.org>
//
// You can use this code in accordance with the GNU General Public License v3.0
// which can be found in the LICENSE file.
//
// Please note that you can use the source code for your own purposes,
// but we do not give any warranty. For more information, refer to the GPLv3.

package json

import "testing"

type TestStruct struct {
  FirstField         int
  SecondBeepBupField string `json:"SecondField"`
  hiddenField        int
}

var obj = TestStruct{
  FirstField:         15,
  SecondBeepBupField: "Beep-Bup",
  hiddenField:        3,
}

var jsn = `{"FirstField":15,"SecondField":"Beep-Bup"}`

func TestToJson(t *testing.T) {
  j := ToJson(obj)

  if !equals(j, []byte(jsn)) {
    t.Fatalf("ToJson incorrect, got: %s, want: %s", string(j), jsn)
  }
}

func TestFromJson(t *testing.T) {
  o := TestStruct{}
  err := FromJson([]byte(jsn), &o)
  if err != nil {
    t.Fatal(err.Error())
  }

  if o.FirstField != obj.FirstField ||
    o.SecondBeepBupField != obj.SecondBeepBupField ||
    o.hiddenField == obj.hiddenField {

    t.Fatal("FromJson incorrect, got:", o, "want:", obj)
  }
}

func equals(a []byte, b []byte) bool {
  if len(a) != len(b) {
    return false
  }

  for i := 0; i < len(a); i++ {
    if a[i] != b[i] {
      return false
    }
  }

  return true
}
