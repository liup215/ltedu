package com.ltedu.app.ui.theme

import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.darkColorScheme
import androidx.compose.material3.lightColorScheme
import androidx.compose.runtime.Composable
import androidx.compose.ui.graphics.Color

private val LTEduBlue = Color(0xFF1565C0)
private val LTEduBlueLight = Color(0xFF5E92F3)
private val LTEduBlueDark = Color(0xFF003C8F)

private val LightColorScheme = lightColorScheme(
    primary = LTEduBlue,
    onPrimary = Color.White,
    primaryContainer = LTEduBlueLight,
    secondary = Color(0xFF03DAC5),
    background = Color(0xFFF5F5F5),
    surface = Color.White,
    error = Color(0xFFB00020)
)

private val DarkColorScheme = darkColorScheme(
    primary = LTEduBlueLight,
    onPrimary = Color.Black,
    primaryContainer = LTEduBlueDark,
    secondary = Color(0xFF03DAC5),
    background = Color(0xFF121212),
    surface = Color(0xFF1E1E1E),
    error = Color(0xFFCF6679)
)

@Composable
fun LTEduTheme(
    darkTheme: Boolean = false,
    content: @Composable () -> Unit
) {
    val colorScheme = if (darkTheme) DarkColorScheme else LightColorScheme

    MaterialTheme(
        colorScheme = colorScheme,
        content = content
    )
}
