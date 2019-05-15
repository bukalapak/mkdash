# mkdash
Creating dashboard in a dash

## Description
MKDash is a cli application that can create a standard dashboard from a template for your services


## Owner

SRE (sre@bukalapak.com)

## Contact and On-Call Information

See [Contact and On-Call Information](https://bukalapak.atlassian.net/wiki/display/INF/Contact+and+On-Call+Information)

### Installing

#### Linux 64

1. Download file with browser (do not use wget or curl, because mkdash is a private repo). Choose **mkdash_linux_amd64**

   ```
   https://github.com/bukalapak/mkdash/releases/latest
   ```

2. Give permission to execute

   ```
   sudo chmod +x ~/Downloads/mkdash_linux_amd64
   ```

3. Move to /usr/local/bin

   ```
   sudo mv ~/Downloads/mkdash_linux_amd64 /usr/local/bin/mkdash
   ```


#### OSX 64

1. Download file with browser (do not use wget or curl, because mkdash is a private repo). Choose **mkdash_darwin_amd64**

   ```
   https://github.com/bukalapak/mkdash/releases/latest
   ```

2. Give permission to execute

   ```
   sudo chmod +x ~/Downloads/mkdash_darwin_amd64
   ```

3. Move to /usr/local/bin

   ```
   sudo mv ~/Downloads/mkdash_darwin_amd64 /usr/local/bin/mkdash
   ```


#### Windows 64

1. Download file (choose **mkdash_windows_amd64.exe**)

   ```
   https://github.com/bukalapak/mkdash/releases/latest
   ```

#### Other Architecture

See https://github.com/bukalapak/mkdash/releases/latest

#### Usage
	```
	mkdash <project-name> <service-name>
	# project-name is your project name for the service
	# ervice-name is your service name
	mkdash chronos api #example
	```	

when it successful, it will create a dashboard on grafana with the name `MKDash - <project-name> - <servicename>`
