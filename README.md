## wordgen

Easy wordlist generator tool for backup files.

## Installation
```
go install github.com/rix4uni/wordgen@latest
```

## Download prebuilt binaries
```
wget https://github.com/rix4uni/wordgen/releases/download/v0.0.1/wordgen-linux-amd64-0.0.1.tgz
tar -xvzf wordgen-linux-amd64-0.0.1.tgz
rm -rf wordgen-linux-amd64-0.0.1.tgz
mv wordgen ~/go/bin/wordgen
```
Or download [binary release](https://github.com/rix4uni/wordgen/releases) for your platform.

## Compile from source
```
git clone --depth 1 github.com/rix4uni/wordgen.git
cd wordgen; go install
```

## Usage
```
Usage of wordgen:
  -e string
        Comma separated list of extensions
  -path string
        Path or list of paths to add (comma-separated or a file)
  -silent
        Silent mode.
  -version
        Print the version of the tool and exit.
```

## Usage Examples
**Single URL, Without extensions:**
```
▶ echo "https://pentestingdorks.netlify.app" | wordgen -path "github,admin"

https://pentestingdorks.netlify.app/github
https://pentestingdorks.netlify.app/admin
```

**Single URL, With extensions:**
```
▶ echo "https://pentestingdorks.netlify.app" | wordgen -path "github,admin" -e ".zip,.php"

https://pentestingdorks.netlify.app/github.zip
https://pentestingdorks.netlify.app/github.php
https://pentestingdorks.netlify.app/admin.zip
https://pentestingdorks.netlify.app/admin.php
```

**Multiple URLs with paths file, Without extensions:**
```
▶ cat subs.txt | wordgen -path wordlist.txt

https://pentestingdorks.netlify.app/github
https://pentestingdorks.netlify.app/admin
```

**Multiple URLs with paths file, Multiple URLs, With extensions:**
```
▶ cat subs.txt | wordgen -path wordlist.txt -e ".zip,.php"

https://pentestingdorks.netlify.app/github.zip
https://pentestingdorks.netlify.app/github.php
https://pentestingdorks.netlify.app/admin.zip
https://pentestingdorks.netlify.app/admin.php
```