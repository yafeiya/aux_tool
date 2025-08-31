<template>
  <div v-if="show" class="modal-backdrop">
    <div class="modal-content">
      <div class="modal-header">
        <h3>{{ title }}</h3>
        <button class="close-btn" @click="$emit('close')">×</button>
      </div>

      <div class="modal-body">
        <template v-if="type === 'model'">
          <div class="form-group">
            <label>选择大模型API</label>
            <select v-model="localModelApi" class="form-control">
              <option value="gpt-3.5">GPT-3.5</option>
              <option value="gpt-4">GPT-4</option>
              <option value="claude-3">Claude-3</option>
              <option value="gemini-pro">Gemini Pro</option>
            </select>
          </div>
        </template>

        <template v-else-if="type === 'history'">
          <div class="form-group">
            <label>消息数量上限</label>
            <input
              type="number"
              v-model.number="localHistoryCount"
              min="5"
              max="50"
              class="form-control"
            />
            <small>设置范围: 5-50</small>
          </div>
        </template>
      </div>

      <div class="modal-footer">
        <button class="cancel-btn" @click="$emit('close')">取消</button>
        <button class="save-btn" @click="saveSettings">保存</button>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState } from "vuex";

export default {
  name: "SettingsModal",
  props: {
    show: {
      type: Boolean,
      default: false,
    },
    type: {
      type: String,
      required: true,
      validator: (value) => ["model", "history"].includes(value),
    },
  },
  data() {
    return {
      localModelApi: this.$store.state.settings.modelApi,
      localHistoryCount: this.$store.state.settings.historyMessageCount,
    };
  },
  computed: {
    ...mapState({
      modelApi: (state) => state.settings.modelApi,
      historyMessageCount: (state) => state.settings.historyMessageCount,
    }),
    title() {
      return this.type === "model" ? "大模型API设置" : "历史消息数设置";
    },
  },
  watch: {
    show(newVal) {
      if (newVal) {
        // Reset local values when modal is opened
        this.localModelApi = this.modelApi;
        this.localHistoryCount = this.historyMessageCount;
      }
    },
  },
  methods: {
    saveSettings() {
      if (this.type === "model") {
        this.$store.dispatch("updateModelApi", this.localModelApi);
      } else if (this.type === "history") {
        this.$store.dispatch("updateHistoryCount", this.localHistoryCount);
      }
      this.$emit("close");
    },
  },
};
</script>

<style scoped>
.modal-backdrop {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-content {
  background-color: white;
  border-radius: 10px;
  width: 400px;
  max-width: 90%;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 15px 20px;
  border-bottom: 1px solid #eee;
}

.close-btn {
  background: none;
  border: none;
  font-size: 24px;
  cursor: pointer;
  color: #999;
}

.modal-body {
  padding: 20px;
}

.form-group {
  margin-bottom: 15px;
}

label {
  display: block;
  margin-bottom: 8px;
  font-weight: bold;
}

.form-control {
  width: 100%;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 5px;
  font-size: 16px;
}

small {
  display: block;
  margin-top: 5px;
  color: #666;
}

.modal-footer {
  padding: 15px 20px;
  display: flex;
  justify-content: flex-end;
  gap: 10px;
  border-top: 1px solid #eee;
}

.cancel-btn,
.save-btn {
  padding: 8px 16px;
  border-radius: 5px;
  cursor: pointer;
  font-size: 14px;
}

.cancel-btn {
  background-color: #f5f5f5;
  border: 1px solid #ddd;
  color: #333;
}

.save-btn {
  background-color: #1a91ff;
  border: none;
  color: white;
}
</style>
