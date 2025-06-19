# WebDash
Fast, interactive terminal application for launching your favorite websites with beautiful UI. Easily customizable for any collection of web resources - documentation, tools, social media, or entertainment.

# 🐹 Go Resources CLI

## ✨ Features

### 🎯 **Interactive Navigation**
- **Keyboard Navigation**: Use arrow keys (↑/↓) or vim-style keys (k/j) to browse through resources
- **Instant Access**: Press Enter, Space, or click to open resources directly in your default browser

### 📚 **Curated Go Resources**
The application comes pre-loaded with 5 essential Go learning resources:

1. **Go.dev** - Official Go website with comprehensive documentation, tutorials, and packages
2. **Go by Example** - Hands-on introduction to Go using annotated example programs
3. **Effective Go** - Essential guide for writing clear, idiomatic Go code
4. **Go Tour** - Interactive tour covering Go's basics and advanced features
5. **Awesome Go** - Curated list of awesome Go frameworks, libraries, and software

### 🎨 **Beautiful Terminal UI**
- **Modern Design**: Clean, colorful interface with professional styling
- **Visual Feedback**: Clear indication of selected items with colors and symbols
- **Responsive Layout**: Adapts to different terminal sizes
- **Non-intrusive**: Uses alternate screen buffer, preserving your terminal history


## 🚀 **Quick Start**

### Installation
```bash
# Clone or create project directory
mkdir go-resources-cli && cd go-resources-cli

# Copy the source code to main.go
# Copy the go.mod file

# Install dependencies
go mod tidy

# Run the application
go run main.go
```

### Usage
```bash
# Launch the application
go run main.go

# Navigation:
# ↑/↓ or k/j - Navigate through resources
# Mouse wheel - Scroll through list
# Enter/Space - Open selected resource
# Mouse click - Select and open resource
# q/Ctrl+C/Esc - Exit application
```

## 🔧 **Customization**

### **Adapt for Your Own Resources**
This application is designed to be easily customizable for any collection of websites or resources you frequently visit. Simply modify the `resources` slice in the code:

```go
var resources = []Resource{
    {
        Name:        "Your Favorite Site",
        URL:         "https://yoursite.com",
        Description: "Description of what this site offers",
    },
    // Add more resources...
}
```

### **Perfect for Various Use Cases:**
- **Development Resources**: Documentation sites, tutorials, code repositories
- **News & Blogs**: Your daily reading list, tech blogs, news sources  
- **Social Media**: Quick access to different platforms
- **Work Tools**: Internal tools, dashboards, project management sites
- **Learning Materials**: Online courses, reference materials, cheat sheets
- **Entertainment**: Streaming services, gaming sites, forums

### **Easy Styling Customization**
Modify the `lipgloss` styles to match your preferred color scheme:
- Change colors by updating hex values in style definitions
- Adjust padding and margins for different layouts
- Customize fonts and text formatting

## 🏗️ **Technical Details**

### **Built With**
- **Go**: Fast, reliable systems programming language
- **Bubble Tea**: Powerful TUI framework based on Elm architecture
- **Lipgloss**: Terminal styling library for beautiful interfaces

---

**Perfect for developers, students, and anyone who wants quick, organized access to their frequently visited websites through a beautiful command-line interface!**
