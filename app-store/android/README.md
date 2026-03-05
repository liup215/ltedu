# Android – Google Play Store

This directory contains metadata and assets for the LTEdu Android app submission
to the Google Play Store.

## Metadata Files

| File | Max Length | Description |
|------|-----------|-------------|
| `metadata/*/title.txt` | 50 chars | App title |
| `metadata/*/short_description.txt` | 80 chars | Short description shown in search results |
| `metadata/*/full_description.txt` | 4000 chars | Full store listing description |
| `metadata/*/changelogs/default.txt` | 500 chars | What's new in this release |

## Asset Requirements

### App Icon
- **High-resolution icon**: 512×512 px, PNG, 32-bit color, no alpha
- Place at: `assets/icon-512.png`

### Feature Graphic
- Size: 1024×500 px, JPEG or 24-bit PNG, no alpha
- Place at: `assets/feature-graphic.png`

### Screenshots
- **Minimum**: 2 screenshots per device type
- **Maximum**: 8 screenshots per device type
- **Format**: JPEG or 24-bit PNG, no alpha
- **Size**: Each side between 320 px and 3840 px; longest side ≤ 2× shortest side
- Device types: Phone, 7-inch tablet, 10-inch tablet

Place screenshots at:
```
assets/screenshots/phone/
assets/screenshots/tablet-7/
assets/screenshots/tablet-10/
```

## Supported Locales

| Directory | Language |
|-----------|----------|
| `en-US`   | English (United States) |
| `zh-CN`   | Simplified Chinese |

## Release Signing

The release keystore is stored securely in GitHub Secrets. See the root
[README.md](../README.md) for the list of required secrets.

## Automated Workflow

See `.github/workflows/android-publish.yml` for the automated build and
Play Store submission workflow.

## Google Play Console Setup

1. Create an app in [Google Play Console](https://play.google.com/console).
2. Complete the store listing using the metadata files in this directory.
3. Set up content rating questionnaire.
4. Configure data safety section.
5. Create a service account and download the JSON key for CI/CD.
6. Store the JSON key as the `GOOGLE_PLAY_SERVICE_ACCOUNT_JSON` secret.
