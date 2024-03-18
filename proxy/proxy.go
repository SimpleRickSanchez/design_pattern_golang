package proxy

import "fmt"

// Image 是抽象主题接口，定义了图片对象的行为
type Image interface {
	Display()
}

// RealImage 是真实主题类，实现了 Image 接口
type RealImage struct {
	filename string
}

// Display 方法显示图片内容
func (img *RealImage) Display() {
	fmt.Printf("Displaying %s\n", img.filename)
}

// ProxyImage 是代理类，也实现了 Image 接口
type ProxyImage struct {
	realImage *RealImage
	loaded    bool
}

// Display 方法决定是否加载真实图片并显示
func (img *ProxyImage) Display() {
	if !img.loaded {
		img.realImage = &RealImage{"proxied_image.jpg"}
		img.loaded = true
	}
	img.realImage.Display()
}
