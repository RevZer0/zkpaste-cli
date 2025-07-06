# ZK.paste CLI tool

## Installation

Download the lastest release for your platform from the Releases section

## Usage

### Use the the help

```

$ zkpaste --help 


```

### Create the paste

```

$ zkpaste create "this is the paste" 


```

or stream stdin through the pipe

```

$ cat /etc/passwd | zkpaste create 


```

### Create paste with password

```

$ zkpaste create "password protected paste" --password s3cur3pwd 


```

### Set the TTL option

```

$ zkpaste create "this is the paste" --ttl 10m


```

Use the help for all the available TTL options

### Limit the views count

```

$ zkpaste create "this is the paste" --views 5 


```

### Read the paste 

```

$ zkpaste read https://zkpaste.com/paste/f4748e87-c573-463b-bafc-00fc284fece1#H8q01JOWJSJmE9IPhAvGnarQbS27Q9fl/oDWHWxOSQY=


```

### Delete the paste

```

$ zkpaste delete https://zkpaste.com/paste/f4748e87-c573-463b-bafc-00fc284fece1#H8q01JOWJSJmE9IPhAvGnarQbS27Q9fl/oDWHWxOSQY=


```
