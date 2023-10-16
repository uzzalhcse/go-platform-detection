# Go Platform Detection

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

This Go package provides a robust device detection mechanism based on User-Agent headers. It is designed to help web applications distinguish between normal browsers, mobile devices, and tablets.

## Features

-   Detects devices based on User-Agent headers.
-   Compatible with the [Gin](https://github.com/gin-gonic/gin) web framework.
-   Provides an easy-to-use middleware for integrating device detection into your Gin application.
-   Supports identification of normal browsers, mobile devices, and tablets.

## Installation
Install the package using `go get`:

```
go get github.com/uzzalhcse/go-platform-detection
```

## Usage

1.  Import the package:
 ```
import "github.com/uzzalhcse/go-platform-detection"
```
2.  Add the middleware to your Gin router:

    ```
    r := gin.Default()
    r.Use(platform.ResolveDevice())
    ```

3.  Retrieve the detected device type in your handlers:

    ```
    device := platform.GetDevice(c)
    ```


## Example

```
package main

import (
"github.com/gin-gonic/gin"
"github.com/uzzalhcse/go-platform-detection"
)

func main() {
r := gin.Default()

	// Use the platform detection middleware
	r.Use(platform.ResolveDevice())

	r.GET("/hello", func(c *gin.Context) {
		device := platform.GetDevice(c)
		c.JSON(200, gin.H{
			"message":   "Hello World!",
			"device":    device,
			"platform":  device.GetPlatform(),
			"is_normal": device.IsNormal(),
			"is_mobile": device.IsMobile(),
			"is_tablet": device.IsTablet(),
		})
	})

	r.Run()
}
```

## Contributing

Contributions are welcome! For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

This project is licensed under the MIT License - see the [LICENSE](https://chat.openai.com/c/LICENSE) file for details.

## Author

[Uzzal Hosen](https://github.com/uzzalhcse)