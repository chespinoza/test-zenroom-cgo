package main

/*
#cgo CFLAGS: -I./include
#cgo LDFLAGS: -Llib -lzenroom
#include "zenroom.h"
*/
import "C"

func main() {
	script := `
json = cjson()
keys = json.decode(KEYS)

data = DATA
nonce=randombytes(32)

encrypt_session = exchange_session_x25519(
	  decode_b58(keys.device_secret_key),
	  decode_b58(keys.user_public_key))

enc = encrypt_norx(encrypt_session, nonce, data)

print (encode_b58(nonce) .. encode_b58(enc))`

	keys := `
{
	"user_public_key": "H2ZYZR6QjvZEEoKDBEMeC6rdQoQafPotfHniq9HxS7TQ",
	"device_secret_key": "6ssnMqWxotEfi4zsZmCv9qrtPpU7R1hDojtpfvAQoQKa"
  }  
`
	data := `{"val:1"}`

	C.zenroom_exec(C.CString(script), nil, C.CString(keys), C.CString(data), 1)
}
