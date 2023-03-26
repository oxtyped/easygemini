package main

import (
	"github.com/oxtyped/easygemini/cmd"
)

//func main2() {
//	//C: Opens connection
//	//S: Accepts connection
//	//C/S: Complete TLS handshake (see section 4)
//	//C: Validates server certificate (see 4.2)
//	//C: Sends request (one CRLF terminated line) (see section 2)
//	//S: Sends response header (one CRLF terminated line), closes connection under non-success conditions (see 3.1 and 3.2)
//	//S: Sends response body (text or binary data) (see 3.3)
//	//S: Closes connection (including TLS close_notify, see section 4)
//	//C: Handles response (see 3.4)

//	config, err := loadConfig()
//	if err != nil {
//		log.Fatal(err)
//	}

//	content := config.Server().Plugin("content")

//	var path string
//	content.MapDirective("file", func(op rubyfile.Option) {
//		fmt.Printf("its file: %s\n\n", op.Args)
//		if len(op.Args) > 1 {
//			panic("eh")
//		}

//		path = op.Args[0]
//	})

//	//for _, v := range opt {
//	//	fmt.Printf("name of option is %s\n", v.Name)
//	//	fmt.Printf("name of args is %#v\n", v.Args)

//	//}

//}
//func loadConfig() (*rubyfile.Corefile, error) {

//	a, err := rubyfile.New("Rubyfile")
//	if err != nil {
//		return nil, err
//	}

//	fmt.Printf("servers is %#v\n", a.Servers)
//	for _, v := range a.Servers {
//		for _, vv := range v.Plugins {
//			fmt.Printf("plugins for %s is %#v\n", v.String(), vv.Name)

//		}

//	}

//	fmt.Printf("%#v", a)
//	return a, nil

//}

func main() {

	cmd.Execute()

}
