<template>
  <div class="app-container">
    <Sidebar @open-settings="handleSettingsOpen" />

    <div class="main-section">
      <div class="top-section">
        <div class="chat-column">
          <div class="content-header">AI对话</div>
          <ChatSection />
        </div>

        <div class="code-column">
          <div class="content-header">代码生成</div>
          <CodeSection />
        </div>
      </div>

      <div class="mindmap-container">
        <div class="section-header">思维导图</div>
        <MindMap />
      </div>
    </div>

    <SettingsModal
      :show="showSettingsModal"
      :type="settingsType"
      @close="showSettingsModal = false"
    />

    <UploadNotification
      :show="notification.show"
      :type="notification.type"
      :title="notification.title"
      :message="notification.message"
      @close="clearNotification"
    />
  </div>
</template>

<script>
import Sidebar from "./components/Sidebar.vue";
import ChatSection from "./components/ChatSection.vue";
import MindMap from "./components/MindMap.vue";
import CodeSection from "./components/CodeSection.vue";
import SettingsModal from "./components/SettingsModal.vue";
import UploadNotification from "./components/UploadNotification.vue";
import { mapState } from "vuex";

export default {
  name: "App",
  components: {
    Sidebar,
    ChatSection,
    MindMap,
    CodeSection,
    SettingsModal,
    UploadNotification,
  },
  data() {
    return {
      showSettingsModal: false,
      settingsType: "model",
    };
  },
  computed: {
    ...mapState({
      isUploading: (state) => state.knowledgeBase.isUploading,
      uploadError: (state) => state.knowledgeBase.lastError,
      knowledgeBaseFiles: (state) => state.knowledgeBase.files,
      notification: (state) => state.notification,
    }),
  },
  watch: {
    isUploading(newVal, oldVal) {
      // 只有当状态从false变为true时才显示通知，避免初始加载时触发
      if (newVal && oldVal === false) {
        this.$store.dispatch("showNotification", {
          type: "info",
          title: "文件上传中",
          message: "正在处理您的文件上传请求...",
        });
      }
    },
    uploadError(newVal) {
      if (newVal) {
        this.$store.dispatch("showNotification", {
          type: "error",
          title: "上传失败",
          message: newVal,
        });
      }
    },
    knowledgeBaseFiles(newVal, oldVal) {
      if (newVal.length > oldVal.length) {
        const latestFile = newVal[newVal.length - 1];
        this.$store.dispatch("showNotification", {
          type: "success",
          title: "上传成功",
          message: `文件 ${latestFile.name} 已成功上传到知识库`,
        });
      }
    },
  },
  methods: {
    handleSettingsOpen(type) {
      this.settingsType = type;
      this.showSettingsModal = true;
    },

    clearNotification() {
      this.$store.commit("CLEAR_NOTIFICATION");
    },
  },
};
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  font-family: Arial, sans-serif;
}

.app-container {
  display: flex;
  height: 100vh;
  width: 100%;
  padding: 10px;
  background-color: #f5f5f5;
}

.main-section {
  flex: 1;
  display: flex;
  flex-direction: column;
  margin-left: 10px;
}

.top-section {
  height: 70%;
  display: flex;
  flex: 1;
  margin-bottom: 10px;
}

.chat-column {
  flex: 1;
  display: flex;
  flex-direction: column;
  border: 1px solid #e0e0e0;
  border-radius: 15px;
  margin-right: 10px;
  background-color: #fff;
  overflow: hidden;
}

.code-column {
  width: 40%;
  display: flex;
  flex-direction: column;
  border: 1px solid #e0e0e0;
  border-radius: 15px;
  background-color: #fff;
  overflow: hidden;
}

.content-header {

  text-align: center;
  padding: 15px;
  font-size: 18px;
  font-weight: bold;
  border-bottom: 1px solid #e0e0e0;
}

.mindmap-container {
  height: 35%;
  border: 1px solid #e0e0e0;
  border-radius: 15px;
  background-color: #fff;
  overflow: hidden;
}

.section-header {
  padding: 10px 15px;
  font-weight: bold;
  border-bottom: 1px solid #e0e0e0;
}
</style>
