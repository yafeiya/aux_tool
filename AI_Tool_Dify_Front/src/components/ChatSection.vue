<template>
  <div class="chat-section">
    <div class="chat-messages" ref="messagesContainer">
      <div
          v-for="(message, index) in parsedMessages"
          :key="index"
          :class="['message', message.isUser ? 'user-message' : 'ai-message']"
      >
        <div v-if="message.isUser" class="message-content">
          {{ message.text }}
        </div>

        <div v-else>
          <div v-if="message.thinking" class="thinking-section">
            <div class="thinking-header" @click="toggleThinking(index)">
              <span class="thinking-icon" :class="{ expanded: expandedThinking[index] }">▶</span>
              <span>思考过程</span>
            </div>
            <div v-show="expandedThinking[index]" class="thinking-content">
              <div v-html="formatMarkdown(message.thinking)"></div>
            </div>
          </div>

          <div class="message-content" v-html="formatMarkdown(message.text)"></div>
        </div>
      </div>

      <!-- 加载中提示 -->
      <div v-if="loading" class="loading-message ai-message">
        <em>
          <span class="spinner"></span>
          AI 正在思考...
        </em>
      </div>
    </div>

    <div class="chat-input-container">
      <input
          v-model="userInput"
          placeholder="请输入..."
          @keyup.enter="sendMessage"
          class="chat-input"
      />
      <button class="send-btn" @click="sendMessage">发送</button>
    </div>
  </div>
</template>

<script>
import { mapState } from "vuex";
import { marked } from "marked";

export default {
  name: "ChatSection",
  data() {
    return {
      userInput: "",
      expandedThinking: {},
      loading: false,           // 控制加载提示
      lastUserMsgCount: 0,      // 记录发送时的用户消息数
    };
  },
  computed: {
    ...mapState(["chatHistory"]),

    parsedMessages() {
      return this.chatHistory.map((msg) => {
        let rawText = "";
        let isUser = true;

        if (typeof msg === "string") {
          rawText = msg;
          isUser = true;
        } else if (msg && typeof msg === "object") {
          rawText = msg.text || "";
          isUser = msg.isUser !== undefined ? msg.isUser : true;
        }

        const thinkRegex = /<think>([\s\S]*?)<\/think>/i;
        const match = rawText.match(thinkRegex);
        let thinking = null;
        let text = rawText;

        if (match) {
          thinking = match[1].trim();
          text = rawText.replace(thinkRegex, "").trim();
        }

        return {
          text,
          thinking,
          isUser,
          rawText,
        };
      });
    },
  },
  methods: {
    async sendMessage() {
      if (!this.userInput.trim()) return;

      this.loading = true;

      const input = this.userInput.trim();  // 先缓存输入
      this.userInput = "";                   // 立即清空输入框

      try {
        await this.$store.dispatch("sendMessage", input);
      } catch (e) {
        console.error(e);
      } finally {
        this.loading = false;
      }

      this.$emit("message-sent", input);
      this.$nextTick(() => {
        this.scrollToBottom();
      });
    },

    toggleThinking(index) {
      this.expandedThinking[index] = !this.expandedThinking[index];
    },

    formatMarkdown(text) {
      if (!text) return "";
      return marked(text);
    },

    scrollToBottom() {
      if (this.$refs.messagesContainer) {
        const container = this.$refs.messagesContainer;
        container.scrollTop = container.scrollHeight;
      }
    },
  },
  watch: {
    chatHistory() {
      this.$nextTick(() => {
        this.scrollToBottom();
      });
    }
  },
  mounted() {
    this.scrollToBottom();
  },
};
</script>

<style scoped>
.chat-section {
  height: 90%;
  flex: 1;
  display: flex;
  flex-direction: column;
}

.chat-messages {
  flex: 1;
  padding: 15px;
  overflow-y: auto;
}

.message {
  margin-bottom: 15px;
  padding: 12px 15px;
  border-radius: 8px;
  max-width: 80%;
  animation: fadeIn 0.3s ease;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.user-message {
  background-color: #e3f2fd;
  margin-left: auto;
  color: #0d47a1;
  border-bottom-right-radius: 4px;
}

.ai-message {
  background-color: #f5f5f5;
  margin-right: auto;
  color: #333;
  border-bottom-left-radius: 4px;
}

.message-content {
  line-height: 1.5;
}

/* Markdown 样式 */
.message-content :deep(h1),
.message-content :deep(h2),
.message-content :deep(h3),
.message-content :deep(h4),
.message-content :deep(h5),
.message-content :deep(h6) {
  margin-top: 0.8em;
  margin-bottom: 0.5em;
}

.message-content :deep(p) {
  margin: 0.5em 0;
}

.message-content :deep(ul),
.message-content :deep(ol) {
  padding-left: 1.5em;
  margin: 0.5em 0;
}

.message-content :deep(pre) {
  background-color: #f0f0f0;
  padding: 0.8em;
  border-radius: 4px;
  overflow-x: auto;
  margin: 0.7em 0;
}

.message-content :deep(code) {
  font-family: monospace;
  background-color: #f0f0f0;
  padding: 2px 4px;
  border-radius: 3px;
}

.message-content :deep(blockquote) {
  border-left: 4px solid #e0e0e0;
  padding-left: 1em;
  margin: 0.7em 0;
  color: #666;
}

.message-content :deep(a) {
  color: #2196f3;
  text-decoration: none;
}

.message-content :deep(a:hover) {
  text-decoration: underline;
}

/* 思考过程样式 */
.thinking-section {
  margin-bottom: 10px;
  border-radius: 6px;
  border: 1px solid #e0e0e0;
  overflow: hidden;
}

.thinking-header {
  display: flex;
  align-items: center;
  padding: 8px 12px;
  background-color: #f5f5f5;
  cursor: pointer;
  font-weight: 500;
  color: #757575;
  transition: background-color 0.2s;
}

.thinking-header:hover {
  background-color: #eeeeee;
}

.thinking-icon {
  margin-right: 8px;
  transition: transform 0.3s ease;
}

.thinking-icon.expanded {
  transform: rotate(90deg);
}

.thinking-content {
  padding: 12px;
  background-color: #f9f9f9;
  border-top: 1px solid #e0e0e0;
  font-size: 0.95em;
  color: #555;
}

.chat-input-container {
  display: flex;
  padding: 15px;
  border-top: 1px solid #e0e0e0;
  background-color: white;
}

.chat-input {
  flex: 1;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 5px;
  margin-right: 10px;
  outline: none;
  transition: border-color 0.3s;
}

.chat-input:focus {
  border-color: #64b5f6;
}

.send-btn {
  padding: 10px 20px;
  background-color: #a7d1ff;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  color: #0056b3;
  transition: background-color 0.2s;
}

.send-btn:hover {
  background-color: #90c2ff;
}

/* 新增加载中提示样式 */
.loading-message {
  font-style: italic;
  color: #999;
  margin-bottom: 15px;
  max-width: 80%;
  padding: 12px 15px;
  border-radius: 8px;
  background-color: #f0f0f0;
}

.spinner {
  display: inline-block;
  width: 16px;
  height: 16px;
  margin-right: 8px;
  border: 2px solid #999;
  border-top-color: transparent;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  vertical-align: middle;
}

@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}
</style>
