package main

/*
#cgo CFLAGS: -I./include
#cgo LDFLAGS: -Llib -lzenroom
#include "zenroom.h"
*/
import "C"

func main() {
	script := `
	-- generate a private keyring and other fictional public keys

	-- run with: zenroom keygen.zen
	
	-- any combination of public and private keys generated this way and
	-- exchanged among different people will lead to the same secret which
	-- is then usable for asymmetric encryption.
	
	octet = require'octet'
	json = require'json'
	ecdh = require'ecdh'
	keyring = ecdh.new()
	keyring:keygen()
	
	recipients={'jaromil','francesca','jim','mark','paulus','mayo'}
	keys={}
	for i,name in ipairs(recipients) do
	   kk = ecdh.new()
	   kk:keygen()
	   keys[name] = kk:public():base64()
	   assert(kk:checkpub(kk:public()))
	end
	
	
	keypairs = json.encode({
		  keyring={public=keyring:public():base64(),
				   secret=keyring:private():base64()},
		  recipients=keys})
	print(keypairs)
	
	
	`

	keys := `
{
	"user_public_key": "H2ZYZR6QjvZEEoKDBEMeC6rdQoQafPotfHniq9HxS7TQ",
	"device_secret_key": "6ssnMqWxotEfi4zsZmCv9qrtPpU7R1hDojtpfvAQoQKa"
  }  
`
	data := `{"val:1"}`

	C.zenroom_exec(C.CString(script), nil, C.CString(keys), C.CString(data), 1)
}
