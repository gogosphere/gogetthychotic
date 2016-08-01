# getThycoticCreds

This is meant to offer some of the same flexibility in accessing your Secrets on Thychotic Secret Server. Alone, it is a way to both search for credentials and retrieve credentials and secrets stored in Thycotic Secret Server. Currently it offers full AD Authentication using Thycotic as the pivot point to the local domain, as well as searching capabilities. Secrets are stored and retrieved via a SecretID, you may text search for your secrets name to locate the ID and then lookup all values stored with the secret itself (Description, Name, URL, Password, Email, etc).

I originally wrote this for folks here, so I didn't bother with a config for it.  You'll need to edit two values in the code, the hosturl may be different but you can look at your SS config for guidance:

```
const (
  // update these (or replace them with a Flags package which ever is better for the user)
  domain  = "yourlocal auth active directory domain"
  hosturl = "https://yoursecretserver.domain.com/SecretServer/webservices/SSWebservice.asmx"
```


## Installation

Clone the repo
```
go run gogetSecret.go
-or-
get build gogetSecret.go
```
Then you can execute the binary, it is written in a way that you can cross compile on any platform you wish and just distribute the binaries. 

## Usage

Simply execute the proper binary. 

```
MYMAC:getThycoticCreds whancock$ go run ./gogetSecret.go
AD Username: whancock
Password: *************

(S)earch or (L)ookup S
Criteria: AWS
```

In the above example I'm searching for all instances of AWS in the name of the secret, and I'm returned 
all secretIDs and secrets that I have access to that are associated with AWS in the name or description.

```
22 AWS Account Secrets
... snip ...
2967 AWS Palto Login
```

Now that I have the secretID I can perform a (L)ookup:
```
MYMAC:getThycoticCreds whancock$ ./gogetSecret-darwin
AD Username: whancock
Password: *****************************************

(S)earch or (L)ookup l
Criteria: 22
https:url to whatever
email attched to the secret
password ***************
NOTES
2016/07/26 05:28:27 END OF RESULTS: SecretID
```


## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D

## History

Version .1 released 7/26/2016

## Credits

Bill Hancock

## License
MIT
