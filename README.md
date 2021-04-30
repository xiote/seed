# seed

`seed` is package that implement SEED encryption for Go. 

It's based on [https://github.com/geeksbaek/seed](https://github.com/geeksbaek/seed)

## What is SEED

From [base project](https://github.com/geeksbaek/seed)

>SEED is a block cipher developed by the Korea Internet & Security Agency (KISA). It is used broadly throughout South Korean industry, but seldom found elsewhere. It gained popularity in Korea because 40-bit encryption was not considered strong enough, so the Korea Information Security Agency developed its own standard. However, this decision has historically limited the competition of web browsers in Korea, as no major SSL libraries or web browsers supported the SEED algorithm, requiring users to use an ActiveX control in Internet Explorer for secure web sites.
>
>On April 1, 2015 the Ministry of Science, ICT and Future Planning (MSIP) announced its plan to remove the ActiveX dependency from at least 90 percent of the country's top 100 websites by 2017. Instead, HTML5-based technologies will be employed as they operate on many platforms, including mobile devices. Starting with the private sector, the ministry plans to expand this further to ultimately remove this dependency from public websites as well.
>
>[Read more from Wikipedia](https://en.wikipedia.org/wiki/SEED)

## Disclaimer

Currently, only 128-bit encryption is supported.

## Usage

Base project is cool enough but a little hard to use. So i made it a little more convenient to use.

```go
package main

import (
	"github.com/zajann/seed"
)

func main() {
    // You can use key to string, lenght SHOULD be 16
    key := "thisiskeyfortest"
    
    // Init seed with key
    if err := seed.InitECB(key); err != nil {
        panic(err)
    }
    
    plainText := "A test plain text"
    
    // Encrypte
    encrypted, err := seed.ECBEncryptAll([]byte(plainText))
    if err != nil {
        panic(err)
    }
    
    // Decrypte
    decrypted, err := seed.ECBDecryptAll(encrypted)
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("PlanText: %s\n", plainText)
    fmt.Printf("Encrypted: %v\n", encrypted)
    fmt.Printf("Decrypted: %s\n", string(decrypted))
    
    // Output:
    // PlanText: A test plain text
	// Encrypted: [50 125 186 158 176 188 101 173 157 70 48 202 210 245 109 152 67 196 4 121 225 244 141 128 220 135 147 116 218 226 156 151]
	// Decrypted: A test plain text
}
```

Once you initialize `seed` with `Init(key sring)`, you can use seed globally.

## License

[MIT](https://github.com/zajann/seed/blob/master/LICENSE)