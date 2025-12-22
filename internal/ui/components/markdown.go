package components

import (
	"regexp"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

// MarkdownToHTML converts markdown content to HTML
func MarkdownToHTML(md string) string {
	// Process custom shortcodes first
	md = processYouTubeShortcodes(md)
	md = processImageShortcodes(md)
	
	// Create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse([]byte(md))

	// Create HTML renderer with options  
	htmlFlags := html.CommonFlags | html.HrefTargetBlank | html.Safelink
	opts := html.RendererOptions{
		Flags: htmlFlags,
	}
	renderer := html.NewRenderer(opts)

	return string(markdown.Render(doc, renderer))
}

// processYouTubeShortcodes converts {{< youtube VIDEO_ID >}} to iframe
func processYouTubeShortcodes(content string) string {
	re := regexp.MustCompile(`\{\{<\s*youtube\s+([a-zA-Z0-9_-]+)\s*>\}\}`)
	return re.ReplaceAllStringFunc(content, func(match string) string {
		matches := re.FindStringSubmatch(match)
		if len(matches) < 2 {
			return match
		}
		videoID := matches[1]
		return `<div class="video-container my-8" style="position: relative; padding-bottom: 56.25%; height: 0; overflow: hidden; border-radius: 8px;">
<iframe src="https://www.youtube.com/embed/` + videoID + `" style="position: absolute; top: 0; left: 0; width: 100%; height: 100%;" frameborder="0" allowfullscreen></iframe>
</div>`
	})
}

// processImageShortcodes converts {{< image src="path" alt="text" >}} to styled img
func processImageShortcodes(content string) string {
	re := regexp.MustCompile(`\{\{<\s*image\s+src="([^"]+)"\s+alt="([^"]*)"\s*>\}\}`)
	return re.ReplaceAllStringFunc(content, func(match string) string {
		matches := re.FindStringSubmatch(match)
		src := matches[1]
		alt := matches[2]
		return `<div class="image-container my-8">
  <img 
    src="` + src + `" 
    alt="` + alt + `" 
    class="w-full rounded-lg shadow-lg"
    style="max-height: 500px; object-fit: cover;"
  />
</div>`
	})
}
