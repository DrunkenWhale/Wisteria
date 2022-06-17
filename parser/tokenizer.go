package parser

import (
	"log"
	"regexp"
	"strings"
)

// There exist and only exist some operation
// PUT/GET/DELETE/UPDATE a STRING/INT/FLOAT/MAP/LIST
// specially, we must analyze LIST Element to know its type
// and make sure whether MAP's kvs' type are string

func Analyze(sentence string) {
	str := strings.TrimSpace(sentence)
	strRe, err := regexp.Compile(",(?=(?:[^\"]*\"[^\"]*\")*[^\"]*$)")
	if err != nil {
		log.Fatalln(err)
	}
	strArray := strRe.Split(str, -1)

	log.Println(strArray)
}
