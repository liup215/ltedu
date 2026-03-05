# Huawei AppGallery – HarmonyOS

This directory contains metadata and assets for the LTEdu HarmonyOS app
submission to Huawei AppGallery.

## Metadata Files

| File | Max Length | Description |
|------|-----------|-------------|
| `metadata/*/appName.txt` | 64 chars | App name |
| `metadata/*/briefInfo.txt` | 80 chars | Brief introduction shown in search results |
| `metadata/*/appDesc.txt` | 8000 chars | Full app description |
| `metadata/*/newFeatures.txt` | 500 chars | New features in this version |

## Asset Requirements

### App Icon
- **AppGallery icon**: 216×216 px, PNG, transparent background allowed
- Place at: `assets/icon-216.png`

### Feature Graphic
- Size: 1080×504 px, JPEG or PNG
- Place at: `assets/feature-graphic.png`

### Screenshots
- **Minimum**: 3 screenshots
- **Maximum**: 5 screenshots
- **Recommended size**: 1080×1920 px (portrait) or 1920×1080 px (landscape)
- **Format**: JPEG or PNG

Place screenshots at:
```
assets/screenshots/
```

## Supported Locales

| Directory | Language |
|-----------|----------|
| `en-US`   | English |
| `zh-CN`   | Simplified Chinese |

## Release Signing

The release keystore is stored securely in GitHub Secrets. See the root
[README.md](../README.md) for the list of required secrets.

## Automated Workflow

See `.github/workflows/huawei-publish.yml` for the automated build and
AppGallery submission workflow.

## AppGallery Connect Setup

1. Register as a Huawei developer at [AppGallery Connect](https://developer.huawei.com/consumer/en/appgallery).
2. Create an app and obtain the App ID.
3. Complete the store listing using the metadata files in this directory.
4. Create an API client in AppGallery Connect and obtain the Client ID and Client Secret.
5. Store the credentials as GitHub Secrets (`HUAWEI_CLIENT_ID`, `HUAWEI_CLIENT_SECRET`, `HUAWEI_APP_ID`).
6. Configure HarmonyOS SDK signing with the keystore secrets.
