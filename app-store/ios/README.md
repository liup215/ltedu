# iOS – Apple App Store

This directory contains metadata and assets for the LTEdu iOS app submission
to the Apple App Store.

## Metadata Files

| File | Max Length | Description |
|------|-----------|-------------|
| `metadata/*/name.txt` | 30 chars | App name |
| `metadata/*/subtitle.txt` | 30 chars | Short subtitle shown below the app name |
| `metadata/*/description.txt` | 4000 chars | Full store listing description |
| `metadata/*/keywords.txt` | 100 chars | Comma-separated search keywords |
| `metadata/*/release_notes.txt` | 4000 chars | What's new in this release |

## Asset Requirements

### App Icon
- **App Store icon**: 1024×1024 px, PNG, no alpha, no rounded corners (Apple applies rounding)
- Place at: `assets/icon-1024.png`

### Screenshots
- **Minimum**: 3 screenshots per required device
- **Maximum**: 10 screenshots per device
- **Required sizes**:
  - 6.7" iPhone (Super Retina XDR): 1290×2796 px or 2796×1290 px
  - 5.5" iPhone: 1242×2208 px or 2208×1242 px
  - 12.9" iPad Pro: 2048×2732 px or 2732×2048 px

Place screenshots at:
```
assets/screenshots/iphone-6.7/
assets/screenshots/iphone-5.5/
assets/screenshots/ipad-12.9/
```

## Supported Locales

| Directory | Language |
|-----------|----------|
| `en-US`   | English (United States) |
| `zh-CN`   | Simplified Chinese (China mainland) |

## Release Signing

Distribution certificates and provisioning profiles are stored securely in
GitHub Secrets. See the root [README.md](../README.md) for the list of
required secrets.

## Automated Workflow

See `.github/workflows/ios-publish.yml` for the automated build and
App Store submission workflow.

## App Store Connect Setup

1. Create an App ID in the [Apple Developer portal](https://developer.apple.com).
2. Create a distribution certificate and provisioning profile.
3. Register the app in [App Store Connect](https://appstoreconnect.apple.com).
4. Complete the app information, pricing, and availability sections.
5. Generate an App Store Connect API key and store the values as GitHub Secrets.
6. Provide a privacy policy URL in the app's App Store information.
