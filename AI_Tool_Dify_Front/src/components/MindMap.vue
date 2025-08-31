<template>
  <div className="mindmap-container">
    <div ref="mermaidContainer" className="mermaid-diagram">
      <div v-if="renderedSvg" v-html="renderedSvg"></div>
      <pre v-else className="mermaid-code">{{ diagramCode }}</pre>
      <!-- 添加保存按钮 -->
      <div class="actions-bar">
        <button class="save-btn" @click="saveDiagram" v-if="diagramCode">
          <svg viewBox="0 0 24 24" width="16" height="16" class="action-icon">
            <path
                d="M19 9h-4V3H9v6H5l7 7 7-7zM5 18v2h14v-2H5z"
                fill="currentColor"
            />
          </svg>
          保存
        </button>
      </div>
    </div>


  </div>
</template>

<script>
import mermaid from 'mermaid'
import {mapState} from 'vuex'

export default {
  props: {
    // 可以接收单个图表代码作为属性
    code: {
      type: String,
      default: null
    },
    // 可以指定使用 store 中的第几个图表
    index: {
      type: Number,
      default: 0
    }
  },
  computed: {
    ...mapState({
      mindMapArray: state => state.mindMap
    }),
    // 优先使用传入的代码，如果没有则从 store 中获取
    diagramCode() {
      if (this.code) {
        return this.code;
      }

      // 从 store 中获取指定索引的图表代码
      if (this.mindMapArray && this.mindMapArray.length > this.index) {
        return this.mindMapArray[this.index];
      }

      // 如果没有可用的图表代码，返回空字符串
      return '';
    }
  },
  data() {
    return {
      renderedSvg: null
    }
  },
  watch: {
    // 当图表代码变化时重新渲染
    diagramCode: {
      immediate: true,
      handler(newCode) {
        if (newCode) {
          this.renderDiagram();
        }
      }
    }
  },
  mounted() {
    mermaid.initialize({
      startOnLoad: false,
      theme: 'default',
      securityLevel: 'loose'
    });
  },
  methods: {
    async renderDiagram() {
      await this.$nextTick();
      const container = this.$refs.mermaidContainer;
      if (!container || !this.diagramCode) return;

      try {
        // Mermaid 9.x API: render 返回 {svg, bindFunctions}
        const id = `mermaid-${Date.now()}`;
        const {svg} = await mermaid.render(id, this.diagramCode);
        this.renderedSvg = svg;
      } catch (err) {
        console.error('Mermaid 渲染错误:', err);
        this.renderedSvg = `<p style="color: red;">思维导图渲染失败: ${err.message}</p>`;
      }
    },

    // 添加保存思维导图的方法
    saveDiagram() {
      if (!this.diagramCode) return;

      // 方法1: 保存为SVG文件
      if (this.renderedSvg) {
        const blob = new Blob([this.renderedSvg], { type: 'image/svg+xml' });
        this.downloadFile(blob, 'mindmap.svg');
        return;
      }

      // 方法2: 如果SVG不可用，保存为文本文件
      const blob = new Blob([this.diagramCode], { type: 'text/plain' });
      this.downloadFile(blob, 'mindmap.mmd');
    },

    // 通用下载文件方法
    downloadFile(blob, filename) {
      const url = URL.createObjectURL(blob);
      const a = document.createElement('a');
      a.href = url;
      a.download = filename;
      document.body.appendChild(a);
      a.click();
      document.body.removeChild(a);
      URL.revokeObjectURL(url);
    }
  }
}
</script>

<style scoped>
.mindmap-container {
  padding: 10px;
  background: #fff;
  border-radius: 8px;
  width: 100%;
  height: 100%;
  position: relative;
}

.mermaid-diagram {
  border: 1px solid #ddd;
  padding: 12px;
  border-radius: 6px;
  background-color: #f9f9f9;
  overflow-x: auto;
  height: calc(100% - 40px); /* 为按钮留出空间 */
  text-align: center; /* 备用，SVG 是块级后不影响 */
}

/* 关键：让 svg 宽度撑满并居中 */
.mermaid-diagram svg {
  display: block;
  margin: 0 auto;
  width: 100%;
  height: auto;
  max-width: none;
}

.mermaid-code {
  white-space: pre-wrap;
  font-family: monospace;
  color: #666;
  user-select: text;
}

/* 添加按钮样式 */
.actions-bar {
  margin-top: 10px;
  display: flex;
  justify-content: flex-end;
}

.save-btn {
  display: flex;
  align-items: center;
  gap: 5px;
  padding: 6px 12px;
  border: none;
  background-color: #1a91ff;
  color: white;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
  transition: background-color 0.2s;
}

.save-btn:hover {
  background-color: #0056b3;
}

.action-icon {
  vertical-align: middle;
}
</style>
