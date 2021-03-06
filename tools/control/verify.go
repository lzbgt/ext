package control

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/openpgp"

	"v2ray.com/core/common"
)

const (
	pubkey = `-----BEGIN PGP PUBLIC KEY BLOCK-----
Comment: GPGTools - https://gpgtools.org

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
80yAOiqyRyAtruVaY72n
=CVru
-----END PGP PUBLIC KEY BLOCK-----
`
)

func firstIdentity(m map[string]*openpgp.Identity) string {
	for k := range m {
		return k
	}
	return ""
}

func init() {
	const name = "verify"
	common.Must(RegisterCommand(name, "Verify if a binary is officially signed.", func(args []string) {
		fs := flag.NewFlagSet(name, flag.ContinueOnError)

		sigFile := fs.String("sig", "", "Path to the signature file")

		err := fs.Parse(args)
		switch err {
		case nil:
		case flag.ErrHelp:
			fmt.Println("v2ctl verify [--sig=<sig-file>] file")
			fmt.Println("Verify the file officially signed by V2Ray.")
		default:
			fmt.Fprintln(os.Stderr, "Error parsing arguments:", err)
			return
		}

		target := fs.Arg(0)
		if len(target) == 0 {
			fmt.Fprintln(os.Stderr, "Empty file path.")
			return
		}

		if len(*sigFile) == 0 {
			*sigFile = target + ".sig"
		}

		targetReader, err := os.Open(os.ExpandEnv(target))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error opening file (", target, "):", err)
			return
		}

		sigReader, err := os.Open(os.ExpandEnv(*sigFile))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error opening file (", *sigFile, "): ", err)
			return
		}

		keyring, err := openpgp.ReadArmoredKeyRing(strings.NewReader(pubkey))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error creating keyring:", err)
			return
		}

		entity, err := openpgp.CheckDetachedSignature(keyring, targetReader, sigReader)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error verifying signature:", err)
			return
		}

		fmt.Println("Signed by:", firstIdentity(entity.Identities))
	}))
}
