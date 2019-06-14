package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

const READ_MEGA_BYTE = 1024 * 1024
const READ_UNIT = 2

func usedSHA512() {
	fmt.Println(time.Now(), "sha512 시작")
	bytes, err := ioutil.ReadFile("../test-file/test.dcf")
	if err != nil {
		fmt.Println(err)
		return
	}
	h1 := sha512.Sum512(bytes)

	fmt.Println(h1, time.Now(), "sha512 해싱 완료")
}

func usedSHA256() {
	fmt.Println(time.Now(), "sha256 시작")
	bytes, err := ioutil.ReadFile("../test-file/test.dcf")
	if err != nil {
		fmt.Println(err)
		return
	}
	h1 := sha256.Sum256(bytes)

	fmt.Println(h1, time.Now(), "sha256 해싱 완료")
}

func usedSHA1() {
	fmt.Println(time.Now(), "sha1 시작")
	bytes, err := ioutil.ReadFile("../test-file/test.dcf")
	if err != nil {
		fmt.Println(err)
		return
	}
	h1 := sha1.Sum(bytes)

	fmt.Println(h1, time.Now(), "sha1 해싱 완료")
}

func usedMD5() {
	fmt.Println(time.Now(), "md5 시작")
	bytes, err := ioutil.ReadFile("../test-file/test.dcf")
	if err != nil {
		fmt.Println(err)
		return
	}
	h1 := md5.Sum(bytes)

	fmt.Println(h1, time.Now(), "md5 해싱 완료")
}

//짤라서 읽고 해싱 뜨기
func divideSha512() {
	//읽을 파일 오픈
	file, _ := os.Open("../test-file/test.dcf")
	defer file.Close()

	fmt.Println(time.Now(), "divideSha512 시작")

	buf := make([]byte, READ_MEGA_BYTE*READ_UNIT)
	var h2 string
	for {
		n, err := file.Read(buf)
		//log.Printf("읽은 바이트 수: %d\n", n)

		if n > 0 {

			h1 := sha512.Sum512(buf[:n])
			stringSha512 := string(h1[:])
			//value, _ := fmt.Println(h1, "해싱")
			//s := strconv.Itoa(value)
			h2 = h2 + stringSha512

			//fmt.Println(time.Now(), "sha512 해싱 완료")
		}

		if err == io.EOF {
			log.Printf("다 읽은건가?: %d\n", n)
			b := []byte(h2)
			h3 := sha512.Sum512(b)
			fmt.Println(h3, time.Now(), "sha512 나눠서 더하기 해싱 완료")
			break
		}
	}

	/*if fi.Size >= READ_MEGA_BYTE*READ_UNIT {

	}

	bytesRead, err := file.Read(byteSlice)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of bytes read: %d\n", bytesRead)*/
	//log.Printf("Data read: %s\n", byteSlice)

}

func main() {
	/*s := "Hello, world!"
	h1 := sha512.Sum512([]byte(s)) // 문자열의 SHA512 해시 값 추출
	fmt.Printf("%x\n", h1)

	sha := sha512.New()          // SHA512 해시 인스턴스 생성
	sha.Write([]byte("Hello, ")) // 해시 인스턴스에 데이터 추가
	sha.Write([]byte("world!"))  // 해시 인스턴스에 데이터 추가
	h2 := sha.Sum(nil)           // 해시 인스턴스에 저장된 데이터의 SHA512 해시 값 추출
	fmt.Printf("%x\n", h2)*/

	usedSHA512()

	usedSHA256()

	usedSHA1()

	usedMD5()

	divideSha512()

	/*file, err := os.Open("../test-file/test.dcf")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(fi.Size())

	var data = make([]byte, fi.Size())

	fmt.Println( len(data))

	n, err := file.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, time.Now(), "바이트 읽기 완료")

	h1 := sha512.Sum512(data[:n])

	fmt.Println("%x\n", h1, time.Now(), "해싱 완료")*/
}
