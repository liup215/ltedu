import SwiftUI

// MARK: - AIChatView

struct AIChatView: View {
    @EnvironmentObject var chatVM: AIChatViewModel
    @State private var showSessions = false

    var body: some View {
        NavigationStack {
            Group {
                if let session = chatVM.currentSession {
                    ChatConversationView(session: session)
                        .environmentObject(chatVM)
                } else {
                    emptyChatView
                }
            }
            .navigationTitle("AI Assistant")
            .toolbar {
                ToolbarItem(placement: .navigationBarLeading) {
                    Button(action: { showSessions.toggle() }) {
                        Image(systemName: "list.bullet")
                    }
                }
                ToolbarItem(placement: .navigationBarTrailing) {
                    Button(action: { chatVM.startNewSession() }) {
                        Image(systemName: "square.and.pencil")
                    }
                }
            }
            .sheet(isPresented: $showSessions) {
                ChatSessionsListView()
                    .environmentObject(chatVM)
            }
            .onAppear {
                if chatVM.currentSession == nil && !chatVM.sessions.isEmpty {
                    chatVM.selectSession(chatVM.sessions[0])
                }
            }
        }
    }

    private var emptyChatView: some View {
        VStack(spacing: 20) {
            Image(systemName: "brain.head.profile")
                .font(.system(size: 60))
                .foregroundColor(.blue.opacity(0.7))

            Text("AI Learning Assistant")
                .font(.title2)
                .fontWeight(.bold)

            Text("Ask me about study plans, practice questions, or get personalized learning recommendations.")
                .font(.subheadline)
                .foregroundColor(.secondary)
                .multilineTextAlignment(.center)
                .padding(.horizontal, 40)

            // Suggestion chips
            VStack(spacing: 10) {
                suggestionChip("Create a study plan for me")
                suggestionChip("Give me practice questions")
                suggestionChip("Analyze my progress")
                suggestionChip("What should I review today?")
            }

            Button("Start New Chat") {
                chatVM.startNewSession()
            }
            .font(.headline)
            .frame(maxWidth: .infinity)
            .frame(height: 50)
            .background(Color.blue)
            .foregroundColor(.white)
            .cornerRadius(12)
            .padding(.horizontal, 40)
        }
        .padding()
    }

    private func suggestionChip(_ text: String) -> some View {
        Button(action: {
            chatVM.startNewSession()
            Task { await chatVM.sendMessage(text) }
        }) {
            Text(text)
                .font(.subheadline)
                .padding(.horizontal, 16)
                .padding(.vertical, 10)
                .background(Color.blue.opacity(0.1))
                .foregroundColor(.blue)
                .cornerRadius(20)
        }
    }
}

// MARK: - ChatConversationView

struct ChatConversationView: View {
    @EnvironmentObject var chatVM: AIChatViewModel
    let session: ChatSession

    @State private var scrollProxy: ScrollViewProxy?

    var body: some View {
        VStack(spacing: 0) {
            // Messages
            ScrollViewReader { proxy in
                ScrollView {
                    LazyVStack(spacing: 12) {
                        if chatVM.currentSession?.messages.isEmpty == true {
                            welcomeMessage
                        }
                        ForEach(chatVM.currentSession?.messages ?? []) { message in
                            MessageBubbleView(message: message)
                                .id(message.id)
                        }
                        if chatVM.isGenerating {
                            TypingIndicatorView()
                                .id("typing")
                        }
                    }
                    .padding()
                }
                .onAppear { scrollProxy = proxy }
                .onChange(of: chatVM.currentSession?.messages.count) { _ in
                    if let lastId = chatVM.currentSession?.messages.last?.id {
                        withAnimation { proxy.scrollTo(lastId, anchor: .bottom) }
                    }
                }
                .onChange(of: chatVM.isGenerating) { generating in
                    if generating {
                        withAnimation { proxy.scrollTo("typing", anchor: .bottom) }
                    }
                }
            }

            Divider()

            // Input bar
            ChatInputBar()
                .environmentObject(chatVM)
        }
    }

    private var welcomeMessage: some View {
        VStack(spacing: 8) {
            Image(systemName: "brain.head.profile")
                .font(.system(size: 40))
                .foregroundColor(.blue.opacity(0.6))
            Text("How can I help you today?")
                .font(.headline)
                .foregroundColor(.secondary)
        }
        .padding(.vertical, 40)
    }
}

// MARK: - MessageBubbleView

struct MessageBubbleView: View {
    let message: ChatMessage

    var isUser: Bool { message.role == .user }

    var body: some View {
        HStack(alignment: .bottom, spacing: 8) {
            if isUser { Spacer(minLength: 60) }

            if !isUser {
                Image(systemName: "brain.head.profile")
                    .font(.caption)
                    .padding(6)
                    .background(Color.blue.opacity(0.1))
                    .clipShape(Circle())
            }

            VStack(alignment: isUser ? .trailing : .leading, spacing: 4) {
                Text(message.content)
                    .padding(.horizontal, 14)
                    .padding(.vertical, 10)
                    .background(isUser ? Color.blue : Color(.systemGray6))
                    .foregroundColor(isUser ? .white : .primary)
                    .cornerRadius(18)
                    .cornerRadius(isUser ? 4 : 18, corners: isUser ? .topRight : .topLeft)

                Text(message.timestamp, style: .time)
                    .font(.caption2)
                    .foregroundColor(.secondary)
            }

            if !isUser { Spacer(minLength: 60) }

            if isUser {
                Image(systemName: "person.circle.fill")
                    .font(.caption)
                    .foregroundColor(.blue)
            }
        }
    }
}

// MARK: - TypingIndicatorView

struct TypingIndicatorView: View {
    @State private var animating = false

    var body: some View {
        HStack(alignment: .bottom, spacing: 8) {
            Image(systemName: "brain.head.profile")
                .font(.caption)
                .padding(6)
                .background(Color.blue.opacity(0.1))
                .clipShape(Circle())

            HStack(spacing: 4) {
                ForEach(0..<3) { i in
                    Circle()
                        .fill(Color.secondary)
                        .frame(width: 6, height: 6)
                        .scaleEffect(animating ? 1.2 : 0.8)
                        .animation(
                            .easeInOut(duration: 0.5)
                                .repeatForever()
                                .delay(Double(i) * 0.15),
                            value: animating
                        )
                }
            }
            .padding(.horizontal, 14)
            .padding(.vertical, 12)
            .background(Color(.systemGray6))
            .cornerRadius(18)

            Spacer(minLength: 60)
        }
        .onAppear { animating = true }
    }
}

// MARK: - ChatInputBar

struct ChatInputBar: View {
    @EnvironmentObject var chatVM: AIChatViewModel
    @FocusState private var isFocused: Bool

    var body: some View {
        HStack(spacing: 12) {
            TextField("Message AI assistant...", text: $chatVM.inputText, axis: .vertical)
                .lineLimit(1...5)
                .padding(.horizontal, 14)
                .padding(.vertical, 10)
                .background(Color(.systemGray6))
                .cornerRadius(20)
                .focused($isFocused)
                .onSubmit {
                    sendMessage()
                }

            Button(action: sendMessage) {
                Image(systemName: chatVM.isGenerating ? "stop.circle.fill" : "arrow.up.circle.fill")
                    .font(.system(size: 30))
                    .foregroundColor(canSend ? .blue : .gray)
            }
            .disabled(!canSend)
        }
        .padding(.horizontal)
        .padding(.vertical, 8)
        .background(Color(.systemBackground))
    }

    private var canSend: Bool {
        !chatVM.inputText.trimmingCharacters(in: .whitespacesAndNewlines).isEmpty && !chatVM.isGenerating
    }

    private func sendMessage() {
        let text = chatVM.inputText
        Task { await chatVM.sendMessage(text) }
    }
}

// MARK: - ChatSessionsListView

struct ChatSessionsListView: View {
    @EnvironmentObject var chatVM: AIChatViewModel
    @Environment(\.dismiss) var dismiss

    var body: some View {
        NavigationStack {
            List {
                ForEach(chatVM.sessions) { session in
                    Button(action: {
                        chatVM.selectSession(session)
                        dismiss()
                    }) {
                        VStack(alignment: .leading, spacing: 4) {
                            Text(session.title)
                                .font(.headline)
                                .lineLimit(1)
                            Text(session.updatedAt, style: .relative)
                                .font(.caption)
                                .foregroundColor(.secondary)
                        }
                    }
                    .swipeActions {
                        Button(role: .destructive) {
                            chatVM.deleteSession(session)
                        } label: {
                            Label("Delete", systemImage: "trash")
                        }
                    }
                }
            }
            .navigationTitle("Chat History")
            .toolbar {
                ToolbarItem(placement: .navigationBarTrailing) {
                    Button("Done") { dismiss() }
                }
                ToolbarItem(placement: .navigationBarLeading) {
                    Button("New Chat") {
                        chatVM.startNewSession()
                        dismiss()
                    }
                }
            }
        }
    }
}

// MARK: - Corner Radius Helper

extension View {
    func cornerRadius(_ radius: CGFloat, corners: UIRectCorner) -> some View {
        clipShape(RoundedCorner(radius: radius, corners: corners))
    }
}

private struct RoundedCorner: Shape {
    var radius: CGFloat
    var corners: UIRectCorner

    func path(in rect: CGRect) -> Path {
        let path = UIBezierPath(
            roundedRect: rect,
            byRoundingCorners: corners,
            cornerRadii: CGSize(width: radius, height: radius)
        )
        return Path(path.cgPath)
    }
}
