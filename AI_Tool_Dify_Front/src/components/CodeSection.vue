<template>
  <div class="code-content">
    <div v-if="codeBlocks.length" class="code-list">
      <div
          v-for="(block, index) in codeBlocks"
          :key="index"
          class="code-block"
      >
        <div class="code-header">
          <span class="code-lang">{{ block.lang || 'plaintext' }}</span>
          <button class="save-btn" @click="saveCode(block)">
            <svg viewBox="0 0 24 24" width="16" height="16" class="action-icon">
              <path
                  d="M19 9h-4V3H9v6H5l7 7 7-7zM5 18v2h14v-2H5z"
                  fill="currentColor"
              />
            </svg>
            保存代码
          </button>
        </div>
        <pre
            v-html="highlightCode(block.code, block.lang)"
            :class="`language-${block.lang || 'plaintext'}`"
            class="code-block-pre"
        ></pre>
      </div>
    </div>
    <div v-else class="code-placeholder">代码将在对话后生成</div>
  </div>
</template>

<script>
import { mapState } from "vuex";
import Prism from "prismjs";
import "prismjs/themes/prism.css";

export default {
  name: "CodeSection",
  computed: {
    ...mapState(["codeBlocks"]),
  },
  methods: {
    highlightCode(code, lang) {
      if (!code) return "";
      const language = Prism.languages[lang] || Prism.languages.javascript;
      return Prism.highlight(code, language, lang);
    },
    saveCode(block) {
      if (!block.code) return;

      const blob = new Blob([block.code], { type: "text/plain" });
      const url = URL.createObjectURL(blob);

      const a = document.createElement("a");
      a.href = url;

      // 根据语言扩展名生成文件名
      a.download = `generated-code-${new Date().getTime()}.${this.getExtension(block.lang)}`;
      document.body.appendChild(a);
      a.click();

      setTimeout(() => {
        document.body.removeChild(a);
        URL.revokeObjectURL(url);
      }, 100);

      this.$store.dispatch("showNotification", {
        type: "success",
        title: "代码已保存",
        message: "代码已成功下载到您的计算机",
      });
    },
    getExtension(lang) {
      if (!lang) return "txt";
      switch (lang.toLowerCase()) {
        case "javascript":
        case "js":
          return "js";
        case "python":
          return "py";
        case "java":
          return "java";
        case "csharp":
        case "c#":
          return "cs";
        case "html":
          return "html";
        case "css":
          return "css";
        case "json":
          return "json";
        case "mermaid":
          return "mmd";
        default:
          return "txt";
      }
    },
  },
};
</script>

<style scoped>
.code-content {
  flex: 1;
  padding: 15px;
  overflow: auto;
  background-color: #f8f9fa;
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.code-list {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.code-block {
  background-color: #f1f1f1;
  border-radius: 6px;
  padding: 12px 16px;
  box-shadow: 0 0 4px rgba(0,0,0,0.1);
  position: relative;
}

.code-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.code-lang {
  font-weight: 600;
  color: #1a91ff;
  font-family: monospace;
  text-transform: uppercase;
}

.save-btn {
  display: flex;
  align-items: center;
  gap: 5px;
  background-color: #1a91ff;
  color: white;
  border: none;
  border-radius: 4px;
  padding: 6px 12px;
  font-size: 12px;
  cursor: pointer;
  transition: background-color 0.2s;
}

.save-btn:hover {
  background-color: #0056b3;
}

.action-icon {
  opacity: 0.9;
}

.code-block-pre {
  white-space: pre-wrap;
  font-family: monospace;
  margin: 0;
  padding: 10px;
  border-left: 3px solid #1a91ff;
  color: #333;
  border-radius: 4px;
}

.code-placeholder {
  color: #999;
  font-size: 16px;
  text-align: center;
  width: 100%;
  margin-top: 40px;
}
</style>
