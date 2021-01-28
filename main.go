package main

import(
	"bytes"
	"gopkg.in/gomail.v2"
	"fmt"
	"html/template"
  "github.com/rob121/vhelp"
	"flag"
	"log"
	"strings"
)


var embed_raw string
var embed []string
var attach_raw string
var attach []string
var tmpl string
var to_raw string
var to []string
var from string
var subject string

func main() {
    flag.StringVar(&to_raw,"to","","To Address")
	flag.StringVar(&from,"from","","From Address")
    flag.StringVar(&subject,"subject","","Subject")
	flag.StringVar(&tmpl,"template","template.html","Path to template file")
	flag.StringVar(&embed_raw,"embed","","Comma list of embeddable resources (full path)")
	flag.Parse()

	embed = strings.Split(embed_raw,",")
	attach = strings.Split(attach_raw,",")
	to = strings.Split(to_raw,",")

	vhelp.Load("config")

	conf,cerr := vhelp.Get("config")

	if(cerr!=nil){}

	conf.SetDefault("port",25)

	host := conf.GetString("host")
	user := conf.GetString("user")
	pass := conf.GetString("password")
    port := conf.GetInt("port")

	fmt.Println("Sending Message...")
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)


	for _,e := range embed {
		if(len(e)<1){
			continue
		}
		fmt.Printf("Embedding: %s\n",e)
		m.Embed(e)
	}

	for _,a := range attach {
		if(len(a)<1){
			continue
		}

		fmt.Printf("Attaching: %s\n",a)
		m.Attach(a)
	}

	result := parseTmpl()
    m.SetBody("text/html",result)
	d := gomail.NewDialer(host, port, user, pass)
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Message Sent!")

}

func parseTmpl() string{

	t := template.New(tmpl)

	var err error
	t, err = t.ParseFiles(tmpl)
	if err != nil {
		log.Println(err)
	}

	var data interface{}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl,data); err != nil {
		log.Println(err)
	}

	result := tpl.String()

	return result


}
