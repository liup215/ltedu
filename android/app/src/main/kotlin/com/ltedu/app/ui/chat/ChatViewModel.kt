package com.ltedu.app.ui.chat

import androidx.lifecycle.SavedStateHandle
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.ltedu.app.data.db.entity.ChatMessageEntity
import com.ltedu.app.data.repository.ChatRepository
import com.ltedu.app.data.repository.Result
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.*
import kotlinx.coroutines.launch
import javax.inject.Inject

data class ChatUiState(
    val isSending: Boolean = false,
    val error: String? = null
)

@HiltViewModel
class ChatViewModel @Inject constructor(
    private val chatRepository: ChatRepository,
    savedStateHandle: SavedStateHandle
) : ViewModel() {

    val sessionId: String = checkNotNull(savedStateHandle["sessionId"])

    val messages: StateFlow<List<ChatMessageEntity>> =
        chatRepository.getMessages(sessionId)
            .stateIn(viewModelScope, SharingStarted.WhileSubscribed(5_000), emptyList())

    private val _uiState = MutableStateFlow(ChatUiState())
    val uiState: StateFlow<ChatUiState> = _uiState.asStateFlow()

    fun send(text: String) {
        if (text.isBlank() || _uiState.value.isSending) return
        viewModelScope.launch {
            _uiState.value = ChatUiState(isSending = true)
            when (val result = chatRepository.sendMessage(sessionId, text.trim())) {
                is Result.Success -> _uiState.value = ChatUiState(isSending = false)
                is Result.Error -> _uiState.value = ChatUiState(
                    isSending = false,
                    error = result.message
                )
            }
        }
    }

    fun clearError() {
        _uiState.value = _uiState.value.copy(error = null)
    }

    fun clearSession() {
        viewModelScope.launch {
            chatRepository.clearSession(sessionId)
        }
    }
}
