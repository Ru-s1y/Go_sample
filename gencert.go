package main

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "crypto/x509/pkix"
    "encoding/pem"
    "math/big"
    "net"
    "os"
    "time"
)

func main() {
    max := new(big.Int).Lsh(big.NewInt(1), 128)
    serialNumber, _ := rand.Int(rand.Reader, max)
    subject := pkix.Name{
        Organization:		[]string{"Manning Publications Co."},
        OrganizationalUnit: []string{"Books"},
        CommonName:			"Go Web Programming",
    }

    // 証明書の構成を設定する構造体
    template := x509.Certificate{
        SerialNumber:	serialNumber,
        Subject:			subject,
        NotBefore:		time.Now(),
        NotAfter:		time.Now().Add(365 * 24 * time.Hour),   // 有効期限 1年
        KeyUsage:		x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,   // X.509証明書をサーバ認証で使用する
        ExtKeyUsage:	[]x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},     // X.509証明書をサーバ認証で使用する
        IPAddresses:	[]net.IP{net.ParseIP("127.0.0.1")}, // 証明書が効力を持つアドレス
    }

    // RSAの秘密鍵を生成
    pk, _ := rsa.GenerateKey(rand.Reader, 2048)

    // DER形式のバイトデータのスライスを生成
    derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)
    // 証明書データを符号化してcert.pemを作成
    certOut, _ := os.Create("cert.pem")
    pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
    certOut.Close()

    // 生成鍵をPEM符号化してkey.pemに保存
    keyOut, _ := os.Create("key.pem")
    pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
    keyOut.Close()
}