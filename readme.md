# HTTP Request raw data
get / http/1.1 <br>
host: localhost:8080 <br>
header_key: header_value <br>
--new line. below is request body-- <br>

# HTTP Response raw data
http/1.1 200 ok <br>
content-type: text/html <br>

--new line. below is response data-- <br>
`<h2>hello world</h2>`