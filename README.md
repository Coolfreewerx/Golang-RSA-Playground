# Golang-RSA-Playground
- Encrypt & Decrypt file (`*.csv.zip`) by use public key & private key RSA-4096 (`*.asc`)  
- Encode & Decode encrypted file (*.csv.zip) by use Base64

## Result-Test
#### Encrypt zip file , encode/decode encrypted zip file and decrypt zip file to csv file

	In practice, someone want to send a zip file to Tam will use Tam's public key to encrypt the zip gilr.
	Then, Bring encrypted zip file to Tam by use Base64 to encode the zip file. 
	After that, Tam will decode the zip file by use Base64 and decrypt the zip file by use Tam's private key.
	Finaly, Tam unzip file and got *.csv file.
	
```
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDbHQDPnqbbiJM3NTkeG+4LeFOEVb6wr8magesbCMbP15tJuSFw....
-----END RSA PRIVATE KEY-----

-----BEGIN RSA PUBLIC KEY-----
MIGJAoGBANsdAM+eptuIkzc1OR4b7gt4U4RVvrCvyZqB6xsIxs/Xm0m5IXCp2Eks....
-----END RSA PUBLIC KEY-----

-----BEGIN MESSAGE-----
ZuTnpW5rA9qPdEmmDGE35V+g3iMUpDDo5FTeyvlGGhRmxUG6CnIDTIun9KhbLqJD....
-----END MESSAGE-----


-------------Begin With Base64------------------

Encode message from RSA By Base64 :
LS0tLS1CRUdJTiBNRVNTQUdFLS0tLS0KW....

Decode message from RSA By Base64 :
-----BEGIN MESSAGE-----
ZuTnpW5rA9qPdEmmDGE35V+g3iMUpDDo5FTeyvlGGhRmxUG6CnIDTIun9KhbLqJD....
...
-----END MESSAGE-----

-------------End With Base64------------------

Decrypted success
```
Finally, you will got unzip file (*.csv.zip) in `csv-files folder`

#### Covid-Thailand-2021 API
- swagger
```
http://localhost:1122/swagger/index.html#/
```

### Conference
- [Dataset (COVID-19 in Thailand 2021)](https://data.go.th/en/dataset/covid-19-daily)