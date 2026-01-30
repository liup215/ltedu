# 阿里云 OSS 配置指南

## 1. 获取 AccessKey ID 和 AccessKey Secret

1. 登录 [阿里云控制台](https://home.console.aliyun.com/)。
2. 将鼠标悬停在右上角的头像上，点击 **AccessKey 管理**。
3. 选择 **开始使用子用户 AccessKey** (推荐) 或 **继续使用主账号 AccessKey**。
   - **推荐方式**: 创建一个 RAM 子用户，并授予 `AliyunOSSFullAccess` 权限，然后为该子用户创建 AccessKey。这比使用主账号更安全。
4. 创建成功后，复制保存 `AccessKey ID` 和 `AccessKey Secret`。

## 2. 获取 Bucket Name (存储空间名称)

1. 进入 [OSS 管理控制台](https://oss.console.aliyun.com/)。
2. 点击左侧导航栏的 **Bucket 列表**。
3. 点击 **创建 Bucket** (如果还没有) 或选择现有的 Bucket。
4. **Bucket 名称** 即为配置中的 `Bucket Name`。
5. **读写权限** 建议设置为 **公共读** (Public Read)，以便图片可以被公开访问。

## 3. 获取 Endpoint (地域节点)

1. 在 Bucket 列表中点击您的 Bucket 名称进入概览页。
2. 在 **概览** 页面的 **访问端口** 区域，找到 **Endpoint（地域节点）**。
   - 示例: `oss-cn-hangzhou.aliyuncs.com`
   - 不包含 `http://` 或 `https://` 前缀。

## 4. 获取 CDN URL (可选)

如果您配置了 CDN 加速或绑定了自定义域名：
1. 在 Bucket 概览页的 **域名管理** 中可以绑定自定义域名。
2. 配置 CNAME 指向。
3. 配置成功后，填写您的自定义域名作为 `CDN URL` (包含 `http://` 或 `https://`)。
   - 示例: `https://cdn.example.com`
   - 如果未配置自定义域名，可以留空，系统将使用默认的 OSS 域名。

## 配置示例

- **Access Key ID**: `LTAI5t7...`
- **Access Key Secret**: `F3gS8...`
- **Bucket Name**: `ltedu-assets`
- **Endpoint**: `oss-cn-shanghai.aliyuncs.com`
- **CDN URL**: `https://assets.ltedu.com` (或留空)