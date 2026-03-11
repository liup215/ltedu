<template>
  <div class="flex-1 bg-gray-50 min-h-screen">
    <!-- Header -->
    <section class="bg-gradient-to-r from-indigo-600 to-purple-600 text-white py-16">
      <div class="container mx-auto px-6 text-center">
        <h1 class="text-4xl font-normal mb-4">Download edu-cli</h1>
        <p class="text-xl text-indigo-100 mb-2">The command-line tool for the ltedu platform</p>
        <p class="text-indigo-200 text-sm">Manage questions, sync content, and integrate AI-powered education tools from your terminal</p>
      </div>
    </section>

    <div class="container mx-auto px-6 py-12 max-w-4xl">

      <!-- Quick Install -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-8 mb-8">
        <h2 class="text-2xl font-normal text-gray-800 mb-2">Quick Install</h2>
        <p class="text-gray-500 text-sm mb-6">Install edu-cli using Go (recommended)</p>
        <div class="bg-gray-900 rounded-lg p-4 text-sm font-mono text-green-400 mb-4">
          <span class="text-gray-500"># Install with Go</span><br />
          go install github.com/liup215/ltedu/cli@latest
        </div>
        <p class="text-gray-500 text-sm">Requires Go 1.21 or later.
          <a href="https://go.dev/dl/" target="_blank" rel="noopener noreferrer" class="text-indigo-600 hover:text-indigo-800 underline">Install Go →</a>
        </p>
      </div>

      <!-- Binary Downloads -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-8 mb-8">
        <h2 class="text-2xl font-normal text-gray-800 mb-2">Binary Downloads</h2>
        <p class="text-gray-500 text-sm mb-6">Download pre-built binaries for your platform from the GitHub Releases page</p>
        <a
          href="https://github.com/liup215/ltedu/releases"
          target="_blank"
          rel="noopener noreferrer"
          class="inline-flex items-center gap-2 px-6 py-3 bg-gray-900 text-white rounded-lg text-sm font-normal hover:bg-gray-700 transition"
        >
          <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
            <path d="M12 0C5.37 0 0 5.37 0 12c0 5.31 3.435 9.795 8.205 11.385.6.105.825-.255.825-.57 0-.285-.015-1.23-.015-2.235-3.015.555-3.795-.735-4.035-1.41-.135-.345-.72-1.41-1.23-1.695-.42-.225-1.02-.78-.015-.795.945-.015 1.62.87 1.845 1.23 1.08 1.815 2.805 1.305 3.495.99.105-.78.42-1.305.765-1.605-2.67-.3-5.46-1.335-5.46-5.925 0-1.305.465-2.385 1.23-3.225-.12-.3-.54-1.53.12-3.18 0 0 1.005-.315 3.3 1.23.96-.27 1.98-.405 3-.405s2.04.135 3 .405c2.295-1.56 3.3-1.23 3.3-1.23.66 1.65.24 2.88.12 3.18.765.84 1.23 1.905 1.23 3.225 0 4.605-2.805 5.625-5.475 5.925.435.375.81 1.095.81 2.22 0 1.605-.015 2.895-.015 3.3 0 .315.225.69.825.57A12.02 12.02 0 0 0 24 12c0-6.63-5.37-12-12-12z" />
          </svg>
          GitHub Releases
        </a>
        <div class="mt-6 grid md:grid-cols-3 gap-4">
          <div class="bg-gray-50 rounded-lg p-4 text-center border border-gray-200">
            <div class="text-2xl mb-2">🐧</div>
            <div class="font-normal text-gray-800 text-sm">Linux</div>
            <div class="text-gray-500 text-xs mt-1">amd64 / arm64</div>
          </div>
          <div class="bg-gray-50 rounded-lg p-4 text-center border border-gray-200">
            <div class="text-2xl mb-2">🍎</div>
            <div class="font-normal text-gray-800 text-sm">macOS</div>
            <div class="text-gray-500 text-xs mt-1">Intel / Apple Silicon</div>
          </div>
          <div class="bg-gray-50 rounded-lg p-4 text-center border border-gray-200">
            <div class="text-2xl mb-2">🪟</div>
            <div class="font-normal text-gray-800 text-sm">Windows</div>
            <div class="text-gray-500 text-xs mt-1">amd64</div>
          </div>
        </div>
      </div>

      <!-- Getting Started -->
      <div class="bg-white rounded-lg shadow-sm border border-gray-200 p-8 mb-8">
        <h2 class="text-2xl font-normal text-gray-800 mb-6">Getting Started</h2>
        <ol class="space-y-6">
          <li class="flex gap-4">
            <div class="flex-shrink-0 w-8 h-8 bg-indigo-100 rounded-full flex items-center justify-center text-indigo-600 font-normal text-sm">1</div>
            <div>
              <h3 class="font-normal text-gray-800 mb-1">Configure the backend URL</h3>
              <div class="bg-gray-900 rounded p-3 text-xs font-mono text-green-400">
                edu-cli config set-url https://www.alevel.icu
              </div>
            </div>
          </li>
          <li class="flex gap-4">
            <div class="flex-shrink-0 w-8 h-8 bg-indigo-100 rounded-full flex items-center justify-center text-indigo-600 font-normal text-sm">2</div>
            <div>
              <h3 class="font-normal text-gray-800 mb-1">Authenticate with your CLI token</h3>
              <div class="bg-gray-900 rounded p-3 text-xs font-mono text-green-400">
                edu-cli login --token &lt;your-cli-token&gt;
              </div>
              <p class="text-gray-500 text-xs mt-2">
                Get your CLI token from
                <router-link to="/account/cli-tokens" class="text-indigo-600 hover:text-indigo-800 underline">Personal Center → CLI Tokens</router-link>
              </p>
            </div>
          </li>
          <li class="flex gap-4">
            <div class="flex-shrink-0 w-8 h-8 bg-indigo-100 rounded-full flex items-center justify-center text-indigo-600 font-normal text-sm">3</div>
            <div>
              <h3 class="font-normal text-gray-800 mb-1">Start using edu-cli</h3>
              <div class="bg-gray-900 rounded p-3 text-xs font-mono text-green-400">
                <div><span class="text-gray-500"># Sync questions from a CSV file</span></div>
                <div>edu-cli question sync --file questions.csv</div>
                <div class="mt-2"><span class="text-gray-500"># View available commands</span></div>
                <div>edu-cli --help</div>
              </div>
            </div>
          </li>
        </ol>
      </div>

      <!-- CTA -->
      <div class="bg-indigo-50 rounded-lg border border-indigo-200 p-8 text-center">
        <h3 class="text-lg font-normal text-indigo-800 mb-2">Need your API token?</h3>
        <p class="text-indigo-700 text-sm mb-4">Visit your Personal Center to generate CLI tokens and manage API access.</p>
        <router-link
          to="/account/cli-tokens"
          class="inline-flex items-center px-6 py-3 bg-indigo-600 text-white rounded-lg text-sm hover:bg-indigo-700 transition"
        >
          Get CLI Token →
        </router-link>
      </div>

    </div>
  </div>
</template>

<script setup lang="ts">
// No additional logic needed for this static page
</script>
