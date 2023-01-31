package block

import (
	"bytes"
	"encoding/hex"
	"github.com/nknorg/nkn/v2/common"
	"reflect"
	"testing"
)

var rawHeader = "0ade0208011220da29b937682f9f84d0254d38002ee4a95abf587d4bec65d97db5a6612ad9328c1a20bb46b8e72e98a888d39c43bbc68fb79c2cc26bc83f6cf11ace1fe0e0380fc7cf2220e23f927c7258afc5619f2efbbfda1ba5571232f9c11008501d39c000bd1ec0b728dbb6dd9e0630b680b0023a8001e2a3f566880d7f8ab3a77248835f4fe04cff1f5471a61b585f199d949c790f390ba1923b4ad570232bec4f4ec866b86905738e6473137f1baf4ef21a4b7e1c02ff4a7376a497d786cf47a370523ba6bade4583bb9f45c5beec543a50f236cd09c0702c4419f8e4083f560205e06682d2b1052345e6b71ea56eb7bc6e1ac2e56b42201499620fadd18125056760f5498a481fffae45c2431857f8a127a685954edb144801522088811c3a1ebf787355a08d2ebcba6fd86ef2a19233350bc2fb9c13ac39fddbf85a20e14812772c280e2c4f4b17505aa13a3724c170f976e530a77f5e99d0d1de54a01240af62957879d75a765b96f495fdb7dcb74f4c8d13104ba58c34339ff25f1a8ac2297a5ba9d272f3c6c20ffd1a08eeef7d6265f8a8ff15f636bb5df3fab82d8f01"

func TestHeaderUnsignedSerialization(t *testing.T) {
	check := func(f string, got, want interface{}) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s mismatch: got %v, want %v", f, got, want)
		}
	}

	headerBytes, err := hex.DecodeString(rawHeader)
	if err != nil {
		t.Fatal("decode header error:", err)
	}

	h := new(Header)
	err = h.Unmarshal(headerBytes)
	if err != nil {
		t.Fatal("unmarshal header error:", err)
	}

	buf := bytes.NewBuffer(nil)
	err = h.SerializeUnsigned(buf)
	if err != nil {
		return
	}

	check("SerializeUnsigned", hex.EncodeToString(buf.Bytes()), "0100000020da29b937682f9f84d0254d38002ee4a95abf587d4bec65d97db5a6612ad9328c20bb46b8e72e98a888d39c43bbc68fb79c2cc26bc83f6cf11ace1fe0e0380fc7cf20e23f927c7258afc5619f2efbbfda1ba5571232f9c11008501d39c000bd1ec0b75b5bd7630000000036004c00010000002088811c3a1ebf787355a08d2ebcba6fd86ef2a19233350bc2fb9c13ac39fddbf820e14812772c280e2c4f4b17505aa13a3724c170f976e530a77f5e99d0d1de54a0")
	pk, id, err := h.GetSigner()
	if err != nil {
		t.Fatal("get signer error:", err)
	}
	check("GetSigner", hex.EncodeToString(pk), "88811c3a1ebf787355a08d2ebcba6fd86ef2a19233350bc2fb9c13ac39fddbf8")
	check("GetSigner", hex.EncodeToString(id), "e14812772c280e2c4f4b17505aa13a3724c170f976e530a77f5e99d0d1de54a0")
	info, err := h.GetInfo()
	if err != nil {
		t.Fatal("get info error:", err)
	}
	check("GetInfo", hex.EncodeToString(info), "7b2276657273696f6e223a312c2270726576426c6f636b48617368223a2264613239623933373638326639663834643032353464333830303265653461393561626635383764346265633635643937646235613636313261643933323863222c227472616e73616374696f6e73526f6f74223a2262623436623865373265393861383838643339633433626263363866623739633263633236626338336636636631316163653166653065303338306663376366222c227374617465526f6f74223a2265323366393237633732353861666335363139663265666262666461316261353537313233326639633131303038353031643339633030306264316563306237222c2274696d657374616d70223a313637353035383031312c22686569676874223a343938303739302c2272616e646f6d426561636f6e223a2265326133663536363838306437663861623361373732343838333566346665303463666631663534373161363162353835663139396439343963373930663339306261313932336234616435373032333262656334663465633836366238363930353733386536343733313337663162616634656632316134623765316330326666346137333736613439376437383663663437613337303532336261366261646534353833626239663435633562656563353433613530663233366364303963303730326334343139663865343038336635363032303565303636383264326231303532333435653662373165613536656237626336653161633265353662222c2277696e6e657248617368223a2231343939363230666164643138313235303536373630663534393861343831666666616534356332343331383537663861313237613638353935346564623134222c2277696e6e657254797065223a2254584e5f5349474e4552222c227369676e6572506b223a2238383831316333613165626637383733353561303864326562636261366664383665663261313932333333353062633266623963313361633339666464626638222c227369676e65724964223a2265313438313237373263323830653263346634623137353035616131336133373234633137306639373665353330613737663565393964306431646535346130222c227369676e6174757265223a226166363239353738373964373561373635623936663439356664623764636237346634633864313331303462613538633334333339666632356631613861633232393761356261396432373266336336633230666664316130386565656637643632363566386138666631356636333662623564663366616238326438663031222c2268617368223a2239386332363764393836326435306532646133663764363233666330626663356532366565373837643237616662643339373436653539363134373936656361227d")
	check("GetMessage", hex.EncodeToString(h.GetMessage()), "0100000020da29b937682f9f84d0254d38002ee4a95abf587d4bec65d97db5a6612ad9328c20bb46b8e72e98a888d39c43bbc68fb79c2cc26bc83f6cf11ace1fe0e0380fc7cf20e23f927c7258afc5619f2efbbfda1ba5571232f9c11008501d39c000bd1ec0b75b5bd7630000000036004c00010000002088811c3a1ebf787355a08d2ebcba6fd86ef2a19233350bc2fb9c13ac39fddbf820e14812772c280e2c4f4b17505aa13a3724c170f976e530a77f5e99d0d1de54a0")
	check("GetSignature", hex.EncodeToString(h.GetSignature()), "af62957879d75a765b96f495fdb7dcb74f4c8d13104ba58c34339ff25f1a8ac2297a5ba9d272f3c6c20ffd1a08eeef7d6265f8a8ff15f636bb5df3fab82d8f01")
	progHashes, err := h.GetProgramHashes()
	if err != nil {
		t.Fatal("get program hashes error:", err)
	}
	check("GetProgramHashes", progHashes, []common.Uint160{{241, 191, 91, 34, 190, 204, 94, 229, 157, 120, 219, 141, 174, 113, 123, 63, 114, 163, 224, 1}})
}
