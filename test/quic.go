package test

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"io"
	"log"
	"math/big"
	"time"

	quic "github.com/lucas-clemente/quic-go"
)

const (
	message = "-123- -123- -123- -123- -123- -123- -123- -123- -123- -123- -123- -123- -123- -123-"
	addr    = "localhost:4321"
)

func QuicTest() {
	go func() {
		log.Fatal(echoServer())
	}()

	err := clientMain()
	if err != nil {
		panic(err)
	}
}

func echoServer() error {
	listener, err := quic.ListenAddr(addr, generateTLSConfig(), nil)
	if err != nil {
		return err
	}
	sess, err := listener.Accept(context.Background())
	if err != nil {
		return err
	}
	time.Sleep(3 * time.Second)
	stream, err := sess.AcceptStream(context.Background())
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(loggingWriter{stream}, stream)
	return err
}

func clientMain() error {
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic-echo-example"},
	}
	session, err := quic.DialAddr(addr, tlsConf, nil)
	if err != nil {
		return err
	}
	stream, err := session.OpenStreamSync(context.Background())
	if err != nil {
		return err
	}
	buf := make([]byte, len(message))

	for i := 0; i < 10000; i++ {
		_, err = stream.Write([]byte(message))
		if err != nil {
			return err
		}

		_, err = io.ReadFull(stream, buf)
		if err != nil {
			return err
		}
	}

	return nil
}

type loggingWriter struct {
	io.Writer
}

func (w loggingWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func generateTLSConfig() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}

	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:   []string{"quic-echo-example"},
	}
}
