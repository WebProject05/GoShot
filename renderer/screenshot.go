package renderer

import (
	"context"
	"net/url"
	"os"
	"time"

	"github.com/chromedp/chromedp"
)

func RenderToImage(codeHTML string, filename string, output string) error {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	fullHTML := wrapHTML(codeHTML, filename)

	htmlURL := "data:text/html;charset=utf-8," + url.PathEscape(fullHTML)

	var buf []byte

	err := chromedp.Run(ctx,
		chromedp.EmulateViewport(1400, 900,
			chromedp.EmulateScale(3.0), // retina quality
		),

		chromedp.Navigate(htmlURL),
		chromedp.Sleep(500*time.Millisecond),

		chromedp.FullScreenshot(&buf, 100),
	)

	if err != nil {
		return err
	}

	return os.WriteFile(output, buf, 0644)
}

func wrapHTML(codeHTML string, filename string) string {
	html := `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>GoShot Screenshot</title>
	<style>
		/* Body background & center content */
		body {
			background: #2f416a;
			padding: 40px;
			display: flex;
			justify-content: center;
			align-items: flex-start;
			font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", monospace;
			margin: 0;
		}

		/* Mac-style window container */
		.window {
			background: #1e1e1e;
			border-radius: 20px;
			box-shadow: 0 20px 40px rgba(0,0,0,0.45);
			overflow: visible;
			display: inline-block;
			min-width: 400px;
		}

		/* Title bar with dots */
		.titlebar {
			height: 42px;
			background: #2b2b2b;
			display: flex;
			align-items: center;
			padding: 0 14px;
			border-radius: 20px 20px 0 0;
			position: relative;
		}

		.dots {
			display: flex;
			align-items: center;
		}

		.dot {
			width: 12px;
			height: 12px;
			border-radius: 50%;
			margin-right: 8px;
		}

		.red { background: #ff5f56; }
		.yellow { background: #ffbd2e; }
		.green { background: #27c93f; }

		/* Filename in center */
		.filename {
			position: absolute;
			left: 50%;
			transform: translateX(-50%);
			font-size: 13px;
			color: #d1d5db;
			opacity: 0.9;
			white-space: nowrap;
			user-select: none;
		}

		/* Content area for code */
		.content {
			padding: 24px;
		}

		/* Code styling */
		pre {
			margin: 0;
			font-size: 15px;
			line-height: 1.6;
			white-space: pre-wrap;
			word-break: break-word;
			color: #f8f8f2;
			background: #1e1e1e !important;
		}

		code {
			font-family: "JetBrains Mono", "Fira Code", "Consolas", monospace;
		}
	</style>
</head>
<body>
	<div class="window">
		<div class="titlebar">
			<div class="dots">
				<div class="dot red"></div>
				<div class="dot yellow"></div>
				<div class="dot green"></div>
			</div>
			<div class="filename">` + filename + `</div>
		</div>
		<div class="content">
			` + codeHTML + `
		</div>
	</div>
</body>
</html>`

	return html
}
