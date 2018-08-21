# xtract
xtract text from html, trim spaces for each token by default (Use xtract)

## Usage:

```go
// Extract all text from the page
xtract.Page("https://github.com/golovers")

// Extract the first 20 words from the page
xtract.PageLim("https://github.com/golovers", 20)

// Extract all text from html value
xtract.Value("<div>Text</div>")

// Extract the first word in the html value
xtract.ValueLim("<div>Text and another text</div>", 1)

// Override trim function if needed
xtract.SetTrimFunc(f TrimFunc)
```

