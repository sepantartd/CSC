# ☁️ Cloud Saver CLI

![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)
![Platform](https://img.shields.io/badge/Platform-Linux_|_macOS_|_Windows_|_Termux-orange?style=for-the-badge)

[فارسی](README.md) | **English**

**Cloud Saver** is a lightweight, intelligent command-line tool (CLI) written in Go. It fetches and analyzes your DigitalOcean cloud resources and leverages OpenAI language models to provide color-coded, actionable FinOps cost optimization recommendations directly in your terminal.

---

## 🚀 Features

- 🔍 **Automatic Resource Fetching**: Retrieves Droplets, Volumes, and Managed Databases via the DigitalOcean API.
- 🤖 **AI-Powered Analysis**: Evaluates resource utilization using OpenAI GPT-4o-mini to spot idle or over-provisioned infrastructure.
- 🎨 **Rich Terminal UI**: Displays interactive, color-coded recommendations categorized by severity (Delete, Downsize, Keep).
- 📱 **Termux & Mobile Friendly**: Fully functional on Android via Termux without requiring root access or `sudo`.

---

## 📦 Installation

### Quick Install (via `go install`)
If you have Go installed, you can install the CLI directly:

```bash
go install [github.com/sepantartd/cloud-saver@latest](https://github.com/sepantartd/cloud-saver@latest)
```

### Build from Source

```bash
# Clone the repository
git clone [https://github.com/sepantartd/cloud-saver.git](https://github.com/sepantartd/cloud-saver.git)
cd cloud-saver

# Download dependencies & build binary
go mod download
go build -o cloud-saver main.go
```

---

## ⚙️ Configuration

Create a `.env` file in the root directory of the project and add your API credentials:

```env
DIGITALOCEAN_TOKEN=dop_v1_your_token_here
OPENAI_API_KEY=sk-proj-your_key_here
```

### Obtaining API Keys
1. **DigitalOcean Token**: Log into DigitalOcean -> Go to **API** -> Generate a **Personal Access Token** with Read permissions.
2. **OpenAI API Key**: Visit **platform.openai.com** -> Go to **API Keys** -> Create a new secret key.

---

## 💻 Usage

### Check CLI Version
```bash
./cloud-saver version
```

### Run Cost Optimization Analysis
```bash
./cloud-saver analyze
```

---

## 📄 License

Distributed under the [MIT](LICENSE) License.
