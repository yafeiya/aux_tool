<template>
  <div class="file-upload">
    <input
        type="file"
        ref="fileInput"
        @change="handleFileChange"
        style="display: none"
        accept=".txt,.pdf,.doc,.docx,.md"
    />
    <button class="upload-btn" @click="triggerFileInput" :disabled="isUploading">
      <svg
          class="upload-icon"
          viewBox="0 0 200 200"
          xmlns="http://www.w3.org/2000/svg"
      >
        <path
            d="M30,40 H150 Q170,40 170,60 V150 Q170,170 150,170 H30 Q10,170 10,150 V60 Q10,40 30,40 Z"
            stroke="#1a91ff"
            stroke-width="12"
            fill="none"
        />
        <line
            x1="40"
            y1="80"
            x2="120"
            y2="80"
            stroke="#1a91ff"
            stroke-width="12"
            stroke-linecap="round"
        />
        <line
            x1="40"
            y1="110"
            x2="120"
            y2="110"
            stroke="#1a91ff"
            stroke-width="12"
            stroke-linecap="round"
        />
        <line
            x1="40"
            y1="140"
            x2="80"
            y2="140"
            stroke="#1a91ff"
            stroke-width="12"
            stroke-linecap="round"
        />
        <circle cx="150" cy="120" r="8" fill="#1a91ff" />
        <path
            d="M150,70 L150,40 L180,70 L150,100 L150,70"
            fill="#1a91ff"
            stroke="#1a91ff"
            stroke-width="0"
        />
        <line
            x1="180"
            y1="70"
            x2="180"
            y2="120"
            stroke="#1a91ff"
            stroke-width="8"
            stroke-linecap="round"
        />
        <line
            x1="165"
            y1="120"
            x2="180"
            y2="120"
            stroke="#1a91ff"
            stroke-width="8"
            stroke-linecap="round"
        />
        <line
            x1="190"
            y1="100"
            x2="190"
            y2="120"
            stroke="#1a91ff"
            stroke-width="8"
            stroke-linecap="round"
        />
      </svg>
      <span>{{ isUploading ? '上传中...' : '知识库上传' }}</span>
    </button>
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex';

export default {
  name: "FileUpload",
  props: {
    datasetId: {
      type: String,
      required: true,
    },
    apiKey: {
      type: String,
      required: true,
    },
  },
  computed: {
    ...mapState({
      isUploading: state => state.knowledgeBase.isUploading,
    }),
  },
  methods: {
    ...mapActions(['uploadKnowledgeFile']),
    triggerFileInput() {
      if (!this.isUploading) {
        this.$refs.fileInput.click();
      }
    },
    async handleFileChange(event) {
      const file = event.target.files[0];
      if (!file) return;

      try {
        const resData = await this.uploadKnowledgeFile({
          file,
          datasetId: this.datasetId,
          apiKey: this.apiKey,
        });

        alert("文件上传成功！");
        this.$emit("upload-success", resData);

        // 重置文件输入
        this.$refs.fileInput.value = "";
      } catch (error) {
        console.error("上传错误:", error);
        this.$emit("upload-error", error);
      }
    },
  },
};
</script>

<style scoped>
.file-upload {
  width: 100%;
  margin: 15px 0;
}

.upload-btn {
  width: 100%;
  display: flex;
  align-items: center;      /* 垂直居中 */
  justify-content: center;  /* 水平居中 */
  background: none;
  border: none;
  padding: 8px 10px;
  cursor: pointer;
  color: #1a91ff;
  font-weight: 500;
  border-radius: 5px;
  font-size: 1rem;
  transition: background-color 0.2s;
}

.upload-btn:disabled {
  cursor: not-allowed;
  opacity: 0.6;
}

.upload-btn:hover:not(:disabled) {
  background-color: rgba(26, 145, 255, 0.1);
}

.upload-icon {
  width: 24px;
  height: 24px;
  margin-right: 8px;
}
</style>
