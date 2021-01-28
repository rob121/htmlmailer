# htmlmailer

A Simple html mailer that will send a message from a html template


## Configuration

Create a config.json, this will go in /etc/htmlmailer, /$user/.htmmailer or ./config.json

```
{
  "host": "host",
  "user": "user@mail.com",
  "password": "abc123",
  "port": "25"
}
```


## Usage

```
./htmlmailer --embed=./image1.jpeg,./image2.jpeg --to=me@example.com --from=sender@example.com --subject="Test"
```

The mailer will load template.html file and send it's content, if you need to embed files they can be put in the html via using cid:

```
<img width="100%" src="cid:image1.jpeg" alt="My image here" />
```


## Arguments

```
Usage of ./htmlmailer:
  -attach string
    	Comma list of resources to attach (full path)
  -embed string
    	Comma list of embeddable resources (full path)
  -from string
    	From Address
  -subject string
    	Subject
  -template string
    	Path to template file (default "template.html")
  -to string
    	To Address (Use comma for multiples)
      
```

