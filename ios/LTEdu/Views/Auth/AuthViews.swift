import SwiftUI

// MARK: - LoginView

struct LoginView: View {
    @EnvironmentObject var authViewModel: AuthViewModel
    @State private var username = ""
    @State private var password = ""
    @State private var captchaValue = ""
    @State private var showPassword = false
    @State private var showRegister = false

    var body: some View {
        NavigationStack {
            ScrollView {
                VStack(spacing: 32) {
                    // Logo & title
                    headerSection

                    // Form
                    VStack(spacing: 16) {
                        // Username
                        LTEduTextField(
                            text: $username,
                            placeholder: "Username or Email",
                            icon: "person"
                        )

                        // Password
                        LTEduSecureField(
                            text: $password,
                            placeholder: "Password",
                            showPassword: $showPassword
                        )

                        // Captcha
                        captchaSection

                        // Login button
                        Button(action: {
                            Task { await authViewModel.login(username: username, password: password, captchaValue: captchaValue) }
                        }) {
                            Group {
                                if authViewModel.isLoading {
                                    ProgressView()
                                        .tint(.white)
                                } else {
                                    Text("Sign In")
                                        .fontWeight(.semibold)
                                }
                            }
                            .frame(maxWidth: .infinity)
                            .frame(height: 50)
                            .background(Color.blue)
                            .foregroundColor(.white)
                            .cornerRadius(12)
                        }
                        .disabled(authViewModel.isLoading || username.isEmpty || password.isEmpty || captchaValue.isEmpty)
                    }
                    .padding(.horizontal)

                    // Error message
                    if let error = authViewModel.errorMessage {
                        Text(error)
                            .font(.caption)
                            .foregroundColor(.red)
                            .multilineTextAlignment(.center)
                            .padding(.horizontal)
                    }

                    // Register link
                    Button("Don't have an account? Register") {
                        showRegister = true
                    }
                    .font(.footnote)
                    .foregroundColor(.blue)
                }
                .padding(.vertical, 40)
            }
            .background(Color(.systemGroupedBackground))
            .sheet(isPresented: $showRegister) {
                RegisterView()
                    .environmentObject(authViewModel)
            }
        }
        .task {
            await authViewModel.loadCaptcha()
        }
    }

    // MARK: - Header

    private var headerSection: some View {
        VStack(spacing: 12) {
            Image(systemName: "graduationcap.fill")
                .font(.system(size: 60))
                .foregroundColor(.blue)
            Text("LTEdu")
                .font(.largeTitle)
                .fontWeight(.bold)
            Text("AI-Powered Education Platform")
                .font(.subheadline)
                .foregroundColor(.secondary)
        }
        .padding(.top, 20)
    }

    // MARK: - Captcha Section

    private var captchaSection: some View {
        HStack(spacing: 12) {
            LTEduTextField(
                text: $captchaValue,
                placeholder: "Enter captcha",
                icon: "shield"
            )

            // Captcha image
            Group {
                if !authViewModel.captchaBase64.isEmpty,
                   let imageData = Data(base64Encoded: authViewModel.captchaBase64.components(separatedBy: ",").last ?? ""),
                   let uiImage = UIImage(data: imageData) {
                    Image(uiImage: uiImage)
                        .resizable()
                        .scaledToFit()
                        .frame(width: 100, height: 44)
                        .cornerRadius(8)
                        .onTapGesture {
                            Task { await authViewModel.loadCaptcha() }
                        }
                } else {
                    Button(action: {
                        Task { await authViewModel.loadCaptcha() }
                    }) {
                        Text("Refresh")
                            .font(.caption)
                            .frame(width: 100, height: 44)
                            .background(Color.blue.opacity(0.1))
                            .cornerRadius(8)
                    }
                }
            }
        }
    }
}

// MARK: - RegisterView

struct RegisterView: View {
    @EnvironmentObject var authViewModel: AuthViewModel
    @Environment(\.dismiss) var dismiss

    @State private var username = ""
    @State private var email = ""
    @State private var mobile = ""
    @State private var password = ""
    @State private var passwordConfirm = ""
    @State private var verificationCode = ""
    @State private var showPassword = false
    @State private var codeSent = false
    @State private var successMessage: String?

    var body: some View {
        NavigationStack {
            ScrollView {
                VStack(spacing: 20) {
                    Text("Create Account")
                        .font(.title2)
                        .fontWeight(.bold)
                        .padding(.top)

                    VStack(spacing: 14) {
                        LTEduTextField(text: $username, placeholder: "Username", icon: "person")
                        LTEduTextField(text: $email, placeholder: "Email", icon: "envelope")
                            .keyboardType(.emailAddress)
                            .autocapitalization(.none)
                        LTEduTextField(text: $mobile, placeholder: "Mobile (optional)", icon: "phone")
                            .keyboardType(.phonePad)

                        // Verification code
                        HStack(spacing: 10) {
                            LTEduTextField(text: $verificationCode, placeholder: "Verification Code", icon: "lock.shield")
                                .keyboardType(.numberPad)
                            Button(codeSent ? "Resend" : "Send Code") {
                                Task {
                                    await authViewModel.sendVerificationCode(email: email)
                                    codeSent = true
                                }
                            }
                            .font(.caption)
                            .frame(width: 90, height: 44)
                            .background(Color.blue)
                            .foregroundColor(.white)
                            .cornerRadius(10)
                            .disabled(email.isEmpty)
                        }

                        LTEduSecureField(text: $password, placeholder: "Password (min 6 chars)", showPassword: $showPassword)
                        LTEduSecureField(text: $passwordConfirm, placeholder: "Confirm Password", showPassword: $showPassword)
                    }
                    .padding(.horizontal)

                    if let error = authViewModel.errorMessage {
                        Text(error)
                            .font(.caption)
                            .foregroundColor(.red)
                            .multilineTextAlignment(.center)
                            .padding(.horizontal)
                    }

                    if let success = successMessage {
                        Text(success)
                            .font(.caption)
                            .foregroundColor(.green)
                            .padding(.horizontal)
                    }

                    Button(action: registerTapped) {
                        Group {
                            if authViewModel.isLoading {
                                ProgressView().tint(.white)
                            } else {
                                Text("Create Account").fontWeight(.semibold)
                            }
                        }
                        .frame(maxWidth: .infinity)
                        .frame(height: 50)
                        .background(Color.blue)
                        .foregroundColor(.white)
                        .cornerRadius(12)
                        .padding(.horizontal)
                    }
                    .disabled(authViewModel.isLoading)
                }
                .padding(.bottom, 40)
            }
            .background(Color(.systemGroupedBackground))
            .navigationBarTitleDisplayMode(.inline)
            .toolbar {
                ToolbarItem(placement: .navigationBarLeading) {
                    Button("Cancel") { dismiss() }
                }
            }
        }
    }

    private func registerTapped() {
        Task {
            await authViewModel.register(
                username: username,
                email: email,
                mobile: mobile.isEmpty ? nil : mobile,
                password: password,
                passwordConfirm: passwordConfirm,
                verificationCode: verificationCode
            )
            if authViewModel.errorMessage == nil {
                successMessage = "Account created! Please sign in."
                DispatchQueue.main.asyncAfter(deadline: .now() + 2) {
                    dismiss()
                }
            }
        }
    }
}

// MARK: - Shared Input Components

struct LTEduTextField: View {
    @Binding var text: String
    let placeholder: String
    let icon: String

    var body: some View {
        HStack(spacing: 12) {
            Image(systemName: icon)
                .foregroundColor(.secondary)
                .frame(width: 20)
            TextField(placeholder, text: $text)
                .autocapitalization(.none)
                .disableAutocorrection(true)
        }
        .padding()
        .background(Color(.systemBackground))
        .cornerRadius(10)
        .overlay(
            RoundedRectangle(cornerRadius: 10)
                .stroke(Color(.systemGray4), lineWidth: 1)
        )
    }
}

struct LTEduSecureField: View {
    @Binding var text: String
    let placeholder: String
    @Binding var showPassword: Bool

    var body: some View {
        HStack(spacing: 12) {
            Image(systemName: "lock")
                .foregroundColor(.secondary)
                .frame(width: 20)
            if showPassword {
                TextField(placeholder, text: $text)
                    .autocapitalization(.none)
            } else {
                SecureField(placeholder, text: $text)
            }
            Button(action: { showPassword.toggle() }) {
                Image(systemName: showPassword ? "eye.slash" : "eye")
                    .foregroundColor(.secondary)
            }
        }
        .padding()
        .background(Color(.systemBackground))
        .cornerRadius(10)
        .overlay(
            RoundedRectangle(cornerRadius: 10)
                .stroke(Color(.systemGray4), lineWidth: 1)
        )
    }
}
