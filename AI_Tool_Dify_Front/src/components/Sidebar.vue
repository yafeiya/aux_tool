<template>
  <div class="sidebar">
    <header class="sidebar-header">
      <h2>智能辅助建模工具</h2>
    </header>

    <section class="upload-section">
      <FileUpload
          :datasetId="datasetId"
          :apiKey="apiKey"
          @upload-success="handleUploadSuccess"
          @upload-error="handleUploadError"
      />
    </section>

    <section class="settings-section">
      <button class="settings-btn" @click="showModelSelector = true">大模型API</button>
      <button class="settings-btn" @click="openDifyAdmin">知识库管理</button>
    </section>

    <!-- 大模型选择器弹窗 -->
    <div v-if="showModelSelector" class="model-selector-overlay" @click.self="showModelSelector = false">
      <div class="model-selector">
        <h3>选择大模型</h3>
        <ul class="model-list">
          <li v-for="model in modelOptions" :key="model.id">
            <button class="model-option" @click="selectModel(model)">
              {{ model.name }}
            </button>
          </li>
        </ul>
        <button class="close-btn" @click="showModelSelector = false">关闭</button>
      </div>
    </div>

    <section class="documents-section">
      <div v-if="isUploading" class="status-message">
        <p>加载知识库文档中...</p>
      </div>

      <div v-else-if="uploadError" class="status-message error">
        <p>错误: {{ uploadError }}</p>
      </div>

      <div v-else>
        <h4 class="documents-title">知识库文档列表</h4>

        <ul v-if="documents.length > 0" class="document-list">
          <li v-for="doc in documents" :key="doc.id" class="document-item">
            <span class="document-title">{{ doc.title || doc.name || '未命名文档' }}</span>
            <span class="document-date">{{ formatDocumentInfo(doc) }}</span>
          </li>
        </ul>

        <p v-else class="empty-message">暂无知识库文档</p>
      </div>
    </section>
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex'
import FileUpload from './FileUpload.vue'

export default {
  name: 'Sidebar',
  components: { FileUpload },
  emits: ['open-settings', 'model-selected'],
  data() {
    return {
      datasetId: "2750d917-3f80-47e6-ab14-01f239ebcf27",
      apiKey: "dataset-3qH43ExwpT2qaYylvDNzIiZW",
      difyAdminUrl: "http://localhost/datasets", // Dify管理后台URL
      showModelSelector: false,
      modelOptions: [
        { id: 'gpt-3.5-turbo', name: 'GPT-3.5 Turbo' },
        { id: 'gpt-4', name: 'GPT-4' },
        { id: 'claude-3-opus', name: 'Claude 3 Opus' },
        { id: 'claude-3-sonnet', name: 'Claude 3 Sonnet' },
        { id: 'llama-3', name: 'Llama 3' }
      ]
    }
  },
  computed: {
    ...mapState({
      documents: state => state.knowledgeBase.documents,
      isUploading: state => state.knowledgeBase.isUploading,
      uploadError: state => state.knowledgeBase.lastError,
    }),
  },
  methods: {
    ...mapActions(['fetchKnowledgeDocuments']),
    handleUploadSuccess() {
      this.fetchKnowledgeDocuments()
    },
    handleUploadError(error) {
      console.error("上传错误:", error)
    },
    formatDocumentInfo(doc) {
      if (!doc.created_at) return ''
      return new Date(doc.created_at * 1000).toLocaleDateString()
    },
    selectModel(model) {
      this.$emit('model-selected', model.id)
      this.showModelSelector = false
      console.log(`已选择模型: ${model.name}`)
    },
    openDifyAdmin() {
      // 在新窗口打开Dify管理后台
      window.open(this.difyAdminUrl, '_blank')
    }
  },
  mounted() {
    // 传递 isInitialLoad: true 参数，表示这是组件初始化时的加载
    this.fetchKnowledgeDocuments({ isInitialLoad: true })
  },
}
</script>

<style scoped>
.sidebar {
  padding: 24px 20px;
  background-color: #fff;
  height: 100%;
  display: flex;
  flex-direction: column;
  font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
  color: #333;
}

.sidebar-header h2 {
  font-weight: 700;
  font-size: 1.8rem;
  margin-bottom: 10px;
  color: #222;
  user-select: none;
}

.upload-section {
  display: flex;
  margin-bottom: 10px;
}

.settings-section {
  display: flex;
  flex-direction: column;
  gap: 12px;
  margin-bottom: 20px;
}

.settings-btn {
  background-color: #1a91ff;
  color: white;
  font-weight: 600;
  border: none;
  border-radius: 6px;
  padding: 12px 0;
  cursor: pointer;
  font-size: 1rem;
  transition: background-color 0.25s ease;
  user-select: none;
}

.settings-btn:hover {
  background-color: #0f6cd1;
}

.documents-section {
  flex-grow: 1;
  overflow-y: auto;
}

.status-message {
  text-align: center;
  font-style: italic;
  color: #666;
  padding: 20px 0;
  user-select: none;
}

.status-message.error {
  color: #d9534f;
}

.documents-title {
  font-size: 1.2rem;
  font-weight: 700;
  margin-bottom: 12px;
  border-bottom: 2px solid #1a91ff;
  padding-bottom: 6px;
  user-select: none;
}

.document-list {
  list-style: none;
  padding: 0;
  margin: 0;
  max-height: 300px; /* 限高，超出滚动 */
  overflow-y: auto;
  border: 1px solid #ddd;
  border-radius: 8px;
}

.document-item {
  display: flex;
  justify-content: space-between;
  padding: 10px 14px;
  border-bottom: 1px solid #eee;
  transition: none;
  user-select: none;
}

.document-item:last-child {
  border-bottom: none;
}

.document-item:hover {
  background-color: #f0f8ff;
}

.document-title {
  font-weight: 600;
  color: #1a1a1a;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  max-width: 70%;
}

.document-date {
  color: #888;
  font-size: 0.9rem;
  white-space: nowrap;
  flex-shrink: 0;
  user-select: text;
}

.empty-message {
  text-align: center;
  font-style: italic;
  color: #999;
  padding: 40px 0;
  user-select: none;
}

/* 模型选择器弹窗样式 */
.model-selector-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.model-selector {
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.2);
  padding: 20px;
  width: 320px;
  max-width: 90vw;
}

.model-selector h3 {
  margin-top: 0;
  margin-bottom: 16px;
  font-size: 18px;
  text-align: center;
  color: #333;
}

.model-list {
  list-style: none;
  padding: 0;
  margin: 0 0 16px 0;
}

.model-list li {
  margin-bottom: 8px;
}

.model-option {
  width: 100%;
  padding: 12px;
  border: 1px solid #1a91ff;
  background: white;
  color: #1a91ff;
  border-radius: 6px;
  cursor: pointer;
  font-size: 14px;
  font-weight: 500;
  text-align: center;
  transition: all 0.2s ease;
}

.model-option:hover {
  background-color: #f0f8ff;
}

.close-btn {
  display: block;
  width: 100%;
  padding: 10px;
  background-color: #f0f0f0;
  border: none;
  border-radius: 6px;
  color: #333;
  font-weight: 500;
  cursor: pointer;
  transition: background-color 0.2s;
}

.close-btn:hover {
  background-color: #e0e0e0;
}
</style>
