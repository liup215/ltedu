# App Store Submissions

This directory contains all assets, metadata, and documentation required for
submitting and maintaining LTEdu applications across all major app stores.

## Supported Platforms

| Platform | App Store | Status |
|----------|-----------|--------|
| Android  | Google Play Store | 🔄 Pending submission |
| iOS      | Apple App Store   | 🔄 Pending submission |
| HarmonyOS | Huawei AppGallery | 🔄 Pending submission |

## Directory Structure

```
app-store/
├── android/          # Google Play Store assets and metadata
│   ├── README.md
│   └── metadata/
│       ├── en-US/    # English metadata
│       └── zh-CN/    # Simplified Chinese metadata
├── ios/              # Apple App Store assets and metadata
│   ├── README.md
│   └── metadata/
│       ├── en-US/    # English metadata
│       └── zh-CN/    # Simplified Chinese metadata
└── huawei/           # Huawei AppGallery assets and metadata
    ├── README.md
    └── metadata/
        ├── en-US/    # English metadata
        └── zh-CN/    # Simplified Chinese metadata
```

## Automated Publishing

GitHub Actions workflows handle automated builds and store submissions:

| Workflow | File | Trigger |
|----------|------|---------|
| Android Publish | `.github/workflows/android-publish.yml` | Tag `v*` or manual |
| iOS Publish     | `.github/workflows/ios-publish.yml`     | Tag `v*` or manual |
| Huawei Publish  | `.github/workflows/huawei-publish.yml`  | Tag `v*` or manual |

## Required Secrets

Configure the following secrets in your GitHub repository settings before
running any publish workflow:

### Android
| Secret | Description |
|--------|-------------|
| `ANDROID_KEYSTORE_BASE64` | Base64-encoded release keystore file |
| `ANDROID_KEY_ALIAS` | Key alias within the keystore |
| `ANDROID_KEY_PASSWORD` | Password for the key |
| `ANDROID_STORE_PASSWORD` | Password for the keystore |
| `GOOGLE_PLAY_SERVICE_ACCOUNT_JSON` | Google Play API service account JSON |

### iOS
| Secret | Description |
|--------|-------------|
| `IOS_CERTIFICATE_BASE64` | Base64-encoded Apple distribution certificate (.p12) |
| `IOS_CERTIFICATE_PASSWORD` | Password for the distribution certificate |
| `IOS_PROVISIONING_PROFILE_BASE64` | Base64-encoded provisioning profile |
| `APP_STORE_CONNECT_API_KEY_ID` | App Store Connect API key ID |
| `APP_STORE_CONNECT_API_ISSUER_ID` | App Store Connect API issuer ID |
| `APP_STORE_CONNECT_API_KEY_BASE64` | Base64-encoded App Store Connect API private key (.p8) |

### Huawei AppGallery
| Secret | Description |
|--------|-------------|
| `HUAWEI_KEYSTORE_BASE64` | Base64-encoded release keystore file |
| `HUAWEI_KEY_ALIAS` | Key alias within the keystore |
| `HUAWEI_KEY_PASSWORD` | Password for the key |
| `HUAWEI_STORE_PASSWORD` | Password for the keystore |
| `HUAWEI_CLIENT_ID` | AppGallery Connect API client ID |
| `HUAWEI_CLIENT_SECRET` | AppGallery Connect API client secret |
| `HUAWEI_APP_ID` | App ID from AppGallery Connect |

## Asset Requirements

### Icons
| Platform | Requirement |
|----------|-------------|
| Android  | 512×512 px PNG (Play Store), adaptive icon layers |
| iOS      | 1024×1024 px PNG (App Store), multiple sizes for device |
| Huawei   | 216×216 px PNG (AppGallery) |

### Screenshots
| Platform | Requirement |
|----------|-------------|
| Android  | 2–8 screenshots, JPEG/PNG, min 320px on shortest side |
| iOS      | 3–10 screenshots per device size (6.7", 5.5", iPad) |
| Huawei   | 3–5 screenshots, JPEG/PNG, recommended 1080×1920 px |

Place screenshots in a `screenshots/` subdirectory within each platform folder.

## Compliance Checklist

- [ ] Privacy policy URL configured in each store listing
- [ ] Age rating / content rating declared on all platforms
- [ ] Required permissions justified in store listings
- [ ] Data safety / privacy nutrition labels completed
- [ ] Export compliance (encryption) declarations submitted
- [ ] All app store guidelines reviewed and confirmed

## Update Process

1. Increment the version in the mobile app source.
2. Push a version tag (e.g. `v1.2.0`) to trigger the publish workflows.
3. Workflows build, sign, and submit to each store automatically.
4. Monitor store review queues and respond promptly to reviewer feedback.
5. Track user reviews and ratings via each store's analytics dashboard.
