# 🚀 GitHub Repo Card CLI

A professional, high-performance CLI tool built in **Go** that generates stunning, **Neo-Brutalism** style repository cards for your GitHub projects.

![Repo Card Preview](https://github.com/adeleeeeyyyy/GoCard/raw/main/mockup.png)
*(Replace with actual generated image or screenshot)*

## ✨ Features

- **🎨 Stunning UI**: Modern Neo-Brutalism design with soft pastel colors, rounded corners, and glassmorphism.
- **✨ Professional Icons**: Integrated **Lucide Icons** for the UI and **Devicon/Simple Icons** for programming languages.
- **⚡ Fast & Lightweight**: Powered by Go and server-side rendering.
- **📸 Smart Export**: 
  - **In-Browser PNG**: Download high-quality cards directly from your browser.
  - **Headless CLI PNG**: Export via `--png` flag using `chromedp` (requires Chrome/Chromium).
- **📄 README Preview**: Automatically snippets your repository's README for a complete overview.

## 🛠️ Installation

Ensure you have [Go](https://go.dev/doc/install) (1.20+) installed.

```bash
git clone https://github.com/adeleeeeyyyy/GoCard.git
cd GoCard
go build -o GoCard main.go
```

## 🚀 Usage

Generate a card by simply providing the `owner/repository` name:

```bash
./GoCard gen golang/go
```

### Options

| Flag | Shorthand | Description |
| --- | --- | --- |
| `--output` | `-o` | Specify custom name for the HTML file |

### Examples

```bash
# Generate with custom name
./repocard gen facebook/react -o my-card.html

# Generate with PNG export (requires Chrome)
./repocard gen vercel/next.js --png
```

## 📦 Project Structure

```text
.
├── cmd/                # CLI Command logic (Cobra)
├── internal/
│   ├── github/         # GitHub API Client
│   ├── models/         # Data structures
│   └── render/         # HTML/PNG Rendering logic
├── main.go             # Entry point
└── README.md
```

## 🤝 Contributing

Contributions are welcome! Feel free to open issues or pull requests to improve the design or add features.

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---
Created with ❤️ by **adeleeeeyyyy**
