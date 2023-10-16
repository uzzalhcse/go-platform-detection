package platform

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Constants for device detection
const (
	DefaultKey  = "github.com/uzzalhcse/go-platform-detection"
	Android     = "android"
	Mobile      = "mobile"
	Ipad        = "ipad"
	Iphone      = "iphone"
	Ipod        = "ipod"
	Ios         = "ios"
	Kindle      = "Kindle"
	Silk        = "Silk"
	Unknown     = "Unknown"
	Wap         = "wap"
	XwapProfile = "X-Wap-Profile"
	Profile     = "Profile"
)

// MobileUserAgentPrefixes List of mobile user agent prefixes
var MobileUserAgentPrefixes = []string{
	"w3c ", "w3c-", "acs-", "alav", "alca", "amoi", "audi", "avan", "benq",
	"bird", "blac", "blaz", "brew", "cell", "cldc", "cmd-", "dang", "doco",
	"eric", "hipt", "htc_", "inno", "ipaq", "ipod", "jigs", "kddi", "keji",
	"leno", "lg-c", "lg-d", "lg-g", "lge-", "lg/u", "maui", "maxo", "midp",
	"mits", "mmef", "mobi", "mot-", "moto", "mwbp", "nec-", "newt", "noki",
	"palm", "pana", "pant", "phil", "play", "port", "prox", "qwap", "sage",
	"sams", "sany", "sch-", "sec-", "send", "seri", "sgh-", "shar", "sie-",
	"siem", "smal", "smar", "sony", "sph-", "symb", "t-mo", "teli", "tim-",
	"tosh", "tsm-", "upg1", "upsi", "vk-v", "voda", "wap-", "wapa", "wapi",
	"wapp", "wapr", "webc", "winw", "winw", "xda ", "xda-"}

// MobileUserAgentKeywords List of mobile user agent keywords
var MobileUserAgentKeywords = []string{
	"blackberry", "webos", "ipod", "lge vx", "midp", "maemo", "mmp", "mobile",
	"netfront", "hiptop", "nintendo DS", "novarra", "openweb", "opera mobi",
	"opera mini", "palm", "psp", "phone", "smartphone", "symbian", "up.browser",
	"up.link", "wap", "windows ce"}

// TabletUserAgentKeywords List of tablet user agent keywords
var TabletUserAgentKeywords = []string{"ipad", "playbook", "hp-tablet", "kindle"}

// Device represents a device's type and platform
type Device interface {
	IsNormal() bool
	IsMobile() bool
	IsTablet() bool
	GetPlatform() string
}

// device implements Device interface
type device struct {
	isNormal bool
	isMobile bool
	isTablet bool
	platform string
}

// ResolveDevice Middleware function that resolves the device type based on the User-Agent header.
func ResolveDevice() gin.HandlerFunc {
	return func(c *gin.Context) {
		d := getDeviceType(c.Request.Header)
		c.Set(DefaultKey, d)
		c.Next()
	}
}

// getDeviceType determines the type of device based on the User-Agent header.
func getDeviceType(header http.Header) Device {
	agent := strings.ToLower(header.Get("User-Agent"))

	// Check Tablet User Agent Keywords
	for _, keyword := range TabletUserAgentKeywords {
		if strings.Contains(agent, keyword) {
			switch keyword {
			case Ipad:
				return &device{isTablet: true, platform: Ipad}
			case Kindle:
				return &device{isTablet: true, platform: Kindle}
			default:
				return &device{isTablet: true, platform: Unknown}
			}
		}
	}

	// User Agent Profile detection
	xWapProfile := header.Get(XwapProfile)
	profile := header.Get(Profile)

	if xWapProfile != "" || profile != "" {
		switch {
		case strings.Contains(agent, Android):
			return &device{isMobile: true, platform: Android}
		case strings.Contains(agent, Iphone) || strings.Contains(agent, Ipod) || strings.Contains(agent, Ipad):
			return &device{isMobile: true, platform: Ios}
		default:
			return &device{isMobile: true, platform: Unknown}
		}
	}

	// User Agent Prefix check
	prefix := agent[:4]
	for _, uaPrefix := range MobileUserAgentPrefixes {
		if strings.Contains(prefix, uaPrefix) {
			return &device{isMobile: true, platform: Unknown}
		}
	}

	// Accept Header check
	accept := header.Get("Accept")
	if accept != "" && strings.Contains(accept, Wap) {
		return &device{isMobile: true, platform: Unknown}
	}

	// Check Mobile User Agent Keywords
	for _, keyword := range MobileUserAgentKeywords {
		if strings.Contains(agent, keyword) {
			switch {
			case strings.Contains(agent, Android):
				return &device{isMobile: true, platform: Android}
			case strings.Contains(agent, Iphone) || strings.Contains(agent, Ipod) || strings.Contains(agent, Ipad):
				return &device{isMobile: true, platform: Ios}
			default:
				return &device{isMobile: true, platform: Unknown}
			}
		}
	}

	// Assume 'normal' if mobile or tablet was not identified
	return &device{isNormal: true, platform: Unknown}
}

// GetDevice returns the detected device type from the context.
func GetDevice(c *gin.Context) Device {
	return c.MustGet(DefaultKey).(Device)
}

// IsNormal returns true if a device is normal (neither mobile nor tablet).
func (d *device) IsNormal() bool {
	return d.isNormal
}

// IsMobile returns true if a device is a mobile device.
func (d *device) IsMobile() bool {
	return d.isMobile
}

// IsTablet returns true if a device is a tablet.
func (d *device) IsTablet() bool {
	return d.isTablet
}

// GetPlatform returns the device's platform (e.g., Android, iOS, Kindle, Unknown).
func (d *device) GetPlatform() string {
	return d.platform
}
