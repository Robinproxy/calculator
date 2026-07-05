# 剩余价值计算器 · Residual Value Calculator

<p align="center">
  <img src="static/favicon.svg" width="80" height="80" alt="Calculator Icon">
</p>

<p align="center">
  <a href="https://pages.cloudflare.com/" target="_blank">
    <img src="https://img.shields.io/badge/Deployed%20on-CF%20Pages-F38020?logo=cloudflare&logoColor=white" alt="Cloudflare Pages">
  </a>
  <img src="https://img.shields.io/badge/license-MIT-blue" alt="License">
</p>

<p align="center">
  <b>English</b> · <a href="#中文">中文</a>
</p>

---

<p align="center">
  A lightweight tool for calculating the <strong>residual value</strong> of prepaid services (VPS, hosting, subscriptions).<br>
  Supports prorated daily calculation, real-time exchange rates across 11 currencies, and premium/discount analysis.
</p>

---

## ✨ Features

| Feature | Description |
|---------|-------------|
| ⏱ **Daily Proration** | Computes remaining value based on days elapsed in the billing cycle |
| 💱 **11 Currencies** | USD, CNY, EUR, GBP, JPY, KRW, HKD, TWD, CAD, SGD, AUD |
| 🔄 **Live Exchange Rates** | Fetched from ExchangeRate-API, with local cache fallback |
| 📈 **Premium/Discount** | Compare actual trade price against calculated residual value |
| 🔗 **URL Sharing** | All parameters encoded in URL — share the link, not a screenshot |
| 🌙 **Dark Mode** | Toggle manually or auto-follows system preference |
| 📋 **Markdown Export** | Copy results as a formatted Markdown table |
| 📱 **PWA Ready** | Installable on mobile, works offline via Service Worker |

---

## 🚀 Live Demo

Deploy your own:

[![Deploy to Cloudflare Pages](https://img.shields.io/badge/Deploy%20to-CF%20Pages-F38020?logo=cloudflare)](https://dash.cloudflare.com/?to=/:account/pages/)

```
1. Fork this repo
2. Go to Cloudflare Pages → Create a Page → Connect to Git
3. Build output directory: static
4. Deploy 🎉
```

---

## 🖥 Usage

1. Enter the **renewal amount** and select the **billing cycle**
2. Pick the **service end date** and **transaction date**
3. Choose **source currency** and **target settlement currency**
4. Click **"计算剩余价值" (Calculate)**
5. Optionally enter an **actual trade price** for premium/discount analysis
6. Copy the result as a Markdown table via **"复制文本" (Copy Text)**

```
## 剩余价值评估

| 项目 | 数值 |
|------|------|
| 剩余天数 | 182 天 |
| 剩余价值 | $24.95 (USD) |
| 交易金额 | $20.00 |
| 溢价情况 | -$4.95 (折价) |
```

---

## 🏗 Project Structure

```
calculator/
├── static/
│   ├── index.html        ← Single-page application
│   ├── favicon.svg        ← Favicon
│   ├── manifest.json      ← PWA manifest
│   └── sw.js              ← Service Worker
└── README.md
```

That's it — **one HTML file**, zero dependencies, zero build step.

---

## 📄 License

[MIT](LICENSE)

---

## <a id="中文"></a>中文说明

# 剩余价值计算器

<p align="center">
  一款轻量级的 VPS/服务 **剩余价值计算器**。<br>
  支持按天折算、11 种货币实时汇率、溢价分析，适合 VPS 二手交易场景。
</p>

---

## ✨ 功能一览

| 功能 | 说明 |
|------|------|
| ⏱ **按天折算** | 根据支付周期已过天数，精确计算剩余价值 |
| 💱 **11 种货币** | 美元、人民币、欧元、英镑、日元、韩元、港元、新台币、加元、新加坡元、澳元 |
| 🔄 **实时汇率** | 自动获取 ExchangeRate-API 汇率，断网时使用本地缓存 |
| 📈 **溢价分析** | 输入实际成交价，自动计算是溢价还是折价（红/绿标识） |
| 🔗 **URL 分享** | 所有参数编码在链接中，发链接就能让别人看到你的计算结果 |
| 🌙 **暗黑模式** | 手动切换或自动跟随系统设置 |
| 📋 **Markdown 导出** | 一键复制结果表格，方便粘贴到论坛/Telegram/Discord |
| 📱 **PWA 支持** | 可添加到手机桌面，离线也能查看 |

---

## 🚀 在线演示

一键部署自己的版本：

```
1. Fork 本仓库
2. 打开 Cloudflare Pages → Create a Page → 连接 GitHub
3. 构建输出目录填 static
4. 部署完成 🎉
```

---

## 🖥 使用方法

1. 输入 **续费金额**，选择 **付款周期**（月付/季付/半年付/年付等）
2. 选择 **到期时间** 和 **交易日期**
3. 选择 **原币种** 和 **结算币种**（支持 11 种货币）
4. 点击 **"计算剩余价值"**
5. 可选：输入 **实际成交金额**，查看溢价/折价分析
6. 点击 **"复制文本"**，一键复制 Markdown 表格

```
## 剩余价值评估

| 项目 | 数值 |
|------|------|
| 剩余天数 | 182 天 |
| 剩余价值 | ¥178.50 (CNY) |
| 交易金额 | ¥150.00 |
| 溢价情况 | -¥28.50 (折价) |
```

---

## 🏗 项目结构

```
calculator/
├── static/
│   ├── index.html        ← 单页应用，一切都在这个文件里
│   ├── favicon.svg        ← 网站图标
│   ├── manifest.json      ← PWA 配置文件
│   └── sw.js              ← 离线缓存（Service Worker）
└── README.md              ← 本文件
```

整个项目就 **一个 HTML 文件**，零依赖、零构建步骤。

---

## 📄 许可证

[MIT](LICENSE)