package main

import (
    "fmt"
    "io"

    "crypto/sha1"
    "crypto/sha256"
    "crypto/sha512"
)

func main() {
    /* SHA-1 */
    s1 := sha1.New()
    io.WriteString(s1, "ABCDE")
    fmt.Printf("%x\n", s1.Sum(nil))
    // => "7be07aaf460d593a323d0db33da05b64bfdcb3a5"

    /* SHA-256 */
    s256 := sha256.New()
    io.WriteString(s256, "ABCDE")
    fmt.Printf("%x\n", s256.Sum(nil))
    // => "f0393febe8baaa55e32f7be2a7cc180bf34e52137d99e056c817a9c07b8f239a"

    /* SHA-512 */
    s512 := sha512.New()
    io.WriteString(s512, "ABCDE")
    fmt.Printf("%x\n", s512.Sum(nil))
    // => "9989a8fcbc29044b5883a0a36c146fe7415b1439e995b4d806ea0af7da9ca4390eb92@<dtp>{lb}a604b3ecfa3d75f9911c768fbe2aecc59eff1e48dcaeca1957bdde01dfb"
}
