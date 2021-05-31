## Internet speed tester

Given GoLang module allows user to know current internet upload and download speed. You can use one of two possible speed testing providers: Ookla and Fast.
In case you entered wrong provider name value it returns error.

Example of usage:

```
package main

import "github.com/serenko-portfolio/internet_speed"

func main() {
	upload_speed,download_speed,err := internet_speed.GetInternetSpeed("Ookla")
}

```

Implementation uses one library directly for Ookla and headless chrome emulator for Fast.com. There are no libraries were found that allow to measure upload speed too in a more convenient way.

Module also contains tests. Overall test coverage is about 85%

