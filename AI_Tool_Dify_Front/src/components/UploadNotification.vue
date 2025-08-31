<template>
  <transition name="notification">
    <div v-if="show" :class="['notification', `notification-${type}`]">
      <div class="notification-icon">
        <svg
          v-if="type === 'success'"
          viewBox="0 0 24 24"
          width="24"
          height="24"
        >
          <circle cx="12" cy="12" r="10" fill="#4caf50" />
          <path
            d="M9,16 L5,12 L6.41,10.58 L9,13.17 L14.59,7.58 L16,9"
            fill="white"
          />
        </svg>
        <svg
          v-else-if="type === 'error'"
          viewBox="0 0 24 24"
          width="24"
          height="24"
        >
          <circle cx="12" cy="12" r="10" fill="#f44336" />
          <line x1="8" y1="8" x2="16" y2="16" stroke="white" stroke-width="2" />
          <line x1="16" y1="8" x2="8" y2="16" stroke="white" stroke-width="2" />
        </svg>
        <svg v-else viewBox="0 0 24 24" width="24" height="24">
          <circle cx="12" cy="12" r="10" fill="#2196f3" />
          <circle cx="12" cy="12" r="2" fill="white" />
          <circle cx="12" cy="6" r="2" fill="white" />
          <circle cx="12" cy="18" r="2" fill="white" />
        </svg>
      </div>
      <div class="notification-content">
        <div class="notification-title">{{ title }}</div>
        <div v-if="message" class="notification-message">{{ message }}</div>
      </div>
      <button @click="$emit('close')" class="notification-close">×</button>
    </div>
  </transition>
</template>

<script>
export default {
  name: "UploadNotification",
  props: {
    show: {
      type: Boolean,
      default: false,
    },
    type: {
      type: String,
      default: "info",
      validator: (value) => ["info", "success", "error"].includes(value),
    },
    title: {
      type: String,
      required: true,
    },
    message: {
      type: String,
      default: "",
    },
    autoClose: {
      type: Boolean,
      default: true,
    },
    duration: {
      type: Number,
      default: 3000,
    },
  },
  watch: {
    show(newVal) {
      if (newVal && this.autoClose) {
        setTimeout(() => {
          this.$emit("close");
        }, this.duration);
      }
    },
  },
};
</script>

<style scoped>
.notification {
  position: fixed;
  bottom: 20px;
  right: 20px;
  background-color: white;
  border-radius: 4px;
  box-shadow: 0 3px 10px rgba(0, 0, 0, 0.15);
  display: flex;
  align-items: center;
  padding: 10px 15px;
  max-width: 350px;
  z-index: 1000;
}

.notification-success {
  border-left: 4px solid #4caf50;
}

.notification-error {
  border-left: 4px solid #f44336;
}

.notification-info {
  border-left: 4px solid #2196f3;
}

.notification-icon {
  margin-right: 15px;
}

.notification-content {
  flex: 1;
}

.notification-title {
  font-weight: 500;
  margin-bottom: 2px;
}

.notification-message {
  font-size: 12px;
  color: #666;
}

.notification-close {
  background: none;
  border: none;
  font-size: 20px;
  color: #999;
  cursor: pointer;
  padding: 0 5px;
}

.notification-enter-active,
.notification-leave-active {
  transition: all 0.3s;
}

.notification-enter-from,
.notification-leave-to {
  opacity: 0;
  transform: translateY(30px);
}
</style>
