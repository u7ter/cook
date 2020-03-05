package main

import (
	"github.com/maksimmysak/cook/api"
)

func main() {

	api.Run()

}

//package main
//
//import (
//"fmt"
//"github.com/dafanasev/go-yandex-translate"
//)
//
//func main() {
//
//	tr := translate.New("trnsl.1.1.20170505T201046Z.765061fd7d327f2f.c80d8b95dd956de79d7f9537011fcd3cc802e6e2")
//
//	response, err := tr.GetLangs("en")
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(response.Langs)
//		fmt.Println(response.Dirs)
//	}
//
//	translation, err := tr.Translate("ru", "A lazy dog")
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println(translation.Result())
//	}
//}
