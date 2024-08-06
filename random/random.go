package random

import (
    crypto_rand "crypto/rand"
    "encoding/binary"
    "math/big"
    "math/rand"
    "time"
)

var (
    Digits                = "0123456789"                                                         // Digits: [0-9]
    ASCIILettersLowercase = "abcdefghijklmnopqrstuvwxyz"                                         // Asci Lowerrcase Letters: [a-z]
    ASCIILettersUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"                                         // Ascii Uppercase Letters: [A-Z]
    Letters               = ASCIILettersLowercase + ASCIILettersUppercase                        // Ascii Letters: [a-zA-Z]
    ASCIICharacters       = ASCIILettersLowercase + ASCIILettersUppercase + Digits               // Ascii Charaters: [a-zA-Z0-9]
    Hexdigits             = "0123456789abcdefABCDEF"                                             // Hex Digits: [0-9a-fA-F]
    Octdigits             = "01234567"                                                           // Octal Digits: [0-7]
    Punctuation           = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"                                 // Punctuation and special characters
    Printables            = Digits + ASCIILettersLowercase + ASCIILettersUppercase + Punctuation // Printables
)

// getInt generates a cryptographically-secure random Int.
// Provided max can't be <= 0.
func getInt(max int) int {
    if max <= 0 {
        return 0
    }
    nbig, err := crypto_rand.Int(crypto_rand.Reader, big.NewInt(int64(max)))
    if err != nil {
        return GetIntInsecure(max)
    }
    n := int(nbig.Int64())
    
    return n
}
func init() {
    var rb [4]byte
    _, err := crypto_rand.Read(rb[:])
    seed := time.Now().UnixNano()
    if err == nil {
        seed += int64(binary.LittleEndian.Uint32(rb[:]))
    }
    // rand.Seed(seed)
    insecureRand = rand.New(rand.NewSource(seed))
}

var insecureRand *rand.Rand

//  generate a random integer using a seed of current system time.
func GetIntInsecure(i int) int {
    return insecureRand.Intn(i)
}

// String generates a cryptographically secure string.
func String(n int) string {
    return Random(n, ASCIICharacters)
}



// StringRange generates a secure random string within the given range.
func StringRange(min int, max int) string {
    i := IntRange(min, max)
    return String(i)
}

// IntRange returns a random integer between a given range.
func IntRange(min int, max int) int {
    i := getInt(max - min)
    i += min
    return i
}
func Random(n int, charset string) string {
    var charsetByte = []byte(charset)
    
    s := make([]byte, n)
    
    var mrange int
    for i := range s {
        mrange = getInt(len(charset))
        
        s[i] = charsetByte[mrange]
    }
    
    return string(s)
}


// Bytes generates a cryptographically secure set of bytes.
func Bytes(n int) ([]byte, error) {
    b := make([]byte, n)
    _, err := crypto_rand.Read(b)
    if err != nil {
        return []byte{}, err
    }
    
    return b, nil
}

// Choice makes a random choice from a slice of string.
func Choice(j []string) (string, error) {
    i := getInt(len(j))
    
    return j[i], nil
}
