package main

import (
	"testing"
)

func TestParseDecodedJson(t *testing.T) {
	in := `{"data":
    [{
      "common_name":"リドカイン塩酸塩・アドレナリン(2)注射液",
      "company_name":"アストラゼネカ",
      "medicine_id":"9eef79a8b59fa30aaee2dcb36a0ddda1",
      "pdfIdList":"[\"00000126\"]",
      "product_name":"キシロカイン注射液2%"
    }]
}`
	j := parseDecodedJson(in)
	if "9eef79a8b59fa30aaee2dcb36a0ddda1" != j.MedicineDocuments[0].MedicineId {
		t.Errorf("TEST NG")
	}
}

func TestEvalStringToList(t *testing.T) {
	in := `["012","234"]`
	out := evalStringToList(in)
	if (out[0] != "012") {
		t.Errorf("TEST NG")
	}

	if (out[1] != "234") {
		t.Errorf("TEST NG")
	}

}

func TestFetchPdf(t *testing.T) {
	err := fetchPdf("00000126")
	if err != nil {
		t.Errorf("TEST NG")
	}
}
